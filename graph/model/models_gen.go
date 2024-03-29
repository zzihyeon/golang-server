// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type StuffInterface interface {
	IsStuffInterface()
}

type StuffResult interface {
	IsStuffResult()
}

type UserInterface interface {
	IsUserInterface()
}

type AdminUser struct {
	UID       int       `json:"uid"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Gender    Gender    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Level     int       `json:"level"`
}

func (AdminUser) IsUserInterface() {}

type BizInfo struct {
	BizNum string `json:"biz_num"`
	BizImg string `json:"biz_img"`
}

type BuyerUser struct {
	UID           int       `json:"uid"`
	Email         string    `json:"email"`
	Name          string    `json:"name"`
	Phone         string    `json:"phone"`
	Gender        Gender    `json:"gender"`
	BirthDate     time.Time `json:"birth_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Purchases     []*Stuff  `json:"purchases"`
	ShoppingCarts []*Stuff  `json:"shopping_carts"`
}

func (BuyerUser) IsUserInterface() {}

type Food struct {
	Name      string    `json:"name"`
	Kind      string    `json:"kind"`
	Price     int       `json:"price"`
	ShelfLife time.Time `json:"shelf_life"`
}

func (Food) IsStuffInterface() {}
func (Food) IsStuffResult()    {}

type Furniture struct {
	Name        string `json:"name"`
	Kind        string `json:"kind"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

func (Furniture) IsStuffInterface() {}
func (Furniture) IsStuffResult()    {}

type SellerUser struct {
	UID       int        `json:"uid"`
	Email     string     `json:"email"`
	Name      string     `json:"name"`
	Phone     string     `json:"phone"`
	Gender    Gender     `json:"gender"`
	BirthDate time.Time  `json:"birth_date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	BizInfo   []*BizInfo `json:"bizInfo"`
	Stuff     []*Stuff   `json:"stuff"`
}

func (SellerUser) IsUserInterface() {}

type Stuff struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Available int    `json:"available"`
	Total     int    `json:"total"`
}

type User struct {
	UID       int       `json:"uid"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Gender    Gender    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserInput struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Gender    Gender    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
}

type DeleteUserInput struct {
	UID int64 `json:"uid"`
}

type UpdateUserInput struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Gender    Gender    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
}

type UserInput struct {
	UID int64 `json:"uid"`
}

type Gender string

const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
)

var AllGender = []Gender{
	GenderMale,
	GenderFemale,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
