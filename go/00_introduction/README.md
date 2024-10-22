# Introduction

- Go (sometimes called Golang) is one of the most popular languages today, and is a key part of many enterprise tech stacks.
- Many developers prefer Go to other languages like C++ and Scala because of its memory management model which allows for easier concurrency.
- designed by Robert Griesemer, Rob Pike, and Ken Thompson.
- It is an open-source programming language that makes it easy to build simple, reliable, and efficient software solutions.
- Go is a statically typed and compiled programming language.
- Statically typed means that variable types are explicitly declared and thus are determined at compile time. Whereas, by compiled language we mean a language that translates source code to machine code before execution.

## History

- Go’s year of genesis was 2007 at Google, and it was publicly launched in 2009 with a fully open-source BSD-style license released for the Linux and Mac OS platforms.
- The first Windows-port was announced on November 22 of the same year.
- Go 1.0 (the first production-ready version) was released in March 2012. Since 2012, Go has grown from version 1.1 to 1.12 (March 2019), and work for Go 2.0 is underway!

## Languages that influenced Go

- Go belongs to the C-family, like C++, Java, and C#, and is inspired by many other languages created and used by its designers.
- Go has the features of a dynamic language, so Python and Ruby programmers feel more comfortable with it.

## Languages that influenced Go

[image]

## Why a new language?

- Evolving with computing landscape
- Need for faster software developement
- Need for efficiency and ease of programming

### Evolving with computing landscape #

- Programming languages like C/C++ did not evolve with the computing landscape, so there is a need for a new systems language, appropriate for the needs of our computing era.

### Need for faster software development #

- In contrast to computing power, software development has not become considerably faster or more successful (considering the number of failed projects), whereas applications still grow in size.
- a new low-level language was needed, equipped with higher concepts.

### Need for efficiency and ease of programming #

- Before Go, a developer had to choose between fast execution but slow and inefficient building (like C++) or efficient compilation but not so fast execution (like .NET or Java), or ease of programming but slower execution (like dynamic languages such as Python, Ruby or JavaScript).
- Go is an attempt to combine all the three wishes: efficient and fast compilation, fast execution, and ease of programming.

### Targets of Go #

The main target of Golang’s design was to combine the efficacy, speed, and safety of a statically typed and compiled language with the ease of programming of a dynamic language to make programming more fun again.

other targets

- Support for network communication, concurrency, and parallelization
- Support for excellent building speed
- Support for memory management

### Support for network communication, concurrency, and parallelization #

To get the most out of distributed and multi-core machines excellent support for networked-communication, concurrency, and parallelization. Golang was expected to achieve this target for internal use in Google, and this target is achieved through the concepts of goroutines.

### Support for excellent building speed #

- There was a growing concern to improve the building speed (compilation and linking to produce machine code) of C++ projects, which are heavily used in Google’s infrastructure.
- This concern gave birth to the idea of developing the Go programming language.
- In particular, dependency management is a very important part of software development today. The “header files” of languages caused considerable overhead leading in order to build times of hours for the biggest projects.
- Developers felt the need for clean dependency analysis and fast compilation, which Go language provides with its package model.
- The entire Go standard library compiles in less than 20 seconds. Typical projects compile in half a second.
- This lightning-fast compiling process is even faster than C or Fortran, making compilation a non-issue

### Support for memory management #

- Because memory problems (memory leaks) are a long-time problem of C++, Go’s designers decided that memory management should not be the responsibility of a developer.
- So although Go executes native code, it runs in a small runtime, which takes care of an efficient and fast garbage collection.
- Go also has a built-in runtime reflection capability.

> -- In computer science, reflective programming or reflection is the ability of a process to examine, introspect, and modify its own structure and behavior.

