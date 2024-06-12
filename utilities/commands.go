package utilities

import (
	"PaaS/environment"
	"fmt"
	"os/exec"
	"strings"
)

// https://github.com/Alejandrocuartas/chiper-front.git
func CloneRepo(repoUrl string, repoName string) (
	err error,
) {
	isHttps := strings.Contains(repoUrl, "https://")
	if !isHttps {
		return fmt.Errorf("repoUrl must be https")
	}

	repoUrl = strings.Replace(repoUrl, ".git", "", 1)

	cmd := exec.Command("git", "clone", repoUrl, repoName)
	cmd.Dir = "."
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return fmt.Errorf("error cloning repo: %s", err)
	}
	fmt.Println(string(out))

	return nil

}

func RemoveDir(dir string) error {
	cmd := exec.Command("rm", "-rf", dir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error removing dir: %s", err)
	}
	fmt.Println(string(out))
	return nil
}

func RemoveDockerImage(imageName string) error {
	cmd := exec.Command("docker", "rmi", imageName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error removing docker image: %s", err)
	}
	fmt.Println(string(out))
	return nil
}

func DeployToECS(dir string, deploymentName string) error {

	//build docker image
	cmd := exec.Command("docker", "build", "-t", deploymentName, ".")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return fmt.Errorf("error building docker image: %s", err)
	}
	fmt.Println(string(out))

	//tag docker image
	dockerTag := fmt.Sprintf("%s.dkr.ecr.us-east-1.amazonaws.com/%s:latest", environment.AwsAccountID, deploymentName)
	cmd = exec.Command("docker", "tag", deploymentName+":latest", dockerTag)
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return fmt.Errorf("error tagging docker image: %s", err)
	}
	fmt.Println(string(out))

	//log to aws ecr
	awsCmd := exec.Command("aws", "ecr", "get-login-password", "--region", "us-east-1")

	awsOutput, err := awsCmd.Output()
	if err != nil {
		fmt.Println(string(awsOutput))
		return fmt.Errorf("error getting aws login password: %s", err)
	}
	dockerCmd := exec.Command("docker", "login", "--username", "AWS", "--password-stdin", environment.AwsAccountID+".dkr.ecr.us-east-1.amazonaws.com")
	dockerStdin, err := dockerCmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("error getting docker login password: %s", err)
	}

	if err := dockerCmd.Start(); err != nil {
		return fmt.Errorf("error starting docker login: %s", err)
	}

	if _, err := dockerStdin.Write(awsOutput); err != nil {
		return fmt.Errorf("error writing docker login password: %s", err)
	}

	if err := dockerStdin.Close(); err != nil {
		return fmt.Errorf("error closing docker login password: %s", err)
	}

	if err := dockerCmd.Wait(); err != nil {
		return fmt.Errorf("error waiting for docker login: %s", err)
	}

	//creating ecr repository
	cmd = exec.Command("aws", "ecr", "create-repository", "--repository-name", deploymentName)
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return fmt.Errorf("error creating ecr repository: %s", err)
	}
	fmt.Println(string(out))

	//push image to ecr
	cmd = exec.Command("docker", "push", dockerTag)
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return fmt.Errorf("error pushing image to ecr: %s", err)
	}
	fmt.Println(string(out))

	//create cluster
	cmd = exec.Command("aws", "ecs", "create-cluster", "--cluster-name", deploymentName)
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return fmt.Errorf("error creating cluster: %s", err)
	}
	fmt.Println(string(out))

	//create task definition
	taskDefinitionFile := fmt.Sprintf("file://%s/task-definition.json", dir)
	cmd = exec.Command("aws", "ecs", "register-task-definition", "--cli-input-json", taskDefinitionFile)
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return fmt.Errorf("error creating task definition: %s", err)
	}
	fmt.Println(string(out))

	//run task
	cmd = exec.Command("aws", "ecs", "run-task", "--cluster", deploymentName, "--task-definition", deploymentName, "--launch-type", "FARGATE", "--network-configuration", "awsvpcConfiguration={subnets=[subnet-05ac77cc9408cd3cf,subnet-09ac7e192c9968e48,subnet-0c93fcdd003c4abdf,subnet-05e82bf9b01cdfb76,subnet-0b57d05061517d4a7,subnet-0b7ed245b0a08af17],securityGroups=[sg-0254f0cc15dc0ceaf],assignPublicIp=ENABLED}")
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return fmt.Errorf("error running task: %s", err)
	}
	fmt.Println(string(out))

	RemoveDockerImage(dockerTag)

	return nil

}
