package usecase

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/RyuichiroYoshida/quest-board-project/internal/auth/domain"
	"github.com/RyuichiroYoshida/quest-board-project/utils"
)

type AuthRepository interface {
	CreateUser(user *domain.User) error
	ReadUser(id string) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id string) error
}

type AuthUsecase struct {
	repo AuthRepository
}

func NewAuthUsecase(repo AuthRepository) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
	}
}

// Loginは、ユーザーの認証を行い、必要に応じてユーザーをデータベースに登録する
func (u *AuthUsecase) Login(user *domain.User) error {
	if user == nil || !user.IsValid() {
		utils.LogWarning("authUsecase.Login: user is nil or invalid")
		return errors.New("user is nil or invalid")
	}

	existing, err := u.repo.ReadUser(user.Id)
	if err != nil {
		utils.LogWarning("authUsecase.Login: failed to read user from repository")
		return err
	}
	if existing == nil {
		if err := u.repo.CreateUser(user); err != nil {
			utils.LogWarning("authUsecase.Login: failed to create user")
			return err
		}
	}
	return nil
}

// TODO: Logoutは、ユーザーのセッションを終了する
func (u *AuthUsecase) Logout(id string) error {
	return nil
}

// TODO: セッションを元にユーザー情報を取得する
func (u *AuthUsecase) Me(id string) (*domain.User, error) {
	return u.repo.ReadUser(id)
}

// Discordの認証ページへのリダイレクトURLを生成する
func (u *AuthUsecase) RedirectAuthPage(clientId, redirectUri string, scopes ...string) string {
	authURL := "https://discord.com/api/oauth2/authorize" +
		"?client_id=" + clientId +
		"&redirect_uri=" + url.QueryEscape(redirectUri) +
		"&response_type=code"

	if len(scopes) != 0 {
		authURL += "&scope="
		for i, scope := range scopes {
			authURL += scope
			if i < len(scopes)-1 {
				authURL += "%20"
			}
		}
	}
	return authURL
}

// DiscordのトークンエンドポイントにcodeをPOSTし、アクセストークンを取得
func (u *AuthUsecase) ExchangeCode(code, clientId, clientSecret, redirectUri string) (map[string]any, error) {
	if code == "" || clientId == "" || clientSecret == "" || redirectUri == "" {
		utils.LogWarning("authUsecase.ExchangeCode: missing required parameters")
		return nil, nil
	}

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)

	resp, err := http.PostForm("https://discord.com/api/oauth2/token", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Discordのユーザー情報取得
func (u *AuthUsecase) GetDiscordUserInfo(accessToken string) (map[string]any, error) {
	if accessToken == "" {
		return nil, nil
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://discord.com/api/users/@me", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var userInfo map[string]any
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}
