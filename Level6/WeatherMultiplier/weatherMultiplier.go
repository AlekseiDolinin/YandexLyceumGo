package main

import (
	"fmt"
)

type WeatherCondition int

const (
	Clear WeatherCondition = iota
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	coeff := 1.0
	if weather.Condition == Rain {
		coeff += 0.125
	}
	if weather.Condition == HeavyRain {
		coeff += 0.2
	}
	if weather.Condition == Snow {
		coeff += 0.15
	}
	if weather.WindSpeed > 15 {
		coeff += 0.1
	}
	return coeff
}

func main() {
	fmt.Println(GetWeatherMultiplier(WeatherData{HeavyRain, 10})) //1.2
	fmt.Println(GetWeatherMultiplier(WeatherData{HeavyRain, 20})) //1.3

	fmt.Println(GetWeatherMultiplier(WeatherData{Rain, 10})) //1.125
	fmt.Println(GetWeatherMultiplier(WeatherData{Rain, 20})) //1.225

	fmt.Println(GetWeatherMultiplier(WeatherData{Clear, 10})) //1
	fmt.Println(GetWeatherMultiplier(WeatherData{Clear, 20})) //1.1

	fmt.Println(GetWeatherMultiplier(WeatherData{Snow, 10})) //1.15
	fmt.Println(GetWeatherMultiplier(WeatherData{Snow, 20})) //1.25
}
