package internal

type Info struct {
	Products []Products `json:"products"`
}

// информация о продукте:
// калорийность в ккал;
// остальное - мг 100 гр
type Products struct {
	Id     int    `json:"id"`
	Animal string `json:"animal"`
	Part   string `json:"part"`

	Kkal    float64 `json:"kkal"`
	Protein float64 `json:"protein"`
	Fat     float64 `json:"fat"`
	Carbs   float64 `json:"carbs"`
	Water   float64 `json:"water"`

	Ca     float64 `json:"ca"`
	P      float64 `json:"p"`
	Mg     float64 `json:"mg"`
	Na     float64 `json:"na"`
	K      float64 `json:"k"`
	Fe     float64 `json:"fe"`
	A      float64 `json:"a"`
	Omega3 float64 `json:"omega3"`
	Omega6 float64 `json:"omega6"`
}

// для файла given.json, информация о кошке и продуктах
type GivenProducts struct {
	Cat      CatInfo `json:"cat"`
	Products []Given `json:"products"`
}

// информация о продуктах в смеси
type Given struct {
	// id соответствует id продукта в minerals.json
	Id int `json:"id"`
	// масса продукта в смеси
	Mass float64 `json:"mass"`
}

// информация о кошке: масса в кг
type CatInfo struct {
	Mass float64 `json:"mass"`
}

//

// в мг на кг массы
type Result struct {
	Mass float64 `json:"mass"`

	Kkal    float64 `json:"kkal"`
	Protein float64 `json:"protein"`
	Fat     float64 `json:"fat"`
	Carbs   float64 `json:"carbs"`
	Water   float64 `json:"water"`

	Ca     float64 `json:"ca"`
	P      float64 `json:"p"`
	Mg     float64 `json:"mg"`
	Na     float64 `json:"na"`
	K      float64 `json:"k"`
	Fe     float64 `json:"fe"`
	A      float64 `json:"a"`
	Omega3 float64 `json:"omega3"`
	Omega6 float64 `json:"omega6"`
}
