package value_test

import (
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("TodoItemTitle test", func() {
	
	ginkgo.It("should panic on empty string", func() {
		gomega.Expect(func() { value.NewTodoItemTitle("") }).To(gomega.Panic())
	})

	ginkgo.It("should equal when same value", func() {
		title := value.NewTodoItemTitle("title")
		other := value.NewTodoItemTitle("title")
		gomega.Expect(title == other).To(gomega.BeTrue())
	})

	ginkgo.It("should get value", func() {
		title := value.NewTodoItemTitle("title")
		gomega.Expect(title.Value()).To(gomega.Equal("title"))
	})
})