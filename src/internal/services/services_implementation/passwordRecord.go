package services_implementation

import (
	"src/internal/models"
	"src/internal/repositories"
	"src/internal/services"
)

type passwordRecordServiceImplementation struct {
	passwordRecordRepository repositories.PasswordRecordRepository
}

func NewPasswordRecordServiceImplementation(passwordRecordRepository repositories.PasswordRecordRepository) services.PasswordRecordService {
	return &passwordRecordServiceImplementation{
		passwordRecordRepository: passwordRecordRepository,
	}
}

func (s *passwordRecordServiceImplementation) Set(record *models.PasswordRecord) error {
	err := s.passwordRecordRepository.Set(record)
	if err != nil {
		return err
	}
	return nil
}

func (s *passwordRecordServiceImplementation) Get(username string, service string) (*models.PasswordRecord, error) {
	record, err := s.passwordRecordRepository.Get(username, service)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (s *passwordRecordServiceImplementation) Delete(username string, service string) error {
	err := s.passwordRecordRepository.Delete(username, service)
	if err != nil {
		return err
	}
	return nil
}
