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

type sessionServiceMock struct{}

func (service sessionServiceMock) CreateSession(userId uint32) (string, error) {
	return "xxxxxx", nil
}

func (service sessionServiceMock) GetSession(key string) (string, error) {
	return "session token", nil
}

// FindUserByEmail finds user by email and return user with transaction
func (service sessionServiceMock) FindUserByEmail(email string) (*core.User, error) {
	user := &core.User{
		1,
		"",
		"",
		"email",
		"password",
	}
	return user, nil
}

func (service sessionServiceMock) FindUserById(id uint32) (*core.User, error) {
	user := &core.User{
		1,
		"",
		"",
		"email",
		"password",
	}
	return user, nil
}

func (service sessionServiceMock) ComparePassword(user *core.User, input string) error {
	return nil
}

type sessionServiceMock2 struct{}

func (service sessionServiceMock2) CreateSession(userId uint32) (string, error) {
	return "xxxxxx", nil
}

func (service sessionServiceMock2) GetSession(key string) (string, error) {
	return "session token", nil
}

func (service sessionServiceMock2) FindUserByEmail(email string) (*core.User, error) {
	return nil, nil
}

func (service sessionServiceMock2) FindUserById(id uint32) (*core.User, error) {
	return nil, nil
}

func (service sessionServiceMock2) ComparePassword(user *core.User, input string) error {
	return nil
}

func TestCreateSessionHandler(t *testing.T) {

	t.Run("it returns 201", testCreateSession_return201)
	t.Run("it returns 400", testCreateSession_return400)

}

func testCreateSession_return201(t *testing.T) {
	userJson := fmt.Sprintf(`{"email": "test%s@app465.com","password":"123456"}`, fmt.Sprint(testutils.GenerateRandomNum()))

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	service := &sessionServiceMock{}
	handler := CreateSessionHandler(service)

	result := "1\n"
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, result, rec.Body.String())
	}
}

func testCreateSession_return400(t *testing.T) {
	userJson := fmt.Sprintf(`{"email": "test%s@app465.com","password":"123456"}`, fmt.Sprint(testutils.GenerateRandomNum()))

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	service := &sessionServiceMock2{}
	handler := CreateSessionHandler(service)

	result := "\"email no match\"\n"
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, result, rec.Body.String())
	}
}
