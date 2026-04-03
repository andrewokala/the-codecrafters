# GO Structs
A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

While arrays are used to store multiple values of the same data type into single variable, structs are used to store multiple values of different data types into single variable. 

A struct can be useful for grouping data together to creat records

### Declaring a Struct
To declare a structure in GO, use the `type` and `struct` keywords:

Syntax:

`type struct_name struct {`
 
 ` member1 datatype;`

 ` member2 datatype;`

  `member3 datatype;`

`}` 

***Example***
Here we declare a struct type `Person` wiht the following members: `name`, `age`, `job` and `salary`:

`type Person struct {`

  `name string`

 ` age int`

  `job string`

`  salary int`

`}`

### Access Struct Members
To access any members of a structure, use the dot operator `.` between the structure variable name and the structure member:

***Example***

`type Person struct {`

 ` name string`

 ` age int`

 ` job string`

`  salary int`

`}`

`func main() {`

`  var pers1 Person`

  `pers1.name = "Hege"`

  `pers1.age = 45`

 ` pers1.job = "Teacher"`

  `pers1.salary = 6000`

  `fmt.Println("Name: ", pers1.name)`

  `fmt.Println("Age: ", pers1.age)`

  `fmt.Println("Job: ", pers1.job)`

  `fmt.Println("Salary: ", pers1.salary)`

`}`

***Output***
`Name: Hege`

`Age: 45`

`Job: Teacher`

`Salary: 6000`

### Passing Struct as Function Arguments
You can also pass a structure as a function argument.

`type Person struct {`

 ` name string`

 ` age int`

 ` job string`

`  salary int`

`}`

`func main() {`

`  var pers1 Person`

  `pers1.name = "Hege"`

  `pers1.age = 45`

 ` pers1.job = "Teacher"`

  `pers1.salary = 6000`

`printPerson(pers1)`

`}`

`func printPerson(pers Person) {`

  `fmt.Println("Name: ", pers.name)`

  `fmt.Println("Age: ", pers.age)`

  `fmt.Println("Job: ", pers.job)`

  `fmt.Println("Salary: ", pers.salary)`

`}`

***Output***
`Name: Hege`

`Age: 45`

`Job: Teacher`

`Salary: 6000`
