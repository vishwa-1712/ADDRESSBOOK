package services

import (
	"addressbook/config"
	"addressbook/internal/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// func AuthenticateUser(email string, password string) error {

// 	if email != staticUser.Email {
// 		return fmt.Errorf("INVALID EMAIL")
// 	}

// 	//here it will take the the password
// 	// 	We're checking:
// 	// “Are the passwords the same?” means the entered one which will be bcrypted one and our hashed one.
// 	err := bcrypt.CompareHashAndPassword([]byte(staticUser.Password), []byte(password))
// 	// password dont match
// 	if err != nil {
// 		return fmt.Errorf("INVALID PASSWORD")
// 	}
// 	//password matches
// 	return nil
// }

// func AuthenticateUser(email, password string) bool {
// 	if email != staticUser.Email {
// 		return false
// 	}

// err := bcrypt.CompareHashAndPassword([]byte(staticUser.Password), []byte(password))
// return err == nil
func AuthenticateUser(email string, password string) error {

	var staticUser = models.User{
		Email:    config.GetEnv("STATIC_USER_EMAIL"),
		Password: config.GetEnv("STATIC_USER_PASSWORD"),
	}
	fmt.Println("ENV email:", staticUser.Email)
	fmt.Println("Request email:", email)

	if email != staticUser.Email {
		return fmt.Errorf("INVALID EMAIL")
	}

	testPswd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password")
	}
	err = bcrypt.CompareHashAndPassword(testPswd, []byte(password))
	if err != nil {
		// fmt.Println("Provided password:", testPswd)
		// fmt.Println("Expected hash:", staticUser.Password)
		return fmt.Errorf("INVALID PASSWORD")
	}

	return nil

}
