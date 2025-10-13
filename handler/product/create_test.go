package product

import (
	"bytes"
	"encoding/json"
	"golang-training/logic/product/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProductHandler(t *testing.T) {

	type mockService struct {
		productServ *mocks.Products
	}

	type Args struct {
		request []byte
	}

	tests := []struct {
		name               string
		mocked             mockService
		args               Args
		ctx                *gin.Context
		expectedStatusCode int
		expectedResponse   any
	}{
		{
			name: "should return 201",
			ctx:  nil,
			args: Args{
				request: []byte(`{"name": "test1", "price": 100.00, "discount_price": 10.00}`),
			},
			mocked: mockService{
				productServ: mockLogicProduct(true, nil),
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse: &CreateProductResponse{
				Status:  "success",
				Message: "product successfully created",
			},
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
					product: tt.mocked.productServ,
				}
				h.CreateProductHandler(c)
			})

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			var response *CreateProductResponse
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expectedResponse, response)
		})
	}
}

func TestErrorsOnCreateProductHandler(t *testing.T) {
	type mockService struct {
		productServ *mocks.Products
	}

	type Args struct {
		request []byte
	}

	tests := []struct {
		name               string
		mocked             mockService
		args               Args
		ctx                *gin.Context
		expectedStatusCode int
		expectedResponse   any
	}{
		{
			name: "should return 400",
			ctx:  nil,
			args: Args{
				request: []byte(`{"name": "test1", "price": 0, "discount_price": 10.00}`),
			},
			mocked: mockService{
				productServ: mockLogicProduct(true, nil),
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "should return 400",
			ctx:  nil,
			args: Args{
				request: []byte(`{"name": "test1", "discount_price": 10.00}`),
			},
			mocked: mockService{
				productServ: mockLogicProduct(true, nil),
			},
			expectedStatusCode: http.StatusBadRequest,
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
					product: tt.mocked.productServ,
				}
				h.CreateProductHandler(c)
			})

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			var response map[string]any
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}

			assert.Contains(t, response["error"], "Price")
		})
	}
}

func mockLogicProduct(enableFlag bool, err error) *mocks.Products {
	client := &mocks.Products{}
	if enableFlag {
		client.On("CreateProduct", mock.Anything, mock.Anything).Return(err)
	}

	return client
}
