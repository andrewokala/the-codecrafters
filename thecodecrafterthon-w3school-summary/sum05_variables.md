# GO Variables
Variable are containers for storing data values.

In GO, there are differet types of variables.
1. `int`: stores integers(whole numbers) such as 123 0r -123
2. `float32`: stores floating point numbers, with decimals, such as 20.20 or -20.20
3. `string`: store text, such as "Hello World". String values are surrounded by double quotes
4. `bool`: store values with two states: true or false

### Declaring a Variable
There are two ways to declare variable in GO:
1. With the `var` keyword
***e.g*** the syntax `var variablename = value`
2. With the `:=` sign:
***e.g*** the syntax `variablename := value`. Note: in this case the type of the variable is inferred from the value(meaning the compiler decides the type of the variable, based on the value). It is not possible to declare a variable using  `:=` without assigning value to it.

### Difference Between var and :=
1. `var` can be used inside and outside of a functions
`:=` can only be used inside a function
2. `var` Variable declaration and value assignment can be done sseparately
`:=` Variable declaration and value assignment cannot be done separately (must be done in the same line)

## GO Multiple Variable Declaration
It is possible to declare multiple variables on the same line.

***e.g*** `var a, b, c, d int = 1, 2, 3, 4, 5`

Note: If the `type` keyword is used, it is only possible to store one type of variable per line. If the `type` keyword is not used, you can store different type of variable per line.

***e.g*** `var a, b = 6, "Hello"`

## GO Variable Declaration in Block
Multiple variable declaration can be grouped together into a block for greater readability.
***e.g*** 
    `var (`
    `a int`
    `b int = 1`
    `c string = "Hello"`
    `)`

## GO Variable Naming Rules
A variable can have a short name like x and y or more descriptive name like age, price, carname, etc.

1. A variable name must start with a letter or an underscore character(_)
2. A variable name cannot start with a digit
3. A variable name can only contain alpha-numeric charaacters and underscores (a-z, A-Z, 0-9)
4. Variable names are case sensitive (age, Age and AGE are three different variables)

There are several techniques you can use to make a variable name readable
1. Camel Case: Each word, except the first starts with capital letter. 
***e.g*** `myVariableName = "Andrew"`
2. Pascal Case: Eacch words start with capital letter. 
***e.g*** MyVariableName = "Andrew"
3. Snake Case: Each word is separated by and underscore character. 
***e.g*** my_variable_name = "Andrew"