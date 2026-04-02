# Floating-Point Numbers

## Float32 and Float64

```go
package main
 
import (
    "fmt"
)
 
func main() {
    var f float32 
    var f2 float64 
    f = 12.34567890123456
    f2 = 12.34567890123456
    fmt.Println(f, f2)  // prints "12.345679 12.34567890123456"
}
```

## Loss of Precision with Floating Point Numbers

Loss of precision will occur when a 64-bit floating-point number is converted to 32-bit float.

```go
package main
 
import (
    "fmt"
)
 
func main() {
    var f1 float32
    var f2 float64
     
    f2 = 1.234567890123
    f1 = float32(f2) 
     
    fmt.Println(f1)         // prints "1.2345679"
}
```
