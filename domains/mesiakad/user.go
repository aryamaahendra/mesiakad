package mesiakad

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	RoleID    *uint64   `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Profile struct {
	ID              uint64    `json:"id"`
	Nim             string    `json:"nim"`
	Nip             string    `json:"nip"`
	Nidn            string    `json:"nidn"`
	NoRegBlu        string    `json:"no_reg_blu"`
	Name            string    `json:"name"`
	Sex             string    `json:"sex"`
	Religion        string    `json:"religion"`
	Phone           string    `json:"phone"`
	Address         string    `json:"address"`
	BirthPlace      string    `json:"birth_place"`
	BirthDate       time.Time `json:"birth_date"`
	ProdiID         *uint64   `json:"prodi_id"`
	ConcentrationID *uint64   `json:"concentration_id"`
	PembimbingID    *uint64   `json:"pembimbing_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
