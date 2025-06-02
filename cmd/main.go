package main

// import (
// 	"fmt"
// 	// "strings"
// 	// "unicode/utf8"
// 	// "errors"
// 	// "math/rand"
// 	"time"
// 	"sync"
// )

// func main() {
// fmt.Println("Hello, World!")
// var intNum int
// intNum = 10
// fmt.Println(intNum)

// var floatNum float32 = 10.1
// fmt.Println(floatNum)

// this will not work because int and float are different types
// var result = intNum + floatNum
// fmt.Println(result)

// need to convert int to float only then we can add them or float to int
// var result = float32(intNum) + floatNum
// fmt.Println(result)

// if want to store a single line string use double quotes
// var myString string = "Hello, World!"
// fmt.Println(myString)
// fmt.Println(len(myString))

// for mulitple line string use backticks
// var myString2 string = `
// Hello, World!
// Hello, World!
// Hello, World!
// `
// fmt.Println(myString2)

// if you want to find length of string use len()
// fmt.Println(len(myString))

// if you want to find character at index use [index]
// fmt.Println(myString[0])

// if character contains outside vanillla quotes use double quotes
// var myString2 string = "Hello, \"World\"!"
// fmt.Println(myString2)
// fmt.Println(len(myString2))

// if you want to find length of string use len()
// fmt.Println(len("γ")) //shows 2 because it is a unicode character
// // so if you want to count the characters in string use utf8.RuneCountInString()
// fmt.Println(utf8.RuneCountInString("γ")) //shows 1 because it is a unicode character

// var runeValue rune = 'γ'
// fmt.Println(runeValue) // shows 947
// fmt.Println(string(runeValue)) // shows γ
// fmt.Println(utf8.RuneCountInString(string(runeValue)))

// var myBool bool = true
// fmt.Println(myBool)

// we can define variable and assign value later
// var myInt int //as int, will assign it initial value 0 (its totally depends on the type)
// myInt = 10
// fmt.Println(myInt)

// for bool its false
// for string its ''
// for uint, int, float and runes its 0

// we can drop the type and var and use short declaration operator
// myInt := 10
// fmt.Println(myInt)

// we can also declare multiple variables at once
// myInt, myFloat := 10, 10.1
// fmt.Println(myInt, myFloat)

// const can be used to define constant values
// const myConst int = 10
// fmt.Println(myConst)

// we can also declare multiple constants at once
// const (
// 	myConst1 int = 10
// 	myConst2 int = 20
// )
// fmt.Println(myConst1, myConst2)

// iota is a special constant that is used to assign values to constants in a sequence
// const (
// 	myConst1 int = iota
// 	myConst2 int = iota
// )
// fmt.Println(myConst1, myConst2)

// functions and control structures

// if else
// if else if else
// switch
// for loop
// for loop with range
// for loop with condition
// for loop with break and continue

// var name string = "John" // type is string is enforced by go
// printMe(name)

// var int1, int2 int = 20, 10
// // var int1, int2 int = 20, 0

// // fmt.Println(intDiv(int1, int2))
// result, remainder, err := intDiv(int1, int2)
// if err != nil {
// 	fmt.Println(err.Error())
// } else {
// 	fmt.Println(result, remainder)
// }

// var int1, int2 float64 = 10, 20
// fmt.Println(intDiv(int1, int2))

// arrays

// properties of arrays
// 1. fixed size
// 2. homogeneous (same type)
// 3. contiguous memory allocation
// 4. random access
// 5. fixed size

// var myArray [3]int
// myArray[0] = 1
// myArray[1] = 2
// myArray[2] = 3
// // myArray[3] = 4 // this will throw an error because the array is of size 3
// fmt.Println(myArray)

// we can also declare an array with values
// var myArray2 [3]int = [3]int{1, 2, 3}
// fmt.Println(myArray2)

// slice
// intArray := [3]int{1, 2, 3}
// var intArray [3]int = [3]int{1, 2, 3}
// var intSlice []int = intArray[0:2] // this is a slice of the array
// fmt.Println(intSlice)

// var intSlice []int = []int{1, 2, 3} // explicitly declaring a slice
// fmt.Println(intSlice)
// intSlice = append(intSlice, 4) // appending a new element to the slice
// fmt.Println(intSlice)

// difference between array and slice
// 1. array is a fixed size and slice is a dynamic size
// 2. array is a contiguous memory allocation and slice is a contiguous memory allocation
// 3. array is a fixed size and slice is a dynamic size

// diff between cap and len
// cap is the capacity of the slice
// len is the length of the slice

