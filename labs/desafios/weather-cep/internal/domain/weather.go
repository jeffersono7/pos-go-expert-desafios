package domain

type Weather struct {
	TempC float32
	TempF float32
	TempK float32
}

func NewWeather(tempC float32) Weather {
	return Weather{
		TempC: tempC,
		TempF: (tempC * 1.8) + 32,
		TempK: tempC + 273,
	}
}
