package repositories

import (
	"clean-arc/modules/entities"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type usersRepo struct {
	Db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) entities.UsersRepository {
	return &usersRepo{
		Db: db,
	}
}

func (r *usersRepo) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {

	query := `
	INSERT INTO public.users("username","password","email","role","create_at")
	VALUES ($1, $2, $3, $4, $5)
	RETURNING "id", "username";
	`

	// Initail a user object
	user := new(entities.UsersRegisterRes)

	// Query part
	rows, err := r.Db.Queryx(query, &req.Username, &req.Password, &req.Email, &req.Role, &req.CreateAt)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	//defer r.Db.Close()

	return user, nil
}

func (r *usersRepo) GetAllUsers(req *entities.GetAllUserReq) ([]entities.UsersAllRes, error) {
	query := `SELECT "id", "username", "email", "create_at" FROM public.users;`

	// Query part
	rows, err := r.Db.Queryx(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	var users []entities.UsersAllRes

	for rows.Next() {
		var u entities.UsersAllRes
		if err := rows.StructScan(&u); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		users = append(users, u)

	}
	// defer r.Db.Close()

	return users, nil
}

func (r *usersRepo) FindOneUser(username string) (*entities.UsersPassport, error) {
	query := `
	SELECT
	"id",
	"username",
	"password",
	"role"
	FROM public.users
	WHERE "username" = $1;
	`

	res := new(entities.UsersPassport)
	if err := r.Db.Get(res, query, username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		fmt.Println(err.Error())
		return nil, errors.New("error, user not found")
	}
	return res, nil
}
