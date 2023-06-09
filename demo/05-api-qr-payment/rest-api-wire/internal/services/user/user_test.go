package user_test

import (
	"context"
	mock_user_repo "example/apiwire/internal/repository/user/mock_user"
	"example/apiwire/internal/services/user"
	"example/apiwire/internal/utils/timeutils"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type testUserServiceSuite struct {
	suite.Suite
	mockUserRepo *mock_user_repo.MockIUserRepository
	underTest    user.UserService
}

func (s *testUserServiceSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	s.mockUserRepo = mock_user_repo.NewMockIUserRepository(ctrl)
	s.underTest = user.ProvideUserService(s.mockUserRepo)
}

func (s *testUserServiceSuite) SetupTest() {
	timeutils.Freeze()
}

func (s *testUserServiceSuite) TearDownTest() {
	timeutils.UnFreeze()
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, &testUserServiceSuite{})
}

func (s *testUserServiceSuite) TestUserService_CreateUser() {
	var (
		ctx      = context.TODO()
		username = "test-name"
	)
	s.mockUserRepo.EXPECT().Create(ctx, username).Return(nil)

	err := s.underTest.CreateUser(ctx, username)
	s.Require().NoError(err)
}
