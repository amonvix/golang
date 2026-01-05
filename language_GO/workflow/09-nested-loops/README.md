# Nested Loops and Iteration Composition in Go

This chapter introduces **nested loops** and shows how Go composes multiple iterations to work with combinations, grids, and multi-dimensional logic.

Nested loops allow a program to iterate over one sequence **inside** another, expanding control flow in a structured and predictable way.

---

## Concepts Covered

- Nested `for` loops
- Outer and inner loop relationships
- Iteration order
- Index coordination
- Range-based nested loops
- Readability in multi-level iteration

---

## Purpose

The goal of this chapter is to understand:

- How nested loops execute
- How inner loops relate to outer loops
- How iteration order affects output
- How Go handles multi-level repetition
- How to structure nested logic clearly

This chapter builds directly on single-loop iteration concepts.

---

## Nested Loop Overview

A nested loop places one `for` loop inside another.

For each iteration of the **outer loop**, the **inner loop** runs from start to finish.

This structure allows programs to:

- Traverse grids
- Compare combinations
- Generate paired values
- Model two-dimensional workflows

---

## Range-Based Nested Loops

Go also allows nested loops using `range`.

This form improves readability when iteration depends on a fixed count rather than explicit counters.

Range-based nesting emphasizes intent while preserving full control over iteration.

---

## Example Behavior

- The outer loop advances one step at a time
- The inner loop completes all its iterations for each outer step
- Output reflects the current state of both loop indices
- The process repeats until all combinations are exhausted

Execution order remains consistent and deterministic.

---

## Key Takeaways

- Nested loops multiply iteration scope
- Inner loops restart for each outer iteration
- Clear structure prevents cognitive overload
- Range-based loops improve expressiveness
- Nested iteration enables complex workflows

---

## Notes

Nested loops increase expressive power,  
but they also increase complexity.

Structure and clarity matter more as depth grows.
