# Starting with go lang

### Initialize the project

```shell
go mod init project
```

### Run main file

```shell
go run .
```

### Create a executable file

```shell
go build
```

```shell
./go-hello
```



## Operators

`=`  -> Assignment operator, for existents variables

`:=` -> Declaration operator


eg: 

```go
x := 10
x = 20
```


## Array
> Only used to improve slice performance 


## Slice
> Similar to array but more flexible

The slice has an underlying array, in other words it has a limited length. Array does not add new values, slice does but does not add using a non-existent position.


```go
slice := []int{1,2,3}

slice[3] = 4 // this does not work

slice = append(slice, 4) // this will work
```


### Make
> Slice with 10 items but with 50 items of capacity

```go
  make([]T, len, cap) // cap -> capacity
```


## Map
> Create a list with keys

```go
	friends := map[string]int{
			"alfredo": 5551234,
			"joana": 9996674,
		}

		fmt.Println(friends["miro"]) // this returns 0, because there is no "miro" in the friends map

		hasFriend, ok := friends["miro"]

    // "ok":bool is used to test for key existence
  ``` 


## Function 

```go
// the variadic parameter must be the last one

func name(s string, x ...int) (int, string) {

}
```

## Defer

The defer runs on stack first in, first out

> execute by last

> after running everything


## Methods
> Function attached to a type

Function bound to a type, must be used when it has a context, as it is not possible to call the function directly

### Method chain
> Create a sequence of methods, function output is used for method input

```go
type employee struct {
	name   string
	age    int
	salary int
}

func (e employee) printName() employee {
	fmt.Printf("Name: %s\n", e.name)
	return e
}

func (e employee) printAge() employee {
	fmt.Printf("Age: %d\n", e.age)
	return e
}

func (e employee) printSalary() {
	fmt.Printf("Salary: %d\n", e.salary)
}

func main() {
	emp := employee{name: "Sam", age: 31, salary: 2000}
	emp.printName().printAge().printSalary()
}
```

## Interface and Polymorphism
> Run a function based on interface type


```go
package main

import (
    "fmt"
    "math"
)

// Here’s a basic interface for geometric shapes.
type geometry interface {
    area() float64
    perim() float64
}

// For our example we’ll implement this interface on rect and circle types.
type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

/*
To implement an interface in Go, we just need to implement all the methods in the interface. 
Here we implement geometry on rects.
*/
func (r rect) area() float64 {
    return r.width * r.height
}

func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// The implementation for circles.
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

/*
If a variable has an interface type, then we can call methods that are in the named interface. 
Here’s a generic measure function taking advantage of this to work on any geometry.
*/
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

/*
The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
*/
    measure(r)
    measure(c)
}
```


## Pointers 
> Stores the memory address

```go
const x = 10
const y = &x // get memory address
const z = *y // get content of memory address 
```

- We use it when we deal with a large amount of data and we don't want to copy it during the code (in variables, pass it in functions). We leave it in a fixed place in memory and whoever wants to use it can look for the value directly in the memory address


## Export
> In go, everything with a capital first letter is exported and can be imported


## Tags Struct
> Used for DDD => user <> costumer

```go
type info struct {
	Name 		string 	`json: "Name"`
	LastName	string 	`json: "LastName"`
	Age 		int    	`json: "Age"`
	Job 		string 	`json: "Profession"`
	BankAccount float64 `json: "Balance"`
}
```


## Json

### Encoder vs Marshal

- Marshal needs it to store values ​​in a variable to be used later

```go
user := person{
    Name: "Miro",
    Age: 23,
}

j, err := json.Marshal(user)
```

- Encoder, sends directly to the interface defined in newEncoder


```go
user := person{
    Name: "Miro",
    Age: 23,
}

encoder := json.NewEncoder(os.Stdout)
encoder.Encode(user)

// encoder := json.NewEncoder(v interface{})
// interface is anything
```


## Package

install a package

```shell
go get -u golang.org/x/crypto/bcrypt
```

## Go routine and WaitGroups
Execute independently

```go 
go any() 
any()
```
- Waiting groups are used to synchronize go routines with code, waiting for all functions to run to close the program (func main)

```go
import "sync"

var wg sync.WaitGroups

wg.Add(numberOfGoRoutines)

wg.Wait() // wait for all functions to be executed

defer wg.Done() // close the process
```


## Race condition
> When two or more concurrent processes access a shared resource simultaneously

- To resolve this, it is necessary to create a synchronization to coordinate access to goroutines between shared variables


*Check if has a race condition*

```shell
go run -race file.go
```

### Mutex
> Mutual exclusion, works like binary lock

- Only one goroutine at a time can acquire the mutex lock and proceed until another goroutine releases the lock.

### Atomic
- It provides a set of functions that allow for the safe reading, writing, and manipulation of variables without the need to block access with a mutex


## Comma ok
> ok is used to differentiate the open/closed channel

```go
    v, ok := <- channel

	fmt.Println(v, ok) // 0, true

	v, ok = <- channel

	fmt.Println(v, ok) // 0, false
```


## Error handling
> Not use try/catch

```go
    v, err := random()

    if err != nil {
        log.Println(err)
        return
    }
```

*Recommended to use log instead of fmt*


### Recover
> The recover function interrupts the current panic and returns the value passed to "panic", allowing you to handle the error in a controlled manner.

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Panic:", r)
        }
    }()

    panic("FAILED")
}
```


## Test
Need to have `_test.go` at the end of the file and the function must start with *Test* 

```go
func TestAny(t *testing.T) {
    // code
}
```


*Running test* 

```shell
go test
```

## GoLint

```shell
gofmt -w file.go # fix the formatting       
go fmt -v ./.. # correct formatting of all files
```

## Benchmark
> Allow testing code speed or performance

BET: Benchmarks, Examples and Tests

```shell
go test -bench BenchmarkMultiply
go test -bench . # running to all files
```