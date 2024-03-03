package defer_statment

import "fmt"

type product struct {
	productName string
	unitedPrice int
}

func (p product) save() {
	fmt.Println("Urun kayt edildi : ", p.productName)
	defer Log()
}
func Log() {
	fmt.Println("loglandi")
}
func Detail1() {
	p := product{productName: "loptop", unitedPrice: 50000}
	defer p.save() //defer geciktiriyor ve en  son fonksiyonu calistiriyor.
	p = product{productName: "mouse", unitedPrice: 45}
	p = product{}
	fmt.Println("islem basirili")
	fmt.Println(p.productName)

	
}
