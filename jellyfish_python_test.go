//go:build python

package jellyfish

import (
	"math"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func pythonMetaphone(t *testing.T, s string) string {
	t.Helper()
	out, err := exec.Command("python3", "-c",
		"from jellyfish import metaphone; print(metaphone('"+s+"'), end='')").Output()
	if err != nil {
		t.Fatalf("python3 metaphone(%q): %v", s, err)
	}
	return string(out)
}

func pythonJaroWinkler(t *testing.T, a, b string) float64 {
	t.Helper()
	out, err := exec.Command("python3", "-c",
		"from jellyfish import jaro_winkler_similarity; print(jaro_winkler_similarity('"+a+"','"+b+"'), end='')").Output()
	if err != nil {
		t.Fatalf("python3 jaro_winkler(%q, %q): %v", a, b, err)
	}
	f, err := strconv.ParseFloat(strings.TrimSpace(string(out)), 64)
	if err != nil {
		t.Fatalf("parse float %q: %v", out, err)
	}
	return f
}

func pythonJaro(t *testing.T, a, b string) float64 {
	t.Helper()
	out, err := exec.Command("python3", "-c",
		"from jellyfish import jaro_similarity; print(jaro_similarity('"+a+"','"+b+"'), end='')").Output()
	if err != nil {
		t.Fatalf("python3 jaro(%q, %q): %v", a, b, err)
	}
	f, err := strconv.ParseFloat(strings.TrimSpace(string(out)), 64)
	if err != nil {
		t.Fatalf("parse float %q: %v", out, err)
	}
	return f
}

func TestMetaphonePython(t *testing.T) {
	words := []string{
		"metaphone", "algorithm", "telephone", "christopher",
		"wright", "knight", "psychology", "pneumonia",
		"xylophone", "aesthetic", "gnome", "schadenfreude",
		"", "a", "ab",
	}
	for _, w := range words {
		t.Run(w, func(t *testing.T) {
			got := Metaphone(w)
			want := pythonMetaphone(t, w)
			if got != want {
				t.Errorf("Metaphone(%q) = %q, python = %q", w, got, want)
			}
		})
	}
}

func TestJaroWinklerPython(t *testing.T) {
	pairs := [][2]string{
		{"martha", "marhta"},
		{"dwayne", "duane"},
		{"dixon", "dicksonx"},
		{"jellyfish", "smellyfish"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "abc"},
		{"abc", "def"},
		{"sergey", "sergei"},
		{"zhukov", "joukov"},
	}
	for _, p := range pairs {
		t.Run(p[0]+"_"+p[1], func(t *testing.T) {
			got := JaroWinkler(p[0], p[1])
			want := pythonJaroWinkler(t, p[0], p[1])
			if math.Abs(got-want) > 1e-6 {
				t.Errorf("JaroWinkler(%q, %q) = %.10f, python = %.10f", p[0], p[1], got, want)
			}
		})
	}
}

func TestJaroPython(t *testing.T) {
	pairs := [][2]string{
		{"martha", "marhta"},
		{"dwayne", "duane"},
		{"dixon", "dicksonx"},
		{"jellyfish", "smellyfish"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
	}
	for _, p := range pairs {
		t.Run(p[0]+"_"+p[1], func(t *testing.T) {
			got := Jaro(p[0], p[1])
			want := pythonJaro(t, p[0], p[1])
			if math.Abs(got-want) > 1e-6 {
				t.Errorf("Jaro(%q, %q) = %.10f, python = %.10f", p[0], p[1], got, want)
			}
		})
	}
}
