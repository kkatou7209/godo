package dto

type TodoItemDto struct {
	Id string
	Title string
	Description string
	IsDone bool
	UserId string
}

type AddTodoCommand struct {
	UserId string
	Title string
	Description string
}

type UpdateTodoCommand struct {
	Id string
	Title string
	Description string
	UserId string
}