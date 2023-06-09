package handler_test

import (
	"context"
	"example/apiwire/cmd/api/handler"
	mock_user_service "example/apiwire/internal/services/user/mock_user"
	"example/apiwire/pkg/testutils"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type testUserHandlerSuite struct {
	suite.Suite
	mockUserService *mock_user_service.MockUserService
	underTest       handler.UserHandler
}

func (h *testUserHandlerSuite) SetupSuite() {
	ctrl := gomock.NewController(h.T())
	h.mockUserService = mock_user_service.NewMockUserService(ctrl)
	h.underTest = handler.ProvideUserHandler(h.mockUserService)
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, &testUserHandlerSuite{})
}

func (h *testUserHandlerSuite) TestUserHandler_RegisterUserHandler() {
	var (
		ctx = context.TODO()
	)
	req := handler.UserRequest{
		Name: "test-name",
	}
	res, gCtx := testutils.NewWithRequestContext(http.MethodPost, "/user", testutils.JSON(req))

	h.mockUserService.EXPECT().CreateUser(ctx, req.Name).Return(nil)

	h.underTest.RegisterUserHandler(gCtx)

	h.Equal(http.StatusCreated, res.Code)
	var response map[string]interface{}
	testutils.DecodeJSON(res.Body, &response)
	h.Equal("hi", response["Msg"])
}
