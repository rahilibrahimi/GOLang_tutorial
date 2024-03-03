package error_handling

import (
	"fmt"
	"os"
)

func Error1() {
	f, err := os.Open("error1.txt")
	//nil
	if err != nil {
		fmt.Println("Dosya bulunamadi")
	} else {
		fmt.Println(f.Name())
	}
}
