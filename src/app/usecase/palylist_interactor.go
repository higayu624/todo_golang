package usecase

type PlayListInteractor struct {
	PlayListRepository PlayListRepository
}

func(interactor *PlayListInteractor) GetPlayList(){
	interactor.PlayListRepository.ShowOne()
}