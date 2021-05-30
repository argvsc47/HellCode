package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
)

var program string
var iptr int
var compiled string
var lc int

func PTR_INC() {
	compiled += "inc rsi\n"
}

func PTR_DEC() {
	compiled += "dec rsi\n"
}

func REG_INC() {
	compiled += "inc byte [rsi]\n"
}

func REG_DEC() {
	compiled += "dec byte [rsi]\n"
}

func LOP_BEG() {
	compiled += fmt.Sprintf("loop%v\n", lc)
	compiled += "cmp byte [rsi], 0\n"
	compiled += fmt.Sprintf("jz end%v\n", lc)
	lc += 1
}

func LOP_END() {
	compiled += fmt.Sprintf("end%v\n", lc - 1)
}

func ASCII_I() {
	compiled += "call input\n"
}

func ASCII_O() {
	compiled += "call print\n"
}

func Compile() {
	compiled += ".section bss\nmem:resb 30000\n.section text\n"
	if strings.Contains(program, ".") {
		compiled += "print:\nxor rax, rax\ninc rax\nxor rdi, rdi\ninc rdi\nxor rdx, rdx\ninc rdx\nsyscall\nret\n"
	}
	if strings.Contains(program, ",") {
		compiled += "input:\nxor rax, rax\nxor rdi, rdi\nxor rdx, rdx\ninc rdx\nsyscall\nret\n"
	}
	for iptr = 0;iptr < len(program); iptr++ {
		switch string(program[iptr]) {
			case ">":
				PTR_INC()
			case "<":
				PTR_DEC()
			case "+":
				REG_INC()
			case "-":
				REG_DEC()
			case "[":
				LOP_BEG()
			case "]":
				LOP_END()
			case ",":
				ASCII_I()
			case ".":
				ASCII_O()
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid argmuent count.")
		os.Exit(1)
	}

	file := os.Args[1]
	data, error := ioutil.ReadFile(file)
	
	if error != nil {
		panic(error)
	}
	program = string(data)

	Compile()

	bin := []byte(compiled)
	error = ioutil.WriteFile(os.Args[2], bin, 0644)

	if error != nil {
		panic(error)
	}
}