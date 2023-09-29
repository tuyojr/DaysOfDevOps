package loop

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(`Current number:`, i)
	}

	// the for loop also supports the while loop syntax
	j := 1
	for j <= 5 {
		fmt.Println(`Current number:`, j)
		j++
	}

	// you can also use the for loop to iterate through an array
	numbers := []int{1, 2, 3, 4, 5}

	for index, value := range numbers {
		fmt.Println(`Current index:`, index, `--<>--`, `Current value`, value)
	}

	// you can also use the for loop to iterate through a map
	uchihas := make(map[string]string)
	uchihas[`Sasuke`] = `RinneSharingan`
	uchihas[`Itachi`] = `Tsukuyomi`
	uchihas[`Shisui`] = `Kotoamatsukami`
	uchihas[`Madara`] = `Eternal Mangekyo Sharingan`

	for key, value := range uchihas {
		fmt.Println(`Name:`, key, `--<>--`, `PowerUp:`, value)
	}
}
