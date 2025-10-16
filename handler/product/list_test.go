package product

import (
	"encoding/json"
	"golang-training/logic/product/mocks"
	"golang-training/repository/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	returnListProductData = []model.Product{
		{
			Name:            "test-1",
			Price:           100.00,
			DiscountedPrice: 10.00,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			Name:            "test-2",
			Price:           100.00,
			DiscountedPrice: 10.00,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}
)

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

			assert.Equal(t, response.Message, "successful")
			assert.NotNil(t, response.Message)
			assert.NotNil(t, response.Data)

			tt.fields.product.AssertExpectations(t)
		})
	}
}

func mockListLogicProduct(enableFlag bool, returnProductListData []model.Product, createErr error) *mocks.Products {
	client := &mocks.Products{}
	if enableFlag {
		client.
			On("List", mock.Anything).
			Return(returnProductListData, createErr)
	}
	return client
}
