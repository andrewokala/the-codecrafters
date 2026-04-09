# Go Functions
A function is a block of statements that can be used repeatedly in a program.

To create a function:

1. Use the `func` keyword.
2. Specify a name for the function followed by parenthesis `()`
3. Finally, add code that defines the function should, inside curly braces `{}`

Syntax:

`func FunctionName() {`

`code to be executed`

`}`

Functions are not executed immediatetly. They are kept for later use, and will be executed when they are called.

***Example***
`func myMessage() {`

`  fmt.Println("I just got executed!")`

`}`

`func main() {`

`  myMessage() // call the function`

`}`

***Output***

`I just got executed!`

Calling a function multiple times:

`func myMessage() {`

`  fmt.Println("I just got executed!")`

`}`

`func main() {`

`  myMessage()`

`  myMessage()`

`  myMessage()`


`}`

***Output***

`I just got executed!`

`I just got executed!`

`I just got executed!`

### Naming Rules for GO Functions
1. A function name must start with a letter
2. A function name can only contain alpha-numeric characters and undercore `(A-z, 0-9, and _)`
3. Function names are case-sensitive
4. Give the function a name that reflects what the fucntion does

### Function Parameters and Arguments
Information can be passed to afunctions as a parameter. Paramter act as variables inside function.

Paramters and their types are specified after the function name, inside the parenthesis.

Syntax:

`func FunctionName(param1 type, param2 type, param3 type) {`

`code to be executed`

`}`

***Example***

`func familyName(fname string) {`

`  fmt.Println("Hello", fname, "Refsnes")`

`}`

`func main() {`

`  familyName("Liam")`

`  familyName("Jenny")`

  `familyName("Anja")`

`}`

***Output***
`Hello Liam Refsnes`

`Hello Liam Refsnes`

`Hello Liam Refsnes`

***Note***
When a `parameter` is passed to the function, it is called an `argument`. So, from the example above: `fname` is a parameter, while `Liam`, `Jenny` and `Anja` are `arguments`

### Multiple Parameters
Inside a function, you can add as many parameters as you want:

***Example***

`func familyName(fname string, age int) {`

  `fmt.Println("Hello", age, "year old", fname, "Refsnes")`

`}`

`func main() {`

`  familyName("Liam", 3)`

  `familyName("Jenny", 14)`

`  familyName("Anja", 30)`

`}`

***Output***
`Hello 3 year old Liam Refsnes`

`Hello 14 year old Jenny Refsnes`

`Hello 30 year old Anja Refsnes`

***Note***
When working with multiple paramters, the function call must have the same number of arguments as there are parameters, and the arguments must be passed in the same other.

### Function Return
If you want the function to return a value, you need to define the data ype of the return value, and use the `return` keyword inside the function:

Syntax:

`func FunctionName(param1 type, param2 type, param3 type) type {`

`code to be executed`

`return output`

`}`

***Example***

`func myFunction(x int, y int) int {`

`  return x + y`

`}`

`func main() {`

`  fmt.Println(myFunction(1, 2))`

`}`

***Output***
`3`

### Recursion Function
GO accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.

***Example***
`func testcount(x int) int {`

 ` if x == 11 {`

   ` return 0`

 ` }`
  `fmt.Println(x)`

  `return testcount(x + 1)`

`}`

`func main(){`

  `testcount(1)`

`}`

***Output***
`1`

`2`

`3`

`4`

`5`

`6`

`7`

`8`

`9`

`10`

Recursion is a common mathematical and programming concept. This has the benefit of meaning that you can loop through data to reach a result.