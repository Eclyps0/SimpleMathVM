# SimpleMathVM

## .NET Core 8

```cs
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
```
