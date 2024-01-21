package masker

type Service struct {
	prod Producer
	pres Presenter
}

/* Service constructor */
func NewService(r Presenter, w Producer) *Service {
	var svc *Service = new(Service)
	svc.prod = w
	svc.pres = r
	return svc
}

func (s *Service) Run() error {
	data, err := s.prod.produce()
	if err != nil {
		return err
	}
	err = s.pres.present(data)
	if err != nil {
		return err
	}
	return nil
}
