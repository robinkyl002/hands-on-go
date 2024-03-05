# Generics
Return to overview of notes: [notes.md](../notes.md)

## Type Parameters
Code [here](type-parameters/begin/main.go)

- Support different types for same behavior
  - i.e. adding two numbers together, whether they be integers or floats
- Structure as follows
  - `func methodName[T varType | otherVarType](input1, input2 T) T`
- when calling the method, you can explicitly indicate what type is being input, but it is not necessary
  - With type indicated: `sum[int](3, 2)`
  - Without type indicated: `sum(3, 2)`
- Use `~` next to types in parameterization to allow for custom types based on the primitive types
  - Defines the primitive types as underlying types for the custom types
- Types can have parameter lists
  - Custom list that will work with any parameter type

## Interface Constraints
Code [here](type-sets/begin/main.go)

## Standard Library support
Code [here](standard-library/begin/main.go)