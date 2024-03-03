package ranges

import "fmt"


func Range3()  {
	sozluk := map[string]string {"book: " : "kitap ", "pen: " : "kalem "}

	for k, v := range sozluk {
		fmt.Print(k) // burada string deger donecegi icin print 
		fmt.Println(v) //  bir alt satira yazdiramak icin kullaniyoruz
	}
}
