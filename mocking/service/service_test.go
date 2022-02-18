package service

import (
	"testing"
)

func TestCheckUserExist(t *testing.T) {
	user := User{
		Name:     "Ankur Anand",
		Email:    "anand@example.com",
		UserName: "anand",
	}

	err := Register(user)
	if err == nil {
		t.Error("Expected Register User to throw and error got nil")
	}
}

func TestUserNotExist(t *testing.T) {
	user := User{
		Name:     "Ankur Anand",
		Email:    "asdfasdf",
		UserName: "anand",
	}

	err := Register(user)
	if err != nil {
		t.Errorf("User should not exist")
	}
}

//	Declare a dynamic variable
var myMockedFunction func(string) bool

//	One test implementation of the Checking Interface
type testChecker struct{}

func (check testChecker) userExists(email string) bool {
	return myMockedFunction(email)
}

func TestMocked(t *testing.T) {
	user := User{
		Name:     "Ankur Anand",
		Email:    "asdfasdf",
		UserName: "anand",
	}

	//	Set the implementation variable to the test checker
	regPreCond = testChecker{}

	//	Modify the mocked function according to requirement
	myMockedFunction = func(email string) bool {
		return false
	}

	err := Register(user)
	if err != nil {
		t.Fatal(err)
	}

	//	Modify the mocked function according to requirement
	myMockedFunction = func(email string) bool {
		return true
	}

	err = Register(user)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
