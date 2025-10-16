package value_test

import (
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("UserName test", func() {

	ginkgo.It("should panic on empty string", func() {
		gomega.Expect(func() { value.NewUserName("") }).To(gomega.Panic())
	})

	ginkgo.It("should equal when same value", func() {
		name := value.NewUserName("user_name_1")
		other := value.NewUserName("user_name_1")
		gomega.Expect(name == other).To(gomega.BeTrue())
	})

	ginkgo.It("should get value", func() {
		name := value.NewUserName("user_name_1")
		gomega.Expect(name.Value()).To(gomega.Equal("user_name_1"))
	})
})