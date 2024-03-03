package structs

import (
	"fmt"
)

func Struct1() {
	fmt.Println(product{"Bilgisayr", 4900, "XYT",54})
	fmt.Println(product{"Bilgisayr", 4900, "XYT",54})
	fmt.Println(product{name: "pc", unitPrice: 7000})
	
}

type product struct {
	name      string
	unitPrice float64
	brand     string
	discount int
}
