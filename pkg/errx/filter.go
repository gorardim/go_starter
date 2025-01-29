package errx

import (
	"strings"

	"gorm.io/gorm"
)

func FilterDuplicateErr(err error) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "Duplicate entry") {
		return nil
	}
	return err
}

func IsDuplicateErr(err error) bool {
	if err == nil {
		return false
	}
	if strings.Contains(err.Error(), "Duplicate entry") {
		return true
	}
	return false
}

func FilterRecordNotFoundErr(err error) error {
	if err == nil {
		return nil
	}
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}
