type User struct {
    Name  string
    Email string
}

type ImmutableUserList struct {
    users []User
}

func (ul *ImmutableUserList) Add(user User) *ImmutableUserList {
    newUsers := append([]User(nil), ul.users...)
    newUsers = append(newUsers, user)
    return &ImmutableUserList{users: newUsers}
}
