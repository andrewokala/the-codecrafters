# Go Arrays
Arrays are used to store multiple vvalue of the same data type in a single variable.

### Declaring an Array
We can declare arrays in GO with two ways:
1. with the `var` keyword:
e.g
`var array_name = [lenth]datatype{values}`
2. with the `:=` sign:
e.g
`array_name := [lenth]datatype{values}`

***NOTE:***
The lenth is use to specify the number of element to be stored in the array. In GO, arrays have fixed lenth. The lenth is either defined or infered.

***Example***
`package main`

`import ("fmt)`

`func main() {`
`   var array1 = [3]int{1,2,3}`
`   array2 := [...]int{4,5,6,7,8}`
`   fmt.Println(array1)`
`   fmt.Println(array2)`
`}`
***Output***
`[1 2 3]`
`[4 5 6 7 8]`

***Note***
In the second array declared the lenth will be infered by the compiler at compile time.

### Accessing Element of an Array
Array indexes start at `0`. That means that `[0]` is the first element, `[1]` is the second element.

***Example***
`package main`

`import ("fmt)`

`func main() {`
`   array2 := [...]int{10,25,50}`

`   fmt.Println(array2[0])`
`   fmt.Println(array2[1])`
`}`
***Output***
`10`
`50`

### Changing Elements of an Array
The value of an array can be change by referring the index number.

***Example***
`package main`

`import ("fmt)`

`func main() {`
`   array2 := [...]int{10,25,50}`

`   array2[0] = 30`
`   fmt.Println(array1)`
`}`
***Output***
`[30 25 50]`