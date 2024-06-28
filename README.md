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

## Python 3.x

```py
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
```

## JavaScript
```js
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
```

## Golang
```go
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
```
