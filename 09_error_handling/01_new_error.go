package error_handling

import (
	"errors"
	"fmt"
)

type IUser interface {
	GetEmail() (string, error)
}

type User struct {
	firstName string
	lastName  string
	email     string
	active    bool
}

func (user User) GetEmail() (string, error) {
	if user.email == "" {
		return "", errors.New("Invalid User email")
	} else {
		return user.email, nil
	}
}

func GetUserEmail(user IUser) (string, error) {
	return user.GetEmail()
}

func New_Error_1() {
	user_1 := User{
		firstName: "Foo",
		lastName:  "Bar",
		email:     "foo@example.com",
		active:    true,
	}

	user_2 := User{
		firstName: "John",
		lastName:  "Doe",
		email:     "",
		active:    true,
	}

	user_1_email, user_1_email_error := GetUserEmail(user_1)
	user_2_email, user_2_email_error := GetUserEmail(user_2)

	fmt.Printf("user_1_email: [%v], user_1_email_error: [%v]\n", user_1_email, user_1_email_error) // user_1_email: [foo@example.com], user_1_email_error: [<nil>]
	fmt.Printf("user_2_email: [%v], user_2_email_error: [%v]\n", user_2_email, user_2_email_error) // user_2_email: [], user_2_email_error: [Invalid User email]
}
