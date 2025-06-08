package infrastructure

import (
	"github.com/RyuichiroYoshida/quest-board-project/internal/auth/domain"
	"github.com/RyuichiroYoshida/quest-board-project/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindOrCreateUser(user *domain.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) FindOrCreateUser(user *domain.User) error {
	// domain.User -> models.User への変換
	mUser := models.User{
		Id:     user.Id,
		Name:   user.Name,
		Avatar: user.Avatar,
		// DiscordId, CreatedAtは外部から渡す場合はuserに追加、またはここで生成
	}
	// DiscordIdは必須なので、user.Nameなどに一時的に格納している場合は修正が必要
	// ここでは例としてNameをDiscordIdに流用（本来はdomain.UserにDiscordIdを追加すべき）
	mUser.DiscordId = user.Id // 仮の割り当て
	mUser.CreatedAt = ""      // 必要に応じてセット

	var existing models.User
	err := r.db.Where("discord_id = ?", mUser.DiscordId).First(&existing).Error
	if err == gorm.ErrRecordNotFound {
		return r.db.Create(&mUser).Error
	}
	return err
}
