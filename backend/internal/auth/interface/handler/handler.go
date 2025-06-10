package handler

import (
	"os"

	"github.com/RyuichiroYoshida/quest-board-project/internal/auth/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	LoginDiscord(c *gin.Context)
	ExchangeCode(c *gin.Context)
	Me(c *gin.Context)
	Logout(c *gin.Context)
}

type authHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(u usecase.AuthUsecase) *authHandler {
	return &authHandler{u}
}

func (h *authHandler) LoginDiscord(c *gin.Context) {
	clientId := os.Getenv("DISCORD_CLIENT_ID")
	redirectUri := os.Getenv("DISCORD_REDIRECT_URI")
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
	clientId := os.Getenv("DISCORD_CLIENT_ID")
	clientSecret := os.Getenv("DISCORD_CLIENT_SECRET")
	redirectUri := os.Getenv("DISCORD_REDIRECT_URI")
	if clientId == "" || clientSecret == "" || redirectUri == "" {
		c.JSON(500, gin.H{"error": "missing required environment variables"})
		return
	}

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
