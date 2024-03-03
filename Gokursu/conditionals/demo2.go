package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 9000

	if cekilmekIstenen > hesap {
		fmt.Println("hesptaki para yetersiz")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Paraniz hazirlaniyor")
		fmt.Println("Dikkat!, hesapta para kalmadi")
		// hesap = hesap - cekilmekIstenen
	}  else {
		fmt.Println("Bitti. Hesaptaki para: " +fmt.Sprintf("%v", hesap))
	} 
	
}