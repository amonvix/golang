# Logical Operators in Go

This chapter introduces **logical operators** in Go and shows how they combine boolean expressions to form more expressive conditional logic.

Logical operators allow programs to evaluate multiple conditions together and make decisions based on combined rules.

---

## Concepts Covered

- Logical AND (`&&`)
- Logical OR (`||`)
- Logical NOT (`!`)
- Boolean expressions
- Short-circuit evaluation
- Combining comparisons in conditionals

---

## Purpose

The goal of this chapter is to understand:

- How Go evaluates logical expressions
- How multiple conditions interact in a single decision
- How logical operators refine conditional flow
- How short-circuit behavior affects execution

This chapter builds directly on basic conditionals introduced earlier.

---

## Logical Operators Overview

Go provides three logical operators:

- `&&` (AND)  
  Evaluates to true only if **both** operands are true

- `||` (OR)  
  Evaluates to true if **at least one** operand is true

- `!` (NOT)  
  Inverts the boolean value of an expression

These operators operate on boolean expressions and return boolean results.

---

## Short-Circuit Behavior

Go evaluates logical expressions **from left to right**.

With short-circuit evaluation:

- `&&` stops evaluating as soon as one condition is false
- `||` stops evaluating as soon as one condition is true

This behavior improves efficiency and prevents unnecessary evaluations.

---

## Example Behavior

- A condition checks whether multiple values fall below a reference point
- Another condition evaluates alternative numeric ranges
- A negation checks inequality explicitly
- Each conditional executes only when its logical expression evaluates to true

The program prints different messages depending on the outcome of each logical test.

---

## Key Takeaways

- Logical operators combine conditions into expressive rules
- Short-circuit evaluation affects execution order
- Clear boolean expressions improve readability
- Logical operators extend the power of conditionals
- Complex decisions start with simple boolean logic

---

## Notes

Logical operators turn simple comparisons into meaningful decision logic.

Before loops, concurrency, or branching structures,  
programs must learn how to reason.
