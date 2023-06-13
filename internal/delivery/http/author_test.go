package http_handlers

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/giicoo/golang_CRUD/internal/models"
	"github.com/giicoo/golang_CRUD/internal/service"
	mock_service "github.com/giicoo/golang_CRUD/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestAuthorGet(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthorServices, au models.Author)

	tests := []struct {
		name                 string
		inputBody            string
		inputAuthor          models.Author
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			inputBody:   `{"author_id":1, "name":"Test"}`,
			inputAuthor: models.Author{Author_id: 1, Name: "Test"},
			mockBehavior: func(r *mock_service.MockAuthorServices, au models.Author) {
				r.EXPECT().GetAuthor(au).Return([]models.Author{{Author_id: 1, Name: "Test"}}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[{"author_id":1,"name":"Test"}]` + "\n",
		},
		{
			name:                 "Invalid id",
			inputBody:            `{"author_id":"b", "name":"Test"}`,
			inputAuthor:          models.Author{Name: "Test"},
			mockBehavior:         func(r *mock_service.MockAuthorServices, au models.Author) {},
			expectedStatusCode:   400,
			expectedResponseBody: ErrorJson.ErrorText + "\n",
		},
		{
			name:        "Invalid service",
			inputBody:   `{"author_id":1, "name":"Test"}`,
			inputAuthor: models.Author{Author_id: 1, Name: "Test"},
			mockBehavior: func(r *mock_service.MockAuthorServices, au models.Author) {
				r.EXPECT().GetAuthor(au).Return([]models.Author{}, errors.New("Тест"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: ErrorService.ErrorText + "\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorServices(c)
			test.mockBehavior(repo, test.inputAuthor)

			services := &service.Services{AuthorServices: repo}
			handler := Handler{services}

			r := httprouter.New()
			r.GET("/get/author", handler.getAuthor)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/get/author",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})

	}
}

func TestAuthorAdd(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthorServices, name string)

	tests := []struct {
		name                 string
		inputBody            string
		inputName            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test"}`,
			inputName: "Test",
			mockBehavior: func(r *mock_service.MockAuthorServices, name string) {
				r.EXPECT().NewAuthor("Test").Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "Author add successful with 1 id",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorServices(c)
			test.mockBehavior(repo, test.inputName)

			services := &service.Services{AuthorServices: repo}
			handler := Handler{services}

			r := httprouter.New()
			r.POST("/add/author", handler.addAuthor)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/add/author",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestAuthorDel(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthorServices, author_id int)

	tests := []struct {
		name             string
		inputBody        string
		inputId          int
		mockBehavior     mockBehavior
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:      "OK",
			inputBody: `{"author_id":1}`,
			inputId:   1,
			mockBehavior: func(r *mock_service.MockAuthorServices, author_id int) {
				r.EXPECT().DeleteAuthor(author_id).Return(nil)
			},
			expectedStatus:   200,
			expectedResponse: "Author with 1 id delete successful",
		},
		{
			name:             "Invalid id",
			inputBody:        `{"author_id":"b"}`,
			inputId:          -1,
			mockBehavior:     func(r *mock_service.MockAuthorServices, author_id int) {},
			expectedStatus:   400,
			expectedResponse: "Invalid JSON request\n",
		},
		{
			name:      "OK",
			inputBody: `{"author_id":1}`,
			inputId:   1,
			mockBehavior: func(r *mock_service.MockAuthorServices, author_id int) {
				r.EXPECT().DeleteAuthor(author_id).Return(errors.New(ErrorService.ErrorText))
			},
			expectedStatus:   500,
			expectedResponse: "Service no execute request\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorServices(c)
			test.mockBehavior(repo, test.inputId)

			services := &service.Services{repo}
			handler := Handler{services: services}

			r := httprouter.New()
			r.POST("/del/author", handler.deleteAuthor)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/del/author", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatus)
			assert.Equal(t, w.Body.String(), test.expectedResponse)
		})
	}

}
