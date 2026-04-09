# GO Switch
Use the `switch` statement to select one of many code blocks to be executed.

### Single-Case switch Syntax

`switch expression {`

`case x:`

`code block`

`case y:`

`code block`

`case z:`

`code block`

`default`:

`code block`

1. The expression is evaluated once
2. The value of the `switch` expression is compared with the value of each `case`
3. If there is a match, the associated block of code is executed
4. The `default` keyword is optional. It specifies some code to run if there is no `case` match.

***Note** All the case values should have the same type as the `switch` expression. Otherwise, the compiler will raise error.

### Multi=case switch Syntax

`switch expression {`

`case x, y:`

`code block`

`case v, w:`

`code block`

`case a, b:`

`code block`

`default`:

`code block`