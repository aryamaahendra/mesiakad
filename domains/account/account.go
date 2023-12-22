package account

import (
	"time"

	"github.com/aryamaahendra/mesiakad/domains/mesiakad"
)

type Account struct {
	user    *mesiakad.User
	profile *mesiakad.Profile
	token   string
}

func New(userID uint64) *Account {

	return &Account{
		user:    &mesiakad.User{ID: userID},
		profile: &mesiakad.Profile{},
		token:   "",
	}
}

func (ac *Account) ifUserEmpty() {
	if ac.user == nil {
		ac.user = &mesiakad.User{}
	}
}

func (ac *Account) ifProfileEmpty() {
	if ac.profile == nil {
		ac.profile = &mesiakad.Profile{}
	}
}

func (ac *Account) SetUser(user *mesiakad.User) {
	ac.user = user
}

func (ac *Account) GetUser() *mesiakad.User {
	return ac.user
}

func (ac *Account) SetID(id uint64) {
	ac.ifUserEmpty()
	ac.user.ID = id
}

func (ac *Account) GetID() uint64 {
	return ac.user.ID
}

func (ac *Account) SetEmail(email string) {
	ac.ifUserEmpty()
	ac.user.Email = email
}

func (ac *Account) GetEmail() string {
	return ac.user.Email
}

func (ac *Account) SetUsername(username string) {
	ac.ifUserEmpty()
	ac.user.Username = username
}

func (ac *Account) GetUsername() string {
	return ac.user.Username
}

func (ac *Account) SetPassword(password string) {
	ac.ifUserEmpty()
	ac.user.Password = password
}

func (ac *Account) GetPassword() string {
	return ac.user.Password
}

func (ac *Account) SetRoleID(roleID *uint64) {
	ac.ifUserEmpty()
	ac.user.RoleID = roleID
}

func (ac *Account) GetRoleID() *uint64 {
	return ac.user.RoleID
}

func (ac *Account) SetUserCreateAt(createAt time.Time) {
	ac.ifUserEmpty()
	ac.user.CreatedAt = createAt
}

func (ac *Account) GetUserCreateAt() time.Time {
	return ac.user.CreatedAt
}

func (ac *Account) SetUserUpdatedAt(updateAt time.Time) {
	ac.ifUserEmpty()
	ac.user.UpdatedAt = updateAt
}

func (ac *Account) GetUserUpdatedAt() time.Time {
	return ac.user.UpdatedAt
}

func (p *Account) SetProfileID(id uint64) {
	p.ifProfileEmpty()
	p.profile.ID = id
}

func (p *Account) GetProfileID() uint64 {
	return p.profile.ID
}

func (p *Account) SetNim(nim string) {
	p.ifProfileEmpty()
	p.profile.Nim = nim
}

func (p *Account) GetNim() string {
	return p.profile.Nim
}

func (p *Account) SetNip(nip string) {
	p.ifProfileEmpty()
	p.profile.Nip = nip
}

func (p *Account) GetNip() string {
	return p.profile.Nip
}

func (p *Account) SetNidn(nidn string) {
	p.ifProfileEmpty()
	p.profile.Nidn = nidn
}

func (p *Account) GetNidn() string {
	return p.profile.Nidn
}

func (p *Account) SetNoRegBlu(noRegBlu string) {
	p.ifProfileEmpty()
	p.profile.NoRegBlu = noRegBlu
}

func (p *Account) GetNoRegBlu() string {
	return p.profile.NoRegBlu
}

func (p *Account) SetName(name string) {
	p.ifProfileEmpty()
	p.profile.Name = name
}

func (p *Account) GetName() string {
	return p.profile.Name
}

func (p *Account) SetSex(sex string) {
	p.ifProfileEmpty()
	p.profile.Sex = sex
}

func (p *Account) GetSex() string {
	return p.profile.Sex
}

func (p *Account) SetReligion(religion string) {
	p.ifProfileEmpty()
	p.profile.Religion = religion
}

func (p *Account) GetReligion() string {
	return p.profile.Religion
}

func (p *Account) SetPhone(phone string) {
	p.ifProfileEmpty()
	p.profile.Phone = phone
}

func (p *Account) GetPhone() string {
	return p.profile.Phone
}

func (p *Account) SetAddress(address string) {
	p.ifProfileEmpty()
	p.profile.Address = address
}

func (p *Account) GetAddress() string {
	return p.profile.Address
}

func (p *Account) SetBirthPlace(birthPlace string) {
	p.ifProfileEmpty()
	p.profile.BirthPlace = birthPlace
}

func (p *Account) GetBirthPlace() string {
	return p.profile.BirthPlace
}

func (p *Account) SetBirthDate(birthDate time.Time) {
	p.ifProfileEmpty()
	p.profile.BirthDate = birthDate
}

func (p *Account) GetBirthDate() time.Time {
	return p.profile.BirthDate
}

func (p *Account) SetProdiID(prodiID *uint64) {
	p.ifProfileEmpty()
	p.profile.ProdiID = prodiID
}

func (p *Account) GetProdiID() *uint64 {
	return p.profile.ProdiID
}

func (p *Account) SetConcentrationID(concentrationID *uint64) {
	p.ifProfileEmpty()
	p.profile.ConcentrationID = concentrationID
}

func (p *Account) GetConcentrationID() *uint64 {
	return p.profile.ConcentrationID
}

func (p *Account) SetPembimbingID(pembimbingID *uint64) {
	p.ifProfileEmpty()
	p.profile.PembimbingID = pembimbingID
}

func (p *Account) GetPembimbingID() *uint64 {
	return p.profile.PembimbingID
}

func (p *Account) SetProfileCreatedAt(createdAt time.Time) {
	p.ifProfileEmpty()
	p.profile.CreatedAt = createdAt
}

func (p *Account) GetProfileCreatedAt() time.Time {
	return p.profile.CreatedAt
}

func (p *Account) SetProfileUpdatedAt(updatedAt time.Time) {
	p.ifProfileEmpty()
	p.profile.UpdatedAt = updatedAt
}

func (p *Account) GetProfileUpdatedAt() time.Time {
	return p.profile.UpdatedAt
}
