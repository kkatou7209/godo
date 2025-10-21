package web_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kkatou7209/godo/app"
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/kkatou7209/godo/password"
	"github.com/kkatou7209/godo/persistence/mock"
	"github.com/kkatou7209/godo/web"
	"github.com/kkatou7209/godo/web/data"
	"github.com/kkatou7209/godo/web/handler"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestServer(t *testing.T) {

	RegisterFailHandler(Fail)

	RunSpecs(t, "server test")
}

var (
	ts         *httptest.Server
	userId 	   string
	todoItemId string
	todoRepository *mock.MockTodoItemRepository
	userRepository *mock.MockUserRepository
)

var _ = BeforeSuite(func() {

	app := app.New()

	todoRepository = mock.NewMockTodoItemRepository()

	userRepository = mock.NewMockUserRepository()

	app.
		SetCreateTodoPersistence(todoRepository).
		SetListTodoPersistence(todoRepository).
		SetGetTodoPersistence(todoRepository).
		SetUpdateTodoPersistence(todoRepository).
		SetDeleteTodoPersistence(todoRepository).
		SetCreateUserPersistence(userRepository).
		SetGetUserPersistence(userRepository).
		SetUpdateUserPersistence(userRepository).
		SetPasswordHasher(password.NewBycryptPasswordHasher())

	e := echo.New()
	e.HideBanner = true

	web.MapRoutes(e, app)

	ts = httptest.NewServer(e)
})

var _ = AfterSuite(func() {
	ts.Close()
})

