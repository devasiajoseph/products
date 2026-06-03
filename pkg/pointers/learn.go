package pointers

type User struct {
	Name  string
	Email string
}

func Pointer(user *User) {
	user.Email = "adoniaromal@gmail.com"
}

func NoPointer(user User) {
	user.Email = "adoniaromal@gmail.com"
}

func InitializePointer() User {
	user := User{Name: "Adonia",
		Email: "adonia@gmail.com"}
	//Pointer(&user)
	NoPointer(user)
	return user
}
