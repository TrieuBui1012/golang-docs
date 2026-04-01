Great — here’s both.

# Go variable declarations cheat sheet

## `var` declaration

Use `var` to declare variables.

```go
var i int
var x = 10
var a, b string = "hi", "go"
var p, q float64
```

### Rules

* With a type but no value → gets **zero value**
* With a value but no type → Go **infers** the type
* Can be used **inside or outside functions**

### Zero values

* `int` → `0`
* `float64` → `0`
* `bool` → `false`
* `string` → `""`
* slices/maps/pointers/interfaces/channels/functions → `nil`

Example:

```go
var n int      // 0
var ok bool    // false
var s string   // ""
var nums []int // nil
```

---

## `:=` short declaration

```go
x := 10
name := "Go"
a, b := 1, 2
```

### Rules

* Only works **inside functions**
* Must have an initializer
* Type is always **inferred**
* Equivalent to `var x = ...`

---

## `=` vs `:=`

```go
x := 1 // declare x
x = 2  // assign new value
```

* `:=` declares
* `=` assigns

---

## Redeclaration with `:=`

Allowed only when:

* variable already exists in the **same block**
* same type
* at least **one non-blank name is new**

```go
a, b := 1, 2
a, c := 3, 4 // legal: a reused, c new
```

Not allowed:

```go
a := 1
a := 2 // illegal
```

Use:

```go
a = 2
```

---

## Blank identifier

Use `_` to ignore values.

```go
_, ok := m["key"]
```

Common with:

* map lookups
* multiple return values

---

## `nil` rule

This is illegal:

```go
var x = nil
```

Because Go cannot infer the type.

Legal:

```go
var p *int = nil
var s []int = nil
```

---

## Package level vs local

### Package level

Use `var`:

```go
var Version = "1.0"
```

### Inside function

Usually use `:=`:

```go
func main() {
    x := 5
}
```

---

## Quick rule of thumb

Use `var` for:

* package-level variables
* explicit types
* zero-value declarations

Use `:=` for:

* short local variables inside functions