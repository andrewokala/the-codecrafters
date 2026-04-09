# GO Syntax
In GO all files consist of the following parts:
1. Package declaaration
2. Import packages
3. Functions
4. Statements and expressions

***EXAMPLE***

`package main`

`import "fmt"`

`func main() {`
    `fmt.Println("Hello World")`
`}`

1. Line 1: every GO program is part of a package. It's define with the `package` keyword, the program belong to the `main` package

2. Line 2: `import ("fmt")` lets us import files included in the `fmt` package

3. Line 3: A blank line. GO ignores white spaces. Having white spaces make the code readable.

4. `func main() {}` is a function. every code in the curly braces `{}` will be executable

5. `fmt.Println()` is function available from the `fmt` package. It is used for output/print text.

***NOTE***
In GO any executable code belongs to the main package.
"

### GO Statement

`fmt.Println("Hello WOrld")` is a statement. In GO, statement are separated by ending a line (hitting the Enter key) or by semicolon `;`.