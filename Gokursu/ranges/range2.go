package ranges

import "fmt"

// 1-10 arasindaki sayilarin tek olanlari toplayan program
// for-range

func Range2()  {
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	toplam := 0

	for _, sayi := range nums {
		if sayi%2 != 0 {
			toplam = toplam + sayi

		}
	}
	fmt.Println("Toplam : ", toplam)
}