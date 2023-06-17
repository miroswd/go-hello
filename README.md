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