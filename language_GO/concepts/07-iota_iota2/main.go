package main // main package for an executable

import "fmt" // import fmt for printing

type ByteSize int // define a custom byte-size type

const ( // start first iota block
	_ = iota // skip zero
	a        // iota = 1
	b        // iota = 2
	c        // iota = 3
	d        // iota = 4
	e        // iota = 5
	f        // iota = 6
) // end first iota block

func main() { // program entry point
	iota1() // run first demo
	iota2() // run second demo
} // end main

func iota1() { // function created to print iota values and shifts
	fmt.Print(a, "\n")                   // print a with newline
	fmt.Print(b, "\n")                   // print b with newline
	fmt.Print(c, "\n")                   // print c with newline
	fmt.Print(d, "\n")                   // print d with newline
	fmt.Print(e, "\n")                   // print e with newline
	fmt.Print(f, "\n \n \n")             // print f and extra blank lines
	fmt.Printf("%d \t %b\n", 1, 1)       // print 1 in decimal and binary
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a) // shift 1 by a
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b) // shift 1 by b
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c) // shift 1 by c
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d) // shift 1 by d
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e) // shift 1 by e
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f) // shift 1 by f
} // end iota1

const ( // start byte-size iota block
	_           = iota             // ignore first value
	KB ByteSize = 1 << (10 * iota) // kilobyte as 1<<10
	MB ByteSize = 3 << (10 * iota) // megabyte using 3 as base
	GB ByteSize = 2 << (10 * iota) // gigabyte using 2 as base
	TB ByteSize = 4 << (10 * iota) // terabyte using 4 as base
	PB ByteSize = 5 << (10 * iota) // petabyte using 5 as base
	EB ByteSize = 6 << (10 * iota) // exabyte using 6 as base
) // end byte-size block

func iota2() { // function created to print byte-size values
	fmt.Printf("%d \t\t\t %b\n", KB, KB) // print KB in decimal and binary
	fmt.Printf("%d \t\t %b\n", MB, MB)   // print MB in decimal and binary
	fmt.Printf("%d \t\t %b\n", GB, GB)   // print GB in decimal and binary
	fmt.Printf("%d \t\t %b\n", TB, TB)   // print TB in decimal and binary
	fmt.Printf("%d \t %b\n", PB, PB)     // print PB in decimal and binary
	fmt.Printf("%d \t %b\n", EB, EB)     // print EB in decimal and binary
} // end iota2