> -- a memory leak is a type of resource leak that occurs when a computer program incorrectly manages memory allocations
`

# Characteristics of Go

- Go is essentially an imperative (procedural, structural) language, built with concurrency in mind.
- It is not truly object-oriented like Java and C++ because it doesn’t have the concepts of classes and inheritance.
- it does have a concepts of interfaces, with which much of the polymorphism can be implemented.
- Go has a clear and expressive type-system, but it is lightweight and without hierarchy.
- So in this respect, it could be called a hybrid language.

Some features of modern OOP languages were intentionally left out. Because, object orientation was too heavy often leading to cumbersome development constructing big type-hierarchies, and so not compliant with the speed goal of the language. As per the decision made by the Go-team, the following OOP features are missing from Golang. Although, some of them might still be implemented in its future versions.

- To simplify the design, no function or operator overloading was added.
- Implicit conversions were excluded to avoid the many bugs and confusion arising from this in languages like C/C++.
- No classes and type inheritance is supported in Golang.
- Golang does not support variant types. However, almost the same functionality is realized through interfaces.
- Dynamic code loading and dynamic libraries are excluded.
- Generics are not included (introduced in Go 1.18).
- Exceptions are not included (although recover/panic often goes in that direction).
- Assertions are not included.
- Immutable (unable to change) variables are excluded.
- Golang is a functional language, meaning that functions are the basic building blocks, and their use is very versatile.
- Go is statically typed, making it a safe language that compiles to native code and has a very efficient execution.
- It is also strongly typed, which means according to the principle keep things explicit.
- Implicit type conversions (also called castings or coercions) are not allowed.
- An important thing to note is that Go does have some features of dynamically typed languages (using var keyword).
- That’s why Go also appeals to programmers who left Java and the .Net world for Python, Ruby, PHP, and JavaScript.
- Go has support for cross-compilation, for example, developing and compiling on a Linux-machine for an application that will execute on Windows.
- It is one of the first programming languages that can use UTF-8 not only in strings but also in program code.
- Go is truly an international language because Go source-files are also in UTF-8.

# Uses of Go

There are many uses of Go. Following are some main uses of this language:

- Used as a system programming language
- Used as a general programming language
- Used as an internal support

### Used as a system programming language #

- Go was originally conceived as a systems programming language for the heavy server-centric (Google) world of web servers, storage architecture, and the like.
- For certain domains like high performance distributed systems, Go has already proven to be a more productive language than most others.
- Golang shines in and makes massive concurrency and event-processing easy. So it should be a good fit for the game server and IoT (Internet of Things) development.

### Used as a general programming language #

- Go is also a general programming language, useful for solving text-processing problems, making frontends, or even scripting-like applications
- However, Go is not suited for real-time software because of the garbage collection and automatic memory allocation.

### Used as an internal support #

- Go is being used for some time internally in Google for heavy-duty distributed applications, e.g., parts of Google Maps run on Go.

(Users)
<https://go.dev/wiki/GoUsers#india>

(Go success stories)
<https://go.dev/wiki/SuccessStories>

### Beware of the trap

- If you come to Go and have a background in other contemporary (mostly class or inheritance-oriented languages like Java, C#, Objective C, Python, or Ruby), then beware! You can fall in the trap of trying to program in Go as you did in your previous language.
- However, Go is built on a different model. So if you decide to move code from language X to Golang, you will most likely produce non-idiomatic code that works poorly overall.

### Guiding design principles

- Go tries to reduce typing, clutter, and complexity in coding through a minimal amount of keywords (25). This, together with the clean, regular, and concise syntax, enhances the compilation speed because the keywords can be parsed without a symbol table as its grammar is LALR(1).
- These aspects reduce the number of code lines necessary, even when compared with a language like Java. Additionally, Go has a minimalist approach: there tends to be only one way of getting things done, so reading other people’s code is generally pretty easy, and we all know the code’s readability is of the utmost importance in software engineering.
- The design concepts of the language are orthogonal because they don’t stand in each other’s way, and they don’t add up complexity to one another.
- Golang is completely defined by an explicit specification that can be found here; it is not defined by an implementation as is Ruby, for example. An explicit language specification was required for implementing the two different compilers gc and gccgo, and this in itself was a great help in clarifying the specification.

# Basic Constructs and Elementary Data Types

## Filename

- Go source code is stored in .go files.
- If the name consists of multiple parts, they are separated by underscores ‘_’, like education_platform.go.
- Filenames cannot contain spaces or any other special characters. A source file contains code lines whose length has no intrinsic limits.

## Keyword

- A reserved word, with a special meaning in a programming language, is called a keyword. Below is the set of 25 keywords, or reserved words, used in Go-code:

break | default | func | interface | select
-- | --
case | defer | go | map | struct
chan | else | goto | package | switch
const | fallthrough | if | range | type
coninue | for | import | return | var

## Identifiers

- An identifier is a name assigned by the user to a program element like a variable, a function, a template, and a struct, etc.
- Nearly all things in Go-code have a name or an identifier.
- Like all other languages in the C-family, Go is case-sensitive.
- Valid identifiers begin with a letter (a letter is every letter in Unicode UTF-8) or _and are followed by 0 or more letters or Unicode digits, like X56, group1,_x23, i, and өԑ12.

The following are NOT valid identifiers:

- `1ab` because it starts with a digit
- `case` because it is a keyword in Go
- `a+b` because operators are not allowed

- Apart from the keywords, Go has a set of 36 predeclared identifiers which contain the names of elementary types and some basic built-in functions

append | false | iota | recover
-- | --
bool | float32 | len | string
byte | float64 | make | true
cap | imag | new | uint
close | int | nil | uint8
complex | int8 | panic | uint16
comlex64 | int16 | print | uint32
complex128 | int32 | println | uint64
copy | int64 | real | uintptr

### Blank identifier #

- The _ itself is a special identifier, called the blank identifier.
- _ can be used in declarations or variable assignments (and any type can be assigned to it).
- However, its value is discarded, so it can no longer be used in the code that follows.

### Anonymous

Sometimes it is possible that even functions have no name because it is not really necessary at that point in the code and not having a name even enhances flexibility. Such functions are called `anonymous`.

### The basic structure of a Go program #

- Programs consist of keywords, constants, variables, operators, types and functions.
- It is also important to know the delimiter and punctuation characters that are a part of Golang.

The following delimiters are used in a Go program:

- Parentheses ()
- Braces {}
- Brackets []

The following punctuation characters are used in a Go program:

- .
- ,
- ;
- :
- ...

The code is structured in statements. A statement doesn’t need to end with a ; (like it is imposed on the C-family of languages).

The Go compiler automatically inserts semicolons at the end of statements.

However, if multiple statements are written on one line (a practice which is not encouraged for readability reasons), they must be separated by ;.

## Packages

- A library, module, or namespace in any other language is called a package.
- Packages are a way to structure code.
- A program is constructed as a package which may use facilities from other packages. A package is often abbreviated as ‘pkg’.
- Every Go file belongs to only one package whereas one package can comprise many different Go files.
- Hence, the filename(s) and the package name are generally not the same.
- The package to which the code-file belongs must be indicated on the first line.
- A package name is written in lowercase letters. For example,

``` package main ```

- A standalone executable belongs to main. Each Go application contains one main.
- An application can consist of different packages. But even if you use only package main, you don’t have to stuff all code in 1 big file.
- You can make a number of smaller files, each having package main as the 1st line of code.

## Package dependencies #

- To build a program, the packages, and the files within them must be compiled in the correct order.
- Package dependencies determine the order in which to build the packages.
- Within a package, the source files must all be compiled together.
- The package is compiled as a unit, and by convention, each directory contains one package. If a package is changed and recompiled, all the client programs that use this package must be recompiled too!

## Import keyword #

- A Go program is created by linking set of packages together, with the import keyword. For example, if you want to import a package say fmt, then you do:

```
package main
import "fmt"
```

- import "fmt" tells Go that this program needs functions, or other elements from the package fmt, which implements a functionality for formatted IO.

- Import loads the public declarations from the compiled package; it does not insert the source code.
- If multiple packages are needed, they can each be imported by a separate statement.
For example

```
import "fmt"
import "os"

