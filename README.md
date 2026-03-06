# go-jellyfish

A Go port of the Python [jellyfish](https://github.com/jamesturk/jellyfish) library for phonetic encoding and string comparison.

## Install

```
go get github.com/xyz27900/go-jellyfish
```

## Usage

```go
package main

import (
	"fmt"

	jf "github.com/xyz27900/go-jellyfish"
)

func main() {
	// String similarity
	fmt.Println(jf.Jaro("martha", "marhta"))           // 0.9444...
	fmt.Println(jf.JaroWinkler("martha", "marhta"))     // 0.9611...

	// Phonetic encoding
	fmt.Println(jf.Metaphone("christopher"))            // KRSTFR
	fmt.Println(jf.Metaphone("psychological"))          // SXLJKL
}
```

## API

| Function                           | Description                                                     |
|------------------------------------|-----------------------------------------------------------------|
| `Jaro(a, b string) float64`        | Jaro similarity (0 = no match, 1 = identical)                   |
| `JaroWinkler(a, b string) float64` | Jaro-Winkler similarity (boosts strings with matching prefixes) |
| `Metaphone(s string) string`       | Metaphone phonetic code                                         |

## Differences from the upstream Go port

This port fixes two bugs present in [jamesturk/go-jellyfish](https://github.com/jamesturk/go-jellyfish):

- **Metaphone**: W before a vowel was incorrectly dropped (checked the wrong lookahead position)
- **Jaro-Winkler**: A `len > 3` guard prevented the Winkler prefix boost on short strings — Python jellyfish has no such guard

Both fixes are verified against Python jellyfish via `go test -tags python` (requires `pip install jellyfish`).

## License

MIT — see [LICENSE](LICENSE).
