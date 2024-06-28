class VM {
    constructor(bytecode) {
        this.OpCodes = {
            IL_ADD: 0,
            IL_SUB: 1,
            IL_MUL: 2,
            IL_DIV: 3,
            IL_MOD: 4,
            IL_XOR: 5,
            IL_INTEGER: 6,
            IL_PRINT: 7,
            IL_RET: 8
        };

        this.stack = [];
        this.opcodes = [];
        this.integers = [];

        let i = 0;
        while (i < bytecode.length) {
            const opcode = this.convertToOpcode(bytecode[i]);
            this.opcodes.push(opcode);
            i++;
            if (opcode === this.OpCodes.IL_INTEGER) {
                const value = this.bytesToInt(bytecode.slice(i, i + 4));
                this.integers.push(value);
                i += 4;
            }
        }
    }

    execute() {
        let intIndex = 0;
        for (let i = 0; i < this.opcodes.length; i++) {
            switch (this.opcodes[i]) {
                case this.OpCodes.IL_ADD:
                    {
                        const b = this.stack.pop();
                        const a = this.stack.pop();
                        this.stack.push(a + b);
                        break;
                    }
                case this.OpCodes.IL_SUB:
                    {
                        const b = this.stack.pop();
                        const a = this.stack.pop();
                        this.stack.push(a - b);
                        break;
                    }
                case this.OpCodes.IL_MUL:
                    {
                        const b = this.stack.pop();
                        const a = this.stack.pop();
                        this.stack.push(a * b);
                        break;
                    }
                case this.OpCodes.IL_DIV:
                    {
                        const b = this.stack.pop();
                        const a = this.stack.pop();
                        this.stack.push(Math.floor(a / b));
                        break;
                    }
                case this.OpCodes.IL_MOD:
                    {
                        const b = this.stack.pop();
                        const a = this.stack.pop();
                        this.stack.push(a % b);
                        break;
                    }
                case this.OpCodes.IL_XOR:
                    {
                        const b = this.stack.pop();
                        const a = this.stack.pop();
                        this.stack.push(a ^ b);
                        break;
                    }
                case this.OpCodes.IL_INTEGER:
                    {
                        this.stack.push(this.integers[intIndex]);
                        intIndex++;
                        break;
                    }
                case this.OpCodes.IL_PRINT:
                    {
                        console.log("Result:", this.stack.pop());
                        break;
                    }
                case this.OpCodes.IL_RET:
                    {
                        return;
                    }
                default:
                    throw new Error("Invalid OPCode");
            }
        }
    }

    convertToOpcode(byte) {
        switch (byte) {
            case 0:
                return this.OpCodes.IL_ADD;
            case 1:
                return this.OpCodes.IL_SUB;
            case 2:
                return this.OpCodes.IL_MUL;
            case 3:
                return this.OpCodes.IL_DIV;
            case 4:
                return this.OpCodes.IL_MOD;
            case 5:
                return this.OpCodes.IL_XOR;
            case 6:
                return this.OpCodes.IL_INTEGER;
            case 7:
                return this.OpCodes.IL_PRINT;
            case 8:
                return this.OpCodes.IL_RET;
            default:
                throw new Error("Invalid OPCode");
        }
    }

    bytesToInt(bytes) {
        return new DataView(new Uint8Array(bytes).buffer).getInt32(0, true);
    }
}

function main() {
    // Example bytecode: push 12, push 13, add, push 5, add, print, return
    const bytecode = [];
    bytecode.push(6); // IL_INTEGER
    bytecode.push(...new Uint8Array(new Int32Array([12]).buffer)); // 12
    bytecode.push(6); // IL_INTEGER
    bytecode.push(...new Uint8Array(new Int32Array([13]).buffer)); // 13
    bytecode.push(0); // IL_ADD
    bytecode.push(6); // IL_INTEGER
    bytecode.push(...new Uint8Array(new Int32Array([5]).buffer)); // 5
    bytecode.push(0); // IL_ADD
    bytecode.push(7); // IL_PRINT
    bytecode.push(8); // IL_RET

    new VM(bytecode).execute();;
}
main();
