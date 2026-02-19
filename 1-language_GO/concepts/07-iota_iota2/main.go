package main

import "fmt"

// ByteSize represents a size in bytes, supporting large units via bit shifting.
type ByteSize int

const (
	_ = iota // skip zero value to start enumeration from 1
	a        // iota = 1
	b        // iota = 2
	c        // iota = 3
	d        // iota = 4
	e        // iota = 5
	f        // iota = 6
)

// main is the entry point of the executable, invoking demonstration functions.
func main() {
	iota1()
	iota2()
}

// iota1 prints sequential iota values and demonstrates bit shifting using those values.
func iota1() {
	fmt.Print(a, "\n")
	fmt.Print(b, "\n")
	fmt.Print(c, "\n")
	fmt.Print(d, "\n")
	fmt.Print(e, "\n")
	fmt.Print(f, "\n \n \n")

	// Print decimal and binary representations of 1 and 1 shifted by iota constants.
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)
}

const (
	_           = iota               // ignore zero value to align sizes starting at 1 KB
	KB ByteSize = 1 << (10 * iota)  // kilobyte as 2^10 bytes
	MB ByteSize = 3 << (10 * iota)  // megabyte scaled by 3 * 2^(10*iota)
	GB ByteSize = 2 << (10 * iota)  // gigabyte scaled by 2 * 2^(10*iota)
	TB ByteSize = 4 << (10 * iota)  // terabyte scaled by 4 * 2^(10*iota)
	PB ByteSize = 5 << (10 * iota)  // petabyte scaled by 5 * 2^(10*iota)
	EB ByteSize = 6 << (10 * iota)  // exabyte scaled by 6 * 2^(10*iota)
)

// iota2 prints the defined ByteSize constants in decimal and binary formats.
// Note that MB through EB use non-standard multipliers for demonstration.
func iota2() {
	fmt.Printf("%d \t\t\t %b\n", KB, KB)
	fmt.Printf("%d \t\t %b\n", MB, MB)
	fmt.Printf("%d \t\t %b\n", GB, GB)
	fmt.Printf("%d \t\t %b\n", TB, TB)
	fmt.Printf("%d \t %b\n", PB, PB)
	fmt.Printf("%d \t %b\n", EB, EB)
}