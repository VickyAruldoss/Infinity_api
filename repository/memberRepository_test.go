package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type MemberRepositoryTestSuite struct {
	suite.Suite
	repo MemberRepository
}

func TestMemberRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MemberRepositoryTestSuite))
}

func (suite *MemberRepositoryTestSuite) SetupTest() {
	db := InitDb()
	suite.repo = NewMemberRepository(db)
}

func (suite *MemberRepositoryTestSuite) TestMemberRepository_GetMembers() {
	members, gotErr := suite.repo.GetMembers()
	suite.Suite.NotNil(members)
	suite.Suite.Nil(gotErr)
}

func InitDb() *sqlx.DB {
	host := "localhost"
	port := "5432"
	userName := "test_user"
	password := "password"
	dbName := "infinity_db"
	connectionString := fmt.Sprintf("host= %s port= %s user=%s password =%s dbname = %s sslmode=disable", host, port, userName, password, dbName)
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
