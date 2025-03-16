package secrets

import (
	"fmt"
	"log"

	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

func GenerateSecret(email string) (string, string) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "SendHive",
		AccountName: email,
	})
	if err != nil {
		log.Fatal("Error generating TOTP key:", err)
	}
	secret := key.Secret()
	fmt.Println("Secret Key:", key.Secret())
	fmt.Println("TOTP URL:", key.URL())
	return secret, key.URL()
}

func CampareKey(userCode string, storedSecret string) bool {
	flag := false
	valid := totp.Validate(userCode, storedSecret)
	if valid {
		fmt.Println("Authentication successful!")
		flag = true
		return flag
	} else {
		fmt.Println(" Invalid TOTP code. Try again.")
		return flag
	}
}

func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error while generating the hash: ", err)
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			log.Println("Password doesn't matches the stored password")
			return false, nil
		}
		log.Println("error while comparing the password: ", err)
		return false, err
	}
	return true, nil
}
