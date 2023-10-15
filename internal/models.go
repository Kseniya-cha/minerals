package internal

type Info struct {
	Minerals []Minerals `json:"minerals"`
}

type Minerals struct {
	Name string  `json:"name"`
	Ca   float64 `json:"ca"`
	P    float64 `json:"p"`
	Mg   float64 `json:"mg"`
}
