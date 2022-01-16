package repository

import (
	"log"
	"rest-api/model/entity"
)

// 外部パッケージに後悔するinterface
type TodoRepository interface {
	GetTodos() (todos []entity.TodoEntity, err error)
	InsertTodo(todo entity.TodoEntity) (id int, err error)
	UpdateTodo(todo entity.TodoEntity) (err error)
	DeleteTodo(id int) (err error)
}

// 非公開のTodoRepository
type todoRepository struct {
}

// TodoRepositoryのコンストラクタ。TodoRepository構造体のポインタを返却
func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

// TODO取得処理
func (tr *todoRepository) GetTodos() (todos []entity.TodoEntity, err error) {
	todos = []entity.TodoEntity{}

	//DBから全てのTODOを取得
	rows, err := Db.Query("SELECT id, title, content FROM todo ORDEER BY id DESC ")
	if err != nil {
		log.Print(err)
		return
	}

	//1行ごとTODOにEntityをマッピングし、返却用のスライスに追加
	for rows.Next() {
		todo := entity.TodoEntity{}
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Content)
		if err != nil {
			log.Print(err)
			return
		}
		todos = append(todos, todo)
	}
	return
}

// TODOの追加処理
func (tr *todoRepository) InsertTodo(todo entity.TodoEntity) (id int, err error) {
	//引数で受け取ったEntityを値を元にDBに追加
	_, err = Db.Exec("INSERT INTO todo (title, content) VALUES (?, ?)", todo.Title, todo.Content)
	if err != nil {
		log.Print(err)
		return
	}
	//created_atが最新のTODOのIDを返却
	err = Db.QueryRow("SELECT id FROM todo ORDER BY id DESC LIMIT 1").Scan(&id)
	return
}

// TODOの更新処理
func (tr *todoRepository) UpdateTodo(todo entity.TodoEntity) (err error) {
	//引数で受け取ったEntityの値を元にDBを更新
	_, err = Db.Exec("UPDATE todo SET title = ?, content = ? WHERE id = ?", todo.Title, todo.Content)
	return err
}

// TODOの削除処理
func (tr *todoRepository) DeleteTodo(id int) (err error) {
	//引数で受け取ったIDの値を元にDBから削除
	_, err = Db.Exec("DELETE FROM todo WHERE id = ?", id)
	return
}
