package service

import (
	"sync"

	"github.com/hypertonyc/rpc-incrementor-service/internal/domain"
	"github.com/hypertonyc/rpc-incrementor-service/internal/repository"
)

type IncrementService interface {
	GetNumber() (int32, error)
	IncrementNumber() error
	SetSettings(incrementStep int32, upperLimit int32) error
}

type incrementService struct {
	incrementRepo repository.IncrementRepository
	mutex         sync.Mutex
}

// Create new increment service
func NewIncrementService(repo repository.IncrementRepository) IncrementService {
	return &incrementService{
		incrementRepo: repo,
	}
}

// Get current number value
func (s *incrementService) GetNumber() (int32, error) {
	number, err := s.incrementRepo.GetNumber()
	if err != nil {
		return 0, err
	}

	return number.GetValue(), nil
}

// Increment current number using the current settings. This method is thread-safe
func (s *incrementService) IncrementNumber() error {
	// Lock the mutex to prevent potential race conditions
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Get the current settings
	settings, err := s.incrementRepo.GetSettings()
	if err != nil {
		return err
	}

	// Get the current number
	currentNumber, err := s.incrementRepo.GetNumber()
	if err != nil {
		return err
	}

	// Increment the number
	currentNumber.Increment(settings.GetIncrementStep(), settings.GetUpperLimit())

	// Update the number in the repository
	err = s.incrementRepo.UpdateNumber(currentNumber)
	if err != nil {
		return err
	}

	return nil
}

// Update current settings
func (s *incrementService) SetSettings(incrementStep int32, upperLimit int32) error {
	var settings domain.Settings
	var err error

	err = settings.SetUpperLimit(upperLimit)
	if err != nil {
		return err
	}

	err = settings.SetIncrementStep(incrementStep)
	if err != nil {
		return err
	}

	// Lock the mutex to prevent potential race conditions
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.incrementRepo.UpdateSettings(&settings)
}
