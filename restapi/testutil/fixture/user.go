package fixture

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/kngnkg/tunetrail/restapi/clock"
	"github.com/kngnkg/tunetrail/restapi/model"
)

// Userはテスト用のユーザーを作成する
func User(u *model.User) *model.User {
	random := strconv.Itoa(rand.Int())[:3]

	result := &model.User{
		Id:        uuid.New().String(),
		UserName:  "test" + random,
		Name:      "test" + random,
		IconUrl:   "test" + random,
		Bio:       "test" + random,
		Email:     "test" + random + "@example.com",
		Password:  "test" + random,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if u == nil {
		return result
	}

	// 初期値がある場合は上書きする
	if u.Id != "" {
		result.Id = u.Id
	}
	if u.UserName != "" {
		result.UserName = u.UserName
	}
	if u.Name != "" {
		result.Name = u.Name
	}
	if u.IconUrl != "" {
		result.IconUrl = u.IconUrl
	}
	if u.Bio != "" {
		result.Bio = u.Bio
	}
	if u.Email != "" {
		result.Email = u.Email
	}
	if u.Password != "" {
		result.Password = u.Password
	}
	if !u.CreatedAt.IsZero() {
		result.CreatedAt = u.CreatedAt
	}
	if !u.UpdatedAt.IsZero() {
		result.UpdatedAt = u.UpdatedAt
	}
	return result
}

func CreateUsers(n int) []*model.User {
	fc := &clock.FixedClocker{}

	users := make([]*model.User, n)
	for i := 0; i < n; i++ {
		users[i] = User(&model.User{
			// タイムスタンプを固定する
			CreatedAt: fc.Now(),
			UpdatedAt: fc.Now(),
		})
	}
	return users
}