// var intSlice []int = []int{1, 2, 3}
// fmt.Println(len(intSlice))
// fmt.Println(cap(intSlice))
// intSlice = append(intSlice, 4)
// fmt.Println(len(intSlice))
// fmt.Println(cap(intSlice))	// cap is the capacity of the slice

// appending slices
// var intSlice []int = []int{1, 2, 3}
// var intSlice2 []int = []int{4, 5, 6}
// intSlice = append(intSlice, intSlice2...)
// fmt.Println(intSlice)

// copy slices
// var intSlice []int = []int{1, 2, 3}
// var intSlice2 []int = []int{4, 5, 6}
// copy(intSlice2, intSlice)
// fmt.Println(intSlice2)

// difference between copy and append
// copy is used to copy the elements of one slice to another slice
// append is used to append the elements of one slice to another slice

// make

// can use make to create a slice
// if slice is created without cap then it will be equal to len
// var intSlice []int = make([]int, 3)
// fmt.Println(intSlice)

// can use make to create a slice with a specific length and capacity
// var intSlice []int = make([]int, 3, 5)
// fmt.Println(intSlice)

// maps

// maps are key value pairs
// maps are unordered
// maps are dynamic
// maps are not indexed
// maps are not contiguous

// var myMap map[string]uint8 = make(map[string]uint8)
// myMap["John"] = 1
// myMap["Jane"] = 2
// myMap["Jim"] = 3
// fmt.Println(myMap)

// if map doesnot contain the key then it will return the default value of the type
// var age, ok = myMap["John"]
// if ok {
// 	fmt.Printf("the age is %v \n", age)
// } else {
// 	fmt.Println("Invalid key")
// }

// map returns a boolean value if the key is present in the map
// if key is present then it will return true and the value of the key
// if key is not present then it will return false and the default value of the type
// fmt.Println(myMap["John"])

// delete
// delete(myMap, "John")
// fmt.Println(myMap)

// iterate over map
// for key, value := range myMap {
// 	fmt.Printf("the key is %v and the value is %v \n", key, value)
// }

// rune
// rune is a character in go
// rune is a single character
// rune is a unicode character

// var str = []rune("résumé")
// // fmt.Println(len(str))
// for i,v :=range str{
// 	fmt.Println(i,v)
// }

// string

// var str = []string{"H","e","l","l","o"}
// var result string = ""
// for _,v :=range str{
// 	result += v
// }
// fmt.Println(result)

// its creating new string everytime we concat instead we can use strings.Builder
// var str = []string{"H","e","l","l","o"}
// var strBuilder strings.Builder
// for i := range str{
// 	strBuilder.WriteString(str[i])
// }
// var result string = strBuilder.String()
// fmt.Println(result)

// can not concat number and string
// var num int = 10
// var str string = "Hello"
// var result string = num + str // this will throw an error
// fmt.Println(result)

// }

// func printMe(name string) { // name is a parameter and string is the type of the parameter
// 	fmt.Println("Hello, " + name + "!")
// }

// func intDiv(int1 int, int2 int) (int, int, error) {
// 	var err error
// 	if int2 == 0 {
// 		err = errors.New("division by zero")
// 		return 0, 0, err
// 	}
// 	var result int = int1 / int2
// 	var remainder int = int1 % int2
// 	return result, remainder, nil
// }

// func intDiv(num1 float64, num2 float64) float64 {
// 	return num1 / num2
// }

// struct

// type Designation struct{
// 	department string
// 	role Role
// 	salary float64
// }

// type Role string

// const (
// 	DEVELOPER Role = "developer"
// 	DESIGNER Role = "designer"
// 	MANAGER Role = "manager"
// )

// type Person struct {
// 	name string
// 	age int
// 	designation Designation
// }

// func (p Person) YearlySalary() float64 {
// 	return p.designation.salary * 12
// }

// func main() {
// 	// var person Person = Person{name: "John", age: 20}
// 	// fmt.Printf("the name is %v and the age is %v \n", person.name, person.age)

// 	// var person1 Person = Person{"John", 20}
// 	// fmt.Printf("the name is %v and the age is %v \n", person1.name, person1.age)

// 	var person Person = Person{
// 		name: "John",
// 		age: 20,
// 		designation: Designation{
// 			department: "IT",
// 			role: DEVELOPER,
// 			salary: 100000,
// 		},
// 	}

// 	fmt.Println(person)
// 	fmt.Println(person.YearlySalary())

// }

// // interface
// type SalaryProvider interface{
// 	YearlySalary() float64
// 	GetName() string
// }

// type FullTimeEmployee struct{
// 	Name string
// 	salary float64
// }

