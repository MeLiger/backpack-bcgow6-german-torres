package users

type Service interface {
	GetAll() ([]User, error)
	Store(name, email string, age int, height int, active bool, date string) (User, error)
	Update(id int, name, email string, age int, height int, active bool, date string) (User, error)
	UpdateNameAge(id int, name string, age int) (User, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]User, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(name, email string, age int, height int, active bool, date string) (User, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return User{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, name, email, age, height, active, date)
	if err != nil {
		return User{}, err
	}

	return producto, nil
}

func (s *service) Update(id int, name, email string, age int, height int, active bool, date string) (User, error) {
	return s.repository.Update(id, name, email, age, height, active, date)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
func (s *service) UpdateNameAge(id int, name string, age int) (User, error) {

	return s.repository.UpdateNameAge(id, name, age)
}
