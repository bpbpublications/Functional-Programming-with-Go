type Operation func(int, int) int

ops := map[string]Operation{
    "add": func(a, b int) int { return a + b },
    "sub": func(a, b int) int { return a - b },
}

result := ops["add"](5, 3)  // 8
