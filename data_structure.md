# Data Structures

---

# Array

var name [n]T

```go
    var array [5]string
```

---

# Array auto counting

var name [n]T

```go
    primes := [...]int{2, 3, 5, 7, 11, 13}    

```

---

# Foreach in Array

var name [n]T

```go
    primes := [...]int{2, 3, 5, 7, 11, 13}

    for i := range primes {
        fmt.Println(primes[i])
    }
```

---

# Foreach index & value

var name [n]T

```go
    primes := [...]int{2, 3, 5, 7, 11, 13}

    for i, prime := range primes {
        fmt.Println(i, prime)
    }
```

---

# Slice

var name []T

```go
    var slice []int

    primes := [...]int{2, 3, 5, 7, 11, 13}    
    slice = primes[1:4]

    for i := range slice {
        fmt.Println(slice[i])
    }
```

---

# Zero value of slice is nil

```go
    var s []int

    if s == nil {
        fmt.Println("it's nil")
    }

// make it first
    s := make([]int, 5)
```

---

# append slice

```go
    s := []int{1, 2, 3, 4, 5}
    s = append(s, 6, 7, 8)
```

---

# Slice Structure

![width:650px](./asset/golang-slices-illustration.jpg)

---

# Slice Structure with value

```go
var a = [6]int{10, 20, 30, 40, 50, 60}
var s = [1:4]
```

![width:650px](./asset/golang-slices-length-capacity.jpg)

---

# Tour of Go

https://tour.golang.org
https://tour.golang.org/moretypes/18

---

# Variadic function (Variable number of arguments)

```go
func variadic(nums ...int) {

}
```

```go
variadic()
variadic(1)
variadic(1,2,3,4)
```

---

# Spread operator

```go
func variadic(nums ...int) {
    
}

var slice = []int{1, 3, 5, 7, 9}
variadic(slice...)

```

---

# map[T]T

```go
    var m map[string]string
```

---

# zero value of map is nil

```go
    var m map[string]string

    if m == nil {
        fmt.Println("it's nil")
    }
```

---

# map need home

## make

```go
    m := make(map[string]string)

    if m == nil {
        fmt.Println("it's nil")
    }

    m["a"] = "apple"
    m["b"] = "banana"
    m["c"] = "coconut"
    m["d"] = "durian"
    m["e"] = "elderberry"
    m["f"] = "fig"
    m["g"] = "guava"
```

---

# construct map

```go
    m := map[string]string{
        "a" : "apple",
        "b" : "banana",
        "c" : "coconut",
        "d" : "durian",
        "e" : "elderberry",
        "f" : "fig",
        "g" : "guava",
    }

    for k, v := range m {
        fmt.Println(k, v)
    }
```

---

# delete a key

```go
    m := map[string]string{
        "a" : "apple",
        "b" : "banana",
        "c" : "coconut",
        "d" : "durian",
        "e" : "elderberry",
        "f" : "fig",
        "g" : "guava",
    }

    delete(m, "d")

    for k, v := range m {
        fmt.Println(k, v)
    }
```

---

# Demo: map

open a file oscar_age_male.csv

https://github.com/focusive/go102/tree/master/testdata

print any actors name who got the oscar more than one time

```txt
    Marlon Brando
    Daniel Day-Lewis
    Sean Penn
    Tom Hanks
    Fredric March
    Spencer Tracy
    Gary Cooper
    Jack Nicholson
    Dustin Hoffman
```

---

# Exercise: map represent json

http://www.groupkt.com/post/c9b0ccb9/country-and-other-related-rest-webservices.htm

http://services.groupkt.com/country/get/iso3code/IND

---

# len() with slice

```go
    s := []primes{2, 3, 5, 7, 11, 13}

    fmt.Println(len(s))
```

---

# len() with map

```go
    m := map[string]string{
        "a" : "apple",
        "b" : "banana",
        "c" : "coconut",
        "d" : "durian",
        "e" : "elderberry",
        "f" : "fig",
        "g" : "guava",
    }

    fmt.Println(len(m))
```

---

# Pointer

Go has pointers. A pointer holds the memory address of a value.

The type *T is a pointer to a T value. Its zero value is nil.

    var p *int
The & operator generates a pointer to its operand.

    i := 42
    p = &i
The # operator denotes the pointer's underlying value.

    fmt.Println(*p) // read i through the pointer p
    *p = 21         // set i through the pointer p
This is known as "dereferencing" or "indirecting".

Unlike C, Go has no pointer arithmetic.

---

# Example Pointer

```go
    var p *int
    i := 42
    p = &i

    fmt.Println(*p, i)

    *p = 43
    fmt.Println(*p, i)
```

---

# new

```go
    var p = new(int)
    i := 42
    p = &i

    fmt.Println(*p, i)

    *p = 43
    fmt.Println(*p, i)
```

---

# Type

> type newType T

```go
    type text string
```

> type aliasType = T

```go
    type text = string
```

---

# method on primitive type inspired by Java

```go
    type Int int

    func (i Int) toString() string{
        return strconv.Itoa(int(i))
    }
```

---

# Demo Int with method Set

New Int Type with Method

.String() to convert integer to string
.Set(n int) to set new value

---

# struct type

```go
    type Rectangle struct {
        Width  float64
        Length float64
    }
```

---

# Construct struct instance

```go
    rec := Rectangle{Width: 10, Length: 12}

    area := rec.Width * rec.Length
```

---

# Exercise JWT Token API

pick a Web Framework to make the Login API

request:

```json
    {
        "email": "abc@mail.com",
        "password": "12345678"
    }
```

response:

```json
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
    }
```

---

# method on struct

```go
    type Rectangle struct {
        Width  float64
        Length float64
    }

    func (r Rectangle) Area() float64 {
        return r.width * r.length
    }
```

---

[Next >](./interface.md#1)