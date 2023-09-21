package repository

type repository struct {
	conection string
}

func NewRepository(conection string) *repository {
	return &repository{
		conection,
	}
}

func (s *repository) Post() {

}

func (s *repository) Get() {

}
