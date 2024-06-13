package services

import (
	"PaaS/models"
	"PaaS/repositories"
	"PaaS/utilities"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	uuid "github.com/satori/go.uuid"
)

func CreateApp(data models.CreateApp) (
	r models.CreateAppResponse,
	e error,
) {
	data.DeploymentDirecotry = strings.ReplaceAll(data.DeploymentDirecotry, "/", "")
	app := models.App{
		Name:                data.Name,
		RepositoryUrl:       data.RepositoryUrl,
		UserId:              data.UserId,
		Status:              models.AppStatusPending,
		DeploymentDirecotry: data.DeploymentDirecotry,
	}

	e = repositories.CreateApp(&app)
	if e != nil {
		return r, utilities.ManageError(e)
	}

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
		},
	)
	if err != nil {
		app.Status = models.AppStatusInactive
		repositories.UpdateApp(&app)
		log.Printf("an error happened trying to connect to aws %s", err.Error())
		return r, fmt.Errorf("an error happened trying to connect to aws %s", err.Error())
	}
	sqsClient := sqs.New(sess)

	queueUrl := "https://sqs.us-east-1.amazonaws.com/277469568219/Deployments.fifo"

	messageAttributes := map[string]*sqs.MessageAttributeValue{
		"APP_UUID": {
			DataType:    aws.String("String"),
			StringValue: aws.String(string(app.UUID.String())),
		},
	}

	messageGroupId := app.UUID.String()

	sendMessageOutput, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:               &queueUrl,
		MessageBody:            aws.String("{}"),
		MessageAttributes:      messageAttributes,
		MessageGroupId:         aws.String(messageGroupId),
		MessageDeduplicationId: aws.String(uuid.NewV4().String()),
	})
	if err != nil {
		app.Status = models.AppStatusInactive
		repositories.UpdateApp(&app)
		log.Println("error sqs message: ", err.Error())
		return r, fmt.Errorf("error sqs message: %s", err.Error())
	}

	log.Println(sendMessageOutput)

	r = models.CreateAppResponse{
		UUID:          app.UUID.String(),
		CreatedAt:     app.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:     app.UpdatedAt.Time.Format(time.RFC3339),
		Name:          app.Name,
		Status:        app.Status,
		RepositoryUrl: app.RepositoryUrl,
		DeployUrl:     app.DeployUrl.String,
	}

	return r, e
}

func GetApps(userId uint) (
	apps []models.GetAppsResponse,
	e error,
) {

	apps, e = repositories.GetApps(userId)
	if e != nil {
		return apps, utilities.ManageError(e)
	}

	return apps, e
}
