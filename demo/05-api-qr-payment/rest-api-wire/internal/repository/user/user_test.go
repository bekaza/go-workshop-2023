package user_test

import (
	"context"
	"example/apiwire/internal/repository/user"
	"example/apiwire/internal/utils/timeutils"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type testUserRepoSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock

	userRepo user.IUserRepository
}

func TestUserRepo(t *testing.T) {
	suite.Run(t, &testUserRepoSuite{})
}

func (u *testUserRepoSuite) SetupSuite() {
	dbMock, mock, err := sqlmock.New()
	u.Require().NoError(err)

	dialector := postgres.New(postgres.Config{
		Conn:       dbMock,
		DriverName: "postgres",
	})
	gormMock, err := gorm.Open(dialector, &gorm.Config{})
	u.Require().NoError(err)

	u.db = gormMock
	u.mock = mock
	u.userRepo = user.ProvideUserRepo(gormMock)
}

func (u *testUserRepoSuite) SetupTest() {
	timeutils.Freeze()
}

func (u *testUserRepoSuite) TearDownTest() {
	timeutils.UnFreeze()
}

func (u *testUserRepoSuite) AfterTest(_, _ string) {
	u.Require().NoError(u.mock.ExpectationsWereMet())
}

func (u *testUserRepoSuite) TearDownSuite() {
	sql, err := u.db.DB()
	if err != nil {
		sql.Close()
	}
}

func (u *testUserRepoSuite) TestUserRepo_Create() {
	u.Run("success - create user", func() {
		var (
			id   = "00000000-0000-0000-0000-000000000000"
			name = "test-name"
			ctx  = context.Background()
		)
		u.mock.ExpectBegin()
		u.mock.ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "user_models" ("id","name","created_at") VALUES ($1,$2,$3)`)).
			WithArgs(id, name, timeutils.Now()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		u.mock.ExpectCommit()

		err := u.userRepo.Create(ctx, name)
		u.Require().NoError(err)
	})
}
