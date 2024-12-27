package main

import (
	"fmt"
	"golang-dasar/helper"
	"golang-dasar/database"
	_ "golang-dasar/internal"
	"errors"
)

type Customer struct {
	Name, Address string
	Age int
}

func (co Customer) sayHello() {
	fmt.Println("hello ," , co.Name)
}

type apalah interface {}
func (co Customer) GetName() string {
	return co.Name
}

func (c *Customer) ChangeName(name string) {
	c.Name = name
}

type HasName interface {
	GetName() string
}

func SayHello(hasname HasName) {
	fmt.Println("hello", hasname.GetName())
}

type Address struct {
	City, Provice string
}
func main() {
	
	fmt.Println("test")
	// func for string
	fmt.Println("panjang string adalah =", len("string"))
	fmt.Println("karakter pertama adalah =", string("raihan"[0]))

	//* variable
	var name string = "raihan" // cara ini kurang rekomen

	// bisa juga di declare langsung tanpa tipe data
	var umur = 17
	fmt.Println("nama saya =", name, "umur =", umur)

	// bisa juga declare tanpa 'var'
	alamat := "pratama bintaro"
	fmt.Println(alamat)

	// contoh multi declare
	// semua var di golang harus digunakan
	var (
		firstName apalah
		lastName  string
	)

	firstName = "raihan"
	lastName = "raihan"

	fmt.Println(firstName, lastName)

	//* const (var yg tidak bisa diubah)

	const rumah string = "ciputat" // const harus langsung di initialisasi
	fmt.Println(rumah)

	//* const multi
	const (
		laptop string = "asus"
		spek   string = "ryzen"
	)

	fmt.Println(laptop, spek)

	//* konversi
	var angka32 int32 = 32
	var angka64 int64 = int64(angka32)

	fmt.Println(angka64)

	var randString = "apakek"
	var e uint8 = randString[3] //? tidak akan langsung keluar 'k'

	var eHuruf = string(e) //todo: kudu konversi dlu

	fmt.Println(eHuruf)

	//*  Type declaration

	type NoKtp string

	var noKtp NoKtp = "1112324242"
	fmt.Println(noKtp)
	fmt.Println(NoKtp("222222"))

	//* operasi matematika
	var (
		a = 1
		b = 2
		c = a + b
	)
	fmt.Println(c)

	//* augmented expression
	a += 4
	fmt.Println(a)
	// unary
	b--
	fmt.Println(b)
	b++
	fmt.Println(b)
	b = -b
	fmt.Println(b)
	b = +b
	fmt.Println(b)

	//* perbandingan

	var perbandingan bool = "eko" == "true"
	fmt.Println(perbandingan)

	//* operasi boolean
	var oprasi bool = !(true || false)
	fmt.Println(oprasi)

	var names [3]string
	names[0] = "raihan"
	names[1] = "yusuf"
	names[2] = "firmansyah"
	fmt.Println(names)
	fmt.Println(names[0])

	var values = [3]int{
		90,
		54,
		65,
	}
	fmt.Println(values)
	fmt.Println(len(values))

	//* slice
	var newName = [...]string{"raihan", "yusuf", "firmansyah"}
	var newSlice []string = newName[:]
	fmt.Println(newSlice)

	namaBaru := append(newSlice, "nama belakang")
	fmt.Println(namaBaru)

	var makeSlice []string = make([]string, 2, 5)
	makeSlice[0] = "mainan"
	makeSlice[1] = "baru"
	fmt.Println(makeSlice)

	//* cp slice
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat"}
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	//! hati-hati ketuker
	iniArray := [...]int{1, 2, 3, 4}
	iniSlice := []int{1, 2, 3, 4}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	//* Tipe data Map
	person := map[string]string{
		"nama":  "ucup",
		"kelas": "12",
	}

	fmt.Println(person)

	//* if statement

	kondisi := "baik"

	if kondisi == "baik" {
		fmt.Println("saya baik")
	} else if kondisi == "buruk" {
		fmt.Println("kabar saya buruk")
	} else {
		fmt.Println("biasa biasa saja")
	}

	//* short if
	if length := len(kondisi); length < 7 {
		fmt.Println("string pendek")
	}

	//* switch

	konSwitch := "apel"

	switch konSwitch {
	case "apel":
		fmt.Println("ini buah apel")
	case "jeruk":
		fmt.Println("ini buah jeruk")
	default:
		fmt.Println("tidak tau buah apa")
	}

	//* switch tanpa kondisi

	lenBuah := len(konSwitch)

	switch {
	case lenBuah > 3:
		fmt.Println("nama buah panjang")
	default:
		fmt.Println("nama buah normal")
	}

	//* for loop
	for count := 0; count <= 5; count++ {
		fmt.Println("perulangan ke", count)
	}

	//* for range
	karyawan := [...]string{"pabos", "dedi", "budi", "auah"}

	for index, item := range karyawan {
		fmt.Println(index, "=>", item)
	}

	//* break & continue

	for i := 0; i <= 8; i++ {
		if i == 3 {
			fmt.Println("perulangan di skip")
			continue //? continue akan mengskip perulangan ini dan lanjut ke perulangan berikutnya
		}
		if i == 6 {
			fmt.Println("perulangan berhenti")
			break //? break akan menghentikan perulagan secara total
		}
		fmt.Println("angka => ", i)
	}

	//* function
	sayHello("raihan")
	hasil := returnValue(1, 2)
	fmt.Println(hasil)

	depan, belakang := multiReturn()
	fmt.Println(depan, belakang)

	// mengabaikan return value
	depanAja, _ := multiReturn()
	fmt.Println(depanAja)

	apalah := namedReturn()
	fmt.Println(apalah)

	// varargs
	varargs := sumAll(10, 20, 5, 15)
	fmt.Println(varargs)

	slcParams := []int{5, 15, 5}
	slcResult := sumAll(slcParams...) //? bisa juga menggunakan slice sebagai parameter
	fmt.Println(slcResult)

	// func as parameter
	funcAsParams("anjing", filter)

	// anonymous func

	blackList := func(name string) bool {
		larangan := []string{"anjing", "babi", "monyet"}
		var isTerlarang bool

		for _, i := range larangan {
			if i == name {
				isTerlarang = true
				break
			} else {
				isTerlarang = false
			}
		}

		return isTerlarang
	}

	registerUser("raihan", blackList)
	registerUser("anjing", func(name string) bool {
		larangan := []string{"anjing", "babi", "monyet"}
		var isTerlarang bool

		for _, i := range larangan {
			if i == name {
				isTerlarang = true
				break
			} else {
				isTerlarang = false
			}
		}

		return isTerlarang
	})

	// recursive

	fmt.Println(factorialRecursive(3))

	// closure
	counter := 0
	increment := func() {
		counter++
	}

	increment()
	increment()

	// defer, panic, recover
	startApp(true)

	// struct

	budi := Customer{
		Name: "budi",
		Address: "sfafasa",
		Age: 12,
	}

	budi.sayHello()

	// interface
	SayHello(budi)

	// type assertion
	randomV := random()
	randomString := randomV.(string)
	fmt.Println(randomString)

	switch randomV.(type) {
	case string:
		fmt.Println("string")
	case int :
		fmt.Println("Int")
	default:
		fmt.Println("Unknown")
	}

	// pointer
	address1 := Address{"bogor", "jabar"}
	address2 := address1

	address2.City = "jakarta"
	
	fmt.Println(address1.City)
	fmt.Println(address2.City)

	var address3 *Address = &address1

	address3.City = "jogja"

	fmt.Println(address1.City)
	fmt.Println(address3.City)

	// pointer struct
	raihan := Customer{
	 Name: "raihan",
	 Address: "jakarta",
	 Age: 17,
	}

	fmt.Println(raihan.Name)
	raihan.ChangeName("yusuf")
	fmt.Println(raihan.Name)

	firman := &raihan

	*firman = Customer{
		Name: "firman",
		Address: "bogor",
		Age: 14,
	}

	fmt.Println(raihan.Name)
	fmt.Println(firman.Name)

	// package
	fmt.Println(helper.SayHelper())

	// init
	fmt.Println("current database => ", database.GetDatabase())

	// error

	hasil, err := pembagian(8,0)

	if err == nil {
		fmt.Println("Hasil :", hasil)
	} else {
		fmt.Println("Error :", err.Error())
	}
}

func sayHello(name string) {
	fmt.Println("hello", name)
}

func returnValue(a int, b int) int {
	return a + b
}

func multiReturn() (string, string) {
	return "raihan", "yusuf"
}

func namedReturn() (kodingan string) { // kita juga bisa memberikan nama untuk return valuenya
	kodingan = "koding asik"
	return kodingan
}

func sumAll(number ...int) int {
	total := 0
	for _, num := range number {
		total += num
	}
	return total
}

type Filter func(string) string

func funcAsParams(name string, filter Filter) {
	fmt.Println("name", filter(name))
}

func filter(name string) string {
	switch name {
	case "anjing":
		return "..."
	default:
		return name
	}
}

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("Youre blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func factorialRecursive(i int) int {
	if i == 1 {
		return 1
	} else {
		return i * factorialRecursive(i-1)
	}

}

func endApp() {
	fmt.Println("selesai menjalankan aplikasi")

	message := recover()
	fmt.Println("terjadi panic : ", message)
}

func startApp(error bool) {
	defer endApp()

	fmt.Println("aplikasi start")
	if error {
		panic("ERROR 404")
	}
}

func random() interface{} {
	return "ok"
}

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("tidak bisa dibagi nol")
	} else {
		return nilai / pembagi, nil
	}
}