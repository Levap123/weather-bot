package entity

type Weather struct {
	Location struct {
		CityName string `json:"name"`
	}
	Data struct {
		Values struct {
			Temperature         float32 `json:"temperature"`
			TemperatureApparent float32 `json:"temperatureApparent"`
		}
	}
}
