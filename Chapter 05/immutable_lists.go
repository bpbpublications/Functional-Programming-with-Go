type ImmutableList struct {
    elements []int
}

func NewImmutableList(elements []int) *ImmutableList {
    newList := make([]int, len(elements))
    copy(newList, elements)
    return &ImmutableList{elements: newList}
}

func (l *ImmutableList) Add(element int) *ImmutableList {
    newElements := append([]int(nil), l.elements...)
    newElements = append(newElements, element)
    return &ImmutableList{elements: newElements}
}
