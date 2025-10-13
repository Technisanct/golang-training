package product

import (
	"context"
	"errors"
	"golang-training/repository/model"
	"golang-training/repository/product/mocks"
	"log"
	"testing"
	"time"

	"github.com/benweissmann/memongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductCreate(t *testing.T) {

	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	type Args struct {
		data *model.Product
	}

	payload := &model.Product{
		Name:            "test-1",
		Price:           100.00,
		DiscountedPrice: 10,
	}

	tests := []struct {
		name       string
		ctx        context.Context
		args       Args
		mockReturn error
		wantErr    error
	}{
		{
			name: "should create user successfully",
			ctx:  c,
			args: Args{
				data: payload,
			},
			mockReturn: nil,
			wantErr:    nil,
		},
		{
			name: "should throw error when nil or empty data is passed for creation",
			ctx:  c,
			args: Args{
				data: nil,
			},
			mockReturn: errors.New("write error"),
			wantErr:    errors.New("write error"),
		},
		{
			name: "should throw error when timeout is finished or function is taking too much time than its limits",
			ctx:  c,
			args: Args{
				data: payload,
			},
			mockReturn: context.DeadlineExceeded,
			wantErr:    context.DeadlineExceeded,
		},
	}

	for _, tt := range tests {

		ProductRepoI := mocks.NewProduct(t)

		ProductRepoI.
			On("Create", mock.Anything, tt.args.data).
			Return(tt.mockReturn)

		t.Run(tt.name, func(t *testing.T) {
			err := ProductRepoI.Create(tt.ctx, tt.args.data)
			assert.Equal(t, err, tt.wantErr)
		})

	}
}

func InitInMemoryDb() (*memongo.Server, string) {
	mongoServer, err := memongo.Start("4.0.5")
	if err != nil {
		log.Fatal(err)
	}

	return mongoServer, mongoServer.URI()
}

func TestProductRepoInMemoryDb(t *testing.T) {

	server, uri := InitInMemoryDb()
	defer server.Stop()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOpts)
	assert.NoError(t, err)

	db := client.Database("testdb")
	collection := db.Collection("products")
	var findOne func(name string, product *model.Product) error

	repo := New(db)

	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	payload := &model.Product{
		Name:            "test-1",
		Price:           100.00,
		DiscountedPrice: 10,
	}

	type args struct {
		data *model.Product
	}

	tests := []struct {
		name    string
		ctx     context.Context
		args    args
		wantErr error
	}{
		{
			name: "should create successfully",
			ctx:  c,
			args: args{
				data: payload,
			},
			wantErr: nil,
		},
		{
			name: "should throw error if the data is nil",
			ctx:  c,
			args: args{
				data: nil,
			},
			wantErr: nil,
		},
	}

	findOne = func(name string, product *model.Product) error {
		err := collection.FindOne(context.Background(), bson.M{
			"name": name,
		}).Decode(&product)

		return err
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err = repo.Create(tt.ctx, payload)
			assert.NoError(t, err)

			var product model.Product
			err = findOne("test-1", &product)
			assert.NoError(t, err)
			assert.Equal(t, payload.Name, product.Name)

			err = findOne("test-2", &product)
			assert.Error(t, err)
		})
	}

}
