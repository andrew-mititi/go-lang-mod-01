package main

import (
	"fmt"
	"math"
)

const APP_NAME string = "golang module 1"

type Person struct {
	firstname string
	age int
}

func (p Person) speak() {
 fmt.Printf("I Am %s and am %d years old\n", p.firstname, p.age)
}

// Type shape

type Shape interface {
	area() float64
}
 
type Square struct{
	length, width float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64{
	return (s.length * s.width)
}

func (c Circle) area() float64{
	return math.Pi * math.Pow(c.radius, 2)
}

func calculateInfo(s Shape) float64{
	return s.area()
}


func main(){ 
	// functions
var num_slice = make([]int, 1)
num_slice = append(num_slice, 1,2,3,4,5)

_ = foo(num_slice...)

// defer fmt.Println("Hello world 1")
// defer fmt.Println("Hello world 2")
// fmt.Println(bar())


_ = Person {
	firstname: "Andrew",
	age: 26,
}

_ = Circle{
	radius: 7,
}
_ = Square{
	length: 2,
	width: 3,
}

// fmt.Println(calculateInfo(cirlce))
// fmt.Println(calculateInfo(square))

// fmt.Println(doMath(2, 4, add))

// fmt.Println(doMath(3, 2, subtract))
 _ = returnFunc(1)


 func(f func()){
	fmt.Println("Done Process the data.....")
    f()
 }(sendToQueue)

}



func foo(number ...int) int {
	sum := 0
	for _, num := range number{
      sum += num
	}
	return sum
}

func add(a, b int) int{
	return a + b
}

func returnFunc(num int) func() int{
	num += 2
	return func() int{
		return num
	}
}

func sendToQueue(){
	fmt.Println("Sent data to queue")
}
// func subtract(a, b int) int{
// 	return a - b
// }

// func doMath(a, b int, f func(int, int) int) int {
//   return f(a, b)
// }

