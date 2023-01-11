package main

import (
	"fmt"
	"net/http"
	"time"

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
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	defer db.Close()

	error := db.Create(&Users{
		Name:     "テスト太郎",
		Age:      "18",
		Address:  "東京都千代田区",
		UpdateAt: getDate(),
	}).Error
	if error != nil {
		panic(err.Error())
	} else {
		fmt.Println("データ追加成功")
		return "データ追加成功"
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

type Users struct {
	ID       int
	Name     string `json:"name"`
	Age      string `json:"age"`
	Address  string `json:"address"`
	UpdateAt string `json:"updateAt" sql:"not null;type:date"`
}

func getDate() string {
	const layout = "2006-01-02 15:04:05"
	now := time.Now()
	return now.Format(layout)
}
