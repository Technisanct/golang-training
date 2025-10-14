package product

//
// import (
// 	"context"
// 	"errors"
// 	"golang-training/repository/model"
// 	"golang-training/repository/product/mocks"
// 	"log"
// 	"testing"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// var (
// 	createProductRequest = model.Product{
// 		Name:            "test-1",
// 		Price:           100.00,
// 		DiscountedPrice: 10,
// 	}
// )

// func TestProductCreate(t *testing.T) {

// 	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	type args struct {
// 		request *model.Product
// 		ctx     context.Context
// 	}

// 	tests := []struct {
// 		name       string
// 		args       args
// 		mockReturn error
// 		wantErr    error
// 	}{
// 		{
// 			name: "happy path",
// 			args: args{
// 				request: &createProductRequest,
// 				ctx:     c,
// 			},
// 			mockReturn: nil,
// 			wantErr:    nil,
// 		},
// 		{
// 			name: "should throw error when nil or empty data is passed for creation",
// 			args: args{
// 				request: nil,
// 				ctx:     c,
// 			},
// 			mockReturn: errors.New("write error"),
// 			wantErr:    errors.New("write error"),
// 		},
// 		{
// 			name: "should throw error when timeout is finished or function is taking too much time than its limits",
// 			args: args{
// 				request: &createProductRequest,
// 				ctx:     c,
// 			},
// 			mockReturn: context.DeadlineExceeded,
// 			wantErr:    context.DeadlineExceeded,
// 		},
// 	}
// 	for _, tt := range tests {

// 		ProductRepoI := mocks.NewProduct(t)
// 		ProductRepoI.
// 			On("Create", mock.Anything, tt.args.request).
// 			Return(tt.mockReturn)

// 		t.Run(tt.name, func(t *testing.T) {
// 			err := ProductRepoI.Create(tt.args.ctx, tt.args.request)
// 			assert.Equal(t, err, tt.wantErr)
// 		})

// 	}
// }

// func InitInMemoryDb() (*memongo.Server, string) {
// 	mongoServer, err := memongo.Start("4.0.5")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return mongoServer, mongoServer.URI()
// }

// func TestProductRepoInMemoryDb(t *testing.T) {

// 	server, uri := InitInMemoryDb()
// 	defer server.Stop()

// 	clientOpts := options.Client().ApplyURI(uri)
// 	client, err := mongo.Connect(context.Background(), clientOpts)
// 	assert.NoError(t, err)

// 	db := client.Database("testdb")
// 	collection := db.Collection("products")
// 	repo := New(db)

// 	var findOne func(name string, product *model.Product) error

// 	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	type args struct {
// 		request *model.Product
// 		ctx     context.Context
// 	}

// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr error
// 	}{
// 		{
// 			name: "happy path",
// 			args: args{
// 				request: &createProductRequest,
// 				ctx:     c,
// 			},
// 			wantErr: nil,
// 		},
// 		{
// 			name: "repo err",
// 			args: args{
// 				request: nil,
// 				ctx:     c,
// 			},
// 			wantErr: nil,
// 		},
// 	}

// 	findOne = func(name string, product *model.Product) error {
// 		err := collection.FindOne(context.Background(), bson.M{
// 			"name": name,
// 		}).Decode(&product)

// 		return err
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err = repo.Create(tt.args.ctx, &createProductRequest)
// 			assert.NoError(t, err)

// 			var product model.Product
// 			err = findOne("test-1", &product)
// 			assert.NoError(t, err)
// 			assert.Equal(t, createProductRequest.Name, product.Name)

// 			err = findOne("test-2", &product)
// 			assert.Error(t, err)
// 		})
// 	}
// }
