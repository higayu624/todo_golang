package requests

type PlayListRepository struct{
	RequestHandler RequestHandler
}

func (repository *PlayListRepository) ShowOne(){
	return repository.RequestHandler.Connect()
}