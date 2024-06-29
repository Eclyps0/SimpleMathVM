using System;
using System.Collections.Generic;
using System.Linq;

public class VM
{
    private enum OpCodes
    {
        IL_ADD, IL_SUB, IL_MUL, IL_DIV, IL_MOD, IL_XOR, IL_INTEGER, IL_PRINT, IL_RET
    }

    private Stack<int> stack = new Stack<int>();
    private List<OpCodes> opcodes;
    private List<int> integers;

    public VM(byte[] args)
    {
        opcodes = new List<OpCodes>();
        integers = new List<int>();

        int i = 0;
        while (i < args.Length)
        {
            OpCodes opcode = ConvertToOpcode(args[i]);
            opcodes.Add(opcode);
            i++;

            if (opcode == OpCodes.IL_INTEGER)
            {
                int value = BitConverter.ToInt32(args, i);
                integers.Add(value);
                i += 4;
            }
        }

    }

    public void Execute()
    {
        int intIndex = 0;
        for (int i = 0; i < opcodes.Count; i++)
        {
            switch (opcodes[i])
            {
                case OpCodes.IL_ADD:
                    {
                        int b = stack.Pop();
                        int a = stack.Pop();
                        stack.Push(a + b);
                        break;
                    }
                case OpCodes.IL_SUB:
                    {
                        int b = stack.Pop();
                        int a = stack.Pop();
                        stack.Push(a - b);
                        break;
                    }
                case OpCodes.IL_MUL:
                    {
                        int b = stack.Pop();
                        int a = stack.Pop();
                        stack.Push(a * b);
                        break;
                    }
                case OpCodes.IL_DIV:
                    {
                        int b = stack.Pop();
                        int a = stack.Pop();
                        stack.Push(a / b);
                        break;
                    }
                case OpCodes.IL_MOD:
                    {
                        int b = stack.Pop();
                        int a = stack.Pop();
                        stack.Push(a % b);
                        break;
                    }
                case OpCodes.IL_XOR:
                    {
                        int b = stack.Pop();
                        int a = stack.Pop();
                        stack.Push(a ^ b);
                        break;
                    }
                case OpCodes.IL_INTEGER:
                    {
                        stack.Push(integers[intIndex]);
                        intIndex++;
                        break;
                    }
                case OpCodes.IL_PRINT:
                    {
                        Console.WriteLine("Result: " + stack.Pop());
                        break;
                    }
                case OpCodes.IL_RET:
                    {
                        return;
                    }
                default:
                    throw new Exception("Invalid OPCode");
            }
        }
    }

    private OpCodes ConvertToOpcode(byte arg)
    {
        return arg switch
        {
            0 => OpCodes.IL_ADD,
            1 => OpCodes.IL_SUB,
            2 => OpCodes.IL_MUL,
            3 => OpCodes.IL_DIV,
            4 => OpCodes.IL_MOD,
            5 => OpCodes.IL_XOR,
            6 => OpCodes.IL_INTEGER,
            7 => OpCodes.IL_PRINT,
            8 => OpCodes.IL_RET,
            _ => throw new Exception("Invalid OPCode")
        };
    }
}

class Program
{
    static void Main()
    {
        // Example bytecode: push 12, push 153, add, print, return
        List<byte> bytecode = new List<byte> { };
        bytecode.Add(6); // IL_INTEGER 
        bytecode.AddRange(BitConverter.GetBytes(12)); // 12
        bytecode.Add(6); // IL_INTEGER 
        bytecode.AddRange(BitConverter.GetBytes(153)); // 153 
        bytecode.Add(0); // IL_ADD
        bytecode.Add(7); // IL_PRINT
        bytecode.Add(8); // IL_RET
        new VM(bytecode.ToArray()).Execute(); // 165
    }
}
