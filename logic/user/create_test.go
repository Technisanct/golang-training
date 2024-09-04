package user

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	userMocks "golang-training/repository/user/mocks"
	"testing"
)

func Test_userImpl_Create(t *testing.T) {
	type fields struct {
		user *userMocks.User
	}
	type args struct {
		ctx     context.Context
		request *CreateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "happy path",
			fields: fields{user: mockCreateUserRepo(true, nil)},
			args: args{
				ctx: context.Background(),
				request: &CreateUserRequest{
					Firstname: "fname",
					Lastname:  "lname",
					Email:     "email",
					Phone:     12345 - 67898,
				},
			},
			wantErr: nil,
		},
		{
			name:   "repo error",
			fields: fields{user: mockCreateUserRepo(true, errors.New("failed"))},
			args: args{
				ctx: context.Background(),
				request: &CreateUserRequest{
					Firstname: "fname",
					Lastname:  "lname",
					Email:     "email",
					Phone:     12345 - 67898,
				},
			},
			wantErr: errors.New("failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := userImpl{
				user: tt.fields.user,
			}
			err := u.Create(tt.args.ctx, tt.args.request)
			assert.Equal(t, tt.wantErr, err)

			tt.fields.user.AssertExpectations(t)
		})
	}
}

func mockCreateUserRepo(enableFlag bool, createErr error) *userMocks.User {
	client := &userMocks.User{}
	if enableFlag {
		client.On("Create", mock.Anything, mock.Anything).Return(createErr)
	}

	return client
}
