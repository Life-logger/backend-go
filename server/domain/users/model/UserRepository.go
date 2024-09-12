package users

import (
	"context"
	"database/sql"
)

type (
	UsersRepository interface {
		Save(User)
	}

	usersRepositoryImpl struct {
		*sql.Conn
		tx  *sql.Tx
		ctx context.Context
	}
)

func NewUsersRepository(
	conn *sql.Conn,
	tx *sql.Tx,
	ctx context.Context,
) UsersRepository {
	i := new(usersRepositoryImpl)
	i.Conn = conn
	i.tx = tx
	i.ctx = ctx
	return i
}

// 사용자 정보를 DB에 INSERT
func (s usersRepositoryImpl) Save(user User) {
	sqlStatement := `INSERT INTO Users (user_id, user_name) VALUES (?, ?)`

	_, err := s.tx.ExecContext(s.ctx, sqlStatement, user.UserId(), user.UserEmail(), user.UserName)
	if err != nil {
		panic(err.Error())
	}
}

////////////////////////////////////////////////////////////////////////////

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jmoiron/sqlx"
// )
// import (
// 	"context"
// 	"database/sql"
// )

// type (
// 	CategoriesRepository interface {
// 		Save(Category)
// 	}

// 	categoriesRepositoryImpl struct {
// 		*sql.Conn
// 		tx  *sql.Tx
// 		ctx context.Context
// 	}
// )

// func NewCategoriesRepository(
// 	conn *sql.Conn,
// 	tx *sql.Tx,
// 	ctx context.Context,
// ) CategoriesRepository {
// 	i := new(categoriesRepositoryImpl)
// 	i.Conn = conn
// 	i.tx = tx
// 	i.ctx = ctx
// 	return i
// }

// func (s categoriesRepositoryImpl) Save(category Category) {
// 	sqlStatement := `INSERT INTO CATEGORY
// 		(user_email, color, title)
// 		VALUES (?, ?, ?)`

// 	_, err := s.tx.ExecContext(s.ctx, sqlStatement, category.UserEmail(), category.Color(), category.Title())
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

// var db *sqlx.DB

// func init() {
// 	var err error
// 	db, err = sqlx.Open("mysql", "user:password@/dbname")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// type Users struct {
// 	UserID   int    `db:"user_id"`
// 	UserName string `db:"user_name"`
// }

// func (u *Users) SaveUserInfo(db *sql.DB) error {
// 	query := `INSERT INTO Users (user_id, user_name) VALUES ($1, $2)`
// 	_, err := db.Exec(query, u.UserID, u.UserName)
// 	if err != nil {
// 		return fmt.Errorf("error saving user: %v", err)
// 	}
// 	return nil
// }

// func (u *Users) LoadByEmail(db *sql.DB, email string) error {
// 	query := `SELECT user_id, user_name FROM Users WHERE email = ?`
// 	return db.QueryRow(query, email).Scan(&u.UserID, &u.UserName)
// }
