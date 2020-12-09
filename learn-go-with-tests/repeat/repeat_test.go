package iteration

import (
    "fmt"
    "testing"
)

func TestRepeat(t *testing.T) {
    assertCorrect := func(t *testing.T, expected, actual string) {
        t.Helper()
        if actual != expected {
            t.Errorf("expected %q but got %q", expected, actual)
        }
    }

    t.Run("should repeat 5 times", func(t *testing.T) {
        repeated := Repeat("a", 5)
        expected := "aaaaa"
        assertCorrect(t, expected, repeated)
    })

    t.Run("should repeat 20 times", func(t *testing.T) {
        repeated := Repeat("a", 20)
        expected := "aaaaaaaaaaaaaaaaaaaa"
        assertCorrect(t, expected, repeated)
    })
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 5)
    }
}

func ExampleRepeat() {
    repeated := Repeat("a", 5)
    fmt.Println(repeated)
    // Output: aaaaa
}
