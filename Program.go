package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type OpCode byte

const (
	IL_ADD OpCode = iota
	IL_SUB
	IL_MUL
	IL_DIV
	IL_MOD
	IL_XOR
	IL_INTEGER
	IL_PRINT
	IL_RET
)

type VM struct {
	stack    []int32
	opcodes  []OpCode
	integers []int32
}

func NewVM(bytecode []byte) *VM {
	vm := &VM{}
	var i int
	for i < len(bytecode) {
		opcode := OpCode(bytecode[i])
		vm.opcodes = append(vm.opcodes, opcode)
		i++
		if opcode == IL_INTEGER {
			var value int32
			buf := bytes.NewReader(bytecode[i : i+4])
			binary.Read(buf, binary.LittleEndian, &value)
			vm.integers = append(vm.integers, value)
			i += 4
		}
	}
	return vm
}

func (vm *VM) Execute() {
	var intIndex int
	for _, opcode := range vm.opcodes {
		switch opcode {
		case IL_ADD:
			b := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]
			a := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]
			vm.stack = append(vm.stack, a+b)
		case IL_SUB:
			b := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]
			a := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]
			vm.stack = append(vm.stack, a-b)
		case IL_MUL:
			b := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]
			a := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]
			vm.stack = append(vm.stack, a*b)
		case IL_DIV:
			b := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]
			a := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]
			vm.stack = append(vm.stack, a/b)
		case IL_MOD:
			b := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]
			a := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]
			vm.stack = append(vm.stack, a%b)
		case IL_XOR:
			b := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]
			a := vm.stack[len(vm.stack)-1]
			vm.stack = vm.stack[:len(vm.stack)-1]
			vm.stack = append(vm.stack, a^b)
		case IL_INTEGER:
			vm.stack = append(vm.stack, vm.integers[intIndex])
			intIndex++
		case IL_PRINT:
			fmt.Println("Result:", vm.stack[len(vm.stack)-1])
			vm.stack = vm.stack[:len(vm.stack)-1]
		case IL_RET:
			return
		default:
			panic("Invalid OpCode")
		}
	}
}

func main() {
	// Example bytecode: push 12, push 13, add, push 5, add, print, return
	var bytecode []byte
	bytecode = append(bytecode, byte(IL_INTEGER))    // IL_INTEGER
	bytecode = append(bytecode, int32ToBytes(12)...) //12
	bytecode = append(bytecode, byte(IL_INTEGER))    // IL_INTEGER
	bytecode = append(bytecode, int32ToBytes(13)...) //13
	bytecode = append(bytecode, byte(IL_ADD))        // IL_ADD
	bytecode = append(bytecode, byte(IL_INTEGER))    //IL_INTEGER
	bytecode = append(bytecode, int32ToBytes(5)...)  // 5
	bytecode = append(bytecode, byte(IL_ADD))        // IL_ADD
	bytecode = append(bytecode, byte(IL_PRINT))      // IL_PRINT
	bytecode = append(bytecode, byte(IL_RET))        // IL_RET

	NewVM(bytecode).Execute() // 30
}

func int32ToBytes(i int32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, i)
	return buf.Bytes()
}
