// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/wire"
	"lifelogger/server/config"
	"lifelogger/server/config/domainEvent"
	categories2 "lifelogger/server/domain/categories/model"
	users2 "lifelogger/server/domain/users/model"

	"lifelogger/server/service/users"
	calendar2 "lifelogger/server/domain/calendar/model"

	"lifelogger/server/service/calendar"
	"lifelogger/server/service/categories"
	blocks2 "lifelogger/server/domain/blocks/model"
	"lifelogger/server/service/blocks"
	"lifelogger/server/service/hello"
	"lifelogger/server/util/mattermost"
	"runtime/debug"
)

// Injectors from di.go:

func InjectGetHelloService() hello.GetHelloService {
	getHelloService := hello.NewGetHelloService()
	return getHelloService
}

func InjectCreateCategoryService() (categories.CreateCategoryService, func()) {
	context := provideCtx()
	conn := provideConn(context)
	tx, cleanup := provideTx(context, conn)
	categoriesRepository := categories2.NewCategoriesRepository(conn, tx, context)
	createCategoryService := categories.NewCreateCategoryService(categoriesRepository)
	return createCategoryService, cleanup
}

func InjectCreateBlockService() (blocks.CreateBlockService, func()) {
	context := provideCtx()
	conn := provideConn(context)
	tx, cleanup := provideTx(context, conn)
	blocksRepository := blocks2.NewBlocksRepository(conn, tx, context)
	createCategoryService := blocks.NewCreateBlockService(blocksRepository)
	return createCategoryService, cleanup
}

func InjectCreateUserService() (users.CreateUserService, func()) {
	context := provideCtx()
	conn := provideConn(context)
	tx, cleanup := provideTx(context, conn)
	usersRepository := users2.NewUsersRepository(conn, tx, context)
	createUserService := users.NewCreateUserService(usersRepository)
	return createUserService, cleanup
}


func InjectCreateCalendarService() (calendar.CreateCalendarService, func()) {
	context := provideCtx()
	conn := provideConn(context)
	tx, cleanup := provideTx(context, conn)
	calendarRepository := calendar2.NewCalendarRepository(conn, tx, context)
	createCalendarService := calendar.NewCreateCalendarService(calendarRepository)
	return createCalendarService, cleanup
}

// di.go:

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
			"DB connection failed", fmt.Sprintf("error: %v\nstackTrace: %v",
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
			"DB connection failed", fmt.Sprintf("error: %v\nstackTrace: %v",
				txErr,
				string(debug.Stack()),
			),
		)
	}

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
