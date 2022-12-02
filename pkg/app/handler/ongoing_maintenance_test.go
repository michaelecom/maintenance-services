package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"rimeks.ru/services/pkg/app/service"
	mock_service "rimeks.ru/services/pkg/app/service/mocks"
	"rimeks.ru/services/pkg/app/structures"
)

func TestHandler_createOrder(t *testing.T) {
	type mockBehavior func(r *mock_service.MockOngoingMaintenance, order structures.OrderList)

	tests := []struct {
		name                 string
		inputBody            string
		inputData            structures.OrderList
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"service_market_id":1,"orders":[{"order_number":"000000000000000","car_brand":"brand","car_model":"model","car_number":"number"}]}`,
			inputData: structures.OrderList{
				ServiceMarketID: 1,
				Orders: []structures.OrderData{
					{
						OrderNumber: "000000000000000",
						CarBrand:    "brand",
						CarModel:    "model",
						CarNumber:   "number",
					},
				},
			},
			mockBehavior: func(r *mock_service.MockOngoingMaintenance, order structures.OrderList) {
				r.EXPECT().CreateOrder(order).Return(1, nil)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"id":0}`,
		},
		{
			name:                 "BadRequest",
			inputBody:            `{"service_market_id":"1","orders":[{"order_number":"000000000000000","car_brand":"brand","car_model":"model","car_number":"number"}]}`,
			inputData:            structures.OrderList{},
			mockBehavior:         func(r *mock_service.MockOngoingMaintenance, order structures.OrderList) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Неверный формат переданных данных"}`,
		},
		{
			name:      "InternalServerError",
			inputBody: `{"service_market_id":1,"orders":[{"order_number":"000000000000000","car_brand":"brand","car_model":"model","car_number":"very long number"}]}`,
			inputData: structures.OrderList{
				ServiceMarketID: 1,
				Orders: []structures.OrderData{
					{
						OrderNumber: "000000000000000",
						CarBrand:    "brand",
						CarModel:    "model",
						CarNumber:   "very long number",
					},
				},
			},
			mockBehavior: func(r *mock_service.MockOngoingMaintenance, order structures.OrderList) {
				r.EXPECT().CreateOrder(order).Return(0, errors.New("Внутренняя ошибка сервера при попытке добавления записи"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"Внутренняя ошибка сервера при попытке добавления записи"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_service.NewMockOngoingMaintenance(c)
			test.mockBehavior(store, test.inputData)

			service := &service.Service{OngoingMaintenance: store}
			handler := Handler{service}

			r := gin.New()
			r.POST("/api/ongoing-maintenance/", handler.createOrder)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/ongoing-maintenance/",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(rec, req)

			assert.Equal(t, rec.Code, test.expectedStatusCode)
			assert.Equal(t, rec.Body.String(), test.expectedResponseBody)
		})
	}
}
