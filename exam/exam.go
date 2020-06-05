package main

import (
	"fmt"
	"libhj/testlib"
	"log"
	"math"
	"os"
)

var a = 11

var f float32 = 11.
var i, j, k int = 1, 2, 3

const c int = 10 // define

// Apple=0, Grape=1, Orange=2, ...
const (
	Apple = iota
	Grape
	Orange
)

const s string = "Hi"

// string
const (
	Visa   = "Visa"
	Master = "MasterCard"
	Amex   = "American Express"
)

// Rect struct for test
type Rect struct {
	width, height int
}

//Shape interface
type Shape interface {
	areaInterface() float64
	perimeterInterface() float64
}

// RectIntertest for shape interface
type RectIntertest struct {
	width, height float64
}

// CircleIntertest for shape interface
type CircleIntertest struct {
	radius float64
}

func main() {

	// 	rawLiteral := `아리랑\n아리랑\n아라리요`
	// 	interLiteral := "아리랑 아리랑\n아라리요"

	// 	say(rawLiteral)
	// 	say(interLiteral)

	// 	sayHello := "hello"
	// 	say2(&sayHello)
	// 	say2(&sayHello)

	// 	say3("this", "is", "a", "book")
	// 	say3("hi")

	// 	fmt.Printf("%d %d %d %d %d %d %d %d %d %s %s %s %s", a, int(f), i, j, k, c, Apple, Grape, Orange, s, Visa, Master, Amex)

	// 	if a == 1 {
	// 		fmt.Println("a is 1")
	// 	} else if a == 2 {
	// 		fmt.Println("a is 2")
	// 	} else {
	// 		fmt.Println("a is not 1 or 2")
	// 	}

	// 	if val := i * 2; i < k {
	// 		fmt.Println(val)
	// 	}

	// 	switch val := 1; val {
	// 	case 1:
	// 		fmt.Println("val is 1")
	// 	case 2:
	// 		fmt.Println("val is 2")
	// 	case 3:
	// 		fmt.Println("val is 3")
	// 	default:
	// 		fmt.Println("val is not 123")
	// 	}

	// 	switch val2 := 0; {
	// 	case val2 == 0:
	// 		fmt.Println("val is 0")
	// 	case val2 == 1:
	// 		fmt.Println("val is 1")
	// 	case val2 == 2:
	// 		fmt.Println("val is 2")
	// 	}

	// 	var i interface{} //???? what is interface{}
	// 	switch i.(type) {
	// 	case int:
	// 		println("int")
	// 	case bool:
	// 		println("bool")
	// 	case string:
	// 		println("string")
	// 	default:
	// 		println("unknown")
	// 	}

	// 	sum := 0
	// 	for i := 1; i <= 100; i++ {
	// 		sum += i
	// 	}
	// 	println(sum)

	// 	ex1 := 1
	// 	for ex1 < 100 {
	// 		ex1 *= 2
	// 	}
	// 	println(ex1)

	// L1:
	// 	for {
	// 		if sum == 5050 {
	// 			break L1
	// 		}
	// 	}
	// 	println("break OK")

	// 	names := []string{"홍길동", "이순신", "강감찬"}
	// 	for index, name := range names {
	// 		println(index, name)
	// 	}

	// 	s := sumi(1, 2, 3, 4, 15, 346, 234, 395)
	// 	println(s)

	// 	cnt, s2 := sumo(1, 2, 3, 4, 15, 346, 234, 395)
	// 	println(cnt, s2)

	// 	array()

	//	slice()

	//	mapTest()

	// methodTest()

	errTest()
}

/****** function ******/
func say(msg string) {
	println(msg)
}

func say2(msg *string) {
	println(*msg)
	*msg = "changed"
}

func say3(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

func sumi(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func sumo(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func array() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3

	var a1 = [3]int{1, 2, 3}
	var a2 = [...]int{1, 2, 3}
	var a3 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	print(a[0], a[1], a[2], a1[0], a1[1], a1[2], a2[0], a2[1], a2[2], a3[0][2], a3[1][2], a3[2][2], "\n")

}

func slice() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10

	s := make([]int, 5, 10)
	var n []int
	if n == nil {
		println("Nil Slice")
	} else {
		println("Not Nil Slice")
	}

	subS := []int{0, 1, 2, 3, 4, 5}

	appendS := []int{0, 1}
	appendS = append(appendS, 2)

	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}
	sliceA = append(sliceA, sliceB...)

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)

	print(a[0], a[1], a[2], "\n")
	println(len(s), cap(s), len(n), cap(n))
	fmt.Println(subS)
	fmt.Println(appendS)
	fmt.Println("sliceA + sliceB = ", sliceA)
	fmt.Println(target)
	fmt.Println(len(target), cap(target))

	target = append(target, 3)
	fmt.Println(target[3])
}

func mapTest() {
	var idMap map[int]string
	idMap = make(map[int]string)

	idMap[901] = "Apple"
	idMap[134] = "Grape"
	idMap[777] = "Tomato"

	str := idMap[134]
	fmt.Println(str)

	noData := idMap[999]
	fmt.Println(noData)

	delete(idMap, 777)
	fmt.Println(idMap[777])

	tickers := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Println(tickers[1])
	val, exists := tickers[4]
	fmt.Println(val, exists)
	if !exists {
		fmt.Println("No ticker")
	}

	for key, val := range idMap {
		fmt.Println(key, val)
	}

	song := testlib.GetMusic("Adele")
	fmt.Println(song)
}

func (r Rect) area() int {
	r.width++
	return r.width * r.height
}

func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func (r RectIntertest) areaInterface() float64 {
	return r.width * r.height
}

func (r RectIntertest) perimeterInterface() float64 {
	return 2 * (r.width + r.height)
}

func (c CircleIntertest) areaInterface() float64 {
	return math.Pi * c.radius * c.radius
}

func (c CircleIntertest) perimeterInterface() float64 {
	return 2 * math.Pi * c.radius
}

func methodTest() {

	rect := Rect{10, 20}
	area := rect.area()
	fmt.Println(area)
	area2 := rect.area2()
	fmt.Println(area2)

	//r := RectIntertest{10., 20.}
	c := CircleIntertest{10}

	showArea(c)

	var x interface{}
	x = 1
	fmt.Println(x)

	x = "TOM"
	fmt.Println(x)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.areaInterface()
		fmt.Println(a)
	}
}

func errTest() {
	f, err := os.Open("/home/hjyoun/tmp/test")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(f.Name())
}
