package loops

import "fmt"

func Demo4()  {
	var text string = "Merhaba Merhaba"
	// fmt.Println(text)
	// fmt.Println(text)
	// fmt.Println(text)
	// fmt.Println(text)
	// fmt.Println(text)
	// fmt.Println(text)

	for i := 1; i <= 5; i++ {
		fmt.Println(text)
	}
	fmt.Println("Done!")

}