// func (f FullTimeEmployee) YearlySalary() float64{
// 	return f.salary * 12
// }

// func (f FullTimeEmployee) GetName() string{
// 	return f.Name
// }

// type PartTimeEmployee struct{
// 	Name string
// 	HourlyRate float64
// 	HoursPerDay int
// 	DaysPerYear int
// }

// func (e PartTimeEmployee) YearlySalary() float64 {
// 	return e.HourlyRate * float64(e.HoursPerDay) * float64(e.DaysPerYear)
// }

// func (e PartTimeEmployee) GetName() string{
// 	return e.Name
// }

// func main(){
// 	employee := []SalaryProvider{
// 		FullTimeEmployee{
// 			Name: "John",
// 			salary: 100000,
// 		},
// 		PartTimeEmployee{
// 			Name: "Jane",
// 			HourlyRate: 100,
// 			HoursPerDay: 8,
// 			DaysPerYear: 250,
// 		},
// 	}

// 	for _,v := range employee{
// 		fmt.Println(v.GetName(),v.YearlySalary())
// 	}

// }

// pointer

// func main(){
// 	var p *int32
// 	var i int32 = 10
// 	p = &i
// 	fmt.Println(p)
// 	fmt.Println(*p)

//difference between var p *int32 and var p *int32=new(int32)
// var p *int32 = new(int32)
// *p = 10
// fmt.Println(p)
// fmt.Println(*p)
// *p=12
// fmt.Println(*p,p)

//copying issue with pointers

// var intSlice []int = []int{1,2,3,4,5}
// var intSlice2 []int = intSlice
// intSlice2[0]=100
// fmt.Println(intSlice)
// fmt.Println(intSlice2)

//copying using using make

// var intSlice []int = []int{1,2,3,4,5}
// var intSlice2 []int = make([]int,len(intSlice))
// copy(intSlice2,intSlice)
// intSlice2[0]=100
// fmt.Println(intSlice)
// fmt.Println(intSlice2)

//now using pointers

// 	var intSlice []int = []int{1,2,3,4,5}

// 	var result []int = square(&intSlice)
// 	fmt.Printf("the address of intSlice is %v==>%p \n",&intSlice,intSlice)
// 	fmt.Printf("the address of result is %v==>%p \n",&result,result)
// 	fmt.Println(result)

// }

// func square(thing2 *[]int) []int{
// 	for i,v := range *thing2{
// 		(*thing2)[i] = v * v
// 	}
// 	return *thing2
// }

// switch case

// func main(){
// 	day := time.Now().Weekday()
// 	fmt.Println(day)

// 	switch day{
// 	case time.Saturday,time.Sunday:
// 		fmt.Println("its weekend")
// 	default:
// 		fmt.Println("its weekday")
// 	}
// }

// variadic functions

// func main(){
// 	sum(1,2,3,4,5)
// }

// func sum(nums ...int){
// 	fmt.Println(nums)
// }

// func main(){
// 	nextInt := intSeq()
// 	fmt.Println(nextInt()) //1
// 	fmt.Println(nextInt()) //2
// 	fmt.Println(nextInt()) //3

// 	newInts := intSeq()
//     fmt.Println(newInts()) //1
// 	fmt.Println(newInts()) //2
// 	fmt.Println(newInts()) //3
// }

// func intSeq() func() int{
// 	i:=0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }

//recursion

// func factorial(n int) int{
// 	if n==0{
// 		return 1
// 	}
// 	return n * factorial(n-1)
// }

// func main(){
// 	fmt.Println(factorial(5))
// 	fmt.Println("Now Fib series")
// 	var fib func(n int) int
// 	fib = func(n int) int{
// 		if n<=2{
// 			return 1
// 		}
// 		return fib(n-1) + fib(n-2)
// 	}
// 	fmt.Println(fib(5))
// }

// pointers

// func zeroval(ival int){
// 	ival = 0 //passing by value
// }

// func zeroptr(iptr *int){
// 	*iptr = 0 //dereferencing the pointer
// }

// func main(){
// 	i := 1
// 	fmt.Println("initial:",i)

// 	zeroval(i)
// 	fmt.Println("zeroval:",i)

// 	zeroptr(&i)
// 	fmt.Println("zeroptr:",i)

// 	fmt.Println("pointer:",&i)
// }

// var m = sync.RWMutex{}
// var wg = sync.WaitGroup{}
// var dbData = []string{"user1","user2","user3"}
// var results = []string{}

