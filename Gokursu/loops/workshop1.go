package loops
import "fmt"

// / oyun tahmin etme oyunu
func Demo6() {
	aklimdakiSayi := 80
	tahminEdilenSayi := 0

	
	fmt.Println("Bir sayi tahmin ediniz")
	fmt.Scanln(&tahminEdilenSayi)

	for tahminEdilenSayi != aklimdakiSayi {
		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Lutfen daha kucuk bir sayiyi giriniz")
			fmt.Scanln(&tahminEdilenSayi)

		}else if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Lutfen daha buyuk bir sayi giriniz")
			fmt.Scanln(&tahminEdilenSayi)

		}
	}
	fmt.Println("Bravo Bildiniz!")

}
