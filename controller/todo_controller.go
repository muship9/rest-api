//リクエストを元にモデルの各処理を呼び出しレスポンスを返却する
package controller

import "net/http"

type TodoController interface {
	GetTodos(w http.ResponseWriter, r *http.Request)
	PostTodo(w http.ResponseWriter, r *http.Request)
	PutTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}
