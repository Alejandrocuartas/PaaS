package services

import (
	"PaaS/environment"
	"PaaS/utilities"
)

func Deploy(appIdentifier string) error {

	var err error

	deploymentDirectory := "dist"
	repoName := "deploy-packages"

	err = utilities.CloneRepo("https://github.com/Alejandrocuartas/chiper-front.git", repoName)
	if err != nil {
		panic(err)
	}

	err = utilities.GenerateTemplateFiles(appIdentifier, deploymentDirectory, environment.TaskRole, repoName)
	if err != nil {
		panic(err)
	}

	err = utilities.DeployToECS(repoName, appIdentifier)
	if err != nil {
		panic(err)
	}

	utilities.RemoveDir(repoName)
	utilities.RemoveDockerImage(appIdentifier)

	return nil
}
