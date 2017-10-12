# Learning Go
A collection of various examples of the [Go Programming Language](https://golang.org/).

## Trying It

Most of the examples (with the exception of the unit tests) can be run via `go run`, eg. `go run slices/leftrotation.go`.

## Topics

### Slices

* [leftrotation.go](/slices/leftrotation.go) demostrates two algorithms for rotating an array (or slice, in this case) to the left by an abitrary number of positions. See the [/slices](/slices) directory for more examples with slices.

### Bitwise operators

* [parity.go](/bits/parity.go) determines the parity of the given integer, returning 1 if the number of ones in the binary string is odd, other wise 0 if even. See the [/bits](/bits) directory for more bitwise stuff.

### Maps

* [twosum.go](/maps/twosum.go) uses a map to solve the two-sum problem, which is to find which two (if any) numbers in an array sum up to the given total. See the [/maps](/maps) directory for more map shenanigans.

### Strings

* [compress.go](/strings/compress.go) takes a string and compresses it using [run-length encoding](https://en.wikipedia.org/wiki/Run-length_encoding). See the [/strings](/strings) directory for more string-based functions.

### Structs & Interfaces

* [operate.go](/structs/operate.go) demostrates polymorphism through an abstract `Vehicle` class [vcl.go](/structs/vcl/vcl.go). Try out [eat.go](/structs/eat.go) and [measure.go](/structs/measure.go) to see other examples of the structs in [/structs](/structs) in action. 

### Goroutines

* [race.go](/goroutines/race.go) will spawn a handful of goroutines to show how concurrency is handled in go. See the [/goroutines](/goroutines) directory for more.

### Channels

* [sync.go](/channels/sync.go) uses goroutines and channels to wait for responses from two (simulated) external dependencies. See the [/channels](/channels) directory for more on how to use channels in Go. 

### Unit tests
* [sum.go](/tests/sum.go) demostrates a simple unit test suite. First `cd tests`, then run `go test` to run the test file in that directory. Optionally, add the `-v` flag at the end to get a verbose pass/fail message for each test.
