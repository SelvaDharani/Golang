**\* Integrate GoMock with testify to test a service that depends on an external interface\_**

In weather.go, we define a WeatherAPI interface and a WeatherService that uses this interface.
WeatherService has a method GetWeather that calls the Fetch method of WeatherAPI.

use mockgen to generate a mock implementation of WeatherAPI.

In the test file weather_test.go,we:

Create a mock instance of WeatherAPI.
Set up an expectation that Fetch will be called with "London" and return "Sunny" and nil for the error.
Create an instance of WeatherService using the mock API.
Call GetWeather and assert that no error is returned and the result is as expected.

mockgen -source=weather.go -destination=mock_weather.go -package=main
