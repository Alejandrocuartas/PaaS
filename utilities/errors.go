package utilities

import (
	"errors"
	"strings"
)

func ManageError(err error) error {

	if err == nil {
		return err
	}

	msg := err.Error()

	if strings.Contains(msg, "uix_users_email") {
		return errors.New("email already exists")
	}

	if strings.Contains(msg, "apps_user_id_users_id_foreign") {
		return errors.New("invalid user_id")
	}

	return err
}
