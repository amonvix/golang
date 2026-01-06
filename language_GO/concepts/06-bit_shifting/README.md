# Bit Shifting and Binary Representation in Go

This lesson demonstrates how **bit shifting** works in Go and how integer values are represented in **binary format**.

By shifting a single bit to the left, we can clearly observe how powers of two are formed and how binary values map directly to decimal numbers.

---

## Concepts Covered

- Bitwise left shift operator (`<<`)
- Binary representation of integers
- Relationship between bit shifting and powers of two
- Formatted output with `%d` and `%b`
- Compile-time evaluation of shift operations

---

## Purpose

The goal of this exercise is to understand:

- How left bit shifting works
- How binary values translate to decimal numbers
- Why shifting left multiplies a number by powers of two
- How Go represents integers internally

---

## Explanation

In Go, the left shift operator (`<<`) moves the bits of a number to the left by a specified number of positions.

Shifting the value `1`:

- `1 << 0` → `1`
- `1 << 1` → `2`
- `1 << 2` → `4`
- `1 << 3` → `8`
- `1 << n` → `2ⁿ`

The program prints each shifted value in both **decimal (`%d`)** and **binary (`%b`)** formats, making the transformation explicit.

---

## Example Behavior

Each line of output shows:

- The decimal value after shifting
- The corresponding binary representation

This allows a direct comparison between numeric value growth and bit movement.

---

## Key Takeaways

- Bit shifting is a low-level operation with predictable behavior
- Left shifting multiplies a value by powers of two
- Binary representation makes bitwise operations intuitive
- Shifts are evaluated at compile time in Go
- `%b` is a powerful tool for visualizing bit-level behavior

---

## Notes

This example is intentionally repetitive to reinforce pattern recognition.  
Once the pattern is clear, the behavior becomes obvious.

Bits don’t lie — they just move.
