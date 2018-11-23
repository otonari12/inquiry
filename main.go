package main

import (
	"inquiry-api/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.POST("/inquiry", api.Post())
	e.GET("/inquiry", api.Get())

	// サーバー起動
	e.Start(":1323") //ポート番号指定してね
}
