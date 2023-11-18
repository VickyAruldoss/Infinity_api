package repository

import (
	"log"

	"github.com/infinity-api/model"
	"github.com/jmoiron/sqlx"
)

type MemberRepository interface {
	Get() ([]model.Member, error)
	Insert(model.Member) error
}

type memberRepository struct {
	dbContext *sqlx.DB
}

func NewMemberController(db *sqlx.DB) *memberRepository {
	return &memberRepository{
		dbContext: db,
	}
}

func (repo *memberRepository) Get() ([]model.Member, error) {
	people := []model.Member{}
	err := repo.dbContext.Select(&people, "SELECT * FROM company")
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return people, nil
}

func (repo *memberRepository) Insert(member model.Member) error {
	var err error
	tx := repo.dbContext.MustBegin()
	_, err = tx.NamedExec("INSERT INTO company (id, name,age,address,salary) VALUES (:id, :name, :age,:address,:salary)", member)
	if err != nil {
		log.Panic("before", err)
	}
	err = tx.Commit()
	if err != nil {
		log.Panic("after", err)
	}
	return err
}
