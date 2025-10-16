package value_test

import (
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Password test", func() {

	ginkgo.It("should panic on empty string", func() {
		gomega.Expect(func() { value.NewPassword("") }).To(gomega.Panic())
	})

	ginkgo.It("should equal when same value", func() {
		password := value.NewPassword("password")
		other := value.NewPassword("password")
		gomega.Expect(password == other).To(gomega.BeTrue())
	})

	ginkgo.It("should get value", func() {
		password := value.NewPassword("password")
		gomega.Expect(password.Value()).To(gomega.Equal("password"))
	})
})