package repository

import (
	"github.com/infinity-api/dto"
	"log"

	"github.com/infinity-api/model"
	"github.com/jmoiron/sqlx"
)

type MemberRepository interface {
	Get() ([]model.Member, error)
	Insert(model.Member) error
	GetMembers() ([]dto.MemberDTO, error)
}

type memberRepository struct {
	dbContext *sqlx.DB
}

func NewMemberRepository(db *sqlx.DB) MemberRepository {
	return &memberRepository{
		dbContext: db,
	}
}

func (repo *memberRepository) GetMembers() ([]dto.MemberDTO, error) {
	var people []dto.MemberDTO
	var err = repo.dbContext.Select(&people, "SELECT * FROM infinity.member")
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return people, nil
}

func (repo *memberRepository) Get() ([]model.Member, error) {
	var people []model.Member
	var err = repo.dbContext.Select(&people, "SELECT * FROM member")
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return people, nil
}

func (repo *memberRepository) Insert(member model.Member) error {
	var err error
	tx := repo.dbContext.MustBegin()
	query := "INSERT INTO company (id, name,age,address,salary) VALUES (:id, :name, :age,:address,:salary)"
	_, err = tx.NamedExec(query, member)
	if err != nil {
		log.Panic(`before`, err)
	}
	err = tx.Commit()
	if err != nil {
		log.Panic("after", err)
	}
	return err
}
