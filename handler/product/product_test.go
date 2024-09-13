package product

import (
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
