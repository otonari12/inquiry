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
	e.POST("/api/inquiry", api.Post())
	e.GET("/api/inquiry", api.Get())

	// サーバー起動
	e.Start(":8090") //ポート番号指定してね
}
