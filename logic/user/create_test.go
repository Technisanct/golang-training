package user

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	repoMocks "golang-training/repository/mocks"
	"testing"
)

func Test_userImpl_Create(t *testing.T) {
	type fields struct {
		repo *repoMocks.Repository
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
			fields: fields{repo: mockCreateUserRepo(true, nil)},
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
			fields: fields{repo: mockCreateUserRepo(true, errors.New("failed"))},
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
				repo: tt.fields.repo,
			}
			err := u.Create(tt.args.ctx, tt.args.request)
			assert.Equal(t, tt.wantErr, err)

			tt.fields.repo.AssertExpectations(t)
		})
	}
}

func mockCreateUserRepo(enableFlag bool, createErr error) *repoMocks.Repository {
	client := &repoMocks.Repository{}
	if enableFlag {
		client.On("CreateUser", mock.Anything, mock.Anything).Return(createErr)
	}

	return client
}
