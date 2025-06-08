package handler

import (
	"github.com/RyuichiroYoshida/quest-board-project/internal/auth/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	LoginDiscord(c *gin.Context)
	ExchangeCode(c *gin.Context)
}

type authHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(u usecase.AuthUsecase) *authHandler {
	return &authHandler{u}
}

func (h *authHandler) LoginDiscord(c *gin.Context) {
	clientId := "YOUR_CLIENT_ID"       // 環境変数や設定から取得してください
	redirectUri := "YOUR_REDIRECT_URI" // 環境変数や設定から取得してください
	scopes := []string{"identify"}
	authURL := h.authUsecase.RedirectAuthPage(clientId, redirectUri, scopes...)
	c.Redirect(302, authURL)
}

func (h *authHandler) ExchangeCode(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(400, gin.H{"error": "code is required"})
		return
	}
	clientId := "YOUR_CLIENT_ID"         // 環境変数や設定から取得してください
	clientSecret := "YOUR_CLIENT_SECRET" // 環境変数や設定から取得してください
	redirectUri := "YOUR_REDIRECT_URI"   // 環境変数や設定から取得してください

	tokenResp, err := h.authUsecase.ExchangeCode(code, clientId, clientSecret, redirectUri)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to exchange code"})
		return
	}
	accessToken, ok := tokenResp["access_token"].(string)
	if !ok || accessToken == "" {
		c.JSON(500, gin.H{"error": "failed to get access token"})
		return
	}
	userInfo, err := h.authUsecase.GetDiscordUserInfo(accessToken)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get user info"})
		return
	}
	// ここでuserInfoからUser構造体を生成し、DB登録やJWT発行などを行う
	// 例: user := domain.UserFromDiscord(userInfo)
	//     h.authUsecase.Login(&user)
	//     token := generateJWT(user)
	// 今回は簡易的にuserInfoを返却
	c.JSON(200, gin.H{"user": userInfo, "access_token": accessToken})
}
