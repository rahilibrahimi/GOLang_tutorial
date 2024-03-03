package defer_statment

import "fmt"




func Adefer(sayi int)string {
	defer DeferFunc()

	if sayi% 2== 0  {
		return "cift sayidir"
	}
	if sayi%2 != 0 {
		return "Tek sayidir"
	}
	return "Belli degil"
}
func Test()  {
	result := Adefer(9)
	fmt.Println(result)
}
func DeferFunc()  {
	fmt.Println("Bitti")
}

