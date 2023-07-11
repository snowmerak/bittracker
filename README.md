# bittracker

bit tracking from byte slice in go

## Install

```bash
go get github.com/snowmerak/bittracker
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/snowmerak/bittracker"
)

func main() {
	bt := bittracker.NewBitTracker([]byte{0b01010101, 0b11111111, 0b00000000, 0b10101010})

	result := bt.GetRange(5, 26)
	fmt.Printf("%08b\n", result)
}
```

```bash
0b10111111 0b11100000 0b00010000
```