var _ = Describe("API integration test", Ordered, func() {

	It("should create user", func() {

		u := map[string]any{
			"username": "http-api-test-user",
			"email": "http-api@example.com",
			"password": "http-api-test-pass",
		}

		ju, err := json.Marshal(u)

		if err != nil {
			log.Fatalln(err)
		}

		res, err := http.Post(ts.URL + "/auth/signup", "application/json", bytes.NewBuffer(ju))

		Expect(err).To(BeNil())
		Expect(res.StatusCode).To(Equal(http.StatusCreated))

		defer res.Body.Close()
	})

	Context("after user created", func() {

		It("should login", func() {

			cred := map[string]any{
				"email": "http-api@example.com",
				"password": "http-api-test-pass",
			}

			jcred, err := json.Marshal(cred)

			if err != nil {
				log.Fatalln(err)
			}

			res, err := http.Post(ts.URL + "/auth/login", "application/json", bytes.NewBuffer(jcred))

			Expect(err).To(BeNil())

			defer res.Body.Close()
			
			Expect(res.StatusCode).To(Equal(http.StatusOK))

			var tokenCookie *http.Cookie

			for _, c := range res.Cookies() {
				if c.Name == "x-api-token" {
					tokenCookie = c
				}
			}

			Expect(tokenCookie).ToNot(BeNil())

			body, _ := io.ReadAll(res.Body)

			var payload data.Payload[handler.UserData]

			_ = json.Unmarshal(body, &payload)

			Expect(payload).ToNot(BeNil())
			Expect(payload.Status).To(Equal(data.StatusSuccess))
		})
	})

	Context("after login", func() {

		BeforeAll(func() {
			u, _ := userRepository.GetByEmail(value.NewEmail("http-api@example.com"))
			userId = u.Id().Value()
		})

		It("should get user by ID", func() {

			res, err := http.Get(ts.URL + "/user/" + userId)

			Expect(err).To(BeNil())
			Expect(res.StatusCode).To(Equal(http.StatusOK))

			defer res.Body.Close()

			body, _ := io.ReadAll(res.Body)

			var payload data.Payload[handler.UserData]

			_ = json.Unmarshal(body, &payload)

			Expect(payload).ToNot(BeNil())
			Expect(payload.Status).To(Equal(data.StatusSuccess))
			Expect(payload.Data.Email).To(Equal("http-api@example.com"))
			Expect(payload.Data.Username).To(Equal("http-api-test-user"))
			Expect(payload.Data.Id).To(Equal(userId))
		})

		It("should update user", func() {

			u := map[string]any{
				"id": userId,
				"email": "http-api-test-updated@example.com",
				"username": "http-api-test-user-updated",
			}

			ju, err := json.Marshal(u)

			if err != nil {
				panic(err)
			}

			req, _ := http.NewRequest(http.MethodPut, ts.URL + "/user/" + userId, bytes.NewReader(ju))
			req.Header.Set("Content-Type", "application/json")

			res, err := http.DefaultClient.Do(req)

			Expect(err).To(BeNil())
			Expect(res.StatusCode).To(Equal(http.StatusOK))

			defer res.Body.Close()
		})

		It("should change password", func() {

			p := map[string]any{
				"oldPassword": "http-api-test-pass",
				"newPassword": "http-api-test-pass-updated",
			}

			jp, err := json.Marshal(p)

			if err != nil {
				log.Fatalln(err)
			}
			
			req, _ := http.NewRequest(http.MethodPatch, ts.URL + "/user/" + userId + "/password", bytes.NewReader(jp))
			req.Header.Set("Content-Type", "application/json")

			res, err := http.DefaultClient.Do(req)

			Expect(err).To(BeNil())
			Expect(res.StatusCode).To(Equal(http.StatusOK))
			
			defer res.Body.Close()
		})

		It("should add todo item", func()  {
			
			todo := map[string]any{
				"title":       "todo-test-title",
				"description": "todo-test-description",
			}

			jtodo, err := json.Marshal(todo)

			if err != nil {
				log.Fatal(err)
			}

			res, err := http.Post(ts.URL + "/user/" + userId + "/todo-item", "application/json", bytes.NewBuffer(jtodo))

			Expect(err).To(BeNil())
			Expect(res.StatusCode).To(Equal(http.StatusCreated))
			
			defer res.Body.Close()
		})

		Context("after todo item added", func() {

			It("should list todo items", func() {

				res, err := http.Get(ts.URL + "/user/" + userId + "/todo-items")

				Expect(err).To(BeNil())
				Expect(res.StatusCode).To(Equal(http.StatusOK))

				defer res.Body.Close()

				var payload data.Payload[[]handler.TodoData]

				body, _ := io.ReadAll(res.Body)

				err = json.Unmarshal(body, &payload)

				Expect(err).To(BeNil())
				Expect(payload.Status).To(Equal(data.StatusSuccess))
				Expect(payload.Data).ToNot(BeNil())

				todoItemId = (*payload.Data)[0].Id
			})

			It("should update todo item", func() {

				todo := map[string]any{
					"title":       "todo-test-title-updated",
					"description": "todo-test-description-updated",
				}
	
				jtodo, err := json.Marshal(todo)
	
				if err != nil {
					log.Fatal(err)
				}

				req, err := http.NewRequest(http.MethodPut, ts.URL + "/user/" + userId + "/todo-item/" + todoItemId, bytes.NewBuffer(jtodo))
				req.Header.Set("Content-Type", "application/json")

				if err != nil {
					log.Fatal(err)
				}

				res, err := http.DefaultClient.Do(req)

				Expect(err).To(BeNil())
				// Expect(res.StatusCode).To(Equal(http.StatusOK))
				
				defer res.Body.Close()

				body, _ := io.ReadAll(res.Body)

				var payload data.Payload[any]

				_ = json.Unmarshal(body, &payload)

				log.Print(payload.Message, payload.Errors)
			})

			Context("after todo item updated", func() {

				It("should todo item updated", func() {

					res, err := http.Get(ts.URL + "/user/" + userId + "/todo-items")

					Expect(err).To(BeNil())
					Expect(res.StatusCode).To(Equal(http.StatusOK))

					defer res.Body.Close()

					var payload data.Payload[[]handler.TodoData]

					body, _ := io.ReadAll(res.Body)

					err = json.Unmarshal(body, &payload)

					Expect(err).To(BeNil())
					Expect(payload.Status).To(Equal(data.StatusSuccess))
					Expect(payload.Data).ToNot(BeNil())
					Expect((*payload.Data)[0].Title).To(Equal("todo-test-title-updated"))
					Expect((*payload.Data)[0].Description).To(Equal("todo-test-description-updated"))
				})
			})

			It("should complete todo item", func() {

				req, _ := http.NewRequest(http.MethodPatch, ts.URL + "/user/" + userId + "/todo-item/" + todoItemId + "/complete", nil)
				req.Header.Set("Content-Type", "application/json")

				res, err := http.DefaultClient.Do(req)

				Expect(err).To(BeNil())
				Expect(res.StatusCode).To(Equal(http.StatusOK))

				defer res.Body.Close()

			})

			Context("after todo item completed", func() {

				It("should todo item completed", func() {

					res, err := http.Get(ts.URL + "/user/" + userId + "/todo-items")

					Expect(err).To(BeNil())
					Expect(res.StatusCode).To(Equal(http.StatusOK))

					defer res.Body.Close()

					var payload data.Payload[[]handler.TodoData]

					body, _ := io.ReadAll(res.Body)

					err = json.Unmarshal(body, &payload)

					Expect(err).To(BeNil())
					Expect(payload.Status).To(Equal(data.StatusSuccess))
					Expect(payload.Data).ToNot(BeNil())
					Expect((*payload.Data)[0].IsDone).To(BeTrue())
				})
			})

			It("should uncomplete todo item", func() {

				req, _ := http.NewRequest(http.MethodPatch, ts.URL + "/user/" + userId + "/todo-item/" + todoItemId + "/uncomplete", nil)
				req.Header.Set("Content-Type", "application/json")

				res, err := http.DefaultClient.Do(req)

				Expect(err).To(BeNil())
				Expect(res.StatusCode).To(Equal(http.StatusOK))

				defer res.Body.Close()

			})

			Context("after todo item uncompleted", func() {

				It("should todo item uncompleted", func() {

					res, err := http.Get(ts.URL + "/user/" + userId + "/todo-items")

					Expect(err).To(BeNil())
					Expect(res.StatusCode).To(Equal(http.StatusOK))

					defer res.Body.Close()

					var payload data.Payload[[]handler.TodoData]

					body, _ := io.ReadAll(res.Body)

					err = json.Unmarshal(body, &payload)

					Expect(err).To(BeNil())
					Expect(payload.Status).To(Equal(data.StatusSuccess))
					Expect(payload.Data).ToNot(BeNil())
					Expect((*payload.Data)[0].IsDone).To(BeFalse())
				})
			})

			It("should delete todo item", func() {

				req, _ := http.NewRequest(http.MethodDelete, ts.URL + "/user/" + userId + "/todo-item/" + todoItemId, nil)
				req.Header.Set("Content-Type", "application/json")

				res, err := http.DefaultClient.Do(req)

				Expect(err).To(BeNil())
				Expect(res.StatusCode).To(Equal(http.StatusOK))

				defer res.Body.Close()
			})

			Context("after todo item deleted", func() {

				It("should todo item deleted", func() {

					res, err := http.Get(ts.URL + "/user/" + userId + "/todo-items")

					Expect(err).To(BeNil())
					Expect(res.StatusCode).To(Equal(http.StatusOK))

					defer res.Body.Close()

					var payload data.Payload[[]handler.TodoData]

					body, _ := io.ReadAll(res.Body)

					err = json.Unmarshal(body, &payload)

					Expect(err).To(BeNil())
					Expect(payload.Status).To(Equal(data.StatusSuccess))
					Expect(payload.Data).ToNot(BeNil())
					Expect(*payload.Data).To(BeEmpty())
				})
			})
		})
	})
})