// or 

import "fmt"; import "os"

```

- Go has provided us with a shorter and more elegant way of importing multiple packages known as factoring the keyword. It is stated as:

```
mport (
  "fmt"
  "os"
)
```

## Visibility

- Packages contain all other code objects apart from the blank identifier (_).
- identifiers of code-objects in a package have to be unique which means that there can be no naming conflicts.
- However, the same identifier can be used in different packages. The package name qualifies a package to be different.

### Visibility rule #

Packages expose their code objects to code outside of the package according to the following rule enforced by the compiler

When the identifier (of a constant, variable, type, function, struct field, …) starts with an uppercase letter, like, Group1, then the ‘object’ with this identifier is visible in code outside the package (thus available to client-programs, or ‘importers’ of the package), and it is said to be exported (like public identifiers/variables in OO languages). Identifiers that start with a lowercase letter are not visible outside the package, but they are visible and usable in the whole package (like private identifiers/variables).

> Note: Capital letters can come from the entire Unicode-range, like Greek; not only ASCII letters are allowed.

Importing a package gives access only to the exported objects in that package. Suppose we have an instance of a variable or a function called Object (starts with O so it is exported) in a package pack1. When pack1 is imported in the current package, Object can be called with the usual dot-notation from OO-languages:

> pack1.Object

Packages also serve as namespaces and can help us avoid name-conflicts. For example, variables with the same name in two packages are differentiated by their package name, like pack1.Object and pack2.Object.

A package can also be given another name called an alias. If you name a package then its alias will be used throughout the code, rather than its original name. For example:

`import fm "fmt"`

Now in the code, whenever you want to use fmt, use its alias:fm (not fmt).

> Note: Go has a motto known as “No unnecessary code!”. So importing a package which is not used in the rest of the code is a build-error.

# Functions

The simplest function declaration has the format: `func functionName()`

- Between the mandatory parentheses ( ) no, one, or more parameters (separated by ,) can be given as input to the function. After the name of each parameter variable must come its type.

- The main function as a starting point is required (usually the first function), otherwise the build-error: undefined: main.main occurs.

- The main function has no parameters and no return type (in contrary to the C-family) otherwise, you get the build-error: func main must have no arguments and no return values.

- When the program executes, after initializations the first function called (the entry-point of the application) will be the main.main() (like in C). The program exits immediately and successfully when main.main returns.

- The code or body in functions is enclosed between braces { }. The first { must be on the same line as the declaration otherwise you get the error: syntax error: unexpected semicolon or newline before { ).

```
func func_Name(param1 type1, param2 type2, ...){
  ...
}
```

- If the function is returning an object of type type1, we follow the syntax as:

```
func func_Name(param1 type1, param2 type2, ...) type1 {
  ...
}

