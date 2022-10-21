package manager

import "pinjol-perdana/repository"

type RepositoryManager interface {
	AccountRepository() repository.AccountRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (i *repositoryManager) AccountRepository() repository.AccountRepository {
	return repository.NewAccountRepository(i.infra.SqlDB())
}

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{infra: infra}
}
