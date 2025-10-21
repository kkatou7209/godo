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

var _ = Describe("", Ordered, func() {

	var (
		todoItemId string
	)

	When("add todo item", func() {

		It("should add new todo item", func() {

			todo := map[string]any{
				"title":       "todo-test-title",
				"description": "todo-test-description",
			}

			jtodo, err := json.Marshal(todo)

			if err != nil {
				log.Fatal(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/user/:userId/todo-item", bytes.NewBuffer(jtodo))
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				log.Fatal(err)
			}

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("userId")
			c.SetParamValues(userId.Value())

			err = handler.AddTodoItem(app)(c)

			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusCreated))

			var res data.Payload[any]

			err = json.Unmarshal(rec.Body.Bytes(), &res)

			Expect(err).To(BeNil())
			Expect(res.Status).To(Equal(data.StatusSuccess))
		})
	})

	When("list todo items", func() {

		It("should list todo items", func() {

			req, err := http.NewRequest(http.MethodGet, "/user/:userId/todo-item", nil)
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				log.Fatal(err)
			}

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("userId")
			c.SetParamValues(userId.Value())

			err = handler.ListTodoItems(app)(c)

			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusOK))

			var res data.Payload[[]handler.TodoData]
			
			err = json.Unmarshal(rec.Body.Bytes(), &res)

			Expect(err).To(BeNil())
			Expect(res).ToNot(BeNil())
			Expect(res.Data).ToNot(BeNil())
			Expect(*res.Data).ToNot(BeEmpty())
			Expect(res.Status).To(Equal(data.StatusSuccess))
			Expect((*res.Data)[0].Title).To(Equal("todo-test-title"))
			Expect((*res.Data)[0].Description).To(Equal("todo-test-description"))
			Expect((*res.Data)[0].IsDone).To(BeFalse())

			todoItemId = (*res.Data)[0].Id
		})
	})

	When("update todo item", func() {

		It("should update todo item", func() {

			todo := map[string]any{
				"title":       "todo-test-title-updated",
				"description": "todo-test-description-updated",
			}

			jtodo, err := json.Marshal(todo)

			if err != nil {
				log.Fatal(err)
			}

			req, err := http.NewRequest(http.MethodPut, "/user/:userId/todo-item/:todoItemId", bytes.NewBuffer(jtodo))
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				log.Fatal(err)
			}

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("userId", "todoItemId")
			c.SetParamValues(userId.Value(), todoItemId)

			err = handler.UpdateTodoItem(app)(c)

			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusOK))

			var res data.Payload[any]

			err = json.Unmarshal(rec.Body.Bytes(), &res)

			Expect(err).To(BeNil())
			Expect(res.Status).To(Equal(data.StatusSuccess))
		})
	})

	When("complete todo item", func() {

		It("should complete todo", func() {

			req, err := http.NewRequest(http.MethodPatch, "/user/:userId/todo-item/:todoItemId/complete", nil)
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				log.Fatal(err)
			}

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("userId", "todoItemId")
			c.SetParamValues(userId.Value(), todoItemId)

			err = handler.CompleteTodoItem(app)(c)

			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusOK))

			var res data.Payload[any]

			err = json.Unmarshal(rec.Body.Bytes(), &res)

			Expect(err).To(BeNil())
			Expect(res.Status).To(Equal(data.StatusSuccess))
		})
	})

	When("uncomplete todo item", func() {

		It("should uncomplete todo", func() {

			req, err := http.NewRequest(http.MethodPatch, "/user/:userId/todo-item/:todoItemId/uncomplete", nil)
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				log.Fatal(err)
			}

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("userId", "todoItemId")
			c.SetParamValues(userId.Value(), todoItemId)

			err = handler.UncompleteTodoItem(app)(c)

			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusOK))

			var res data.Payload[any]

			err = json.Unmarshal(rec.Body.Bytes(), &res)

			Expect(err).To(BeNil())
			Expect(res.Status).To(Equal(data.StatusSuccess))
		})
	})

	When("delete todo item", func() {

		It("should delete todo", func() {

			req, err := http.NewRequest(http.MethodDelete, "/user/:userId/todo-item/:todoItemId", nil)
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				log.Fatal(err)
			}

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("userId", "todoItemId")
			c.SetParamValues(userId.Value(), todoItemId)

			err = handler.DeleteTodoItem(app)(c)

			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusOK))

			var res data.Payload[any]

			err = json.Unmarshal(rec.Body.Bytes(), &res)

			Expect(err).To(BeNil())
			Expect(res.Status).To(Equal(data.StatusSuccess))
		})
	})
})