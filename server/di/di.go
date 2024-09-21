//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"database/sql"
	"fmt"
	"runtime/debug"

	"github.com/google/wire"

	blockService "lifelogger/server/service/blocks"
	calendarService "lifelogger/server/service/calendar"
	categoryService "lifelogger/server/service/categories"
	helloService "lifelogger/server/service/hello"
	userService "lifelogger/server/service/users"

	blockModel "lifelogger/server/domain/blocks/model"
	calendarModel "lifelogger/server/domain/calendar/model"
	categoriesModel "lifelogger/server/domain/categories/model"
	userModel "lifelogger/server/domain/users/model"

	"lifelogger/server/config"
	"lifelogger/server/config/domainEvent"
	"lifelogger/server/util/mattermost"
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

func InjectCreateCategoryService() (categoryService.CreateCategoryService, func()) {
	panic(wire.Build(
		superSet,
		categoriesModel.NewCategoriesRepository,
		categoryService.NewCreateCategoryService,
	))
}

func InjectCreateBlockService() (blockService.CreateBlockService, func()) {
	panic(wire.Build(
		superSet,
		blockModel.NewBlocksRepository,
		blockService.NewCreateBlockService,
	))
}

func InjectCreateCalendarService() (calendarService.CreateCalendarService, func()) {
	panic(wire.Build(
		superSet,
		calendarModel.NewCalendarRepository,
		calendarService.NewCreateCalendarService,
	))
}

func InjectCreateUserService() (userService.CreateUserService, func()) {
	panic(wire.Build(
		superSet,
		userModel.NewUsersRepository,
		userService.NewCreateUserService,
	))
}
