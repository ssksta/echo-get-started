package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/", hello) // ローカル環境の場合、http://localhost:1323/ にGETアクセスされるとhelloハンドラーを実行する

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))

}

// ハンドラーを定義
func hello(c echo.Context) error {
	return c.String(http.StatusOK, dbTest())
}

func dbTest() string {
	_, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		return "DB接続成功"
	}
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "go_example"
	PASS := "12345!"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "go_example"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}
