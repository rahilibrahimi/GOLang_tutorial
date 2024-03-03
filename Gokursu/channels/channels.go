package channels

import (
	"fmt"
)

func CiftSayilar(ciftSayicn chan int) {
	total := 0
	for i := 0; i <= 10; i += 2 {
		total = total + 1
	}
	ciftSayicn <- total
}
func TekSayilar(tekSayilarcn chan int) {
	total := 0
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Tek sayi : ", i)
		total = total + 1
	}

	tekSayilarcn <- total
}
