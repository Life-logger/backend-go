package categories

import (
	"context"
	"database/sql"
)

type (
	CategoriesRepository interface {
		Save(Category)
	}

	categoriesRepositoryImpl struct {
		*sql.Conn
		tx  *sql.Tx
		ctx context.Context
	}
)

func NewCategoriesRepository(
	conn *sql.Conn,
	tx *sql.Tx,
	ctx context.Context,
) CategoriesRepository {
	i := new(categoriesRepositoryImpl)
	i.Conn = conn
	i.tx = tx
	i.ctx = ctx
	return i
}

func (s categoriesRepositoryImpl) Save(category Category) {
	sqlStatement := `INSERT INTO CATEGORY
		(user_email, color, title)
		VALUES (?, ?, ?)`

	_, err := s.tx.ExecContext(s.ctx, sqlStatement, category.UserEmail(), category.Color(), category.Title())
	if err != nil {
		panic(err.Error())
	}
}
