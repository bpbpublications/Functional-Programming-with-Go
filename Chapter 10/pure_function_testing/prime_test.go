package prime

import "testing"

func TestIsPrime(t *testing.T) {
    cases := []struct {
        n        int
        expected bool
    }{
        {1, false},
        {2, true},
        {3, true},
        {4, false},
        {5, true},
        {25, false},
        {29, true},
    }

    for _, c := range cases {
        result := IsPrime(c.n)
        if result != c.expected {
            t.Errorf("IsPrime(%d) == %t, expected %t", c.n, result, c.expected)
        }
    }
}
