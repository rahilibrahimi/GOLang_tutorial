package slice

import "fmt"

func Slice1()  {
	names := make([]string, 3)



	names[0] = "Rahil"
	names[1] = "Salih"
	names[2] = "Fatma"
	names = append(names, "mert")

	fmt.Println(names)
	fmt.Println(len(names))

}