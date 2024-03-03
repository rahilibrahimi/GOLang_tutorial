package fonksiyon

import "fmt"


func Toplama(sayi1 int, sayi2 int) int{
	var toplam = sayi1 + sayi2
	return toplam
}
func Cikarma(sayi1 int, sayi2 int) float64{
	var cikrma = sayi1 - sayi2
	return float64(cikrma)
}
func Bolme(sayi1 float32, sayi2 float32) float32 {
	bole := sayi1/sayi2
	return float32(bole)
}
func Carpma(sayi1 int, sayi2 int) int{
	var carp = sayi1 * sayi2
	return carp

}
func Selamlar(kullaniciAdi string) {
	fmt.Println("Merhaba", kullaniciAdi)
	
}