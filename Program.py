import struct

class VM:
    class OpCodes:
        IL_ADD = 0
        IL_SUB = 1
        IL_MUL = 2
        IL_DIV = 3
        IL_MOD = 4
        IL_XOR = 5
        IL_INTEGER = 6
        IL_PRINT = 7
        IL_RET = 8

    def __init__(self, bytecode) -> None:
        self.stack = []
        self.opcodes = []
        self.integers = []
        i = 0
        while i < len(bytecode):
            opcode = self.convert_to_opcode(bytecode[i])
            self.opcodes.append(opcode)
            i += 1
            if opcode == self.OpCodes.IL_INTEGER:
                value = struct.unpack_from('i', bytecode, i)[0]
                self.integers.append(value)
                i += 4

    def execute(self) -> None:
        int_index = 0
        for opcode in self.opcodes:
            if opcode == self.OpCodes.IL_ADD:
                b = self.stack.pop()
                a = self.stack.pop()
                self.stack.append(a + b)
            elif opcode == self.OpCodes.IL_SUB:
                b = self.stack.pop()
                a = self.stack.pop()
                self.stack.append(a - b)
            elif opcode == self.OpCodes.IL_MUL:
                b = self.stack.pop()
                a = self.stack.pop()
                self.stack.append(a * b)
            elif opcode == self.OpCodes.IL_DIV:
                b = self.stack.pop()
                a = self.stack.pop()
                self.stack.append(a // b)
            elif opcode == self.OpCodes.IL_MOD:
                b = self.stack.pop()
                a = self.stack.pop()
                self.stack.append(a % b)
            elif opcode == self.OpCodes.IL_XOR:
                b = self.stack.pop()
                a = self.stack.pop()
                self.stack.append(a ^ b)
            elif opcode == self.OpCodes.IL_INTEGER:
                self.stack.append(self.integers[int_index])
                int_index += 1
            elif opcode == self.OpCodes.IL_PRINT:
                print("Result:", self.stack.pop())
            elif opcode == self.OpCodes.IL_RET:
                return
            else:
                raise Exception("Invalid OPCode")

    def convert_to_opcode(self, byte: bytes) -> OpCodes:
        if byte == 0:
            return self.OpCodes.IL_ADD
        elif byte == 1:
            return self.OpCodes.IL_SUB
        elif byte == 2:
            return self.OpCodes.IL_MUL
        elif byte == 3:
            return self.OpCodes.IL_DIV
        elif byte == 4:
            return self.OpCodes.IL_MOD
        elif byte == 5:
            return self.OpCodes.IL_XOR
        elif byte == 6:
            return self.OpCodes.IL_INTEGER
        elif byte == 7:
            return self.OpCodes.IL_PRINT
        elif byte == 8:
            return self.OpCodes.IL_RET
        else:
            raise Exception("Invalid OPCode")

if __name__ == "__main__":
    # Example bytecode: push 12, push 13, add, push 5, add, print, return
    bytecode = bytearray()
    bytecode.append(6)  # IL_INTEGER
    bytecode.extend(struct.pack('i', 12))  # 12
    bytecode.append(6)  # IL_INTEGER
    bytecode.extend(struct.pack('i', 13))  # 13
    bytecode.append(0)  # IL_ADD
    bytecode.append(6)  # IL_INTEGER
    bytecode.extend(struct.pack('i', 5))  # 5
    bytecode.append(0)  # IL_ADD
    bytecode.append(7)  # IL_PRINT
    bytecode.append(8)  # IL_RET

    VM(bytecode).execute()
