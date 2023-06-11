package entity

// TableName - returns name of the table
func (*Forecast) TableName() string {
	return "forecasts"
}

type Forecast struct {
	ID             int     `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	Country        string  `gorm:"Column:country" json:"country"`
	State          string  `gorm:"Column:state" json:"state"`
	AvgTemperature int     `gorm:"Column:avg_temperature" json:"avg_temperature"`
	MaxTemperature int     `gorm:"Column:max_temperature" json:"max_temperature"`
	MinTemperature int     `gorm:"Column:min_temperature" json:"min_temperature"`
	Description    string  `gorm:"Column:description" json:"description"`
	Humidity       float32 `gorm:"Column:humidity"  json:"humidity"`
	Code           int     `gorm:"Column:code"  json:"code"`
	Wind           float32 `gorm:"Column:wind" json:"wind"`
	UV             int     `gorm:"Column:uv"  json:"uv"`
	Date           string  `gorm:"Column:date" json:"date"`
}

type RequestForecast struct {
	Current      Current      `json:"current"`
	ForecastInfo ForecastInfo `json:"forecast"`
}

type Current struct {
	TempC     float32   `json:"temp_c"`
	Wind      float32   `json:"wind_kph"`
	Condition Condition `json:"condition"`
	Humidty   float32   `json:"humidity"`
	Pressure  int       `json:"pressure_mb"`
	Date      string    `json:"last_updated"`
	UV        float32   `json:"uv"`
}

type ForecastInfo struct {
	ForecastDay []ForecastDay `json:"forecastday"`
}

type ForecastDay struct {
	Date string `json:"date"`
	Day  Day    `json:"day"`
}

type Day struct {
	MaxTempC       float32   `json:"maxtemp_c"`
	AvgTemperature float32   `json:"avgtemp_c"`
	MinTempC       float32   `json:"mintemp_c"`
	Wind           float32   `json:"maxwind_kph"`
	Condition      Condition `json:"condition"`
	AvgHumidity    float32   `json:"avghumidity"`
	UV             float32   `json:"uv"`
}

type Condition struct {
	Text string `json:"text"`
	Code int    `json:"code"`
	Icon string `json:"icon"`
}
