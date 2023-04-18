package fixture

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/kwtryo/tunetrail/api/model"
)

// Userはテスト用のユーザーを作成する
func User(u *model.User) *model.User {
	random := strconv.Itoa(rand.Int())[:3]

	result := &model.User{
		Id:        rand.Int(),
		UserName:  "test" + random,
		Name:      "test" + random,
		Password:  "test" + random,
		Email:     "test" + random + "@example.com",
		IconUrl:   "test" + random,
		Bio:       "test" + random,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if u == nil {
		return result
	}

	// 初期値がある場合は上書きする
	if u.Id != 0 {
		result.Id = u.Id
	}
	if u.UserName != "" {
		result.UserName = u.UserName
	}
	if u.Name != "" {
		result.Name = u.Name
	}
	if u.Password != "" {
		result.Password = u.Password
	}
	if u.Email != "" {
		result.Email = u.Email
	}
	if u.IconUrl != "" {
		result.IconUrl = u.IconUrl
	}
	if u.Bio != "" {
		result.Bio = u.Bio
	}
	if !u.CreatedAt.IsZero() {
		result.CreatedAt = u.CreatedAt
	}
	if !u.UpdatedAt.IsZero() {
		result.UpdatedAt = u.UpdatedAt
	}
	return result
}