// func main(){
// 	t0 := time.Now()
// 	for i:=0;i<len(dbData);i++{
// 		// dbCall(i)  //this will execute sequentially
// 		// go dbCall(i) //this will execute concurrently and do not wait for the previous call to finish
// 		wg.Add(1)
// 		go dbCall(i)
// 	}
// 	wg.Wait()
// 	fmt.Println(t0)
// 	fmt.Println(results)

// }

// func dbCall(i int){
// 	// var delay float32 = rand.Float32() * 2000
// 	var delay float32 = 2000
// 	time.Sleep(time.Duration(delay) * time.Millisecond)
// 	// fmt.Println(dbData[i])
// 	// m.Lock()
// 	// results = append(results, dbData[i])
// 	save(dbData[i])
// 	log()
// 	// m.Unlock()
// 	wg.Done()
// }

// func save(result string){
// 	m.Lock()
// 	results = append(results, result)
// 	m.Unlock()
// }

// func log(){
// 	m.RLock()
// 	fmt.Printf("\nresult for log %v", results)
// 	m.RUnlock()
// }

// import "fmt"

// func main(){
// 	var c = make(chan int)
// 	c <- 1
// 	var i = <- c
// 	fmt.Println(i)
// }

// above give deadlock error

// channels are more like to a list and if its defined must be read
// func main(){
// 	var c = make(chan int) // creates a channel
// 	go process(c) // starts a goroutine and call process

// 	// for i:=0;i<5;i++{
// 	// 	fmt.Println(<-c)
// 	// }
// 	for i:= range c{
// 		fmt.Println(i)
// 	}
// 	// fmt.Println(<-c) // this prints the value from channel directly
// }

// func process(c chan int){
// 	defer close(c)
// 	for i:=0;i<5;i++{
// 		c <- i
// 	}
// 	// c<- 12345 // sets c to 12345
// }

// import (
// 	"fmt"
// 	"time"

// )

// func main(){
// 	var c = make(chan int,5) // creates a channel
// 	go process(c) // starts a goroutine and call process

// 	// for i:=0;i<5;i++{
// 	// 	fmt.Println(<-c)
// 	// }
// 	for i:= range c{
// 		fmt.Println(i)
// 		time.Sleep(time.Second*1)
// 	}
// 	// fmt.Println(<-c) // this prints the value from channel directly
// }

// func process(c chan int){
// 	defer close(c)
// 	for i:=0;i<5;i++{
// 		c <- i
// 	}
// 	fmt.Println("Existing Process func")
// 	// c<- 12345 // sets c to 12345
// }


//select and channels
// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )


// var MAX_FRIES_PRICES float32 = 5
// var MAX_BURGER_PRICES float32 = 3

// func main(){
// 	var friesChannel = make(chan string)
// 	var burgerChannel = make(chan string)
// 	var websites = []string{"walmart.com","costco.com","wholefoods.com"}
// 	for i := range websites{
// 		go checkFriesPrices(websites[i],friesChannel)
// 		go checkBurgerPrices(websites[i],burgerChannel)
// 	}
// 	sendMessage(friesChannel, burgerChannel)
// }

// func checkFriesPrices(website string, friesChannel chan string){
// 	for {
// 		time.Sleep(time.Second*1)
// 		var friesPrice = rand.Float32()*20
// 		if friesPrice <= MAX_FRIES_PRICES{
// 			friesChannel <- website
// 			break
// 		}
// 	}
// }

// func checkBurgerPrices(website string, burgerChannel chan string){
// 	for {
// 		time.Sleep(time.Second*1)
// 		var burgerPrice = rand.Float32()*20
// 		fmt.Println(burgerPrice)
// 		if burgerPrice <= MAX_BURGER_PRICES{
// 			burgerChannel <- website
// 			break
// 		}
// 	}
// }

// func sendMessage(friesChannel chan string, burgerChannel chan string){
// 	// fmt.Printf("Found a deal on fries at %s \n", <-friesChannel)
// 	select{
// 	case website:= <- friesChannel:
// 		fmt.Printf("Text sent: Found a deal at Fries at %v \n",website)
// 	case website:= <- burgerChannel:
// 		fmt.Printf("Email sent: Found a deal at Bruger at %v \n",website)

// 	}
// }



// Generics 

import  "fmt"

func main(){
    var intSlice = []int{1,2,3}
    fmt.Println(sumSlice[int](intSlice))

    var floatSlice = []float32{1.5,2.5}
    fmt.Println(sumSlice[float32](floatSlice))
}

func sumSlice[T int | float32 | float64](slice[]T)T{
    var sum T
    for _, v:=range slice{
        sum+=v
    }
    return sum
}