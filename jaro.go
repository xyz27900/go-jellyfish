package jellyfish

// Jaro computes the Jaro distance between two strings.
//
// Jaro distance is a string-edit distance that gives a floating point response in [0,1] where 0 represents two completely dissimilar strings and 1 represents identical strings.
func Jaro(s1, s2 string) float64 {
	return jaroWinkler(s1, s2, false, false)
}

// JaroWinkler computes the Jaro-Winkler distance between two strings.
//
// Jaro-Winkler is a modification/improvement to Jaro distance, like Jaro it gives a floating point response in [0,1] where 0 represents two completely dissimilar strings and 1 represents identical strings.
//
// See the Jaro-Winkler distance article at Wikipedia (http://en.wikipedia.org/wiki/Jaro-Winkler_distance) for more details.
func JaroWinkler(s1, s2 string) float64 {
	return jaroWinkler(s1, s2, false, true)
}

func jaroWinkler(s1, s2 string, longTolerance, winklerize bool) float64 {
	r1 := []rune(s1)
	r2 := []rune(s2)
	len1 := len(r1)
	len2 := len(r2)

	if len1 == 0 || len2 == 0 {
		return 0
	}

	minLen := max(len1, len2)
	searchRange := (minLen / 2) - 1
	if searchRange < 0 {
		searchRange = 0
	}

	flags1 := make([]bool, len1)
	flags2 := make([]bool, len2)

	// look within search range for matched pairs
	commonChars := 0
	for i, ch := range r1 {
		low := 0
		if i > searchRange {
			low = i - searchRange
		}
		high := len2 - 1
		if i+searchRange < len2 {
			high = i + searchRange
		}
		for j := low; j <= high; j++ {
			if !flags2[j] && r2[j] == ch {
				flags1[i] = true
				flags2[j] = true
				commonChars++
				break
			}
		}
	}

	// short circuit if no characters match
	if commonChars == 0 {
		return 0
	}

	// count transpositions
	k := 0
	transCount := 0
	for i, f1 := range flags1 {
		if f1 {
			j := k
			for ; j < len2; j++ {
				if flags2[j] {
					k = j + 1
					break
				}
			}
			if r1[i] != r2[j] {
				transCount++
			}
		}
	}
	transCount /= 2

	// adjust for similarities in nonmatched characters
	ccf := float64(commonChars)
	weight := (ccf/float64(len1) + ccf/float64(len2) + (ccf-float64(transCount))/ccf) / 3

	// Winkler boost: no minimum-length guard, matching Python jellyfish behavior.
	if winklerize && weight > 0.7 {
		j := min(min(len1, len2), 4)
		i := 0

		for i < j && r1[i] == r2[i] {
			i++
		}

		if i != 0 {
			weight += float64(i) * 0.1 * (1 - weight)
		}

		// TODO: add long_tolerance? optionally adjust for long strings
	}

	return weight
}
