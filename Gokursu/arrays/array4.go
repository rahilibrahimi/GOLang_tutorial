package arrays

import "fmt"


func CokBuyulu()  {
	var numbers [2][3]int
	numbers[0][0] = 1
	numbers[0][1] = 3
	numbers[0][2] = 5
	numbers[1][0] = 2
	numbers[1][1] = 4
	numbers[1][2] = 6
	

	for satir := 0; satir < 2; satir++ { 
		for sutun := 0; sutun < 3; sutun++ {
			fmt.Print(numbers[satir][sutun])
		} 
		fmt.Println(" ")// ic ice dongunun icindeyken yan yana yazmamasi icin

	}
	


	//fmt.Println(numbers[1][2])
}