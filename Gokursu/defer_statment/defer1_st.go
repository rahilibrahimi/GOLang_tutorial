package defer_statment

import "fmt"

func DeferA() {
	fmt.Println("A fonksiyonu calisti ")
}
func DeferB() {
	DeferC()
	DeferA()
	fmt.Println("B fonksiyonu calisti ")
	
}
func DeferC() {
	fmt.Println("C fonksiyonu calisti ")
}
