# Boolean values

Boolean values are those which can be assigned true or false and has the type bool with it.

```go
package main
 
import (
    "fmt"
)
 
func main() {
    var bVal bool   // default is false
    fmt.Printf("bVal: %v\n", bVal)
}
```

In the code above “bVal” is not initialized and thus has a zero-value. The zero-value for a boolean is false.

## Boolean operators in Go

Boolean operators are those operators that compare one value to others.

Below shown six different boolean operators that evaluate to bool.

```go
package main
 
import (
    "fmt"
)
 
func main() {
    // boolean operators
    a := 3
    b := 4
    fmt.Println(a == b)  // false
    fmt.Println(a != b)  // true
    fmt.Println(a < b)   // true
    fmt.Println(a > b)   // false
    fmt.Println(a >= b)  // false
    fmt.Println(a <= b)  // true
}
```

## Conversion from strings

Strings can be converted to boolean and vice-versa using the strconv package as shown below.

```go
package main
 
import (  
    "fmt"
    "strconv"
)
 
func main() { 
    var a bool
     
    s := "True"  // s can be 1, True, TRUE, False, false, 0 etc
    a, _ = strconv.ParseBool(s)  // observe that second value returned is ignored
                                     // here, because it returns true if it can be converted
                                     // else it resturns false
      
    fmt.Println(a) // prints true
}
```

## Strings from a boolean

A boolean can be converted to a string as well using the strconv package.

```go
package main
 
import (  
    "fmt"
    "strconv"
)
 
func main() { 
    var a string
    b := false
    a = strconv.FormatBool(b)
    fmt.Println(a)            // prints "false"
}
```
