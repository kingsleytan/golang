# Packages

- Start with `package main`
- [Package `rand`](https://golang.org/pkg/math/rand/)
- [cryto/rand](https://golang.org/pkg/crypto/rand/#example_Read)

## Note:
- rename `math.pi` to `math.Pi`, careful on calling function by capital letters.
- Short variable declarations. `:=` only can be used within a function.
- Variables without initialization are given zero value, `0`, `false` or `""`.
- Declaring a variable without specifying type, will be inferred from the value on the right hand side.
- `if`, `for` loop are surrounded by curly braces `{}`.
- `defer` statement defers the execution of a function until the surrounding function returns.