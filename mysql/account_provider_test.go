package mysql

import (
	"testing"

	gomocket "github.com/Selvatico/go-mocket"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/SSL/account_management_system/model"
)

func TestUpdateUserprofile(t *testing.T) {
	// db, mock, err := sqlmock.NewWithDSN("sqlmock_db_0")
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// }
	// defer db.Close()

	const (
		account = "account"
		pwd     = "pwd"
		name    = "Patrick"
		mail    = "test@example.com"
		phone   = "0912345678"
		fbid    = "fbu12345"
	)

	// mock.ExpectBegin()
	// mock.ExpectQuery("INSERT INTO `users` (`account_id`,`password`,`user_name`,`mail`,`fb_id`,`is_deleted`) VALUES (?,?,?,?,?,?,?)").
	// 	WithArgs(account, pwd, name, mail, phone, fbid, false)

	user := &model.User{
		AccountID: account,
		Password:  pwd,
		UserName:  name,
		Mail:      mail,
		Phone:     phone,
		FBID:      fbid,
	}

	gomocket.Catcher.Register()
	gormDb, err := gorm.Open(gomocket.DriverName, "")
	gormDb.LogMode(true)
	defer gormDb.Close()

	gomocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO \"users\"").
		WithArgs(account, pwd, "name", mail, phone, fbid, false).
		WithReply(make([]map[string]interface{}, 0))
	// gormDb, err := gorm.Open("sqlmock", "sqlmock_db_0")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	ap := NewAccountProvider(gormDb)
	err = ap.Register(user)
	assert.Nil(t, err)
}
