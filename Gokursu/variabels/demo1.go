package variabels


import "fmt"

//camelCase
//PascallCase
func Demo1() {
	var metin string = "Merhaba Turkiye!"
	fmt.Println(metin)

	var kdv float32 = 256.0345
	fmt.Println(kdv)
	fmt.Println(100 + (100 * 10 / 100))

	var yas int = 22
	fmt.Println(yas)

	var durum bool = false

	var metin1 string = "Rahil"
	var metin2 string = "Ibrahimi"

	durum = metin1 == metin2

	fmt.Println(durum)
	fmt.Println(!durum)
	
}
