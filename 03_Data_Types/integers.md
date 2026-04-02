# Integers

## Signed integers in Go

- int8 (8-bit signed integer whose range is -128 to 127)
- int16 (16-bit signed integer whose range is -32768 to 32767)
- int32 (32-bit signed integer whose range is -2147483648 to 2147483647)
- int64 (64-bit signed integer whose range is -9223372036854775808 to 9223372036854775807) 

## Unsigned integers in Go

- uint8 (8-bit unsigned integer whose range is 0 to 255 )
- uint16 (16-bit unsigned integer whose range is 0 to 65535 )
- uint32 (32-bit unsigned integer whose range is 0 to 4294967295 )
- uint64 (64-bit unsigned integer whose range is 0 to 18446744073709551615 ) 

## Integer Overflow in Golang

If you assign a type and then use a number larger than the types range to assign it, it will fail. Below is a program trying just that.

```go
package main
 
import (
    "fmt"
)
 
func main() {
    var x uint8
        fmt.Println("Throws integer overflow")
    x = 267       // range of uint8 is 0-255
}
```

## Type conversion in Golang

If you convert to a type that has range lower than your current range, data loss will occur. We do typecast by directly using the name of the variable as a function to convert types.

```go
package main
 
import (
    "fmt"
)
 
func main() {
    var x int32
    var y uint32     // range 0 to 4294967295
    var z uint8      // range 0 to 255
    fmt.Println("Type Conversion")
    x = 26700
    y = uint32(x)       // data preserved because number is inside range
    z = uint8(x)        // data loss due to out of range conversion
    fmt.Println(y, z)   // prints 26700 76
}
```