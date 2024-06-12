type ImmutableMap struct {
    data map[string]int
}

func NewImmutableMap(data map[string]int) *ImmutableMap {
    newData := make(map[string]int)
    for k, v := range data {
        newData[k] = v
    }
    return &ImmutableMap{data: newData}
}

func (m *ImmutableMap) Set(key string, value int) *ImmutableMap {
    newData := make(map[string]int)
    for k, v := range m.data {
        newData[k] = v
    }
    newData[key] = value
    return &ImmutableMap{data: newData}
}
