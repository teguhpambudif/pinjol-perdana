package manager

import "pinjol-perdana/usecase"

type UseCaseManager interface {
	AccountUseCase() usecase.AccountUseCase
}

type useCaseManager struct {
	repo RepositoryManager
}

func (u *useCaseManager) AccountUseCase() usecase.AccountUseCase {
	return usecase.NewAccountUseCase(u.repo.AccountRepository())
}

func NewUseCaseManager(repo RepositoryManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
