# GO Constants
Constants are used to store fixed value that cannot be changed
The `const` keyword declares the variable as "constant" which means that it is unchangeable and read only

### Syntax
`const CONSTNAME type = value`
***Note:*** The value of a constant must be assign upon decclaration.

### Constant Rules
1. Constant names follow the same naming rules as variables
2. Constant ames are usually written in uppercase for easy identification
3. Constants can be declared both inside and outside a function

### Constant Type
There are two types of constants:
1. Typed Constants
2. Untyped Constants

Typed Constants: are declared with the type.
***e.g*** `const A int = 1`
Untyped Constants: are declared without a type.
***e.g*** `const A = 1` Note: in this case, the type is inferred from the value

### Multiple Constants Declaration
Multiple constants can be grouped together like variables into a block for readability.
***e.g***
    `const (`
    `A int = 1`
    `B = 3.14`
    `C = "Hi"`
    `)`