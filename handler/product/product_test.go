package product

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang-training/logic/product/contract"
	"golang-training/logic/product/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	validCreateProductRequest   = []byte(`{"name": "test1", "price": 100.00, "discount_price": 10.00}`)
	invalidCreateProductRequest = []byte(`{"name": "test1", "discount_price": 10.00}`)

	fixedObjectId         = primitive.NewObjectID()
	fixedTestTime         = time.Date(2025, time.October, 0, 0, 0, 0, 0, time.UTC)
	returnListProductData = []*contract.Product{
		{
			ID:              fixedObjectId.String(),
			Name:            "test-1",
			Price:           100.000,
			DiscountedPrice: 10.00,
			CreatedAt:       fixedTestTime,
			UpdatedAt:       fixedTestTime,
		},
		{
			ID:              fixedObjectId.String(),
			Name:            "test-1",
			Price:           100.000,
			DiscountedPrice: 10.00,
			CreatedAt:       fixedTestTime,
			UpdatedAt:       fixedTestTime,
		},
	}
	expectedListProductResponse = ListProductResponse{
		Message: "successful",
		Data:    mapLogicToHandler(returnListProductData),
	}
	expectedErrorListProductResponse = ListProductResponse{
		Message: "",
		Data:    nil,
	}
)

func TestCreateProductHandler(t *testing.T) {
	type fields struct {
		product *mocks.Products
	}
	type args struct {
		request []byte
		ctx     *gin.Context
	}

	tests := []struct {
		name               string
		fields             fields
		args               args
		expectedStatusCode int
		expectedResponse   any
	}{
		{
			name: "happy path",
			args: args{
				request: validCreateProductRequest,
				ctx:     nil,
			},
			fields: fields{
				product: mockLogicProduct(true, nil),
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse: &CreateProductResponse{
				Message: "product successfully created",
			},
		},
		{
			name: "logic error",
			args: args{
				ctx:     nil,
				request: validCreateProductRequest,
			},
			fields: fields{
				product: mockLogicProduct(true, errors.New("failed")),
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   &CreateProductResponse{Message: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			payload := bytes.NewBuffer(tt.args.request)
			r, _ := http.NewRequest("POST", "/product/create", payload)

			c.Request = r

			t.Run(tt.name, func(t *testing.T) {
				h := handler{
					product: tt.fields.product,
				}
				h.CreateProduct(c)
			})

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			var response *CreateProductResponse
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expectedResponse, response)
			tt.fields.product.AssertExpectations(t)
		})
	}
}

func TestListHandler(t *testing.T) {
	type fields struct {
		product *mocks.Products
	}
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name               string
		expectedStatusCode int
		expectedResponse   any
		args               args
		fields             fields
	}{
		{
			name: "happy path",
			args: args{
				ctx: nil,
			},
			fields: fields{
				product: mockListLogicProduct(true, returnListProductData, nil),
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   expectedListProductResponse,
		},
		{
			name: "error",
			args: args{
				ctx: nil,
			},
			fields: fields{
				product: mockListLogicProduct(true, nil, errors.New("failed to fetch products")),
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   expectedErrorListProductResponse,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			r, _ := http.NewRequest("GET", "/product/", nil)
			c.Request = r

			t.Run(tt.name, func(t *testing.T) {
				h := handler{
					product: tt.fields.product,
				}
				h.ListProduct(c)
			})

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			var response *ListProductResponse
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expectedResponse, *response)
			tt.fields.product.AssertExpectations(t)
		})
	}
}

func mockLogicProduct(enableFlag bool, err error) *mocks.Products {
	client := &mocks.Products{}
	if enableFlag {
		client.On("Create", mock.Anything, mock.Anything).Return(err)
	}
	return client
}

func mockListLogicProduct(enableFlag bool, returnProductListData []*contract.Product, createErr error) *mocks.Products {
	client := &mocks.Products{}
	if enableFlag {
		client.
			On("List", mock.Anything).
			Return(returnProductListData, createErr)
	}
	return client
}
