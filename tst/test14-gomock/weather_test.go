// weather_test.go
package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetWeather(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAPI := NewMockWeatherAPI(ctrl)
	mockAPI.EXPECT().Fetch("Delhi").Return("Cold", nil)

	weatherService := NewWeatherService(mockAPI)
	weather, err := weatherService.GetWeather("Delhi")

	assert.NoError(t, err)
	assert.Equal(t, "Cold", weather)
}
