import "sort"

type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 30},
    {"Bob", 25},
}

// Sort by age
sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})
