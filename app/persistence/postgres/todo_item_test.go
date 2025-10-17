package postgres_test

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/kkatou7209/godo/app/port/out/dto"
	"github.com/kkatou7209/godo/persistence/postgres"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("todo item repository test", Ordered, func() {

	var todoItemRepository = postgres.NewTodoItemRepository(os.Getenv("TEST_DATABASE_URL"))

	var userId value.UserId

	var todoItemId value.TodoItemId

	BeforeAll(func() {

		userRepository := postgres.NewUserRepository(os.Getenv("TEST_DATABASE_URL"))

		userRepository.Create(&dto.CreateUserCommand{
			UserName: value.NewUserName("test_user"),
			Email: value.NewEmail("todo-test@example.com"),
			Password: value.NewPassword("test-pass"),
		})

		user, err := userRepository.GetByEmail(value.NewEmail("todo-test@example.com"))

		if err != nil {
			panic("error on getting user")
		}

		if user == nil {
			panic("fail to get user")
		}

		userId = user.Id()
	})

	When("create todo item", func()  {
		
		It("should create new todo", func() {

			err := todoItemRepository.Create(&dto.CreateTodoCommand{
				UserId: userId,
				Title: value.NewTodoItemTitle("todo1"),
				Description: value.NewTodoItemDescription("todo creation test"),
			})

			Expect(err).To(BeNil())
		})
	})

	When("todo item created", func()  {
		
		It("should get by user ID", func()  {
			
			todos, err := todoItemRepository.List(userId)

			Expect(err).To(BeNil())
			Expect(todos).To(HaveLen(1))
			Expect(todos[0].Title()).To(Equal(value.NewTodoItemTitle("todo1")))
			Expect(todos[0].Description()).To(Equal(value.NewTodoItemDescription("todo creation test")))
			Expect(todos[0].IsDone()).To(BeFalse())

			todoItemId = todos[0].Id()
		})

		It("should get by ID", func()  {
			
			todo, err := todoItemRepository.Get(todoItemId)

			Expect(err).To(BeNil())
			Expect(todo).To(Not(BeNil()))
			Expect(todo.Title()).To(Equal(value.NewTodoItemTitle("todo1")))
			Expect(todo.Description()).To(Equal(value.NewTodoItemDescription("todo creation test")))
			Expect(todo.IsDone()).To(BeFalse())
		})
	})

	When("update todo item", func()  {
		
		It("should update todo item", func()  {
			
			todo, err := todoItemRepository.Get(todoItemId)

			Expect(err).To(BeNil())
			Expect(todo).To(Not(BeNil()))

			todo.ChangeTitle("todo2")
			todo.ChangeDescription("todo creation test 2")
			todo.Complete()

			err = todoItemRepository.Update(todo)

			Expect(err).To(BeNil())
		})

		Context("after updating", func() {

			It("should todo updated", func() {

				todo, err := todoItemRepository.Get(todoItemId)

				Expect(err).To(BeNil())
				Expect(todo).ToNot(BeNil())
				Expect(todo.Title()).To(Equal(value.NewTodoItemTitle("todo2")))
				Expect(todo.Description()).To(Equal(value.NewTodoItemDescription("todo creation test 2")))
				Expect(todo.IsDone()).To(BeTrue())
			})
		})
	})

	When("delete todo item", func()  {
		
		It("should delete todo item", func()  {
			
			err := todoItemRepository.Delete(todoItemId)

			Expect(err).To(BeNil())

			todo, err := todoItemRepository.Get(todoItemId)

			Expect(err).To(BeNil())
			Expect(todo).To(BeNil())
		})
	})

	AfterAll(func() {
		ctx := context.Background()
		conn, err := pgx.Connect(ctx, os.Getenv("TEST_DATABASE_URL"))
		Expect(err).To(BeNil())
		defer conn.Close(ctx)
		_, err = conn.Exec(ctx, "TRUNCATE todo_items, users")
		Expect(err).To(BeNil())
	})
})