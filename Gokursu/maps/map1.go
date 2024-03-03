package maps

import "fmt"



func Map1()  {
	sozluk := make(map[string]string)

	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pen"] = "Kalem"
	sozluk["pencil"] = "Kursun Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("eleman asayisi: ", len(sozluk))
	delete(sozluk, "pencil")
	fmt.Println("eleman sayisi:", len(sozluk))
	
	deger := sozluk["pen"]

	fmt.Println(deger)

	deger, varMi := sozluk["lamb"]
	fmt.Println(deger)
	fmt.Println("listede olma durumu: ", varMi)

	sozluk2 := map[string]string{"glass":"bardak", "library":"kutuphane"}
	fmt.Println(sozluk2)
}