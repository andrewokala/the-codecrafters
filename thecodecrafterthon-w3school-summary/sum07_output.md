# Go Output Functions
GO has three functions to output text:
1. Print()
2. Println()
3. Printf()

### The Print() Function
The `Print()` function prints its argument with their default format.
***e.g***
`var i,j string = "Hello", "world"`
`fmt.Print(i)`
`fmt.Print(j)`
Output: 
`HelloWorld`

If we want to print arguments in new lines, we need to use `\n`
***e.g***
`var i,j string = "Hello", "world"`
`fmt.Print(i, "\n")`
`fmt.Print(j, "\n")`
Output: 
`Hello`
`World`
Tip: `\n` creates new lines

It is also possible to use only one Print() for printing multiple variables.
***e.g***
`var i,j string = "Hello", "world"`
`fmt.Print(i, "\n",j)`
Output: 
`Hello`
`World`

If we want to add a space between string arguments, we need to use  `" "`.
***e.g***
`var i,j string = "Hello", "world"`
`fmt.Print(i, " ", j)`
Output:
`Hello World`

`Print()` inserts a space between the arguments if neither are strings:
***e.g***
`var i,j = 10, 20
`fmt.Print(i,j)`
Output:
`10 20`

#### The Println() Function
The `Println()` function is similar to `Print()` with the difference that a whitespace is added between the arguments and a new line is added at the end.
***e.g***
`var i,j string = "Hello", "world"`
`fmt.Println(i,j)`
Output: 
`Hello World`

#### The Printf() Function
The `Printf()` function first formats its arguments based on the given formatting verb and then prints them.

We use two formatting verbs:
1. `%v` is used to print the value of the arguments
2. `%T` is used to print the type of the arguments
***e.g***
`var i, string = "Hello"`
`var j int = 5`


`fmt.Printf("i has value: %v and type: %T\n", i, i)`
`fmt.Printf("j has value: %v and type: %T\n", j, j)`
Output:
`i has value: Hello and type: string`
`j has value: 5 and type: int`