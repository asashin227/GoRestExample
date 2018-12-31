package controller

import (
	"net/http"

	"../model"
	restful "github.com/emicklei/go-restful"
)

func FindUser(req *restful.Request, resp *restful.Response) {
	// リクエストを resHello に割り当てる
	id := req.PathParameter("user_id")

	// 適当なエラーハンドリング
	if id == "" {
		// エラーの内容をJSONで返す
		resp.WriteHeaderAndJson(http.StatusInternalServerError, &model.Error{Error: "", Code: 0}, restful.MIME_JSON)
		return
	}

	var user = model.User{Id: id, Name: "test", Age: 0, Address: "hogehoge"}
	// 結果を出力する (自動で application/json が指定される)
	resp.WriteAsJson(&user)
}
