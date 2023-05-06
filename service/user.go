package service

import (
	"github.com/zzihyeon/golang-server.git/graph/model"
)

func GetUser() ([]model.User, error) {
	return []model.User{
		{
			UID:    1,
			Email:  "test@sk.com",
			Name:   "zzihyeon",
			Phone:  "010-0000-0000",
			Gender: model.GenderMale,
		},
	}, nil
}

func GetAllUser() ([]model.UserInterface, error) {
	return []model.UserInterface{
		model.AdminUser{
			UID:    1,
			Email:  "admin@sk.com",
			Name:   "zzihyeon",
			Phone:  "010-0000-0000",
			Gender: model.GenderMale,
			Level:  0,
		},
		model.SellerUser{
			UID:    2,
			Email:  "seller@sk.com",
			Name:   "zzihyeon",
			Phone:  "010-0000-0000",
			Gender: model.GenderMale,
			BizInfo: []*model.BizInfo{{
				BizNum: "1234",
				BizImg: "http://img.png",
			}},
			Stuff: []*model.Stuff{{
				ID:        "1",
				Name:      "cheeze",
				Available: 10,
				Total:     12,
			}},
		},
		model.BuyerUser{
			UID:    3,
			Email:  "buyer@sk.com",
			Name:   "zzihyeon",
			Phone:  "010-0000-0000",
			Gender: model.GenderMale,
			Purchases: []*model.Stuff{{
				ID:        "1",
				Name:      "cheeze",
				Available: 10,
				Total:     12,
			}},
			ShoppingCarts: []*model.Stuff{{
				ID:        "1",
				Name:      "cheeze",
				Available: 10,
				Total:     12,
			}},
		},
	}, nil
}

func GetAdminUser() ([]*model.AdminUser, error) {
	return []*model.AdminUser{
		{
			UID:    1,
			Email:  "test@sk.com",
			Name:   "zzihyeon",
			Phone:  "010-0000-0000",
			Gender: model.GenderMale,
			Level:  0,
		},
	}, nil
}

func GetSellerUser() ([]*model.SellerUser, error) {
	return []*model.SellerUser{
		{
			UID:    1,
			Email:  "test@sk.com",
			Name:   "zzihyeon",
			Phone:  "010-0000-0000",
			Gender: model.GenderMale,
			BizInfo: []*model.BizInfo{{
				BizNum: "1234",
				BizImg: "http://img.png",
			}},
			Stuff: []*model.Stuff{{
				ID:        "1",
				Name:      "cheeze",
				Available: 10,
				Total:     12,
			}},
		}}, nil
}

func GetBuyerUsers() ([]*model.BuyerUser, error) {
	return []*model.BuyerUser{
		{
			UID:    1,
			Email:  "test@sk.com",
			Name:   "zzihyeon",
			Phone:  "010-0000-0000",
			Gender: model.GenderMale,
			Purchases: []*model.Stuff{{
				ID:        "1",
				Name:      "cheeze",
				Available: 10,
				Total:     12,
			}},
			ShoppingCarts: []*model.Stuff{{
				ID:        "1",
				Name:      "cheeze",
				Available: 10,
				Total:     12,
			}},
		}}, nil
}
