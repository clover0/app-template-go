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

func TestUserStore(t *testing.T) {
	config := testutils.SetUpConfig("../config")
	testutils.SetupDBWithConfig(config)

	t.Run("it creates successfully", testUserStore_CreateSuccess)
	t.Run("it counts successfully", testUserStore_CountsSuccess)

}

func testUserStore_CreateSuccess(t *testing.T) {
	userStoreFn := New()
	user := &core.User{
		Email:    "test" + fmt.Sprint(testutils.GenerateRandomNum()),
		Password: "test_password",
	}
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

func testUserStore_CountsSuccess(t *testing.T) {
	num := testutils.GenerateRandomNum()
	userStoreFn := New()
	user1 := &core.User{
		Email:    "test" + fmt.Sprint(testutils.GenerateRandomNum()),
		Password: "test_password" + fmt.Sprint(num),
	}
	user2 := &core.User{
		Email:    "test" + fmt.Sprint(testutils.GenerateRandomNum()),
		Password: "test_password" + fmt.Sprint(num),
	}
	err := db.Transact(dbtest.GetCon(), func(tx *sqlx.Tx) error {
		res, err := userStoreFn(tx).Create(user1)
		if res == 0 || err != nil {
			t.Error("failed create user1")
			return err
		}
		res, err = userStoreFn(tx).Create(user2)
		if res == 0 || err != nil {
			t.Error("failed create user2")
			return err
		}
		act, err := userStoreFn(tx).Count("password", "test_password"+fmt.Sprint(num))
		if act != 2 {
			t.Fail()
			return err
		}

		return err
	})
	if err != nil {
		t.Error(err)
	}
}
