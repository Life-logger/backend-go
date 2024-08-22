//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"database/sql"
	"fmt"
	"runtime/debug"

	"github.com/google/wire"

	helloService "taskbuddy.io/taskbuddy/service/hello"

	"taskbuddy.io/taskbuddy/config"
	"taskbuddy.io/taskbuddy/config/domainEvent"
	"taskbuddy.io/taskbuddy/util/mattermost"
)

var (
	superSet = wire.NewSet(
		provideCtx,
		provideConn,
		provideTx,
	)
)

func provideCtx() context.Context {
	ctx := context.Background()
	processor := domainEvent.NewEventProcessor()

	ctx = context.WithValue(ctx, "eventProcessor", processor)
	processor.SetCtx(ctx)

	return ctx
}

func provideConn(ctx context.Context) *sql.Conn {
	db := config.GetDBInstance()

	conn, err := db.Conn(ctx)
	if err != nil {
		mattermost.Noti(
			"DB connection failed",
			fmt.Sprintf("error: %v\nstackTrace: %v",
				err,
				string(debug.Stack()),
			),
		)
	}

	return conn
}

func provideTx(ctx context.Context, conn *sql.Conn) (*sql.Tx, func()) {
	tx, txErr := conn.BeginTx(ctx, nil)
	if txErr != nil {
		mattermost.Noti(
			"DB connection failed",
			fmt.Sprintf("error: %v\nstackTrace: %v",
				txErr,
				string(debug.Stack()),
			),
		)
	}

	// recover를 cleanup에서 하기 위해 cleanup은 이거 하나만 써야함
	cleanup := func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}

		if err := tx.Commit(); err != nil {
			tx.Rollback()
			fmt.Println(err)
		}

		if err := conn.Close(); err != nil {
			fmt.Println(err)
		}

		processor := ctx.Value("eventProcessor").(domainEvent.EventProcessor)
		processor.ProcessAndRemove()
	}

	return tx, cleanup
}

func InjectGetHelloService() helloService.GetHelloService {
	wire.Build(
		helloService.NewGetHelloService,
	)
	return nil
}
