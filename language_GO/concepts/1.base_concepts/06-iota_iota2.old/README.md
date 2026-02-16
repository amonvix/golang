# Iota, Bit Shifting, and Byte Size Constants in Go

This lesson demonstrates how Go uses **`iota`** to generate sequential constants and how it combines naturally with **bit shifting** to model byte-size units in a clean, scalable way.

The example is divide into two parts:

1. Basic `iota` sequencing and bit shifting
2. Practical byte-size constants (KB → EB)

---

## Concepts Covered

- `iota` for sequential constants
- Skipping values in `iota` blocks
- Bitwise left shift operator (`<<`)
- Relationship between shifts and powers of two
- Custom named types
- Compile-time constant evaluation
- Modeling real-world units with constants

---

## Purpose

The goal of this exercise is to understand:

- How `iota` works across constant blocks
- How constants increment automatically
- How `iota` integrates with bit shifting
- How Go models memory sizes in a safe way and expressively
- Why constants remain flexible and type-safe

---

## Explanation

### Iota Basics

`iota` is a special identifier used inside `const` blocks.  
It starts at `0` and increments by `1` for each line.

It's possible to intentionally skip values to maintain precise control over numbering.

---

### Iota with Bit Shifting

Bit shifting pairs naturally with `iota`:

- `1 << a` represents `2ⁿ`
- Increasing `iota` increases the shift distance
- This creates predictable exponential growth

The example prints each value in **decimal** and **binary**, making the pattern explicit.

---

### Byte Size Constants

A second `iota` block models byte sizes:

- `KB = 1 << (10 * iota)`
- Each step increases by a power of `1024`
- Values scale cleanly from kilobytes to exabytes

A custom `ByteSize` type improves readability and intent without affecting performance.

---

## Example Behavior

- Sequential constants (`a` to `f`) increase automatically
- Bit shifts grow values exponentially
- Byte-size constants scale predictably
- The **compiler** resolves all values at compile time
- No runtime overhead or implicit conversions

---

## Key Takeaways

- `iota` is ideal for ordered, related constants
- Bit shifting expresses powers of two precisely
- Constants can represent massive values safely
- Custom types improve semantic clarity
- Go favors explicit patterns over hidden magic

---

## Notes

This pattern is widely used in real-world Go codebases for memory sizes, flags, and configuration values.

Simple syntax. Strong guarantees. No surprises.
