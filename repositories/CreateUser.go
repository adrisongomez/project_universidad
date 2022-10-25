package user

type CreateUser struct {
	name     string
	email    string
	password string
}

func NewCreateUserPayload(name string, email string, password string) *CreateUser {
    return &CreateUser{
        name: name,
        email: email,
        password: password,
    }
}

func (u *CreateUser) GetName() string {
    return u.name
}

func (u *CreateUser) GetEmail() string {
    return u.email
}
func (u *CreateUser) GetPassword() string {
    return u.password
}
