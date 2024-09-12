package blocks

import (
	"context"
	"database/sql"
)

type (
	BlocksRepository interface {
		Save(Blocks)
	}

	blocksRepositoryImpl struct {
		*sql.Conn
		tx  *sql.Tx
		ctx context.Context
	}
)

func NewBlocksRepository(
	conn *sql.Conn,
	tx *sql.Tx,
	ctx context.Context,
) BlocksRepository {
	i := new(blocksRepositoryImpl)
	i.Conn = conn
	i.tx = tx
	i.ctx = ctx
	return i
}

func (s blocksRepositoryImpl) Save(blocks Blocks) {
	sqlStatement := `INSERT INTO Blocks
		(start_time, end_time, duration, color, background_image_url, block_memo)
		VALUES (?, ?, ?, ?, ?, ?)`

	_, err := s.tx.ExecContext(s.ctx, sqlStatement, blocks.StartTime(), blocks.EndTime(),
		blocks.Duration(), blocks.Color(), blocks.BackgroundImageUrl(), blocks.BlockMemo())
	if err != nil {
		panic(err.Error())
	}
}
