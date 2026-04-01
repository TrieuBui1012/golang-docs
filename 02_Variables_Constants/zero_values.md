# Working with Constants in Go

## Overview

Constants in Go are immutable values known at compile time. They are useful for fixed values such as mathematical numbers, configuration limits, status codes, and enum-like values.

## Declaring Constants

### Basic constants

```go
const Pi = 3.14159

const (
    StatusOK    = 200
    StatusError = 500
)

const (
    Timeout time.Duration = 30 * time.Second
    MaxSize int           = 1024
    Debug   bool          = false
)
```

### Untyped constants

Untyped constants are flexible because Go can use them with compatible types as needed.

```go
const (
    Pi    = 3.14159
    Hello = "Hello, World"
    Yes   = true
)

var f float64 = Pi
var f32 float32 = Pi
var c complex128 = Pi
```

## Using `iota`

### Basic `iota`

`iota` is commonly used to generate related constant values.

```go
const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

### Advanced `iota` patterns

#### Bit shifting

```go
const (
    KB = 1 << (10 * iota)
    MB
    GB
    TB
)
```

#### Skipping values

```go
const (
    _ = iota
    DebugLevel = 1 << iota
    Info
    Warning
    Error
)
```

#### Offsets and formulas

```go
const (
    Offset = 2*iota + 1
    _
    Value
    _
    Result
)
```

## Constant Types

### Numeric constants

```go
const (
    MaxInt = 9223372036854775807
    MinInt = -9223372036854775808

    Pi = 3.14159265359
    E  = 2.71828182845

    I    = 1i
    TwoI = 2i
)
```

### String constants

```go
const (
    Prefix   = "go_"
    Suffix   = "_temp"
    EmptyStr = ""

    MultiLine = `Line 1
Line 2
Line 3`
)
```

### Boolean constants

```go
const (
    True  = true
    False = false
    Debug = true
)
```

## Constant Expressions

Constants can be built from other constant expressions.

```go
const (
    Hundred = 10 * 10
    Million = Hundred * Hundred * Hundred

    FullName = "John" + " " + "Doe"
    IsValid  = true && !false
)
```

## Best Practices

### Naming

Use clear, descriptive names.

```go
const (
    MaxConnections = 100
    DefaultTimeout = 30 * time.Second
    DatabaseURL    = "postgres://localhost:5432"
)
```

Avoid names that are vague or too short.

```go
const (
    Max = 100
    TO  = 30
    URL = "..."
)
```

### Group related constants

```go
const (
    StatusOK           = 200
    StatusCreated      = 201
    StatusAccepted     = 202
    StatusBadRequest   = 400
    StatusUnauthorized = 401
    StatusForbidden    = 403
)

const (
    MaxRetries     = 3
    RetryDelay     = 5 * time.Second
    RequestTimeout = 30 * time.Second
)
```

### Choose typed vs untyped carefully

Use typed constants when type safety matters.

```go
const (
    Timeout time.Duration = 30 * time.Second
    MaxSize int           = 1024 * 1024
)
```

Use untyped constants when flexibility is more useful.

```go
const (
    Pi    = 3.14159
    Debug = true
)
```

## Common Patterns

### Enumerated constants

```go
type Direction int

const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) String() string {
    return [...]string{"North", "East", "South", "West"}[d]
}
```

### Bitmask constants

```go
type Permission int

const (
    Read Permission = 1 << iota
    Write
    Execute
)

func hasPermission(p Permission) bool {
    return p&Write != 0
}
```

### Version constants

```go
const (
    Major = 1
    Minor = 2
    Patch = 3
)
```

To build a numeric or string version, combine those values in regular code.

## Performance Notes

- Constants are evaluated at compile time.
- Constant expressions do not add runtime overhead.
- Constants can make compiler optimizations easier.

## Common Mistakes

### Trying to modify a constant

```go
const MaxSize = 1024
// MaxSize = 2048 // compilation error

var maxSize = 1024
maxSize = 2048
```

### Using non-constant expressions

```go
// const Timeout = time.Now() // invalid
const Timeout = 30 * time.Second
```
