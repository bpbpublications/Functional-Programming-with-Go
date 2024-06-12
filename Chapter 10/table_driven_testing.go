func TestAddTableDriven(t *testing.T) {
    var tests = []struct {
        a, b   int // inputs
        expect int // expected result
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, -2, -3},
    }

    for _, tt := range tests {
        testname := fmt.Sprintf("%d+%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := Add(tt.a, tt.b)
            if ans != tt.expect {
                t.Errorf("got %d, want %d", ans, tt.expect)
            }
        })
    }
}
