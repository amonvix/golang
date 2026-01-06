# For Range and Data Structure Iteration in Go

This chapter introduces the **`for range`** loop and shows how Go iterates over common data structures such as **slices** and **maps**.

`for range` provides a clear and idiomatic way to traverse collections without manual index management.

---

## Concepts Covered

- `for range` loops
- Slice iteration
- Map iteration
- Index and value binding
- Key-value traversal
- Idiomatic iteration patterns
- Readable data processing

---

## Purpose

The goal of this chapter is to understand:

- How Go iterates over collections
- How `for range` exposes indices, keys, and values
- How iteration differs between slices and maps
- Why `for range` improves clarity and safety
- How Go encourages explicit data traversal

This chapter completes the iteration section of the workflow.

---

## For Range Overview

The `for range` loop iterates over elements in a collection.

Depending on the data structure, it provides:

- An **index and value** for slices
- A **key and value** for maps

Go handles traversal mechanics internally, allowing the code to focus on intent.

---

## Iterating Over Slices

When ranging over a slice:

- The first variable receives the index
- The second variable receives the value at that index

This pattern supports sequential data processing with minimal boilerplate.

---

## Iterating Over Maps

When ranging over a map:

- The first variable receives the key
- The second variable receives the associated value

Iteration order is **not guaranteed**, reinforcing that maps represent unordered data.

---

## Example Behavior

- The program iterates over a slice and prints index-value pairs
- The program iterates over a map and prints key-value pairs
- Each iteration produces a single, well-defined output
- No manual indexing or bounds checks are required

---

## Key Takeaways

- `for range` is the idiomatic way to iterate in Go
- Slices provide ordered iteration
- Maps provide key-value access without ordering guarantees
- Explicit variable binding improves readability
- Iteration logic stays focused on data, not mechanics

---

## Notes

`for range` reflects Go’s philosophy of simplicity and clarity.

Once control flow is established,  
idiomatic iteration turns logic into readable data processing.
