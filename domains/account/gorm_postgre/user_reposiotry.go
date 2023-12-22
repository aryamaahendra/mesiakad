package gormpostgre

import (
	"errors"
	"time"

	"github.com/aryamaahendra/mesiakad/domains/account"
	"github.com/aryamaahendra/mesiakad/pkgs/meerrors"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64
	Email     string
	Username  string
	Password  string
	RoleID    *uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) ToAggregate(acc *account.Account) {
	acc.SetID(u.ID)
	acc.SetUsername(u.Username)
	acc.SetEmail(u.Email)
	acc.SetPassword(u.Password)
	acc.SetRoleID(u.RoleID)
	acc.SetUserCreateAt(u.CreatedAt)
	acc.SetUserUpdatedAt(u.UpdatedAt)
}

func (m *GormPostgreRepository) GetAllUsers() ([]account.Account, error) {
	var users []User
	var accounts []account.Account

	err := m.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		acc := account.Account{}
		user.ToAggregate(&acc)

		accounts = append(accounts, acc)
	}

	return accounts, nil
}

func (m *GormPostgreRepository) getByUsername(username string, user *User) error {

	err := m.db.Where("username = ?", username).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return meerrors.ErrRecordNotFound
		}

		return err
	}

	return nil
}

func (m *GormPostgreRepository) GetUserByUsername(username string) (account.Account, error) {
	var user User
	acc := account.Account{}

	err := m.getByUsername(username, &user)
	if err != nil {
		return acc, err
	}

	user.ToAggregate(&acc)
	return acc, nil
}

func (m *GormPostgreRepository) CreateUser(acc *account.Account) error {
	user := &User{
		Username: acc.GetUsername(),
		Email:    acc.GetEmail(),
		Password: acc.GetPassword(),
		RoleID:   acc.GetRoleID(),
	}

	return m.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(user).Error
		if err != nil {
			return err
		}

		acc.SetID(user.ID)
		return nil
	})
}

func (m *GormPostgreRepository) UpdateUser(username string, acc *account.Account) error {
	user := &User{}

	err := m.getByUsername(username, user)
	if err != nil {
		return err
	}

	user.Username = acc.GetUsername()
	user.Email = acc.GetEmail()
	user.RoleID = acc.GetRoleID()

	return m.db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", user.ID).Save(user).Error
	})
}

func (m *GormPostgreRepository) DeleteUser(username string) error {
	user := &User{}

	err := m.getByUsername(username, user)
	if err != nil {
		return err
	}

	return m.db.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&User{}, user.ID).Error
	})
}
