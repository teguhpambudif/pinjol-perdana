package usecase

import (
	"pinjol-perdana/model"
	"pinjol-perdana/repository"
	"pinjol-perdana/utils"
)

type AccountUseCase interface {
	RegisterNewAccount(newAccount *model.Account) error
	GetAccountById(id string) (model.Account, error)
	Upload(addFile *model.Account) error
}

type accountUseCase struct {
	repo repository.AccountRepository
}

func (a *accountUseCase) RegisterNewAccount(newAccount *model.Account) error {
	newAccount.Id = utils.GenerateId()
	return a.repo.Register(newAccount)
}

func (a *accountUseCase) GetAccountById(id string) (model.Account, error) {
	return a.repo.Get(id)
}

func (a *accountUseCase) Upload(addFile *model.Account) error {
	return a.repo.Upload(addFile)
}

func NewAccountUseCase(repo repository.AccountRepository) AccountUseCase {
	ac := new(accountUseCase)
	ac.repo = repo
	return ac
}
