package main

import (
	"net/http"

	router "./router"
)

func main() {
	// ルーティングの設定を行う
	router.GetUserRoute()

	// 8080ポートで待ち受ける
	http.ListenAndServe(":8080", nil)
}
