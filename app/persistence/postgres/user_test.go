package postgres_test

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/kkatou7209/godo/app/domain/entity"
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/kkatou7209/godo/app/port/out/dto"
	"github.com/kkatou7209/godo/persistence/postgres"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("postgres repository test", Ordered, func() {
	var userRepository *postgres.UserRepository
	var userId value.UserId
	var user *entity.User
	var err error

	BeforeAll(func() {
		userRepository = postgres.NewUserRepository(os.Getenv("TEST_DATABASE_URL"))
	})

	When("create user", func() {
		It("should create new user", func() {
			err = userRepository.Create(&dto.CreateUserCommand{
				UserName:  value.NewUserName("user01"),
				Email:     value.NewEmail("test@example.com"),
				Password:  value.NewPassword("test-password-01"),
			})
			Expect(err).To(BeNil())
		})
	})

	When("user is created", func() {
		It("can get user by its email", func() {
			user, err = userRepository.GetByEmail(value.NewEmail("test@example.com"))
			Expect(err).To(BeNil())
			Expect(user).To(Not(BeNil()))
			Expect(user.Email()).To(Equal(value.NewEmail("test@example.com")))
			Expect(user.Password()).To(Equal(value.NewPassword("test-password-01")))
			userId = user.Id()
		})

		It("can get user by id", func() {
			fetched, err := userRepository.GetById(userId)
			Expect(err).To(BeNil())
			Expect(fetched).To(Not(BeNil()))
			Expect(fetched.Email()).To(Equal(value.NewEmail("test@example.com")))
			Expect(fetched.Password()).To(Equal(value.NewPassword("test-password-01")))
		})

		When("updating user", func() {
			
			It("should have updated values", func() {
				user.ChangeEmail("another@example.com")
				user.ChangePassword("test-password-02")
				err = userRepository.Update(user)
				
				Expect(err).To(BeNil())
				updatedUser, _ := userRepository.GetById(userId)
				Expect(updatedUser.Email()).To(Equal(value.NewEmail("another@example.com")))
				Expect(updatedUser.Password()).To(Equal(value.NewPassword("test-password-02")))
			})
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
