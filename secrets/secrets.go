package secrets

import (
	"fmt"
	"log"

	"github.com/pquerna/otp/totp"
)

func GenerateSecret() string {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "SendHive",
		AccountName: "admin@sendHive.com",
	})
	if err != nil {
		log.Fatal("Error generating TOTP key:", err)
	}
	secret := key.Secret()
	fmt.Println("Secret Key:", key.Secret())
	fmt.Println("TOTP URL:", key.URL())
	return secret
}

func CampareKey(userCode string, storedSecret string) bool {
	flag := false
	valid := totp.Validate(userCode, storedSecret)
	if valid {
		fmt.Println("Authentication successful!")
		flag = true
	} else {
		fmt.Println(" Invalid TOTP code. Try again.")
	}
	return flag
}
