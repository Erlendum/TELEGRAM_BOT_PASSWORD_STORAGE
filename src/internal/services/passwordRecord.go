package services

import "src/internal/models"

type PasswordRecordService interface {
	Set(record *models.PasswordRecord) error
	Get(username string, service string) (*models.PasswordRecord, error)
	Delete(username string, service string) error
}
