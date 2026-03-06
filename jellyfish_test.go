package jellyfish

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"
	"testing"
)

// Test vectors from upstream Python jellyfish: https://github.com/jamesturk/jellyfish

func TestMetaphone(t *testing.T) {
	f, err := os.Open("testdata/metaphone.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	for _, rec := range records {
		input, want := rec[0], rec[1]
		got := Metaphone(input)
		if got != want {
			t.Errorf("Metaphone(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestJaroWinkler(t *testing.T) {
	f, err := os.Open("testdata/jaro_winkler.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	for _, rec := range records {
		a, b := rec[0], rec[1]
		want, _ := strconv.ParseFloat(rec[2], 64)
		got := JaroWinkler(a, b)
		if math.Abs(got-want) > 0.001 {
			t.Errorf("JaroWinkler(%q, %q) = %.4f, want %.4f", a, b, got, want)
		}
	}
}

func TestJaro(t *testing.T) {
	f, err := os.Open("testdata/jaro_distance.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	for _, rec := range records {
		a, b := rec[0], rec[1]
		want, _ := strconv.ParseFloat(rec[2], 64)
		got := Jaro(a, b)
		if math.Abs(got-want) > 0.001 {
			t.Errorf("Jaro(%q, %q) = %.4f, want %.4f", a, b, got, want)
		}
	}
}
