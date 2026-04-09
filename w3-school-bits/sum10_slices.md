# GO Slices
Slices are similar to arrays, but are more powerful and flexible
Slices are also used to store multiplt values of the same type in a single variable
The lenth of a slice can grow and shrink as you see fit.

There are several ways to create a slice:
1. Using the `[]datatype{values}` format
2. Create a slice from an array
3. Using the `make()` function

### Creating Slice with []datatype{values}

`myslice := []int{}`

The code above declares an empty slice of `0` lenth and `0` capacity
To initialize the slice when declaring, us this:
`myslice := []int(1,2,3)`
You can check the lenth of the slice using the `len()` function and also the capacity using `cap()` function.

### Creating a Slice from an Array

`var myarray = [...]int{1,2,3,4,5}`
`myslice := myarray[2:4]`
***Output***
`[3 4]`

### Creating Slice with The make() Function

`myslice := make([]string, 5, 7)`
***Note***
If the capacity parameter is not defined, it will be equal to the lenth