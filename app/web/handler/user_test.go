package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/kkatou7209/godo/web/data"
	"github.com/kkatou7209/godo/web/handler"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("user handler test", func() {

	When("get user", func()  {
		
		It("should get user", func() {

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/user/%s", userId.Value()), nil)

			if err != nil {
				log.Fatalln(err)
			}

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("userId")
			c.SetParamValues(userId.Value())

			err = handler.GetUserById(app)(c)

			Expect(err).To(BeNil())
			Expect(rec.Body).ToNot(BeNil())
			Expect(rec.Code).To(Equal(http.StatusOK))

			res :=  new(data.Payload[handler.UserData])

			err = json.Unmarshal(rec.Body.Bytes(), &res)
			
			Expect(err).To(BeNil())
			Expect(res.Status).To(Equal(data.StatusSuccess))
			Expect(res.Data.Id).To(Equal(userId.Value()))
			Expect(res.Data.Email).To(Equal("handler-test@example.com"))
			Expect(res.Data.Username).To(Equal("handler-test-user"))
		})
	})

	When("update user", func() {

		It("should update user", func() {

			u := map[string]any{
				"email": "handler-test-updated@example.com",
				"username": "handler-test-user-updated",
			}

			ju, err := json.Marshal(u)

			if err != nil {
				panic(err)
			}

			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/user/%s", userId.Value()), bytes.NewBuffer(ju))
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				log.Fatalln(err)
			}

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("userId")
			c.SetParamValues(userId.Value())

			err = handler.UpdateUser(app)(c)

			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusOK))

			var res data.Payload[any]

			err = json.Unmarshal(rec.Body.Bytes(), &res)

			Expect(err).To(BeNil())
			Expect(res.Status).To(Equal(data.StatusSuccess))
		})

		Context("after updating", func() {

			It("user should be updated", func() {

				req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/user/%s", userId.Value()), nil)

				if err != nil {
					log.Fatalln(err)
				}

				rec := httptest.NewRecorder()

				c := e.NewContext(req, rec)
				c.SetParamNames("userId")
				c.SetParamValues(userId.Value())

				err = handler.GetUserById(app)(c)

				Expect(err).To(BeNil())
				Expect(rec.Body).ToNot(BeNil())
				Expect(rec.Code).To(Equal(http.StatusOK))

				res :=  new(data.Payload[handler.UserData])

				err = json.Unmarshal(rec.Body.Bytes(), &res)
				
				Expect(err).To(BeNil())
				Expect(res.Status).To(Equal(data.StatusSuccess))
				Expect(res.Data.Id).To(Equal(userId.Value()))
				Expect(res.Data.Email).To(Equal("handler-test-updated@example.com"))
				Expect(res.Data.Username).To(Equal("handler-test-user-updated"))
			})
		})
	})

	When("change password", func() {

		It("should change password", func() {

			p := map[string]any{
				"oldPassword": "handler-test-pass",
				"newPassword": "handler-test-pass-updated",
			}

			jp, err := json.Marshal(p)

			if err != nil {
				log.Fatalln(err)
			}

			req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("/user/%s", userId.Value()), bytes.NewBuffer(jp))
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				log.Fatalln(err)
			}

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("userId")
			c.SetParamValues(userId.Value())

			err = handler.ChangeUserPassword(app)(c)

			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusOK))

			var res data.Payload[any]

			err = json.Unmarshal(rec.Body.Bytes(), &res)

			Expect(err).To(BeNil())
			Expect(res.Status).To(Equal(data.StatusSuccess))
		})
	})
})