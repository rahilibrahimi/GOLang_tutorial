package main

import (
	"deneme/arrays"
	"deneme/channels"
	"deneme/conditionals"
	"deneme/fonksiyon"
	"deneme/goroutines"
	"deneme/loops"
	"deneme/maps"
	"deneme/pointer"
	"deneme/ranges"
	"deneme/slice"
	"deneme/structs"
	"deneme/variabels"

	//"deneme/interfaces"
	"deneme/defer_statment"
	"deneme/error_handling"
	"fmt"
	"time"
)

func main() {
	variabels.Demo1()
	fmt.Println()
	conditionals.Demo1()
	fmt.Println()
	conditionals.Demo2()
	fmt.Println()
	conditionals.Demo3()
	fmt.Println()
	loops.Demo4()
	fmt.Println()
	loops.Demo5()
	fmt.Println()
	// loops.Demo6()
	// fmt.Println()
	// loops.Demo66()
	// fmt.Println()
	// loops.Demo43()
	// fmt.Println()
	loops.Demo11()
	arrays.Demoarray()
	arrays.Cities()
	arrays.Cities1()
	arrays.CokBuyulu()
	slice.Slice1()
	slice.Slice2()
	var result = fonksiyon.Bolme(19, 13)
	fmt.Println(result)
	fonksiyon.Selamlar("Rahil")
	fonksiyon.Selamlar("Ibrahimi")

	sonuc1, sonuc2, sonuc3, _ := fonksiyon.DortIslem(10, 6)

	fmt.Println("Toplam : ", sonuc1)
	fmt.Println("Cikarma : ", sonuc2)
	fmt.Println("Carpma : ", sonuc3)
	//fmt.Println("Bolme : ", sonuc4)
	fonksiyon.Vardiacfon()
	fmt.Println()

	maps.Map1()
	ranges.Range1()
	ranges.Range2()
	ranges.Range3()

	sayi := 20
	pointer.Pointer1(&sayi) //burada &sayi bellekteki adressini gosteriyor
	fmt.Println("Maindeki sayi : ", sayi)

	sayilar := []int{1, 2, 3}
	pointer.Demo2(sayilar)
	fmt.Println("Maindeki sayi : ", sayilar[0])

	structs.Struct1()
	structs.Struct2()

	go goroutines.CiftSayilar()
	go goroutines.TekSayilar()

	time.Sleep(3 * time.Second)
	fmt.Println("Main Bitti")

	ciftSayiCn := make(chan int) //channel make fonksiyonu ile yazilir
	tekSayiCn := make(chan int)  //Ayni make fonksiyonu ile yazilir

	go channels.CiftSayilar(ciftSayiCn) /// her ne kadar acynchoron calisiyo olsa da biz burda bir
	// channel mantigi kullanmamiz ve ciftsayilarin(ciftSayic chan in ) olarak tanimlarim
	go channels.TekSayilar(tekSayiCn)

	ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, tekSayiCn
	carpim := ciftSayiToplam * <-tekSayiToplam

	// carpim := ciftSayiCn * tekSayiCn boyle bir islem yapamazsiniz
	fmt.Println("Carpim : ", carpim)

	// interfaces.DemoShape()
	// interfaces.CalculateMountlyPym()

	defer_statment.DeferA()
	defer_statment.DeferFunc()
	defer_statment.Detail1()
	error_handling.Error1()

}
