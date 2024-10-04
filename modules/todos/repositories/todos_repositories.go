package repositories

import (
	"clean-arc/modules/entities"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"time"
)

type todoRepo struct {
	Db          *sqlx.DB
	redisClient *redis.Client
}

func NewTodosRepository(db *sqlx.DB, redisClient *redis.Client) entities.TodosRepository {
	return &todoRepo{
		Db:          db,
		redisClient: redisClient,
	}
}

func (r *todoRepo) GetAllTodos(req *entities.TodosReq) ([]entities.TodosRes, error) {
	key := "repositories:GetAllTodos"
	versionKey := "repositories:TodosVersion"

	// Fetch the version from Redis
	cachedVersion, err := r.redisClient.Get(context.Background(), versionKey).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	// Fetch cached todos from Redis
	todosJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		// Unmarshal cached todos
		var cachedTodos []entities.TodosRes
		err = json.Unmarshal([]byte(todosJson), &cachedTodos)
		if err == nil {
			// If cached todos are present and the version matches, return cached todos
			if cachedVersion != "" {
				return cachedTodos, nil
			}
		}
	}

	// Query the database for todos and their last modification timestamp
	query := `SELECT * FROM public.todos ORDER BY id;`
	rows, err := r.Db.Queryx(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var todos []entities.TodosRes
	for rows.Next() {
		var t entities.TodosRes

		if err := rows.StructScan(&t); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		todos = append(todos, t)
	}

	// Marshal todos to JSON
	data, err := json.Marshal(todos)
	if err != nil {
		return nil, err
	}

	// Get the current timestamp/version for caching
	newVersion := time.Now().Format(time.RFC3339)
	redisClearTime := time.Second * 10

	// Set todos and version in Redis
	err = r.redisClient.Set(context.Background(), key, string(data), redisClearTime).Err()
	if err != nil {
		return nil, err
	}
	err = r.redisClient.Set(context.Background(), versionKey, newVersion, redisClearTime).Err()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *todoRepo) AddTodos(req *entities.TodosReq) (*entities.TodosRes, error) {
	query := `INSERT INTO public.todos (created_at, title, complete) VALUES ($1,$2,$3) RETURNING id;`
	var id int
	err := r.Db.QueryRow(query, req.CreatedAt, req.Title, req.Complete).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// Fetch the newly created todo to return it
	newTodo, err := r.GetTodoByID(id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return newTodo, nil
}

func (r *todoRepo) UpdateAllTodos(req []entities.TodosReq) ([]entities.TodosRes, error) {
	query := `UPDATE public.todos SET title = $1, complete = $2 WHERE id = $3 RETURNING id, title, complete;`
	var updatedTodos []entities.TodosRes

	// Start a transaction
	tx, err := r.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback() // Rollback if any error occurs
		} else {
			err = tx.Commit() // Commit if no error
		}
	}()

	// Loop through each todo in the request array
	for _, reqs := range req {
		var updatedTodo entities.TodosRes

		// Execute the query and return the updated todo
		row := tx.QueryRow(query, reqs.Title, reqs.Complete, reqs.ID)
		err := row.Scan(&updatedTodo.ID, &updatedTodo.Title, &updatedTodo.Complete)

		if err != nil {
			if err == sql.ErrNoRows {

				fmt.Printf("No rows updated for ID %d\n", reqs.ID)
				continue
			}
			fmt.Println(err.Error()) // Log the error for debugging
			return nil, err
		}

		// Append the updated todo to the response slice
		updatedTodos = append(updatedTodos, updatedTodo)
	}

	// Commit the transaction if all queries succeed
	if err != nil {
		return nil, err
	}

	return updatedTodos, nil
}

func (r *todoRepo) DeleteTodos(id int64) error {
	query := `DELETE FROM public.todos WHERE id = $1;`
	_, err := r.Db.Exec(query, id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

// GetTodoByID retrieves a single todo by its ID
func (r *todoRepo) GetTodoByID(id int) (*entities.TodosRes, error) {
	query := `SELECT * FROM public.todos WHERE id = $1;`
	var todo entities.TodosRes
	err := r.Db.Get(&todo, query, id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &todo, nil
}
