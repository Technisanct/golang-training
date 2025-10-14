package product

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang-training/logic/product/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	validCreateProductRequest   = []byte(`{"name": "test1", "price": 100.00, "discount_price": 10.00}`)
	invalidCreateProductRequest = []byte(`{"name": "test1", "discount_price": 10.00}`)
)

func TestCreateProductHandler(t *testing.T) {
	type repo struct {
		product *mocks.Products
	}
	type Args struct {
		request []byte
		ctx     *gin.Context
	}

	tests := []struct {
		name               string
		fields             repo
		args               Args
		expectedStatusCode int
		expectedResponse   any
	}{
		{
			name: "happy path",
			args: Args{
				request: validCreateProductRequest,
				ctx:     nil,
			},
			fields: repo{
				product: mockLogicProduct(true, nil),
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse: &CreateProductResponse{
				Message: "product successfully created",
			},
		},
		{
			name: "logic error",
			args: Args{
				ctx:     nil,
				request: validCreateProductRequest,
			},
			fields: repo{
				product: mockLogicProduct(true, errors.New("failed")),
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   &CreateProductResponse{Message: ""},
		},
		{
			name: "validation error",
			args: Args{
				ctx:     nil,
				request: invalidCreateProductRequest,
			},
			fields: repo{
				product: mockLogicProduct(false, nil),
			},
			expectedStatusCode: http.StatusBadRequest,
			// expectedResponse: ,
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

			if tt.name == "validation error" {
				var errResponse string
				if err := json.Unmarshal(w.Body.Bytes(), &errResponse); err != nil {
					assert.Contains(t, strings.ToLower(errResponse), "price")
				}
				return
			}

			var response *CreateProductResponse
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expectedResponse, response)
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
