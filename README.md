# Go Learning

## I've categorised it into phases

### *Phase 1*

#### Includes the following topics

+ Variables and data types
+ Basic Syntax
+ Control Structure (If/Else, Switch Case)
+ The versatile For Loop
+ Pointers
+ Very basic recursion

### *Phase 2*

#### Includes slightly more advanced topics

+ Structs
  + Fields
  + Methods
  + Mutation Methods
  + Creation/Constructor Functions (This is more of a pattern)
  + Struct Embeddings (A similar thing like Inheritance)
    + Can be anonymous
    + Can be "key type" representation too both vary
  + Bufio Package for NewReader() to read long line inputs instead of using Scan with limited OneWord Input
  + Strings Package for TrimSuffix() and More Function like ToLower(), ParseFloat() etc.
  + encoding/json package for Json Conversion and the Marshal function
  + Struct Tags [involved within backticks after type definition in a field eg. `json:"hello"`]
  
+ Interfaces
  + Types that define a contract or a set of methods that a type must implement.
  + Allows for a level of abstraction and flexibility in code since you can write functions that accept arguments of different types as long as they implement the required methods.
  + `Implementation`: Any type that implements the methods defined in the interface is said to satisfy the interface.
  
  ```
    type Dog struct{}
    func (d Dog) Speak() string {
      return "Woof"
    }
    type Cat struct{}
    func (c Cat) Speak string {
      return "Meow"
    }
  ```

  + `Flexibility with Functions`: Interfaces allow functions to accept any type that satisfies the interface.
  + `Use Cases`;
    + Polymorphism : Enabling different types to be treated as the same type, allowing for flexible and reusable code.
    + Decoupling Code : Reducing dependencies between components in a system, making it easier to test and maintain.
    + Dependency Injection : Allowing for the specification of behaviors without needing to know concrete implementations.
