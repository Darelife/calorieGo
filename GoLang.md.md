# A few points about GoLang

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
