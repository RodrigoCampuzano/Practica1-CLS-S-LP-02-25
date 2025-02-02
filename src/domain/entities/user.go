// En este archvio se crean los metodos a usar

package entities

type User struct {
	ID int32
	Name string
	Email string
}

func NewUser(name string, email string) *User {
	return &User{Name: name, Email: email}
}

func (u *User) UpdateUser(name string, email string){
	u.Name = name
	u.Email = email
}