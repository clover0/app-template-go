package handler

import (
	"auth465/core"
	"auth465/testutils"

	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type userServiceMock struct{}

func (u *userServiceMock) CheckDuplicateEmail(user *core.User) (bool, error) {
	return true, nil
}

func (u *userServiceMock) Register(user *core.User) error {
	return nil
}

func TestCreateUserHandler(t *testing.T) {

	t.Run("it returns 201 success", TestCreateUserHandler_return201)

}

func TestCreateUserHandler_return201(t *testing.T) {
	userJson := fmt.Sprintf(`{"email": "test%s@app465.com","password":"123456"}`, fmt.Sprint(testutils.GenerateRandomNum()))

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	service := &userServiceMock{}
	handler := CreateUserHandler(service)

	result := "1\n"
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, result, rec.Body.String())
	}
}
