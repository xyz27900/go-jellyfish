# go-jellyfish

Go port of Python [jellyfish](https://github.com/jamesturk/jellyfish) — phonetic encoding and string similarity.

## Exported API

- `Jaro(s1, s2 string) float64`
- `JaroWinkler(s1, s2 string) float64`
- `Metaphone(s string) string`

## Bugs fixed vs upstream Go port

- **Metaphone W-before-vowel**: original checked `nextnext` instead of `next`, dropping W before a vowel
- **JaroWinkler short-string boost**: upstream Go port had a `len > 3` guard that prevented the Winkler prefix boost on short strings; Python jellyfish has no such guard

## Test data

`testdata/` contains CSVs from the Python jellyfish repo plus regression cases.

## Python parity tests

```
go test -v -tags python ./...
```

Requires `pip install jellyfish`.

## Dependencies

Only `golang.org/x/text` for Unicode normalization.
