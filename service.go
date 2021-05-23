package main

import (
	"net/http"

	router "GoRestExample/router"
)

func main() {
	// ルーティングの設定を行う
	router.GetUserRoute()

	// 8080ポートで待ち受ける
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	} else {
		println("start!")
	}
}
