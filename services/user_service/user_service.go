package userservice

import (
	"errors"
	"time"

	"github.com/aryamaahendra/mesiakad/domains/account"
	"github.com/aryamaahendra/mesiakad/domains/mesiakad"
	"github.com/aryamaahendra/mesiakad/pkgs/meerrors"
	"github.com/aryamaahendra/mesiakad/pkgs/utils"
)

type UserService struct {
	account account.Repository
}

func New(cfgs ...ServiceConfig) (*UserService, error) {
	service := &UserService{}

	for _, cfg := range cfgs {
		err := cfg(service)
		if err != nil {
			return nil, err
		}
	}

	return service, nil
}

func (us *UserService) GetAll() ([]User, error) {
	accounts, err := us.account.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var users []User
	for _, account := range accounts {
		users = append(users, *account.GetUser())
	}

	return users, nil
}

func (us *UserService) GetByUsername(username string) (User, error) {
	account, err := us.account.GetUserByUsername(username)
	if err != nil {
		return User{}, err
	}

	return *account.GetUser(), nil
}

func (us *UserService) CreateUser(data CreateUser) (User, error) {
	hashPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return User{}, err
	}

	acc := account.Account{}
	acc.SetUser(&mesiakad.User{
		Username: data.Username,
		Email:    data.Email,
		Password: hashPassword,
		RoleID:   data.RoleID,
	})

	err = us.account.CreateUser(&acc)
	if err != nil {
		return User{}, err
	}

	return *acc.GetUser(), nil
}

func (us *UserService) UpdateUser(username string, data UpdateUser) (User, error) {
	acc, err := us.account.GetUserByUsername(username)
	if err != nil {
		return User{}, err
	}

	acc.SetUsername(data.Username)
	acc.SetEmail(data.Email)
	acc.SetRoleID(data.RoleID)

	err = us.account.UpdateUser(username, &acc)
	if err != nil {
		return User{}, err
	}

	return *acc.GetUser(), nil
}

func (us *UserService) DeleteUser(username string) error {
	err := us.account.DeleteUser(username)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) AuthAPI(auth UserAuth) (AuthorizedUser, error) {
	var authUser AuthorizedUser

	acc, err := us.account.GetUserByUsername(auth.Username)
	if err != nil {
		if errors.Is(err, meerrors.ErrRecordNotFound) {
			return authUser, meerrors.ErrUnuthorized
		}

		return authUser, err
	}

	isPasswordMatch := utils.CheckPasswordHash(auth.Password, acc.GetPassword())
	if !isPasswordMatch {
		return authUser, meerrors.ErrUnuthorized
	}

	token, err := utils.NewToken(acc.GetUsername())
	if err != nil {
		return authUser, err
	}

	authUser.Token = token
	authUser.User = *acc.GetUser()

	return authUser, nil
}

func (us *UserService) UpdateProfileStudent(user User, student ProfileStudent) (account.Account, error) {
	var acc account.Account

	birthDate, err := time.Parse("2006-01-02", student.BirthDate)
	if err != nil {
		return account.Account{}, err
	}

	acc.SetUser(&user)
	acc.SetNim(student.Nim)
	acc.SetName(student.Name)
	acc.SetConcentrationID(student.ConcentrationID)
	acc.SetPembimbingID(student.PembimbingID)
	acc.SetProdiID(student.ProdiID)
	acc.SetSex(student.Sex)
	acc.SetPhone(student.Phone)
	acc.SetReligion(student.Religion)
	acc.SetAddress(student.Address)
	acc.SetBirthPlace(student.BirthPlace)
	acc.SetBirthDate(birthDate)

	err = us.account.CreateOrUpdateProfile(&acc)
	if err != nil {
		return account.Account{}, err
	}

	return acc, err
}

func (us *UserService) UpdateProfileLecturer(user User, lecturer ProfileLecturer) (account.Account, error) {
	var acc account.Account

	birthDate, err := time.Parse("2006-01-02", lecturer.BirthDate)
	if err != nil {
		return account.Account{}, err
	}

	acc.SetUser(&user)
	acc.SetNip(lecturer.Nip)
	acc.SetNidn(lecturer.Nidn)
	acc.SetNoRegBlu(lecturer.NoRegBlu)
	acc.SetName(lecturer.Name)
	acc.SetProdiID(lecturer.ProdiID)
	acc.SetSex(lecturer.Sex)
	acc.SetPhone(lecturer.Phone)
	acc.SetReligion(lecturer.Religion)
	acc.SetAddress(lecturer.Address)
	acc.SetBirthPlace(lecturer.BirthPlace)
	acc.SetBirthDate(birthDate)

	err = us.account.CreateOrUpdateProfile(&acc)
	if err != nil {
		return account.Account{}, err
	}

	return acc, err
}

func (us *UserService) GetPofile(user User) (account.Account, error) {
	acc, err := us.account.GetProfileByUserID(user.ID)
	if err != nil {
		return account.Account{}, err
	}

	return acc, err
}
