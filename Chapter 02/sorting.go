import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }

    // Sorting by age in ascending order
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age < people[j].Age
    })

    fmt.Println("Sorted by Age (ascending):", people)

    // Sorting by name in descending order
    sort.Slice(people, func(i, j int) bool {
        return people[i].Name > people[j].Name
    })

    fmt.Println("Sorted by Name (descending):", people)
}
