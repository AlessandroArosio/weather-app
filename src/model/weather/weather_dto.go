package weather

type WeatherResponse struct {
	Count int    `json:"count"`
	Data  []Data `json:"data"`
}
type Weather struct {
	Icon        string `json:"icon"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
type Data struct {
	CityName string  `json:"city_name"`
	Temp     float64 `json:"temp"`
}
