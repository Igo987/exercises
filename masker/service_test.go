package masker

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
	err error
}
type mockPresenter struct {
	err error
}

type mockProducer struct {
	err error
}

func (p *mockProducer) produce() ([]string, error) {
	return nil, p.err
}

func (p *mockPresenter) present(data []string) error {
	return p.err
}

func (s *mockService) produce() ([]string, error) {
	args := s.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (s *mockService) present(any []string) error {
	args := s.Called(any)
	return args.Error(0)
}

func (p *mockService) Run() error {
	if p.err != nil {
		return p.err
	}
	return nil
}

func TestNewService(t *testing.T) {
	r := &mockPresenter{}
	w := &mockProducer{}

	svc := NewService(r, w)

	assert.Equal(t, w, svc.prod, "prod must have a method produce() that returns []string and an error")
	assert.Equal(t, r, svc.pres, "pres must have a method present(s []string) that returns an error")

}

// TestRun tests the Run method of the Service struct.
func TestService_Run(t *testing.T) {
	myMock := &mockService{}
	// Test case 1: Successful execution
	myMock.On("produce").Return([]string{
		"http://************",
		"there is not a single link here",
		"And here is the link to http://*****************",
	}, nil)
	myMock.On("present").Return(nil)
	err := myMock.Run()
	assert.NoError(t, err)

	// Test case 2: Error producing data
	expectedProducerError := errors.New("error producing data")
	myMock.On("produce").Return(nil, err)
	err = myMock.Run()
	assert.Error(t, expectedProducerError, err)

	// Test case 3: Error presenting data
	expectedError := errors.New("error presenting data")
	myMock.On("present", []string{
		"http://google.com",
		"there is not a single link here",
		"And here is the link to http://stackoverflow.com",
	}).Return(expectedError)
	err = myMock.Run()
	assert.Error(t, expectedError, err)
}
