package main

import (
	"encoding/json"
	"io"
	"minerals/internal"
	"minerals/pkg/calculations"
	"os"
)

func main() {

	f, err := os.Open("./internal/info/minerals.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var info internal.Info
	jsonInfo, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonInfo, &info)
	if err != nil {
		panic(err)
	}

	err = calculations.CalcFromGiven(info)
	if err != nil {
		panic(err)
	}
}
