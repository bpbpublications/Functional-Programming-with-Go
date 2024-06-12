func (m Result) Bind(f func(interface{}) Monad) Monad {
    if m.Err != nil {
        return m // Propagate the error without executing the function
    }
    return f(m.Value)
}
