package users

import (
   "database/sql"
   "fmt"
   "log"

   _ "github.com/go-sql-driver/mysql"
   "github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
   var err error
   db, err = sqlx.Open("mysql", "user:password@/dbname")
   if err != nil {
      log.Fatalln(err)
   }
}

type Users struct {
   UserID   int    `db:"user_id"`
   UserName string `db:"user_name"`
}

func (u *Users) SaveUserInfo(db *sql.DB) error {
   query := `INSERT INTO Users (user_id, user_name) VALUES ($1, $2)`
   _, err := db.Exec(query, u.UserID, u.UserName)
   if err != nil {
      return fmt.Errorf("error saving user: %v", err)
   }
   return nil
}

func (u *Users) LoadByEmail(db *sql.DB, email string) error {
   query := `SELECT user_id, user_name FROM Users WHERE email = ?`
   return db.QueryRow(query, email).Scan(&u.UserID, &u.UserName)
}
