package jellyfish

// Metaphone calculates the metaphone code for a string.
// Ported from Python jellyfish (https://github.com/jamesturk/jellyfish).
func Metaphone(s string) string {
	r := normalize(s)
	// work in lowercase like Python
	lower := make([]rune, len(r))
	for i, c := range r {
		if c >= 'A' && c <= 'Z' {
			lower[i] = c + 32
		} else {
			lower[i] = c
		}
	}
	r = lower
	rlen := len(r)

	if rlen == 0 {
		return ""
	}

	// skip first character if s starts with these
	if rlen > 1 {
		switch {
		case r[0] == 'k' && r[1] == 'n',
			r[0] == 'g' && r[1] == 'n',
			r[0] == 'p' && r[1] == 'n',
			r[0] == 'w' && r[1] == 'r',
			r[0] == 'a' && r[1] == 'e':
			r = r[1:]
			rlen--
		}
	}

	star := rune('*') // sentinel for "no character"

	next := func(i int) rune {
		if i+1 < rlen {
			return r[i+1]
		}
		return star
	}
	nextnext := func(i int) rune {
		if i+2 < rlen {
			return r[i+2]
		}
		return star
	}
	isVow := func(c rune) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	var result []rune

	for i := 0; i < rlen; i++ {
		c := r[i]
		nx := next(i)
		nnx := nextnext(i)

		// skip doubles except cc
		if c == nx && c != 'c' {
			continue
		}

		switch {
		case isVow(c):
			if i == 0 || r[i-1] == ' ' {
				result = append(result, c)
			}
		case c == 'b':
			if !(i != 0 && r[i-1] == 'm') || nx != star {
				result = append(result, 'b')
			}
		case c == 'c':
			if (nx == 'i' && nnx == 'a') || nx == 'h' {
				result = append(result, 'x')
				i++
			} else if nx == 'i' || nx == 'e' || nx == 'y' {
				result = append(result, 's')
				i++
			} else {
				result = append(result, 'k')
			}
		case c == 'd':
			if nx == 'g' && (nnx == 'i' || nnx == 'e' || nnx == 'y') {
				result = append(result, 'j')
				i += 2
			} else {
				result = append(result, 't')
			}
		case c == 'f' || c == 'j' || c == 'l' || c == 'm' || c == 'n' || c == 'r':
			result = append(result, c)
		case c == 'g':
			if nx == 'i' || nx == 'e' || nx == 'y' {
				result = append(result, 'j')
			} else if nx == 'h' && nnx != star && !isVow(nnx) {
				i++
			} else if nx == 'n' && nnx == star {
				i++
			} else {
				result = append(result, 'k')
			}
		case c == 'h':
			if i == 0 || isVow(nx) || !isVow(r[i-1]) {
				result = append(result, 'h')
			}
		case c == 'k':
			if i == 0 || r[i-1] != 'c' {
				result = append(result, 'k')
			}
		case c == 'p':
			if nx == 'h' {
				result = append(result, 'f')
				i++
			} else {
				result = append(result, 'p')
			}
		case c == 'q':
			result = append(result, 'k')
		case c == 's':
			if nx == 'h' {
				result = append(result, 'x')
				i++
			} else if nx == 'i' && (nnx == 'o' || nnx == 'a') {
				result = append(result, 'x')
				i += 2
			} else {
				result = append(result, 's')
			}
		case c == 't':
			if nx == 'i' && (nnx == 'o' || nnx == 'a') {
				result = append(result, 'x')
			} else if nx == 'h' {
				result = append(result, '0')
				i++
			} else if !(nx == 'c' && nnx == 'h') {
				result = append(result, 't')
			}
		case c == 'v':
			result = append(result, 'f')
		case c == 'w':
			if i == 0 && nx == 'h' {
				i++
				result = append(result, 'w')
			} else if isVow(nx) {
				result = append(result, 'w')
			}
		case c == 'x':
			if i == 0 {
				if nx == 'h' || (nx == 'i' && (nnx == 'o' || nnx == 'a')) {
					result = append(result, 'x')
				} else {
					result = append(result, 's')
				}
			} else {
				result = append(result, 'k')
				result = append(result, 's')
			}
		case c == 'y':
			if isVow(nx) {
				result = append(result, 'y')
			}
		case c == 'z':
			result = append(result, 's')
		case c == ' ':
			if len(result) > 0 && result[len(result)-1] != ' ' {
				result = append(result, ' ')
			}
		}
	}

	// uppercase the result
	for i, c := range result {
		if c >= 'a' && c <= 'z' {
			result[i] = c - 32
		}
	}

	return string(result)
}
