package user

import (
	"errors"
	mock_service "forum/internal/adapters/repository/mocks"
	entities "forum/internal/model"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler1_regPage(t *testing.T) {
	type mockBehavior func(s *mock_service.MockService, user entities.User)

	testTable := []struct {
		name               string
		inputUser          entities.User
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name: "ok",
			inputUser: entities.User{
				Name:      "almas",
				Login:     "almassagyndyk2@gmail.com",
				Password1: "19891042aA",
				Password2: "19891042aA",
			},
			mockBehavior: func(s *mock_service.MockService, user entities.User) {
				s.EXPECT().CreateUser(user).Return(nil)
			},
			expectedStatusCode: 200,
		},

		{
			name: "notEnoughUserData",
			inputUser: entities.User{
				Login:     "almassagyndyk2@gmail.com",
				Password1: "19891042aA",
				Password2: "19891042aA",
			},
			mockBehavior: func(s *mock_service.MockService, user entities.User) {
				s.EXPECT().CreateUser(user).Return(errors.New("Error")) //mock service - return error -><
			},
			expectedStatusCode: 400,
		},

		{
			name: "incorrectUserData",
			inputUser: entities.User{
				Name:      "almas",
				Login:     "almassagyndyk2@gmail.com",
				Password1: "19891",
				Password2: "19891",
			},
			mockBehavior: func(s *mock_service.MockService, user entities.User) {
				s.EXPECT().CreateUser(user).Return(errors.New("Error")) //mock service - return error -><
			},
			expectedStatusCode: 400,
		},
		//handler -> mockService(data) -> response.Body ->error user mode, status ->
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Arrange
			//Init MockService
			c := gomock.NewController(t)
			defer c.Finish()
			service := mock_service.NewMockService(c)
			testCase.mockBehavior(service, testCase.inputUser)
			// service.EXPECT().CreateUser(testCase.inputUser).Return(nil)
			//Init handler
			mux := http.NewServeMux()
			handler := NewHandler(service)
			handler.Register(mux)

			data := url.Values{}
			data.Set("uName", testCase.inputUser.Name)
			data.Set("eMail", testCase.inputUser.Login)
			data.Set("password1", testCase.inputUser.Password1)
			data.Set("password2", testCase.inputUser.Password2)

			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", "/regPage", strings.NewReader(data.Encode()))

			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
			//Act
			mux.ServeHTTP(w, req)
			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)

		})
	}

}

func TestHandler2_regPage(t *testing.T) {

	type Test2 struct {
		name      string
		inputUser entities.User

		expectedStatusCode int
	}

	test := Test2{
		name: "missingUserData",
		inputUser: entities.User{
			Name:      "almas",
			Login:     "almassagyndyk2@gmail.com",
			Password1: "19891042aA",
			Password2: "19891042aA",
		},
		expectedStatusCode: 400,
	}

	t.Run(test.name, func(t *testing.T) {
		//Arrange
		c := gomock.NewController(t)
		defer c.Finish()
		service := mock_service.NewMockService(c)
		// Здесь ненужен MockBehavior т.к данные о юзере не доходят до MockBehavior, если она будет то выдаст ошибку
		// testCase.mockBehavior(service, testCase.inputUser)
		handler := NewHandler(service)
		mux := http.NewServeMux()
		handler.Register(mux)
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/regPage", nil)
		t.Logf("method is %s", req.Method)

		//Act
		mux.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, test.expectedStatusCode, w.Code)
		// assert.Equal(t, testCase.expectedRequestBody, w.Body)

	})

}

func TestHandler3_regPage(t *testing.T) {

	type Test3 struct {
		name               string
		inputUser          entities.User
		expectedStatusCode int
	}

	test := Test3{
		name: "r.FormValueIncorrect",
		inputUser: entities.User{
			Name:      "almas",
			Login:     "almassagyndyk2@gmail.com",
			Password1: "19891042aA",
			Password2: "19891042aA",
		},
		expectedStatusCode: 400,
	}

	t.Run(test.name, func(t *testing.T) {
		//Arrange
		c := gomock.NewController(t)
		defer c.Finish()
		service := mock_service.NewMockService(c)
		// Здесь ненужен MockBehavior т.к данные о юзере не доходят до MockBehavior, если она будет то выдаст ошибку
		// service.EXPECT().CreateUser(testCase.inputUser).Return(nil)
		handler := NewHandler(service)
		mux := http.NewServeMux()
		handler.Register(mux)

		data := url.Values{}
		data.Set("uame", test.inputUser.Name)
		data.Set("eMil", test.inputUser.Login)
		data.Set("pasword1", test.inputUser.Password1)
		data.Set("paswod2", test.inputUser.Password2)

		w := httptest.NewRecorder()

		req := httptest.NewRequest("POST", "/regPage", strings.NewReader(data.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		//Act
		mux.ServeHTTP(w, req)

		//Assert
		assert.Equal(t, test.expectedStatusCode, w.Code)

	})

}
