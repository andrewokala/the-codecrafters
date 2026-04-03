# GO Conditions
Conditional statements are used to perform different actions basaed on the different conditions.

GO Conditions can be either `true` of `false`

Go supports operations from the `Comparison Operators`

1. Less than `<`
2. Less than or equal `<=`
3. Greater than `>`
4. Greater than or equal `>=`
5. Equal to `==`
6. Not equal to `!=`

GO supporst operations from the `Logical Operatos=rs`

1. Logical AND `&&`
2. Logical OR `||`
3. Logical NOT `!`

GO has the following conditional staterments:

1. Use `if` to specify a block of code to be executed, if the specified condition is true
2. Use `else` to specify a block of code to be executed, if the same condition is false
3. Use `else if` to specify a new condition to test, if the first condition is false
4. Use `switch` to specify many alternatives blocks of code to be executed

### if Statement

Syntax:

`if condition {`
    `code to be executed if condition is true`
`}`

***Note*** that `if` is in lowercase. Uppercase letters will generate an error.

### if else Statement

Syntax:

`if conditon {`
    `code to be executed if condition is true`
`} else {`
    `code to be executed if condition is false`
`}`

### else if Statement

Syntax:

`if condition1 {`
    `code to be executed if condition1 is true`
`} else if condition2 {`
    `code to be executed if condition1 is false and condition2 is true`
`} else {`
    `code to be executed if both of the conditions are false`
`}`

### Nested if Statement
You can have `if` statements inside `if` statements, this is called nested if.

Syntax:

`if condition1 {`
    `code to be executed if condition1 is true`
    `if condition2 {`
        `code to be executed if condition1 and condition2 are true`
    `}`
`}`