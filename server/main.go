package main

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/import-ai/ai-know/server/config"
	"github.com/import-ai/ai-know/server/db"
	"github.com/import-ai/ai-know/server/routes"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// AIKnow API
//
//	@title			AIKnow API
//	@version		1.0
//	@description ## Sidebar 多级列表
//	@description
//	@description 列表的每一项是一个`entry`，主要字段是`id`，`type`和`title`。
//	@description
//	@description - `id`: 全局唯一标识
//	@description - `type`: 合法取值为`note`，`group`或`link`
//	@description - `title`: 标题
//	@description
//	@description 一个`entry`下可以嵌套子`entry`，形成树形结构。
//	@description
//	@description ### 查询流程
//	@description
//	@description 1. 调用`Get Workspace`拿到 Private 和 Team Space 分别对应的最外层`entry id`。
//	@description 2. 调用`Get Sub-Entries`，传参`entry id`，拿到该`entry`直接嵌套的子`entry`列表。
//	@description 3. 如果子`entry`继续嵌套子`entry`（`has_sub_entries`为`true`），递归调用`Get Sub-Entries`。
//	@description
//	@description ### 拖拽/移动流程
//	@description
//	@description 调用`Update Entry`，更新`entry`的`parent`和`position`。

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()
	config.InitFromEnv()

	if err := db.Init(context.Background()); err != nil {
		log.Fatal().Err(err).Send()
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Err(err).Msg("http error")
			return fiber.DefaultErrorHandler(c, err)
		},
	})
	routes.RegisterRoutes(app)
	if err := app.Listen(config.ListenAddr()); err != nil {
		log.Fatal().Err(err).Send()
	}
}
