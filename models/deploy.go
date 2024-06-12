package models

import "errors"

type Deploy struct {
	AppIdentifier string `json:"app_identifier"`
}

func (d Deploy) Validate() error {
	if d.AppIdentifier == "" {
		return errors.New("app_identifier is required")
	}
	return nil
}
