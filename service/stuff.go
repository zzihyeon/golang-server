package service

import (
	"time"

	"github.com/zzihyeon/golang-server.git/graph/model"
)

func GetAllStuff() ([]model.StuffResult, error) {
	return []model.StuffResult{
		model.Food{
			Name:      "버팔로 윙",
			Kind:      "chicken",
			Price:     1000,
			ShelfLife: time.Now(),
		},
		model.Furniture{
			Name:        "허먼 밀러",
			Kind:        "chair",
			Price:       2000000,
			Description: "매우 편한 의자",
		},
	}, nil
}
