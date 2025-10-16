package value_test

import (
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Email test", func() {

	ginkgo.It("should panic on empty string", func() {
		gomega.Expect(func() { value.NewEmail("") }).To(gomega.Panic())
	})

	ginkgo.It("should equal on same value", func() {
		email := value.NewEmail("example@test.com")
		other := value.NewEmail("example@test.com")
		gomega.Expect(email == other).To(gomega.BeTrue())
	})

	ginkgo.It("should panic on invalid email", func() {
		gomega.Expect(func() { value.NewEmail("invalid.com") }).To(gomega.Panic())
	})
})