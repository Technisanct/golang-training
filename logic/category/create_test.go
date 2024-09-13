package category

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	categoryMocks "golang-training/repository/category/mocks"
	"testing"
)

func TestImpl_Create(t *testing.T) {
	type fields struct {
		categoryRepo *categoryMocks.Category
	}
	type args struct {
		ctx     context.Context
		request *CreateCategoryRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "happy path",
			fields: fields{categoryRepo: mockCreateCategoryRepo(true, nil)},
			args: args{
				ctx: context.Background(),
				request: &CreateCategoryRequest{
					Name: "name",
				},
			},
			wantErr: nil,
		},
		{
			name:   "repo error",
			fields: fields{categoryRepo: mockCreateCategoryRepo(true, errors.New("failed"))},
			args: args{
				ctx: context.Background(),
				request: &CreateCategoryRequest{
					Name: "name",
				},
			},
			wantErr: errors.New("failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Impl{
				categoryRepo: tt.fields.categoryRepo,
			}
			err := u.Create(tt.args.ctx, tt.args.request)
			assert.Equal(t, tt.wantErr, err)

			tt.fields.categoryRepo.AssertExpectations(t)
		})
	}
}

func mockCreateCategoryRepo(enableFlag bool, createErr error) *categoryMocks.Category {
	client := &categoryMocks.Category{}
	if enableFlag {
		client.On("Create", mock.Anything, mock.Anything).Return(createErr)
	}
	return client
}
