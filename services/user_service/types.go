package userservice

import (
	"errors"

	"github.com/aryamaahendra/mesiakad/domains/mesiakad"
)

var (
	ErrUserUnuthorized = errors.New("Unuthorized")
)

type User = mesiakad.User

type CreateUser struct {
	Email    string  `json:"email" validate:"required,email,max=64"`
	Username string  `json:"username" validate:"required,max=32"`
	Password string  `json:"password" validate:"required"`
	RoleID   *uint64 `json:"role_id" validate:"required"`
}

type UpdateUser struct {
	Email    string  `json:"email" validate:"required,email,max=64"`
	Username string  `json:"username" validate:"required,max=32"`
	RoleID   *uint64 `json:"role_id" validate:"required"`
}

type UserAuth struct {
	Username string `json:"username" validate:"required,max=32"`
	Password string `json:"password" validate:"required"`
}

type AuthorizedUser struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type ProfileStudent struct {
	Nim             string  `json:"nim" validate:"required,max=20"`
	ConcentrationID *uint64 `json:"concentration_id" validate:"required,numeric,min=1"`
	PembimbingID    *uint64 `json:"pembimbing_id" validate:"required,numeric,min=1"`
	Name            string  `json:"name" validate:"required,max=64"`
	Sex             string  `json:"sex" validate:"required,max=20"`
	Religion        string  `json:"religion" validate:"required,max=20"`
	Phone           string  `json:"phone" validate:"required,max=20"`
	Address         string  `json:"address" validate:"required,max=255"`
	BirthPlace      string  `json:"birth_place" validate:"required,max=32"`
	BirthDate       string  `json:"birth_date" validate:"required"`
	ProdiID         *uint64 `json:"prodi_id" validate:"required,min=1"`
}

type ProfileLecturer struct {
	Nip        string  `json:"nip" validate:"required,numeric,max=28"`
	Nidn       string  `json:"nidn" validate:"required,numeric,max=28"`
	NoRegBlu   string  `json:"no_reg_blu" validate:"required,numeric,max=28"`
	Name       string  `json:"name" validate:"required,max=64"`
	Sex        string  `json:"sex" validate:"required,max=20"`
	Religion   string  `json:"religion" validate:"required,max=20"`
	Phone      string  `json:"phone" validate:"required,max=20"`
	Address    string  `json:"address" validate:"required,max=255"`
	BirthPlace string  `json:"birth_place" validate:"required,max=32"`
	BirthDate  string  `json:"birth_date" validate:"required"`
	ProdiID    *uint64 `json:"prodi_id" validate:"required,min=1"`
}
