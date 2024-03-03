package arrays

import "fmt"

func Cities()  {
	var Cities [6]string
	Cities[0] = "Ankara"
	Cities[1] = "Bolu"
	Cities[2] = "Duzce"
	Cities[3] = "Sakarya"
	Cities[4] = "Gebze"
	Cities[5] = "Istanbul"
	fmt.Println(Cities)
	
	for i := 0; i < 5; i++ {
		fmt.Println(Cities[i])
	}
	
}