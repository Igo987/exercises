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
	assert.NotNil(t, svc, "service must not be nil")
	assert.NoError(t, svc.Run())

}

func TestService_Run(t *testing.T) {
	/* tabular testing */

	testCases := []struct {
		input    []string
		expected []string
		err      error
	}{
		{
			input:    []string{"http://google.com", "there is not a single link here", "And here is the link to http://stackoverflow.com"},
			expected: []string{"http://************", "there is not a single link here", "And here is the link to http://*****************"},
			err:      nil,
		},
		{
			input:    []string{"http://google.com", "there is not a single link here"},
			expected: []string{"http://************", "there is not a single link here"},
			err:      nil,
		},
		{
			input:    []string{"http://google.com", "there is not a single link here", "And here is the link to http://stackoverflow.com"},
			expected: nil,
			err:      errors.New("error producing data"),
		},
		{
			input:    []string{},
			expected: nil,
			err:      nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		myMock := &mockService{
			err: tc.err,
		}
		myMock.On("produce").Return(tc.input, tc.err)
		myMock.On("present", tc.expected).Return(tc.err)
		err := myMock.Run()
		assert.Equal(t, err, tc.err)
	}
}
