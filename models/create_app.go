package models

import (
	"errors"
	"strings"
)

type CreateApp struct {
	Name                string `json:"name"`
	RepositoryUrl       string `json:"repository_url"`
	UserId              uint   `json:"user_id"`
	DeploymentDirecotry string `json:"deployment_directory"`
}

func (c CreateApp) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.RepositoryUrl == "" {
		return errors.New("repository_url is required")
	}
	if !strings.Contains(c.RepositoryUrl, "https://") {
		return errors.New("repository_url must start with https://")
	}
	if !strings.Contains(c.RepositoryUrl, ".git") {
		return errors.New("repository_url must end with .git")
	}

	repoParts := strings.Split(c.RepositoryUrl, ".")
	if repoParts[len(repoParts)-1] != "git" {
		return errors.New("repository_url must end with .git")
	}

	if c.UserId == 0 {
		return errors.New("user_id is required")
	}
	if c.DeploymentDirecotry == "" {
		return errors.New("deployment_directory is required")
	}
	return nil
}

type CreateAppResponse struct {
	UUID          string `json:"uuid"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	RepositoryUrl string `json:"repository_url"`
	DeployUrl     string `json:"deploy_url"`
}
