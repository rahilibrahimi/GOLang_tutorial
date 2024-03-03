package slice

import (

	"fmt"
)

func Slice2()  {
	sehirler := []string {"Ankara, Istanbul, Izmir"}
	fmt.Println(sehirler)
	sehirlerKopya := make([]string, len(sehirler))
	fmt.Println(sehirlerKopya)
	copy(sehirlerKopya, sehirler)
	sehirler = append(sehirler, "Bolu")
	fmt.Println(sehirler)
	fmt.Println(sehirlerKopya)

	fmt.Println(sehirler [1:2])
	fmt.Println(sehirler [1:])
	fmt.Println(sehirler [:2])

	

}