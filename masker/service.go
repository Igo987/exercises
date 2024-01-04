package masker

import "fmt"

type Service struct {
	prod Producer
	pres Presenter
}

/* Service constructor */
func NewService(r Presenter, w Producer) *Service {
	svc := new(Service)
	svc.prod = w
	svc.pres = r
	return svc
}

func (s *Service) Run() error {
	data, err := s.prod.produce()
	if err != nil {
		return fmt.Errorf("error opening the file: %w", err)
	}
	err = s.pres.present(data)
	if err != nil {
		return fmt.Errorf("error presenting data: %w", err)
	}
	return nil
}
