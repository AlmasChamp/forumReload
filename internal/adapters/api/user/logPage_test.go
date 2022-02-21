package user

import (
	"errors"
	mock_service "forum/internal/adapters/repository/mocks"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_logPage1(t *testing.T) {
	type serviceBehavior func(s *mock_service.MockService, email string, password string)

	type User struct {
		name               string
		email              string
		password           string
		mockBehavior       serviceBehavior
		expectedStatusCode int
	}

	test := User{
		name:     "ok",
		email:    "almassagyndyk2@gmail.com",
		password: "19891042aA",
		mockBehavior: func(s *mock_service.MockService, email string, password string) {
			s.EXPECT().LogInUser(email, password).Return(&http.Cookie{}, nil)
		},
		expectedStatusCode: 200,
	}

	t.Run(test.name, func(t *testing.T) {
		//Arrange
		//Init MockService
		c := gomock.NewController(t)
		defer c.Finish()
		service := mock_service.NewMockService(c)
		test.mockBehavior(service, test.email, test.password)
		//Init Handler
		mux := http.NewServeMux()
		handler := NewHandler(service)
		handler.Register(mux)
		// InsertData into r.FormValue
		data := url.Values{}
		data.Set("eMail", test.email)
		data.Set("password", test.password)

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/logPage", strings.NewReader(data.Encode()))

		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
		// Act
		mux.ServeHTTP(w, r)
		// Assert
		assert.Equal(t, test.expectedStatusCode, w.Code)
	})

}

func TestHandler_logPage2(t *testing.T) {
	type serviceBehavior func(s *mock_service.MockService, email string, password string)

	type User struct {
		name               string
		email              string
		password           string
		mockBehavior       serviceBehavior
		expectedStatusCode int
	}

	test := User{

		name:     "r.FormValueIncorrect",
		email:    "almassagyndyk2@gmail.com",
		password: "19891042aA",
		mockBehavior: func(s *mock_service.MockService, email string, password string) {
			s.EXPECT().LogInUser(email, password).Return(&http.Cookie{}, errors.New("Error"))
		},
		expectedStatusCode: 404,
	}

	t.Run(test.name, func(t *testing.T) {
		//Arrange
		//Init MockService
		c := gomock.NewController(t)
		defer c.Finish()
		service := mock_service.NewMockService(c)
		test.mockBehavior(service, test.email, test.password)
		//Init Handler
		mux := http.NewServeMux()
		handler := NewHandler(service)
		handler.Register(mux)
		// InsertData into r.FormValue
		data := url.Values{}
		data.Set("eMail", test.email)
		data.Set("password", test.password)

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/logPage", strings.NewReader(data.Encode()))

		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
		// Act
		mux.ServeHTTP(w, r)
		// Assert
		assert.Equal(t, test.expectedStatusCode, w.Code)
	})

}
