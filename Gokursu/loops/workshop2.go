package loops

import "fmt"

// bravo bildiniz 3 tahmin : super
// 1-3 super, 4-6 guzel ve 6<fena degil
func Demo66() {
	aklimdakiSayi := 90
	tahminEdilenSayi := 0
	tahminSayisi := 0

	fmt.Println("Bir sayi Tahmin ediniz")
	fmt.Scanln(&tahminEdilenSayi)
	tahminSayisi = tahminSayisi + 1

	for aklimdakiSayi != tahminEdilenSayi {
		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("daha buyuk bir sayi giriniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1

		} else if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("daha kucuk bir sayi giriniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}
	}

	basariDurumu := ""
	if tahminSayisi > 0 && tahminSayisi <= 3 {
		basariDurumu = "Super"
	}else if tahminSayisi <= 6 {
		basariDurumu = "Guzel"
	}else{
		basariDurumu = "Fena Degil"
	}

	fmt.Printf("Bravo dogru tahmin edebildiniz. %v tahmin : %v", tahminSayisi, basariDurumu)
}
