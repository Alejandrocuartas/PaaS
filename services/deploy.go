package services

import (
	"PaaS/environment"
	"PaaS/repositories"
	"PaaS/utilities"
	"fmt"

	"github.com/jinzhu/gorm"
)

func Deploy(appIdentifier string) error {

	var err error

	app, err := repositories.GetAppByUuid(appIdentifier)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return fmt.Errorf("app with uuid %s not found", appIdentifier)
		}
		return err
	}
	fmt.Println(app)

	dir := appIdentifier

	err = utilities.CloneRepo(app.RepositoryUrl, dir)
	if err != nil {
		utilities.RemoveDir(dir)
		return err
	}

	err = utilities.GenerateTemplateFiles(app.UUID.String(), app.DeploymentDirecotry, environment.TaskRole, dir)
	if err != nil {
		utilities.RemoveDir(dir)
		return err
	}

	err = utilities.DeployToECS(dir, app.UUID.String())
	if err != nil {
		utilities.RemoveDir(dir)
		return err
	}

	utilities.RemoveDir(dir)
	utilities.RemoveDockerImage(app.UUID.String())

	return nil
}
