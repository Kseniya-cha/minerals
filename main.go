package main

import (
	"encoding/json"
	"fmt"
	"io"
	"minerals/internal"
	"os"
)

func main() {
	var info internal.Info

	f, err := os.Open("")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	jsonInfo, err := io.ReadAll(f)

	err = json.Unmarshal(jsonInfo, &info)

	fmt.Println(info)
}
