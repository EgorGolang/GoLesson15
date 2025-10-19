package controller

import (
	mock_contracts "GoLessonFifteen/internal/contracts/mocks"
	"GoLessonFifteen/internal/errs"
	"GoLessonFifteen/internal/models"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_CreateEmployee(t *testing.T) {
	type mockBehaviour func(s *mock_contracts.MockServiceI, p models.Employee)

	var testTable = []struct {
		name                 string
		inputBody            string
		inputUser            models.Employee
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			inputBody: `
					{
						"name": "PashaKach12gfd",
						"email": "PashaKach123r5@kac.com",
						"age": 21
					}`,
			inputUser: models.Employee{
				Name:  "PashaKach12gfd",
				Email: "PashaKach123r5@kac.com",
				Age:   21,
			},
			mockBehaviour: func(s *mock_contracts.MockServiceI, p models.Employee) {
				s.EXPECT().CreateUser(p).Return(nil)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"message":"Employee created successfully!"}`,
		},
		{
			name:                 "Empty fields",
			inputBody:            `{}`,
			mockBehaviour:        func(s *mock_contracts.MockServiceI, p models.Employee) {},
			expectedStatusCode:   http.StatusUnprocessableEntity,
			expectedResponseBody: `{"error":"invalid field value"}`,
		},
		{
			name: "Service failure",
			inputBody: `
					{
						"name": "Pasha",
						"email": "PashaKach123r5@kac.com",
						"age": 21
					}`,
			inputUser: models.Employee{
				Name:  "Pasha",
				Email: "PashaKach123r5@kac.com",
				Age:   21,
			},
			mockBehaviour: func(s *mock_contracts.MockServiceI, p models.Employee) {
				s.EXPECT().CreateUser(p).Return(errs.ErrInvalidUserName)
			},
			expectedStatusCode:   http.StatusUnprocessableEntity,
			expectedResponseBody: `{"error":"invalid employee name, min 4 symbols"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			svc := mock_contracts.NewMockServiceI(ctrl)
			testCase.mockBehaviour(svc, testCase.inputUser)

			handler := NewController(svc)
			r := gin.New()
			r.POST("/users", handler.CreateEmployee)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(testCase.inputBody))
			r.ServeHTTP(w, req)
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
