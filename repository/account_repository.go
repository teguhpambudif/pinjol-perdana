package repository

import (
	"pinjol-perdana/model"
	"pinjol-perdana/utils"

	"github.com/jmoiron/sqlx"
)

type AccountRepository interface {
	Register(newAccount *model.Account) error
	Get(id string) (model.Account, error)
	Upload(addFile *model.Account) error
}

type accountRepository struct {
	db *sqlx.DB
}

func (a *accountRepository) Register(newAccount *model.Account) error {
	_, err := a.db.NamedExec(utils.REGISTER_ACCOUNT, newAccount)
	if err != nil {
		panic(err)
	}
	return nil
}

func (a *accountRepository) Get(id string) (model.Account, error) {
	var account model.Account
	err := a.db.Get(&account, utils.GET_ACCOUNT, id)
	if err != nil {
		return account, err
	}
	return account, nil
}

func (a *accountRepository) Upload(addFile *model.Account) error {
	_, err := a.db.Exec(utils.UPLOAD_KTP, addFile)
	if err != nil {
		panic(err)
	}
	return nil
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	repo := new(accountRepository)
	repo.db = db
	return repo
}
