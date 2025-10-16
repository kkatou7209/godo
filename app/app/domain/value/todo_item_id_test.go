package value_test

import (
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("TodoItemId test", func() {

    ginkgo.It("should panic on empty string", func() {
        gomega.Expect(func() { value.NewTodoItemId("") }).To(gomega.Panic())
    })

    ginkgo.It("should equal when same value", func() {
        id := value.NewTodoItemId("item-123")
        other := value.NewTodoItemId("item-123")
        gomega.Expect(id == other).To(gomega.BeTrue())
    })

    ginkgo.It("should get value", func() {
        id := value.NewTodoItemId("item-123")
        gomega.Expect(id.Value()).To(gomega.Equal("item-123"))
    })
})


