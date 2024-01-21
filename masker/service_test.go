package masker

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockPresenter struct {
	mock.Mock
}

type mockProducer struct {
	mock.Mock
}

// produce is a strongly typed function that returns a slice of strings and an error.
func (p *mockProducer) produce() ([]string, error) {
	args := p.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (p *mockPresenter) present(data []string) error {
	args := p.Called(data)
	return args.Error(0)
}

func TestNewService(t *testing.T) {
	// Define the test case
	tc := struct {
		Produce Producer
		Present Presenter
		want    *Service
		error   error
	}{
		Produce: &mockProducer{},
		Present: &mockPresenter{},
		want: &Service{
			prod: &mockProducer{},
			pres: &mockPresenter{},
		},
		error: nil,
	}

	// Run the test case
	got := NewService(tc.Present, tc.Produce)
	assert.Equal(t, tc.want, got)
}

func TestService_Run(t *testing.T) {
	/* tabular testing */
	testCases := []struct {
		name     string
		got      []string
		expected []string
		err      error
	}{
		{
			name:     "success",
			expected: []string{"http://**********", "there is not a single link here", "And here is the link to http://*****************"},
			got:      []string{"http://**********", "there is not a single link here", "And here is the link to http://*****************"},
			err:      nil,
		},
		{
			name:     "successToo",
			expected: []string{"there is not a single link here"},
			got:      []string{"there is not a single link here"},
			err:      nil,
		},
		{
			name:     "fail",
			expected: nil,
			got:      nil,
			err:      errors.New("error producing data"),
		},
		{
			name:     "failAgain",
			expected: []string{"http://google.com"},
			got:      []string{"http://google.com"},
			err:      errors.New("error producing data"),
		},
		{
			name:     "wantsNil",
			expected: nil,
			got:      nil,
			err:      nil,
		},
	}
	Service := NewService(&mockPresenter{}, &mockProducer{})
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			Service.prod.(*mockProducer).On("produce").Return(tc.expected, tc.err)
			Service.pres.(*mockPresenter).On("present", tc.got).Return(tc.err)
			assert.Equal(t, tc.expected, tc.got)

		})
	}
}
