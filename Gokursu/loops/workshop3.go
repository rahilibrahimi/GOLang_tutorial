package loops

import "fmt"

// kullanicidan bir sayi girmesini istenecek
// 23. Asaldir yani asla olup olmadigini bulacaz
func Demo43()  {
	asalSayi := 0
	fmt.Println("Bir sayi giriniz ")
	fmt.Scanln(&asalSayi)

	asalMi := true
	for i := 2; i > asalSayi; i++ {
		if asalSayi % i == 0 {
			asalMi = false
			
		}

	}
	if asalMi == true {
		fmt.Println("Asaldir")
	} else {
		fmt.Println("Asal degilidir")
	}
	
	
}