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
  + Type

+ Generics
  + Interfaces has certain limitations of type-safety, overcomplicating codebase and a couple more
  + Generics come into picture to help us from that. Generics in Go are a feature that allows you to define functions and data structures that can operate on different types without being explicitly tied to specific ones. It enhances code reusability and provides better type safety in comparison to Interfaces.
  ```
  func Add[T int|string|float64](a, b T) T {
    return a+b
  }
  ```
  + Flexibility: With generics, you can write a single function that can work with different types, without needing to overload or create multiple functions for each type

+ Collections
  + `Array`: Works pretty much like how arrays work in C++, Java, except for the syntanctical differences ofcourse.
    + Fixed Sized DataStructures with specific length
    + Slices on the other hand are dynamically sized, They can grow and shrink as needed! 
  + `Slice`: Very very interesting concept in Go. 
    + Basically works more-or-less like slices in Python but the catch here is, slices are essentially a view into an array. 
    + When you create a slice, Go manages the underlying array automatically.
  + `Using Slices to build Dynamic Array` :
    + You operate on the slice, and if you append to it, Go may allocate a new array if the current on can't accommodate additional elements
    + `append()` function can be used to add elements to a slice. If slice exceeds its capacity, Go allocates a new larger array and copies the content over, which is managed behind the scenes.
  + `Memory Management` : Go Takes care of memory when slices grow. 
  + `Removing Elements`: To remove an element from the slice, you can create a new slice that omits the unwanted elements. 
  + `Performance`: Since slices are resizable, they incorporate increase flexibility compared to arrays. 
  + `Reference behaviour`: Slices are not references by default. If you want to pass a slice to a function reference, you'll have to use the pointer semantics explicitly.
  + #### Maps:
    + A key value data store! Also, provides mutation capabilities and more flexibility with data types as compared to Struct which is of a fixed type pre-defined on compilation.
    + 