package password

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword 生成密码
func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ValidatePassword 验证密码
func ValidatePassword(password, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

func ValidatePayPassword(password string) error {
	// only support 6 digits password can not be the same digits and can not be continuous
	if len(password) != 6 {
		return fmt.Errorf("PASSWORD_MUST_BE_SIX_DIGITS")
	}
	// check if all digits are numbers
	match, _ := regexp.MatchString("^[0-9]+$", password)
	if !match {
		return fmt.Errorf("PASSWORD_MUST_BE_NUMBERS")
	}
	// check if all digits are the same
	allSame := true
	for i := 0; i < len(password)-1; i++ {
		if password[i] != password[i+1] {
			allSame = false
			break
		}
	}
	if allSame {
		return fmt.Errorf("PASSWORD_CAN_NOT_BE_ALL_SAME")
	}
	// check if all digits are continuous
	isContinuous := true
	for i := 0; i < len(password)-1; i++ {
		if password[i]+1 != password[i+1] {
			isContinuous = false
			break
		}
	}
	// check if all digits are continuous in reverse order
	if !isContinuous {
		isContinuous = true
		for i := 0; i < len(password)-1; i++ {
			if password[i]-1 != password[i+1] {
				isContinuous = false
				break
			}
		}
	}
	if isContinuous {
		return fmt.Errorf("PASSWORD_CAN_NOT_BE_CONTINUOUS")
	}
	return nil
}
