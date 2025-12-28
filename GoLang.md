# A few points about GoLang [^1]

1. Statically typed language (types can't be changed (not without type conversion))
2. Strongly typed language (can't add int with a string)
3. It's compiled
4. Compilation is fast
5. GoRoutines (builtin concurrency)
6. Simple syntax
7. Garbage Collector

Packages have a bunch of go files
Modules have a bunch of packages

```bash
go mod init <global url : typically github (github.com/abc/abc)>
```

go.mod has the name of the module, go version, and requirements


## Introduction

`./cmd/tutorial/main.go`

We'll have a `package main` line written in the top, signifying the package. Also, `main` is a special package name, that makes the compiler look for the entry point function (main function)

### Functions

> [!info] "fmt" package
> It's a standard library used for formatted input and output operations

```go
import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

```bash
go build cmd/tutorial/main.go
./main

or

go run cmd/tutorial/main.go
```

### Variables

```go
package main

import "fmt"

func main() {
	var intNum int = 32767
	// int, int8, int16, int32, int64, uint, ..., float32, float64
	// string
	
	// we can't perform any arithmetic operation between any 2 types, like int + float32 won't work. I'll have to do this
	//  floatNum32 + float32(intNum32)
	
	var myString string = "Hello World" // single line
	var myString2 string = `Hello
	World`
	
	fmt.Println(len("test")) // 4, but this isn't the number of chars, it's the number of bytes. For ascii, it's 1 byte per char only, but go uses utf8, so, it's slightly variable, for chars outside it
	
	// for actual length, 
	// import "unicode/utf8"
	// fmt.Println(utf8.RuneCountInString("Test"))
	// rune type = character in utf8 
	var myRune rune = 'a'
	fmt.Println(myRune) // 97
	// bool -> true / false
}
```

we don't need to initialize values btw. Go assigns default values too 
For int, int8, ..., uint, uint8,..., float32, ..., rune = 0
For string, it's an empty string, and for booleans, it's false

We can also infer the type, 
```go
var myString = "hello"

// we can also, remove the var keyword

myString := "hello"

// we can also do this

var var1, var2 int = 1, 2
var var3, var4 = 1, 2
var5, var6 := 1, 2
fmt.Println(var1,var2) // 1 2
```

Instead of var, we can use const as well btw. Const = can't be changed. We have to initialize the value explicitly.


# Functions & Control Structures

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello"
	printMe(printValue)
	var num int = 11
	var den int = 2
	var result = intDivsion(num, den)
	fmt.Println(result)
	var res, rem, err = intDivisionBoth(num, den)
	if err != nil {
		fmt.Printf(err.Error())
	} else if rem == 0 {
		fmt.Printf("The result of the int division is %v", res)
	} else {
		fmt.Printf("The result of the int division is %v with remainder %v", res, rem)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(num int, den int) int {
	var result int = num/den
	return result
}

func intDivisionBoth(num int, den int) (int, int, error) {
	var err error // defualt = nil
	if den == 0 {
		err = errors.New("Can't divide by zero")
		return 0, 0, err
	}
	var result int = num/den
	var rem = num%den
	return result, rem, err
}
```

# Arrays, Slices, Maps, and Loops

## Arrays

```go
package main
import "fmt"

func main() {
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0]) // 0
	fmt.Println(intArr[1:3]) // [123 0]
	// index 1 and 2
	
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
}
```

```go
package main
import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{1,2,3}
	fmt.Println(intArr)
}
```

```go
package main
import "fmt"

func main() {
	intArr := [...]int32{1,2,3}
	fmt.Println(intArr)
}
```

## Slices

They are wrappers around arrays. They are literally the golang equivalent of vectors in c++.
They are dynamically sized arrays. However, a slice is a reference into an underlying array, while the std::vector object in c++ owns the underlying storage. 

Slice = <pointer, length, capacity>. It's just a descriptor

```go
func main() {
	// by ommitting the length value in the array, we'll get a slice
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice) // [4 5 6]
	
	fmt.Printf("length: %v, capacity: %v\n", len(intSlice), cap(intSlice))
	// 3 3 : [4 5 6]
	
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice) // [4 5 6 7]
	// It checks if the array has enough space for more values. 
	// If not, a new array is created with enough space, and the data is copied
	
	fmt.Printf("length: %v, capacity: %v", len(intSlice), cap(intSlice))
	// 4 6 : [4 5 6 7 * *]
	
	
	
	// we can also append more values
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // intSlice2... : spread operator
	fmt.Println(intSlice)
	
	
	// We can use the make function to also create a slice
	// type, length, (and optionally, capacity)
	var intSlice3 []int32 = make(int32[], 3, 8)
}
```

## Maps

```go
func main() {
	var myMap map[string]uint8 = make(map[string]uint8) 
	// key: string, value: uint8
	
	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(myMap2["Adam"]) // 23
	
	fmt.Println(myMap2["Jason"]) 
	// doesn't exist : we'll get the default value = 0
	
	var age, ok = myMap2["Adam"] // maps return an optional 2nd value (bool)
	// ok = true : the value is in the map, and if it's false, the value isn't
	
	delete(myMap2, "Adam") // if the key doesn't exist, it does nothing : no-op
}
```

## Loops

```go
func main() {
	
	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	for name := range myMap2 {
		fmt.Println(name)
	}
	for name, age := range myMap2 {
		fmt.Println(name, age)
	}
	
	
	intArr := [...]int32{1,2,3}
	for i, v := range intArr { // i = index, v = value
		fmt.Println(i, v)
	}
}
```

```go

// while loops are implemented via the for keyword
i := 0
for i<10 {
	fmt.Println(i)
	i = i + 1
}

// or

i := 0
for {
	if i >= 10 {
		break
	}
	fmt.Println(i)
	i = i + 1
}

// or use a regular for loop

for i:=0; i<10; i++ {
	fmt.Println(i)
}
```
[^1]:[Learn Go](https://www.youtube.com/watch?v=8uiZC0l4Ajw) 