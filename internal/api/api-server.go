package api

type Server interface {
	Start() error
}

type server struct {
}

func (s server) Start() error {
	//TODO implement me
	panic("implement me")
}

func NewServer() Server {
	return &server{}
}
