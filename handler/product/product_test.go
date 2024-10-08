package product

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-training/logic/product"
	productMocks "golang-training/logic/product/mocks"
	"net/http"
	"net/http/httptest"

	"testing"
	"time"
)

var (
	validUUID = "uuid-123"

	validData = &Product{
		ID:              "",
		UUID:            "",
		ProductName:     "",
		Price:           0,
		DiscountedPrice: 0,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
	}

	validLogicResponse = &product.Product{
		ID:              "",
		UUID:            "",
		Name:            "",
		Price:           0,
		DiscountedPrice: 0,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
	}
)

func Test_handler_Get(t *testing.T) {
	type fields struct {
		product *productMocks.Products
	}
	type args struct {
		ctx  *gin.Context
		uuid string
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
				product: mockLogicProduct(validLogicResponse, nil),
			},
			args: args{
				ctx:  nil,
				uuid: validUUID,
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   validData,
		},
		{
			name: "logic error",
			fields: fields{
				product: mockLogicProduct(nil, errors.New("failed")),
			},
			args: args{
				ctx:  nil,
				uuid: validUUID,
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   validData,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req, _ := http.NewRequest("GET", "/product/"+tt.args.uuid, nil)
			c.Request = req

			h := handler{
				product: tt.fields.product,
			}
			h.Get(c)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			var response *Product
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expectedResponse, response)

			tt.fields.product.AssertExpectations(t)

		})
	}
}

func mockLogicProduct(product *product.Product, err error) *productMocks.Products {
	client := &productMocks.Products{}
	client.On("Get", mock.Anything, mock.Anything).Return(product, err)

	return client
}

var (
	validReq = []byte(`{"productName":"name","price":123.00}`)
)

func Test_handler_CreateProduct(t *testing.T) {
	type fields struct {
		product *productMocks.Products
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
				product: mockProductLogic(true, nil),
			},
			args: args{
				c:       nil,
				request: validReq,
			},
			expectedResponse:   &CreateProductResponse{Message: "created successfully"},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "logic error",
			fields: fields{
				product: mockProductLogic(true, errors.New("failed")),
			},
			args: args{
				c:       nil,
				request: validReq,
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
			r, _ := http.NewRequest("POST", "/product", payload)

			c.Request = r

			t.Run(tt.name, func(t *testing.T) {
				h := handler{
					product: tt.fields.product,
				}
				h.CreateProduct(c)
			})

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			var res *CreateProductResponse
			if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expectedResponse, res)

			tt.fields.product.AssertExpectations(t)

		})
	}
}

func mockProductLogic(enableFlag bool, err error) *productMocks.Products {
	client := &productMocks.Products{}
	if enableFlag {
		client.On("Create", mock.Anything, mock.Anything).Return(err)
	}

	return client
}

func Test_handler_Delete(t *testing.T) {
	type fields struct {
		product *productMocks.Products
	}
	type args struct {
		c    *gin.Context
		uuid string
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
				product: mockDeleteProductLogic(true, nil),
			},
			args: args{
				c:    nil,
				uuid: "uuid-123",
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   &DeleteProductResponse{Message: "deleted successfully"},
		},
		{
			name: "logic error",
			fields: fields{
				product: mockDeleteProductLogic(true, errors.New("failed")),
			},
			args: args{
				c:    nil,
				uuid: "uuid-123",
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   &DeleteProductResponse{Message: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req, _ := http.NewRequest("DELETE", "/product/"+tt.args.uuid, nil)
			c.Request = req

			h := handler{
				product: tt.fields.product,
			}
			h.Delete(c)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			var response *DeleteProductResponse
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expectedResponse, response)

			tt.fields.product.AssertExpectations(t)
		})
	}
}

func mockDeleteProductLogic(enableFlag bool, err error) *productMocks.Products {
	client := &productMocks.Products{}
	if enableFlag {
		client.On("Delete", mock.Anything, mock.Anything).Return(err)
	}

	return client
}
