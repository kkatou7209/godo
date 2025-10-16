package entity_test

import (
	"github.com/kkatou7209/godo/app/domain/entity"
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("User test", func() {
	
	ginkgo.It("should equal when same user", func() {
		user := entity.NewUser(
			value.NewUserId("1"),
			value.NewUserName("user_name_1"),
			value.NewEmail("example@test.com"),
			value.NewPassword("password"),
		)
		other := entity.NewUser(
			value.NewUserId("1"),
			value.NewUserName("user_name_1"),
			value.NewEmail("example@test.com"),
			value.NewPassword("password"),
		)
		gomega.Expect(user.Is(other)).To(gomega.BeTrue())
	})

	ginkgo.It("should rename user", func() {
		user := entity.NewUser(
			value.NewUserId("1"),
			value.NewUserName("user_name_1"),
			value.NewEmail("example@test.com"),
			value.NewPassword("password"),
		)
		user.Rename("user_name_2")
		gomega.Expect(user.UserName() == value.NewUserName("user_name_2")).To(gomega.BeTrue())
	})

	ginkgo.It("should change email", func() {
		user := entity.NewUser(
			value.NewUserId("1"),
			value.NewUserName("user_name_1"),
			value.NewEmail("example@test.com"),
			value.NewPassword("password"),
		)
		user.ChangeEmail("example@test2.com")
		gomega.Expect(user.Email() == value.NewEmail("example@test2.com")).To(gomega.BeTrue())
		gomega.Expect(user.Email() == value.NewEmail("example@test.com")).To(gomega.BeFalse())
	})
})