package main

import "fmt"

type ByteSize int

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB ByteSize = 3 << (10 * iota)
	GB ByteSize = 2 << (10 * iota)
	TB ByteSize = 4 << (10 * iota)
	PB ByteSize = 5 << (10 * iota)
	EB ByteSize = 6 << (10 * iota)
)

func main() {
	fmt.Printf("%d \t\t\t %b\n", KB, KB)
	fmt.Printf("%d \t\t %b\n", MB, MB)
	fmt.Printf("%d \t\t %b\n", GB, GB)
	fmt.Printf("%d \t\t %b\n", TB, TB)
	fmt.Printf("%d \t %b\n", PB, PB)
	fmt.Printf("%d \t %b\n", EB, EB)
}
