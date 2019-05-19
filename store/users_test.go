package users

import (
	"auth465/core"
	"auth465/db"
	"auth465/db/dbtest"
	"auth465/testutils"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestUserStore_Create(t *testing.T) {
	config := testutils.SetUpConfig("../config")
	testutils.SetupDBWithConfig(config)

	t.Run("it creates successfully", testUserStore_CreateSuccess)
}

func testUserStore_CreateSuccess(t *testing.T) {
	userStoreFn := New()
	user := &core.User{
		Email:    "test" + fmt.Sprint(testutils.GenerateRandomNum()),
		Password: "test_password",
	}
	fmt.Println(dbtest.GetCon())
	err := db.Transact(dbtest.GetCon(), func(tx *sqlx.Tx) error {
		res, err := userStoreFn(tx).Create(user)
		if res == 0 || err != nil {
			t.Error("failed create user")
		}
		return err
	})
	if err != nil {
		t.Error(err)
	}
}