// or

func func_Name(param1 type1, param2 type2, ...) ret1 type1 {
  ...
}

```

- where ret1 is a variable of type type1 to be returned. So a general function returning multiple variables looks like:

```
func func_Name(param1 type1, param2 type2, ...) (ret1 type1, ret2 type2, ...) {
...
}
```

- Smaller functions can be written on one line like:

`func Sum(a, b int) int { return a + b }`

- Let’s create the main function now as an entry point.

```
package main
import "fmt"

func main(){
}
```

- Hello World

```
package main // making package for standalone executable
import "fmt" // importing a package   

func main() { // making an entry point 
    // printing using fmt functionality
    fmt.Println("Hello World")
} // exiting the program
```

# Comments

- Explanation of source code added to a program as a text note is called a comment.
- Comments are un-compilable. They are just for the understanding of the user.
- n Go, a one-line comment starts with //. A multi-line or block-comment starts with /*and ends with*/, where nesting is not allowed.
- examle

```
package main
import "fmt" // Package implementing formatted I/O.
func main() {
    fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
}
```

# Naming things in Go

- Clean, readable code and simplicity are major goals of Go development.
- Therefore, the names of things in Go should be short, concise, and evocative.
- Long names with mixed caps and underscores which are often seen e.g., in Java or Python code, sometimes hinder readability.
- Names should not contain an indication of the package.
- A method or function which returns an object is named as a noun, no Get… is needed.
- To change an object, use SetName. If necessary, Go uses MixedCaps or mixedCaps rather than underscores to write multiword names.

# Overview of Data Types

## Types

- Variables contain data, and data can be of different data types or types for short.
- Go is a statically typed language. It means the compiler must know the types of all the variables, either because they were explicitly indicated, or because the compiler can infer the type from the code context.
- A type defines the set of values and the set of operations that can take place on those values.

Types | Examples

--- | ---

elementary (or primitive) | `int`, `float`, `bool`, `string`
structured (or composite) | `struct`, `array`, `slice`, `map`, channel`
interfaces | They describe the behavior of a type

- A structured type, which has no real value (yet), has the value nil, which is also the default value for these types. To declare a variable, var keyword is used as:

`var var1 type1`

var1 is the variable name, and type1 is the type of var1.

- Functions can also be of a certain type. The (return) type of function is the type of variable which is returned by it. This type is written after the function name and its optional parameter-list, like:

`func FunctionName (a typea, b typeb) typeFunc`

- So, you can see that typeFunc is the (return) type of the function, FunctionName.
- The returned variable var1 of type typeFunc appears somewhere in the function in the statement as:

`return var1`

- A function can have more than one return variables. In this case, the return-types are separated by comma(s) and surrounded by ( ), like:

`func FunctionName (a typea, b typeb) (t1 type1, t2 type2)`

Two variables are returned with type type1 and type2 respectively. For such a case, the return statement takes the form:

