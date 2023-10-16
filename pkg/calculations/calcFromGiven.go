package calculations

import (
	"encoding/json"
	"fmt"
	"io"
	"minerals/internal"
	"os"
)

func CalcFromGiven(info internal.Info) error {
	f, err := os.Open("./given.json")
	if err != nil {
		return err
	}
	defer f.Close()

	var given internal.GivenProducts
	jsonGiven, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonGiven, &given)
	if err != nil {
		return err
	}

	var result internal.Result

	for _, product := range info.Products {
		for _, givenProduct := range given.Products {
			if product.Id == givenProduct.Id {
				fmt.Println(product.Animal, product.Part+":", givenProduct.Mass, "г")
				result.Mass += givenProduct.Mass
				result.Kkal += product.Kkal * givenProduct.Mass / 100
				result.Protein += product.Protein * givenProduct.Mass / 100 / given.Cat.Mass / 1000
				result.Fat += product.Fat * givenProduct.Mass / 100 / given.Cat.Mass / 1000
				result.Carbs += product.Carbs * givenProduct.Mass / 100 / given.Cat.Mass / 1000
				result.Water += product.Water * givenProduct.Mass / 100 / given.Cat.Mass / 1000

				result.Ca += product.Ca * givenProduct.Mass / 100 / given.Cat.Mass / 1000
				result.P += product.P * givenProduct.Mass / 100 / given.Cat.Mass / 1000
				result.Mg += product.Mg * givenProduct.Mass / 100 / given.Cat.Mass / 1000
				result.Na += product.Na * givenProduct.Mass / 100 / given.Cat.Mass / 1000
				result.K += product.K * givenProduct.Mass / 100 / given.Cat.Mass / 1000
				result.Fe += product.Fe * givenProduct.Mass / 100 / given.Cat.Mass / 1000
				result.A += product.A * givenProduct.Mass / 100 / given.Cat.Mass / 1000
				result.Omega3 += product.Omega3 * givenProduct.Mass / 100 / given.Cat.Mass / 1000
				result.Omega6 += product.Omega6 * givenProduct.Mass / 100 / given.Cat.Mass / 1000
			}
		}
	}

	fmt.Println("")
	fmt.Printf("Масса порции: %0.1f г\n", result.Mass)
	fmt.Printf("Масса кошки: %0.1f кг\n", given.Cat.Mass)
	fmt.Println("")
	fmt.Printf("Калорийность: %0.1f ккал\n", result.Kkal)
	fmt.Printf("Белок: %0.1f г/кг\n", result.Protein)
	fmt.Printf("Жир: %0.1f г/кг\n", result.Fat)
	fmt.Printf("Углеводы: %0.2f г/кг\n", result.Carbs)
	fmt.Printf("Влага: %0.1f г/кг (%0.1f%%)\n", result.Water, result.Water*given.Cat.Mass/result.Mass*100)
	fmt.Println("")
	fmt.Printf("Ca: %0.3f г/кг\n", result.Ca)
	fmt.Printf("P: %0.3f г/кг\n", result.P)
	fmt.Printf("Mg: %0.3f г/кг\n", result.Mg)
	fmt.Printf("Na: %0.3f г/кг\n", result.Na)
	fmt.Printf("K: %0.3f г/кг\n", result.K)
	fmt.Printf("Fe: %0.3f г/кг\n", result.Fe)
	fmt.Printf("A: %0.3f г/кг\n", result.A)
	fmt.Printf("Omega3: %0.3f г/кг\n", result.Omega3)
	fmt.Printf("Omega6: %0.3f г/кг\n", result.Omega6)

	return nil
}
