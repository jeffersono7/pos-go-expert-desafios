package controller

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/service"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetWeather(t *testing.T) {
	cases := []struct {
		description    string
		input          string
		setup          func(*mocks.CepClientMock, *mocks.WeatherClientMock)
		expectedStatus int
		expectedBody   string
	}{
		{
			description: "when all valid returns weather",
			input:       "11122290",
			setup: func(cepClient *mocks.CepClientMock, weatherClient *mocks.WeatherClientMock) {
				cepClient.On("GetCEP", mock.Anything, "11122290").Return(service.CepResp{Localidade: "city", Estado: "df"}, nil)

				weatherClient.On("GetTemp", mock.Anything, "city df").Return(service.WeatherResp{
					Current: struct {
						TempC float32 "json:\"temp_c\""
					}{TempC: 12},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "{\"temp_C\":12,\"temp_F\":53.6,\"temp_K\":285}\n",
		},
		{
			description:    "when invalid zipcode",
			input:          "234",
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   "{\"message\":\"invalid zipcode\"}\n",
		},
		{
			description: "when cep not found",
			input:       "33322211",
			setup: func(cep *mocks.CepClientMock, weather *mocks.WeatherClientMock) {
				cep.On("GetCEP", mock.Anything, "33322211").Return(service.CepResp{}, fmt.Errorf("not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "{\"message\":\"can not find zipcode\"}\n",
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			t.Parallel()

			cepClient := &mocks.CepClientMock{}
			weatherClient := &mocks.WeatherClientMock{}
			if tt.setup != nil {
				tt.setup(cepClient, weatherClient)
			}
			weatherService := service.NewWeatherService(cepClient, weatherClient)

			weatherController := NewWeatherController(weatherService)

			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
			defer cancel()

			req := httptest.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("/cep/%s/weather", tt.input), nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("cep", tt.input)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			w := httptest.NewRecorder()

			weatherController.GetWeather(w, req)

			resp := w.Result()
			defer resp.Body.Close()
			respBody, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			assert.Equal(t, tt.expectedBody, string(respBody))
		})
	}
}
