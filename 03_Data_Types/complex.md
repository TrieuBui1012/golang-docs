# Complex Numbers

There are two complex types in Go. The complex64 and the complex128.
## Initializing Complex Numbers

Initializing complex numbers is really easy. You can use the constructor or the initialization syntax as well.

```go
c1 := complex(10, 11) // constructor init
c2 := 10 + 11i        // complex number init syntax
```

## Parts of a Complex Number

There are two parts of a complex number. The real and imaginary part. We use functions to get those.

```go
c := complex(23, 31)
realPart := real(c)    // gets real part
imagPart := imag(c)    // gets imaginary part
```

## Operations on a Complex Number

A complex variable can do any operation like addition, subtraction, multiplication, and division.

Let’s see an example to perform mathematical operations on the complex numbers.

```go
package main
 
import (
    "fmt"
)
 
func main() {
    c1 := complex(2, 3)
    c2 := 4 + 5i // complex initializer syntax a + ib
    c3 := c1 + c2 // addition just like other variables
    fmt.Println("Add: ", c3) // prints "Add: (6+8i)"
    re := real(c3)           // get real part
    im := imag(c3)           // get imaginary part
    fmt.Println(re, im)      // prints 6 8
}
```