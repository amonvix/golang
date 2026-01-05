# For Loops and Iteration Patterns in Go

This chapter introduces the **`for` statement** in Go and shows how it unifies all looping constructs into a single, flexible mechanism.

Go replaces traditional `for`, `while`, and `do-while` loops with one consistent structure.

---

## Concepts Covered

- The `for` statement
- Counter-based loops
- Condition-based loops
- Infinite loops
- Loop control with `break`
- Loop control with `continue`
- Iteration patterns
- Range-based iteration over integers

---

## Purpose

The goal of this chapter is to understand:

- How Go models repetition
- How a single construct replaces multiple loop types
- How to control loop execution explicitly
- How to stop or skip iterations safely
- How iteration fits into program workflow

This chapter establishes repetition as a core control mechanism.

---

## For Loop Overview

Go provides only one looping keyword: `for`.

It supports three primary forms:

- Counter-based loops with initialization, condition, and post statement
- Condition-only loops (while-style)
- Infinite loops with explicit termination

This design reduces language complexity while preserving expressiveness.

---

## Iteration Patterns

The example demonstrates several common patterns:

- Iterating over a fixed range of values
- Looping based on a condition
- Breaking out of an infinite loop
- Skipping iterations using `continue`
- Filtering values during iteration

Each pattern emphasizes explicit control over execution.

---

## Example Behavior

- A loop counts through a fixed range
- A condition-based loop runs until a value changes
- An infinite loop exits using `break`
- A loop skips odd values using `continue`
- Only valid iterations produce output

Execution remains predictable and controlled.

---

## Key Takeaways

- Go uses a single loop construct
- Explicit control replaces implicit behavior
- `break` and `continue` shape loop execution
- Iteration logic remains readable and intentional
- Simple loops form the basis for complex workflows

---

## Notes

Iteration is the engine of most programs.

Once decisions exist,  
repetition turns logic into behavior.
