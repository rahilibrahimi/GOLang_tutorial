package arrays

import "fmt"

func Cities1()  {
	cities := [6]int{1,23,4,5,6,7}
	fmt.Println(cities)
	enBuyuk := cities[0]

	for i := 0; i<len(cities); i++ {
		if cities[i]>enBuyuk{		
			enBuyuk = cities[i]
		}
	} 

	fmt.Println("En buyuk", enBuyuk)

}