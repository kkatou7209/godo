package value_test

import (
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("TodoItemDescription test", func() {
	
	ginkgo.It("should panic on empty string", func() {
		gomega.Expect(func() { value.NewTodoItemDescription("") }).To(gomega.Panic())
	})

	ginkgo.It("should equal when same value", func() {
		description := value.NewTodoItemDescription("description")
		other := value.NewTodoItemDescription("description")
		gomega.Expect(description == other).To(gomega.BeTrue())
	})

	ginkgo.It("should get value", func() {
		description := value.NewTodoItemDescription("description")
		gomega.Expect(description.Value()).To(gomega.Equal("description"))
	})
})