package gormpostgre

import (
	"errors"
	"time"

	"github.com/aryamaahendra/mesiakad/domains/account"
	"github.com/aryamaahendra/mesiakad/pkgs/meerrors"
	"gorm.io/gorm"
)

type Profile struct {
	ID              uint64 `gorm:"primaryKey"`
	UserID          uint64
	User            *User `gorm:"foreignKey:UserID"`
	Nim             string
	Nip             string
	Nidn            string
	NoRegBlu        string
	Name            string
	Sex             string
	Religion        string
	Phone           string
	Address         string
	BirthPlace      string
	BirthDate       time.Time
	ProdiID         *uint64
	ConcentrationID *uint64
	PembimbingID    *uint64
	Pembimbing      *User `gorm:"foreignKey:PembimbingID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (up *Profile) TableName() string {
	return "user_profiles"
}

func (p *Profile) ToAggregate(acc *account.Account) {
	acc.SetNim(p.Nim)
	acc.SetNip(p.Nip)
	acc.SetNidn(p.Nidn)
	acc.SetNoRegBlu(p.NoRegBlu)
	acc.SetName(p.Name)
	acc.SetSex(p.Sex)
	acc.SetReligion(p.Religion)
	acc.SetPhone(p.Phone)
	acc.SetAddress(p.Address)
	acc.SetBirthPlace(p.BirthPlace)
	acc.SetBirthDate(p.BirthDate)
	acc.SetProdiID(p.ProdiID)
	acc.SetConcentrationID(p.ConcentrationID)
	acc.SetPembimbingID(p.PembimbingID)
	acc.SetProfileCreatedAt(p.CreatedAt)
	acc.SetProfileUpdatedAt(p.UpdatedAt)
}

func (p *Profile) Fill(acc *account.Account) {
	p.UserID = acc.GetID()
	p.Nim = acc.GetNim()
	p.Nip = acc.GetNip()
	p.Nidn = acc.GetNidn()
	p.NoRegBlu = acc.GetNoRegBlu()
	p.Name = acc.GetName()
	p.Sex = acc.GetSex()
	p.Religion = acc.GetReligion()
	p.Phone = acc.GetPhone()
	p.Address = acc.GetAddress()
	p.BirthPlace = acc.GetBirthPlace()
	p.BirthDate = acc.GetBirthDate()
	p.ProdiID = acc.GetProdiID()
	p.ConcentrationID = acc.GetConcentrationID()
	p.PembimbingID = acc.GetPembimbingID()
	p.CreatedAt = acc.GetProfileCreatedAt()
	p.UpdatedAt = acc.GetProfileUpdatedAt()
}

func (m *GormPostgreRepository) CreateOrUpdateProfile(acc *account.Account) error {
	var user User
	err := m.db.Where("id = ?", acc.GetID()).First(&user).Error
	if err != nil {
		return err
	}

	var profile Profile
	err = m.db.Where("user_id = ?", user.ID).First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			profile.Fill(acc)
			return m.db.Transaction(func(tx *gorm.DB) error {
				err := tx.Create(&profile).Error
				if err != nil {
					return err
				}

				acc.SetProdiID(&profile.ID)
				return nil
			})
		}

		return err
	}

	profile.Fill(acc)
	err = m.db.Where("id = ?", profile.ID).Save(profile).Error
	return err
}

func (m *GormPostgreRepository) GetProfileByUserID(ID uint64) (account.Account, error) {
	var acc account.Account
	var profile Profile

	err := m.db.Where("user_id = ?", ID).Preload("User").First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return acc, meerrors.ErrRecordNotFound
		}

		return acc, err
	}

	profile.User.ToAggregate(&acc)
	profile.ToAggregate(&acc)

	return acc, nil
}
