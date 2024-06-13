package models

type GetAppsResponse struct {
	UUID                string `json:"uuid" gorm:"column:uuid"`
	CreatedAt           string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt           string `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt           string `json:"deleted_at" gorm:"column:deleted_at"`
	Name                string `json:"name" gorm:"column:name"`
	Status              string `json:"status" gorm:"column:status"`
	RepositoryUrl       string `json:"repository_url" gorm:"column:repository_url"`
	DeployUrl           string `json:"deploy_url" gorm:"column:deploy_url"`
	DeploymentDirecotry string `json:"deployment_directory" gorm:"column:deployment_directory"`
}
