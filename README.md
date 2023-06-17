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
