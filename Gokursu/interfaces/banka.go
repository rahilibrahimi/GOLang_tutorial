package interfaces

import "fmt"

type Mortgage struct {
	creditPayment float64
	address       string
	rate          int
}
type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               int
}
type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPayment * float64(m.rate) / 100 / 12
}
func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * float64(c.rate) / 100 / 12
}

func CalculateMountlyPym(credits []CreditCalculator) float64 {
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}

	return paymentTotal
}

func DemoAdd() {
	credit1 := Mortgage{rate: 10, creditPayment: 1000500.00555, address: "Toyota"}
	car3 := Car{rate: 10, creditPaymentTotal: 100000, carInfo: "Nissan"}

	credits := []CreditCalculator{credit1, car3}
	total := CalculateMountlyPym(credits)

	fmt.Println("Toplam Odem : ", total)
}

// veri yapilari + Algortima Analizi!!!
// json dosya parse 
// json+ verileri dosyada
// ne yazdigimi bilmem sart
// reciever function return
// farklilik
// opencv
//
