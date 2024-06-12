package utilities

import (
	"fmt"
	"os"
	"strings"
)

const (
	DOCKER_IGNORE = `node_modules
{{.DEPLOYMENT_DIRECTORY}}`

	DOCKERFILE = `FROM node:14 as build

WORKDIR /

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:alpine

COPY --from=build /{{.DEPLOYMENT_DIRECTORY}} /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]`

	TASK_DEFINITION = `{
    "family": "{{.NAME}}",
    "networkMode": "awsvpc",
    "taskRoleArn": "{{.TASK_ROLE}}",
    "executionRoleArn": "{{.TASK_ROLE}}",
    "containerDefinitions": [
        {
            "name": "{{.NAME}}",
            "image": "277469568219.dkr.ecr.us-east-1.amazonaws.com/{{.NAME}}:latest",
            "essential": true,
            "portMappings": [
                {
                    "containerPort": 80,
                    "hostPort": 80
                }
            ],
            "logConfiguration": {
                "logDriver": "awslogs",
                "options": {
                    "awslogs-create-group": "true",
                    "awslogs-group": "awslogs-{{.NAME}}",
                    "awslogs-region": "us-east-1",
                    "awslogs-stream-prefix": "awslogs-{{.NAME}}"
                }
            }
        }
    ],
    "requiresCompatibilities": [
        "FARGATE"
    ],
    "cpu": "256",
    "memory": "512"
}`
)

func GenerateTemplateFiles(name string, deploymentDirectory string, taskRole string, dir string) error {

	var err error

	dockerIgnoreContent := DOCKER_IGNORE
	dockerIgnoreContent = strings.ReplaceAll(dockerIgnoreContent, "{{.DEPLOYMENT_DIRECTORY}}", deploymentDirectory)

	dockerfileContent := DOCKERFILE
	dockerfileContent = strings.ReplaceAll(dockerfileContent, "{{.DEPLOYMENT_DIRECTORY}}", deploymentDirectory)

	taskDefinitionContent := TASK_DEFINITION
	taskDefinitionContent = strings.ReplaceAll(taskDefinitionContent, "{{.NAME}}", name)
	taskDefinitionContent = strings.ReplaceAll(taskDefinitionContent, "{{.TASK_ROLE}}", taskRole)

	err = os.WriteFile(dir+"/Dockerfile", []byte(dockerfileContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing Dockerfile: %w", err)
	}

	err = os.WriteFile(dir+"/.dockerignore", []byte(dockerIgnoreContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing .dockerignore: %w", err)
	}

	err = os.WriteFile(dir+"/task-definition.json", []byte(taskDefinitionContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing task-definition.json: %w", err)
	}

	return nil
}
