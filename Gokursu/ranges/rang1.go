package ranges

import "fmt"

func Range1()  {
	renkler := [] string {"kirmizi","mavi", "sari", "siyah"}
	for i := 0; i < len(renkler); i++{//for ile yazildi
		fmt.Println(renkler)
	}
	for _, renk := range renkler{
		fmt.Println(renk)
		
	}
}