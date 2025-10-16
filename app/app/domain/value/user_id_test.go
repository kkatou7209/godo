package value_test

import (
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)


var _ = ginkgo.Describe("UserId test", func() {

	ginkgo.It("should panic on empty string", func() {
		gomega.Expect(func() { value.NewUserId("")}).To(gomega.Panic())
	})

	ginkgo.It("should equal when same value", func() {
		id := value.NewUserId("100")
		other := value.NewUserId("100")
		gomega.Expect(id == other).To(gomega.BeTrue())
	})

	ginkgo.It("should get value", func() {
		id := value.NewUserId("100")
		gomega.Expect(id.Value()).To(gomega.Equal("100"))
	})
})