type Monad interface {
    Bind(func(interface{}) Monad) Monad
    Return(interface{}) Monad
}
