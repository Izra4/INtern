package sdk

import "golang.org/x/crypto/bcrypt"

func Hashing(pass string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashed := string(password)
	return hashed, nil
}

func ValidateHash(hashed, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
	return err
}
