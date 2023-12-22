package utils

import (
	"os"
	"time"

	"github.com/aryamaahendra/mesiakad/domains/account"
	"github.com/aryamaahendra/mesiakad/pkgs/meerrors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Student struct {
	Nim             string    `json:"nim"`
	ConcentrationID *uint64   `json:"concentration_id"`
	PembimbingID    *uint64   `json:"pembimbing_id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Username        string    `json:"username"`
	RoleID          *uint64   `json:"role"`
	Sex             string    `json:"sex"`
	Religion        string    `json:"religion"`
	Phone           string    `json:"phone"`
	Address         string    `json:"address"`
	BirthPlace      string    `json:"birth_place"`
	BirthDate       time.Time `json:"birth_date"`
	ProdiID         *uint64   `json:"prodi_id"`
}

type Lecturer struct {
	Nip        string    `json:"nip"`
	Nidn       string    `json:"nidn"`
	NoRegBlu   string    `json:"no_reg_blu"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	RoleID     *uint64   `json:"role"`
	Sex        string    `json:"sex"`
	Religion   string    `json:"religion"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	BirthPlace string    `json:"birth_place"`
	BirthDate  time.Time `json:"birth_date"`
	ProdiID    *uint64   `json:"prodi_id"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewToken(username string) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "dsiakad",
		"sub": username,
	})

	token, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(token string) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))
	t, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, meerrors.ErrUnuthorized
		}

		return key, nil
	})

	if err != nil {
		return "", err
	}

	if !t.Valid {
		return "", meerrors.ErrUnuthorized
	}

	username, err := t.Claims.GetSubject()
	if err != nil {
		return "", err
	}

	return username, nil
}

func ProfileResponse(acc account.Account) (any, error) {
	if *acc.GetRoleID() == 2 {
		student := Student{
			Nim:             acc.GetNim(),
			ConcentrationID: acc.GetConcentrationID(),
			PembimbingID:    acc.GetPembimbingID(),
			Name:            acc.GetName(),
			Email:           acc.GetEmail(),
			Username:        acc.GetUsername(),
			RoleID:          acc.GetRoleID(),
			Sex:             acc.GetSex(),
			Religion:        acc.GetReligion(),
			Phone:           acc.GetPhone(),
			Address:         acc.GetAddress(),
			BirthPlace:      acc.GetBirthPlace(),
			BirthDate:       acc.GetBirthDate(),
			ProdiID:         acc.GetProdiID(),
		}
		return student, nil
	} else if *acc.GetRoleID() == 3 {
		lecturer := Lecturer{
			Nip:        acc.GetNip(),
			Nidn:       acc.GetNidn(),
			NoRegBlu:   acc.GetNoRegBlu(),
			Name:       acc.GetName(),
			Email:      acc.GetEmail(),
			Username:   acc.GetUsername(),
			RoleID:     acc.GetRoleID(),
			Sex:        acc.GetSex(),
			Religion:   acc.GetReligion(),
			Phone:      acc.GetPhone(),
			Address:    acc.GetAddress(),
			BirthPlace: acc.GetBirthPlace(),
			BirthDate:  acc.GetBirthDate(),
			ProdiID:    acc.GetProdiID(),
		}
		return lecturer, nil
	}

	return nil, meerrors.ErrRecordNotFound
}
