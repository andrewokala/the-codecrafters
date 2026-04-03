# GO Loops
The `for` loop loops through a block of code a specified number of times.

The `for` loop is the only loop available in GO

Each execution of loop is called `iteration`

The `for` loop can take three statements

Syntax:

`for statement1; statement2; statement3 {`

`code to be executed for each iteration`

`}`

`Statement1` Initializes the loop counter value

`statement2` Evaluated for each loop iteration. if it evaluates to TRUE, tthe loop continues. if it evaluetes to FALse, the loop ends.

`statement3` Increases the loop counter value.

***Note*** These statements don't need to be present as loops arguments. However, they need to be present in the code in some form.

### The continue Statement in Loops

`for i:=0; i < 5; i++ {`

`    if i == 3 {`

`      continue`

`    }`

`   fmt.Println(i)`

`  }`

***Output***

`0`

`1`

`2`

`4`

### The break Statement in Loops

`for i:=0; i < 5; i++ {`

`    if i == 3 {`

`      break`

`    }`

`   fmt.Println(i)`

`  }`

***Output***

`0`

`1`

`2`

***Note*** `continue` and `break` are usuallly used with conditions

### Nested Loops

`adj := [2]string{"big", "tasty"}`

`  fruits := [3]string{"apple", "orange", "banana"}`

`  for i:=0; i < len(adj); i++ {`

`    for j:=0; j < len(fruits); j++ {`

`      fmt.Println(adj[i],fruits[j])`
      
`    }`
    
` }`

***Output***

`big apple`

`big orange`

`big banane`

`tasty appple`

`tasty orange`

`tasty banana`

### The Range Keyword

The `range` keyword is used to more easily iterate through the elements of an array, slice or map. It returns both the index and the value.

Syntax:

`for index, value := range array|slice|map {`

`code to be executed for each iteration`

`}`

***Example***

`fruits := [3]string{"apple", "orange", "banana"}`

`  for index, value := range fruits {`

`     fmt.Printf("%v\t%v\n", index, value)`

`  }`

***Output***

`0  apple`

`1  orange`

`2  banana`

To onl show the value or the index, you can omit the other output using an underscore `_`.

Omiting the Index:

`fruits := [3]string{"apple", "orange", "banana"}`

`  for _, value := range fruits {`

`     fmt.Printf("%v\n", val)`

`  }`

***Output***

`apple`

`orange`

`banana`

Omiting the Value:

`fruits := [3]string{"apple", "orange", "banana"}`

`  for index, _ := range fruits {`

`     fmt.Printf("%v\n", index)`

`  }`

***Output***

`0`

`1`

`2`