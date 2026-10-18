# Golang Programs

A categorized collection of **original, standalone Go learning examples** inspired by the topic progression in the [GeeksforGeeks 100 Days of Golang guide](https://www.geeksforgeeks.org/blogs/100-days-of-golang/).

This repository contains small programs for learning and practice—not full projects. Examples are independently written for this collection; they are not copied from the article.

## Requirements

- Go 1.22 or later recommended.
- A terminal and any editor (VS Code works well).

## How to run

Each example is its own small Go program. Open a program folder and run:

```bash
go run main.go
```

For example:

```bash
cd basics/01_hello_world
go run main.go
```

Run the examples individually. Do not run `go run .` from a category directory because it contains multiple independent `main` packages in separate example folders.

## Categories


### Basics

- [Hello World](basics/01_hello_world/main.go)
- [Variables and Constants](basics/02_variables_constants/main.go)
- [Basic Data Types](basics/03_data_types/main.go)
- [Type Conversion](basics/04_type_conversion/main.go)
- [Operators](basics/05_operators/main.go)
- [If Else](basics/06_if_else/main.go)
- [Switch Statement](basics/07_switch/main.go)
- [For Loop](basics/08_for_loop/main.go)
- [Break and Continue](basics/09_break_continue/main.go)
- [Functions](basics/10_functions/main.go)
- [Multiple Return Values](basics/11_multiple_return_values/main.go)
- [Variadic Function](basics/12_variadic_function/main.go)
- [Anonymous Function](basics/13_anonymous_function/main.go)
- [Closure](basics/14_closure/main.go)
- [Defer](basics/15_defer/main.go)
- [Pointers](basics/16_pointers/main.go)
- [String Operations](basics/17_strings/main.go)
- [Runes and Bytes](basics/18_runes_and_bytes/main.go)
- [Arrays](basics/19_arrays/main.go)
- [Slices](basics/20_slices/main.go)
- [Maps](basics/21_maps/main.go)
- [Error Handling](basics/22_error_handling/main.go)
- [Recursive Factorial](basics/23_recursion_factorial/main.go)
- [Fibonacci Sequence](basics/24_fibonacci/main.go)
- [Sort a Slice](basics/25_sort_slice/main.go)
- [Write and Read a File](basics/26_file_write_read/main.go)
- [Generics](basics/27_generics/main.go)
- [Goroutine and Channel](basics/28_goroutine_channel/main.go)

### Data Structures and Algorithms

- [Reverse an Array](data-structures/01_array_reverse/main.go)
- [Linear Search](data-structures/02_linear_search/main.go)
- [Binary Search](data-structures/03_binary_search/main.go)
- [Bubble Sort](data-structures/04_bubble_sort/main.go)
- [Stack](data-structures/05_stack/main.go)
- [Queue](data-structures/06_queue/main.go)
- [Singly Linked List](data-structures/07_singly_linked_list/main.go)
- [Binary Tree Inorder Traversal](data-structures/08_binary_tree_traversal/main.go)
- [Breadth-First Search](data-structures/09_breadth_first_search/main.go)
- [Depth-First Search](data-structures/10_depth_first_search/main.go)
- [Frequency Counter](data-structures/11_frequency_map/main.go)
- [Find Minimum and Maximum](data-structures/12_min_max/main.go)
- [Balanced Parentheses](data-structures/13_parentheses_check/main.go)
- [Set Using a Map](data-structures/14_set_using_map/main.go)

### OOP Concepts in Go

- [Structs](oops/01_structs/main.go)
- [Methods](oops/02_methods/main.go)
- [Pointer Receiver](oops/03_pointer_receiver/main.go)
- [Encapsulation with Unexported Fields](oops/04_encapsulation/main.go)
- [Interface](oops/05_interface/main.go)
- [Polymorphism with Interfaces](oops/06_polymorphism/main.go)
- [Composition and Embedded Types](oops/07_composition/main.go)
- [Type Assertion and Type Switch](oops/08_interface_type_switch/main.go)
- [Custom String Representation](oops/09_custom_stringer/main.go)
- [Interface-Based Dependency Injection](oops/10_dependency_injection/main.go)

## Learning notes

- Go is not a class-based language, but structs, methods, interfaces, and composition support many object-oriented design patterns.
- Examples favor the standard library and avoid external dependencies.
- Some examples print map entries in an unspecified order because Go map iteration order is not guaranteed.
- These are learning examples; review input validation and edge cases before reusing code in production.

## Source roadmap

The topic selection follows the fundamentals, functions, strings, arrays/slices/maps, structs/methods, pointers, interfaces, generics, and concurrency themes in the GeeksforGeeks guide. This is an original practice collection, not a complete reproduction of the 100-day article.
