package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	userMocks "golang-training/logic/user/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	validCreateUserRequestHandler   = []byte(`{"firstname": "fname", "lastname": "lname", "email": "email", "phone": 1234}`)
	invalidCreateUserRequestHandler = []byte(`{"fname": "fname", "lastname": "lname", "email": "email", "phone": 1234}`)
)

func Test_handler_CreateUser(t *testing.T) {
	type fields struct {
		user *userMocks.Users
	}
	type args struct {
		c       *gin.Context
		request []byte
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		expectedStatusCode int
		expectedResponse   interface{}
	}{
		{
			name: "happy path",
			fields: fields{
				user: mockLogicUser(true, nil),
			},
			args: args{
				c:       nil,
				request: validCreateUserRequestHandler,
			},
			expectedResponse:   &CreateUserResponse{Message: "user created successfully"},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "logic error ",
			fields: fields{
				user: mockLogicUser(true, errors.New("failed")),
			},
			args: args{
				c:       nil,
				request: validCreateUserRequestHandler,
			},
			expectedResponse:   &CreateUserResponse{Message: ""},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Assemble
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			payload := bytes.NewBuffer(tt.args.request)
			r, _ := http.NewRequest("POST", "/user", payload)

			c.Request = r

			t.Run(tt.name, func(t *testing.T) {
				h := handler{
					user: tt.fields.user,
				}
				h.CreateUser(c)
			})

			// Assert
			assert.Equal(t, tt.expectedStatusCode, w.Code)

			var response *CreateUserResponse
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expectedResponse, response)

			tt.fields.user.AssertExpectations(t)
		})
	}
}

func mockLogicUser(enableFlag bool, err error) *userMocks.Users {
	client := &userMocks.Users{}
	if enableFlag {
		client.On("Create", mock.Anything, mock.Anything).Return(err)
	}

	return client
}
