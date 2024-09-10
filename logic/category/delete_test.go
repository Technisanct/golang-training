package category

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	categoryMocks "golang-training/repository/category/mocks"
	"testing"
)

func TestImpl_Delete(t *testing.T) {
	type fields struct {
		categoryRepo *categoryMocks.Category
	}
	type args struct {
		ctx  context.Context
		uuid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "happy path",
			fields: fields{categoryRepo: mockDeleteCategoryRepo(true, nil)},
			args: args{
				ctx:  context.Background(),
				uuid: "123",
			},
			wantErr: nil,
		},
		{
			name:   "repo error",
			fields: fields{categoryRepo: mockDeleteCategoryRepo(true, errors.New("failed"))},
			args: args{
				ctx:  context.Background(),
				uuid: "234",
			},
			wantErr: errors.New("failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Impl{
				categoryRepo: tt.fields.categoryRepo,
			}
			err := c.Delete(tt.args.ctx, tt.args.uuid)
			assert.Equal(t, tt.wantErr, err)

			tt.fields.categoryRepo.AssertExpectations(t)
		})
	}
}

func mockDeleteCategoryRepo(enableFlag bool, createErr error) *categoryMocks.Category {
	client := &categoryMocks.Category{}
	if enableFlag {
		client.On("Delete", mock.Anything, mock.Anything).Return(createErr)
	}
	return client
}
