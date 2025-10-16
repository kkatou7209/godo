package entity_test

import (
	"github.com/kkatou7209/godo/app/domain/entity"
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("TodoItem test", func() {
	
	ginkgo.It("should equal when same todo item", func() {
		todo := entity.NewTodoItem(
			value.NewTodoItemId("1"),
			value.NewTodoItemTitle("title"),
			value.NewTodoItemDescription("description"),
			false,
			value.NewUserId("1"),
		)
		other := entity.NewTodoItem(
			value.NewTodoItemId("1"),
			value.NewTodoItemTitle("title2"),
			value.NewTodoItemDescription("description2"),
			false,
			value.NewUserId("1"),
		)
		gomega.Expect(todo.Is(other)).To(gomega.BeTrue())
	})

	ginkgo.It("should complete todo item", func() {
		todo := entity.NewTodoItem(
			value.NewTodoItemId("1"),
			value.NewTodoItemTitle("title"),
			value.NewTodoItemDescription("description"),
			false,
			value.NewUserId("1"),
		)
		todo.Complete()
		gomega.Expect(todo.IsDone()).To(gomega.BeTrue())
	})

	ginkgo.It("should uncomplete todo item", func() {
		todo := entity.NewTodoItem(
			value.NewTodoItemId("1"),
			value.NewTodoItemTitle("title"),
			value.NewTodoItemDescription("description"),
			true,
			value.NewUserId("1"),
		)
		todo.Uncomplete()
		gomega.Expect(todo.IsDone()).To(gomega.BeFalse())
	})

	ginkgo.It("should change title of todo item", func() {
		todo := entity.NewTodoItem(
			value.NewTodoItemId("1"),
			value.NewTodoItemTitle("title"),
			value.NewTodoItemDescription("description"),
			false,
			value.NewUserId("1"),
		)
		todo.ChangeTitle("title2")
		gomega.Expect(todo.Title() == value.NewTodoItemTitle("title2")).To(gomega.BeTrue())
	})

	ginkgo.It("should change description of todo item", func() {
		todo := entity.NewTodoItem(
			value.NewTodoItemId("1"),
			value.NewTodoItemTitle("title"),
			value.NewTodoItemDescription("description"),
			false,
			value.NewUserId("1"),
		)
		todo.ChangeDescription("description2")
		gomega.Expect(todo.Description() == value.NewTodoItemDescription("description2")).To(gomega.BeTrue())
	})
})