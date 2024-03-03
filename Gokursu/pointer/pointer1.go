package pointer

import "fmt"


func Pointer1(sayi *int)  {
	*sayi = *sayi +1 
	fmt.Println("Pointerdeki sayi : ", sayi)
}
func Demo2(sayilar []int) {
	sayilar[0] = 100
	fmt.Println("Demodaki sayi : ", sayilar[0])

	
}