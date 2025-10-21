package handler_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/kkatou7209/godo/web/data"
	"github.com/kkatou7209/godo/web/handler"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("auth handler test", Ordered, func() {

	When("signup user", func()  {
		
		It("should create user", func() {

			u := map[string]any{
				"username": "auth-test-user",
				"email": "auth-api@example.com",
				"password": "auth-api-test-pass",
			}

			ju, err := json.Marshal(u)

			if err != nil {
				log.Fatalln(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(ju))

			if err != nil {
				log.Fatalln(err)
			}

			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			err = handler.SignUp(app)(c)

			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusCreated))

			var res data.Payload[any]

			err = json.Unmarshal(rec.Body.Bytes(), &res)

			Expect(err).To(BeNil())
			Expect(res.Status).To(Equal(data.StatusSuccess))
		})
	})

	When("user created", func() {

		It("should login", func() {

			cred := map[string]any{
				"email": "auth-api@example.com",
				"password": "auth-api-test-pass",
			}

			jcred, err := json.Marshal(cred)

			if err != nil {
				log.Fatalln(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(jcred))
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				log.Fatalln(err)
			}

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			err = handler.Login(app)(c)

			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusOK))

			var res data.Payload[any]

			err = json.Unmarshal(rec.Body.Bytes(), &res)

			Expect(err).To(BeNil())
			Expect(res.Status).To(Equal(data.StatusSuccess))

			var tokenCookie *http.Cookie

			for _, cookie := range rec.Result().Cookies() {
				if cookie.Name == "x-api-token" {
					tokenCookie = cookie
					break
				}
			}

			Expect(tokenCookie).ToNot(BeNil())
			Expect(tokenCookie.Value).ToNot(BeEmpty())
		})
	})
})