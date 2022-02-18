package service

import (
	"fmt"
	"log"

	"github.com/access2content/go-practice/mocking/database"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

// Declare an interface that we want to mock
type registrationPreChecker interface {
	userExists(string) bool
}

//	Create a variable that will be used for implementation
var regPreCond registrationPreChecker

//	One custom implementation of the interface
type customChecker struct{}

func (check customChecker) userExists(email string) bool {
	return database.UserExists(email)
}

func init() {
	//	Assign the implementation variable to custom interface implementation
	regPreCond = customChecker{}
}

// Register a user only if the user has not been previously registered.
func Register(newUser User) error {
	// CHECK if user is already registered using the interface implementation variable
	userExists := regPreCond.userExists(newUser.Email)
	if userExists {
		return fmt.Errorf("email '%s' already registered", newUser.Email)
	}

	log.Println(newUser)
	return nil
}
