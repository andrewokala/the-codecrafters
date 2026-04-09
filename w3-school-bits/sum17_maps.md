# GO Maps

Maps are used to store data values iin key:value pairs.

A map is an unordered and changeable collection that does not allow duplicates.

The default value of a map is `nil`.

Maps hold references to an undelying has table.

### Creating Maps with var and :=

Syntax:

`var a = map[KeyType]ValueType{key1:value1, key2:value2,...}`
`b := map[KeyType]ValueType{key1:value1, key2:value2,...}`

***Example***
`func main() {`

  `var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}`
 
  `b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}`

  `fmt.Printf("a\t%v\n", a)`

 ` fmt.Printf("b\t%v\n", b)`

`}`

***Output***
`a   map[brand:Ford model:Mustang year:1964]`

`b   map[Bergen:2 Oslo:1 Stavanger:4 Trondheim:3]`

***Note:*** 
The order of the map elements defined in the code is different from the way that they are stored. The data are stored in a way to have efficient data retrieval from the map.

### Creating Map with the make() Funtion

Syntax:

`var a = make(map[KeyType]ValueType)`

`b := make(map[KeyType]ValueType)`

***Example***
`func main() {`

  `var a = make(map[string]string) // The map is empty now`

  `a["brand"] = "Ford"`

 ` a["model"] = "Mustang"`

`  a["year"] = "1964"`

  `b := make(map[string]int)`

  `b["Oslo"] = 1`

  `b["Bergen"] = 2`

  `b["Trondheim"] = 3`

  `b["Stavanger"] = 4`

  `fmt.Printf("a\t%v\n", a)`

  `fmt.Printf("b\t%v\n", b)`

`}`

***Output***
`a   map[brand:Ford model:Mustang year:1964]`

`b   map[Bergen:2 Oslo:1 Stavanger:4 Trondheim:3]`

### Updating and Adding to Map
Syntax:

` map_name[key] = value`

***Example***
`func main() {`

`  var a = make(map[string]string)`

  `a["brand"] = "Ford"`

  `a["model"] = "Mustang"`
  
  `a["year"] = "1964"`

`  fmt.Println(a)`

  `a["year"] = "1970"`

  `a["color"] = "red"`

  `fmt.Println(a)`

`}`

***Output***
`map[brand:Ford model:Mustang year:1964]`

`map[brand:Ford color:red model:Mustang year:1970]`

### Removing an Element from Map
Removing element from a map is done using `delete()` function.

Syntax:

` delete(map_name, key)`

***Example***
`func main() {`

  `var a = make(map[string]string)`

  `a["brand"] = "Ford"`

  `a["model"] = "Mustang"`

  `a["year"] = "1964"`

 ` fmt.Println(a)`

`  delete(a,"year")`

`  fmt.Println(a)`

`}`

***Output***
`map[brand:Ford model:Mustang year:1964]`

`map[brand:Ford model:Mustang]`

### Checking for Specific Element in a Map
You can check if certian key exists in a map

Syntax:

` val, ok :=map_name[key]`

If you only want to check the existence of a certain key, you can use the blank identifier `_` in place of val.

***Exaxmple***
`func main() {`

`  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}`

  `val1, ok1 := a["brand"]`

  `val2, ok2 := a["color"]`

  `val3, ok3 := a["day"]`

 ` _, ok4 := a["model"]`

  `fmt.Println(val1, ok1)`

  `fmt.Println(val2, ok2)`

  `fmt.Println(val3, ok3)`

 ` fmt.Println(ok4)`

`}`

***Output***
`Ford true`

 `false`

 `true`

`true`