`return var1, var2`

- We can also have user-defined data types.  it is possible to have an alias for data types similar to what we have for packages. For example

`type IZ int`

- Now to declare an integer variable, we have to use an alias like:
`var a IZ = 5`

- If you have more than one type to define, you can use the factored keyword form, as in:

```
type (
  IZ int
  FZ float32
  STR string
)
```

## Conversions

- Sometimes a value needs to be converted into a value of another type called type-casting.
- Go does not allow implicit conversion, which means Go never does such a conversion by itself.
- The conversion must be done explicitly as valueOfTypeB = typeB(valueOfTypeA).

```
package main
import "fmt"

func main(){
    var number float32 = 5.2         // Declared a floating point variable
    fmt.Println(number)              // Printing the value of variable
    fmt.Println(int(number))         // Printing the type-casted result
}
```

# Constants

- A value that cannot be changed by the program is called a constant. This data can only be of type boolean, number (integer, float, or complex) or string.

## Explicit and implicit typing

In Go, a constant can be defined using the keyword const as:

`const identifier [type] = value`

Here, identifier is the name, and type is the type of constant. Following is an example of a declaration:

`const PI = 3.14159`

You may have noticed that we didn’t specify the type of constant PI here. It’s perfectly fine because the type specifier [type] is optional because the compiler can implicitly derive the type from the value. Let’s look at another example of implicit typing:

`const B = "hello"`

The compiler knows that the constant B is a string by looking at its value. However, you can also write the above declaration with explicit typing as:

`const B string = "hello"`

> Remark: There is a convention to name constant identifiers with all uppercase letters, e.g., const INCHTOCM = 2.54. This improves readability.

## Typed and untyped constants #

- Constants declared through explicit typing are called typed constants, and constants declared through implicit typing are called untyped constants.
- A value derived from an untyped constant becomes typed when it is used within a context that requires a typed value. For example:

```
var n int  
f(n + 5)   // untyped numeric constant "5" becomes typed as int, because n was int.
```

## Compilation #

- Constants must be evaluated at compile-time. A const can be defined as a calculation, but all the values necessary for the calculation must be available at compile time. See the case below:

`const C1 = 2/3 //okay`

Here, the value of c1 was available at compile time. But the following will give an error:

`const C2 = getNumber() //not okay`

Because the function `getNumber()` can’t provide the value at compile-time. A constant’s value should be known at compile time according to the design principles where the function’s value is computed at run time. So, it will give the build error: `getNumber() used as value`.

## Overflow

Numeric constants have no size or sign. They can be of arbitrarily high precision and do not overflow:

```
const Ln2= 0.693147180559945309417232121458\
176568075500134360255254120680009
const Log2E= 1/Ln2 // this is a precise reciprocal
const BILLION = 1e9 // float constant
const HARD_EIGHT = (1 << 100) >> 97
```

We used \ (backslash) in declaring constant Ln2. It can be used as a continuation character in a constant.

## Multiple assignments

- The assignments made in one single assignment statement are called multiple assignments. Go allows different ways of multiple assignments. Let’s start with a simple example:

`const BEEF, TWO, C = "meat", 2, "veg"`

- As you can see, we made 3 constants. All of them are untyped constants. Let’s look at another method where all the constants are named first, and then their values are written if needed. For example:

`const MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY int= 1, 2, 3, 4, 5, 6`

- As you can see, the constants, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY and SATURDAY are typed constants because their type (int) is mentioned explicitly, and they have the values 1, 2, 3, 4, 5 and 6 respectively.

## Enumerations

- Listing of all elements of a set is called enumeration. Constants can be used for enumerations. For example:

```
const (
  UNKNOWN = 0
  FEMALE = 1
  MALE = 2
)
```

- UNKNOWN, FEMALE and MALE are now aliases for 0, 1 and 2. Interestingly value iota can be used to enumerate the values. Let’s enumerate the above example with iota:

```
const (
  UNKNOWN = iota
  FEMALE = iota
  MALE = iota
)
```

- The first use of iota gives 0.
- Whenever iota is used again on a new line, its value is incremented by 1;
- so UNKNOWN gets 0, FEMALE gets 1 and MALE gets 2.
- Remember that a new const block or declaration initializes iota back to 0. The above notation can be shortened, making no difference as:

```
const (
  UNKNOWN = iota
  FEMALE
  MALE
)
```

# Variables
