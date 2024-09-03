<h1 align=center> <a href="https://en.wikipedia.org/wiki/Yoruba_language">Èdè Yorùbá</a></h1>

## **Table of Contents**
- [Introduction](#introduction)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Hello World](#hello-world)
- [Language Syntax](#language-syntax)
  - [Variables](#variables)
  - [Data Types](#data-types)
  - [Arithmetic Operations](#arithmetic-operations)
  - [Expressions Grouping](#expressions-grouping)
  - [Control Structures](#control-structures)
    - [Conditional Statements](#conditional-statements)
    - [Loop Structures](#loop-structures)
  - [Functions](#functions)
  - [Comments](#comments)
- [Builtin Functions](#builtin-functions)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)
- [Future Additions](#future-additions)
- [Resources](#resources)

## **Introduction**
**yoruba** is a high-level, general-purpose programming language that integrates the syntax and semantics of the Yoruba language.

## **Getting Started**
### **Installation**
To get started with **yoruba**, follow these steps:
1. **Clone the repository:**
   ```bash
   git clone https://github.com/Ayobami0/yourba.git
   ```
2. **Run the interpreter:**
   ```bash
   # yoruba files have the extension .yrb
   ./yoruba file.yrb
   ```

### **Hello World**
Here's how you can write and run your first program in **yoruba**:

```yoruba
pe ko pelu "Aye! E kaabo si Èdè Yorùbá!" pa
```

To run this code, save it in a file called `kiki.yrb` and execute:

```bash
./yoruba kiki.yrb
```

## **Language Syntax**
### **Variables**
Variables are declared using the `jeki` keyword, followed by the variable name, assignment operator `je`, and the value.
```yoruba
jeki age je 25
jeki name je "Ayobami"
```

### **Data Types**
- **Integers**: Whole numbers (`10`, `-5`)
- **Strings**: Text enclosed in single or double quotes (`'Oruko mi ni Ayobami'`, `"Oluwa ni"`)
- **Booleans**: `ooto` (true) and `eke` (false)

### **Arithmetic Operations**
Arithmetic operations :
- **Addition (`+`)**
- **Subtraction (`-`)**
- **Multiplication (`*`)**
- **Division (`/`)**

Example:

```yoruba
jeki sum je 3 + 5
jeki product je 2 * 5
```

### **Expressions Grouping**
Expressions can be grouped together using {} for better readability
```yoruba
[[variables]]
jeki a je {1 + 4}
[[loop statements]]
titi {x baje y} se pari
[[if statement]]
ti {x kobaje y} lehinna pari
[[function calls]]
pe ko pelu {x, y, z} pa
[[nested function calls]]
pe ko pelu {x, y, pe ka pelu {a, b, c}} pa
```

### **Control Structures**
#### **Conditional Statements**
Conditional statements use `ti` and `abi` for `if`, and `else` blocks.

```yoruba
ti age baje 18 lehinna [[if age == 18 then]]
    ko "O ti dagba!" [[write "O ti dagba"]]
abi [[else, yoruba has no else if, so we'd use nested if else statements]]
    ti age baje 13 se 
        pe ko pelu "O ti n dagba!" pa
    abi
        pe ko pelu "O tun kere!" pa
pari [[end]]
```

#### **Loop Structures**
Loops can be implemented using the `titi` keyword. `yoruba` has no while or for loops (for now).

**Until Loop (`titi ... baje ... se ... pari`)**
  ```yoruba
  jeki x je 0
  titi x baje 5 se [[until x is equal 5 do]]
      pe ko pelu x pa
      jeki x je x + 1
  pari [[end]]
  ```
**Infinite loops**
  ```yoruba
  jeki x je 0
  titi eke se [[until false, this would cause the loop to repeat infinitely]]
      pe ko pelu x pa
      jeki x je x + 1
    ti x baje 100 lehinna
        fo [[break]]
    pari
  pari [[end]]
  ```
Loops can be broken using `fo`

### **Functions**
Functions are defined using the `ise` keyword, followed by the function name, parameters, and the function body. The function returns a value using `da` and `pada`.

```yoruba
ise sum a, b se [[ `ise` translate to work or function in english, se is do ]]
    da a + b pada [[ the value to be returned is placed between da and pada keywords ]]
pari [[ end ]]

[[
Call a function with pe and pelu.
pe sum pelu 10, 5 pa translates to call sum with 10 and 5 then kill/terminate
]]
jeki result je pe sum pelu 10, 5 pa
```

### **Comments**
Comments provide explanations or notes within the code and can be represented as:

- **Single-line comments**: `[[ This is a comment ]]`
- **Multi-line comments**:
  ```yoruba
  [[
  This is a multi-line comment.
  It can span multiple lines.
  ]]
  ```
Comments can also appear any where, as long as it doesn't break keywords
  ```yoruba
    [[at start]]jeki x [[between keywords]] je 50 [[at end]]
  ```

### **Whitespaces**
Whitespaces has no syntantical meaning in `yoruba`
  ```yoruba
    [[This is the same as...]]
    jeki x je 0 titi x baje 100 se ko x jeki x je x + 1 pari ko "x waje", x pa

    [[...doing this]]
    jeki x je 0
    titi x baje 100 se
        ko x
        jeki x je x + 1
    pari
    ko pe pelu "x waje ", x pa
  ```

## **Builtin Functions**
Currently, `yoruba` has just two built in functions
- ko (write): accepts any number of arguments and writes them out too STDOUT
- ka (read): accepts one or no arguments. If an arguments is supplied, it writes out the supplied argument then it reads from the STDIN
                if not it only reads from the STDIN

## **Examples**
Here are some examples to help you get started with **yoruba**:

### **Simple Addition**
```yoruba
ise add_numbers a, b se
    da a + b pada
pari

jeki result je pe add_numbers pelu 5, 10 pa
pe ko pelu {"The result is", result} pa
```

### **Factorial**
```yoruba
jeki number je pe ka pelu "Enter a number: "

jeki factorial je 1
jeki i je 1
titi i baje number se
    jeki factorial je factorial * i
pari
pe ko pelu {"Factorial of", number, "is", factorial} pa
```
## **Contributing**
Contributions are welcome! Please read the [CONTRIBUTING.md](CONTRIBUTING.md) file for guidelines on how to contribute to the project.

## **License**
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## **Future Additions**
 - For and While loops
 - Data structures (hash maps, list)
 - More built in functions

## **Resources**
- [Thorsten Ball's - Writing an interpreter in go](https://interpreterbook.com)
- [Vaughan R. Pratt Proceedings on Top down operator precedence](https://dl.acm.org/doi/10.1145/512927.512931)
- [Pratt Parsers Expression Parsing Made Easy](www.oilshell.org/blog/2016/11/01.html)
- [Douglas Crockford's video on programming language design](https://youtu.be/Nlqv6NtBXcA?si=dwdxClDaQWpTqGrL)
- [Wren Lang](https://github.com/wren-lang/wren)
---
