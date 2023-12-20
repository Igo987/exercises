package masker

import "log"

type Service struct {
	prod Producer
	pres Presenter
}

// Service constructor
func NewService(r Presenter, w Producer) *Service {
	svc := new(Service)
	svc.prod = w
	svc.pres = r
	return svc
}

func (s *Service) Run() error {
	data, err := s.prod.produce()
	if err != nil {
		log.Println("error opening the file", err)
		return err
	}
	err = s.pres.present(data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
