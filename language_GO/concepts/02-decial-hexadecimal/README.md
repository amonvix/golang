# 02 – Decimal, Binary and Hexadecimal (fmt.Printf) ![Go](https://img.shields.io/badge/Language-Go-blue)

This example shows how Go allows the same integer value to be **represented and displayed in different numeric bases** using `fmt.Printf` format verbs.

The goal here is not advanced math, but understanding **how Go formats values in output**, which is essential for debugging, logging, and low-level reasoning.

## Code overview

```go
number := 42

fmt.Printf("42 as binary, %b \n", number)
fmt.Printf("42 as hexadecimal, %#x \n", number)
```

## Practical examples

- fmt.Printf("%v\n", number) // default value
- fmt.Printf("%T\n", number) // value type
- fmt.Printf("%#v\n", number) // Go-syntax representation

## Why this matters

In Go (and in computing in general), the value never changes — only the way you observe it does.

These representations are constantly used in:

- debugging

- logs

- bit flags

- permissions

- memory addresses

- protocols

Being able to read them correctly avoids confusion and subtle bugs.

## Important notes

- %b, %d, %x do not change the value — only how it is displayed

- The # flag in %#x adds the 0x prefix, making hexadecimal explicit

- fmt.Printf does not append a newline automatically (\n is manual)
