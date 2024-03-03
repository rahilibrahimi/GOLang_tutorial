package conditionals

import "fmt"




func Demo3() {
	var a, b, c  int = 340, 50, 90
	
	var enBuyuk int = a

	if b > enBuyuk {
		enBuyuk = b
	} 
	if c > enBuyuk {
		enBuyuk = c
	}
	fmt.Printf(" en buyuk degisken: %v\n", enBuyuk)
}