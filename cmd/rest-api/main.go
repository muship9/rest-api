package main

import (
	"net/http"
	"rest-api/controller"
	"rest-api/model/repository"
)

// コンストラクタインジェクションをDIで行う　https://qiita.com/yoshinori_hisakawa/items/a944115eb77ed9247794
var tr = repository.NewTodoRepository()
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	// localhost:8080/todos/に届いたリクエストをHandleTodosRequestで処理するように設定
	http.HandleFunc("/todos/", ro.HandleTodosRequest)
	// サーバ起動
	server.ListenAndServe()
}
