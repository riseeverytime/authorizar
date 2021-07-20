package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yauthdev/yauth/server/constants"
	"github.com/yauthdev/yauth/server/db"
	"github.com/yauthdev/yauth/server/enum"
	"github.com/yauthdev/yauth/server/session"
	"github.com/yauthdev/yauth/server/utils"
)

// Defining the Graphql handler
func VerifyEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		errorRes := gin.H{
			"message": "invalid token",
		}
		token := c.Query("token")
		if token == "" {
			c.JSON(400, errorRes)
			return
		}

		_, err := db.Mgr.GetVerificationByToken(token)
		if err != nil {
			c.JSON(400, errorRes)
			return
		}

		// verify if token exists in db
		claim, err := utils.VerifyVerificationToken(token)
		if err != nil {
			c.JSON(400, errorRes)
			return
		}

		user, err := db.Mgr.GetUserByEmail(claim.Email)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		// update email_verified_at in users table
		db.Mgr.UpdateVerificationTime(time.Now().Unix(), user.ID)
		// delete from verification table
		db.Mgr.DeleteToken(claim.Email)

		userIdStr := fmt.Sprintf("%d", user.ID)
		refreshToken, _, _ := utils.CreateAuthToken(utils.UserAuthInfo{
			ID:    userIdStr,
			Email: user.Email,
		}, enum.RefreshToken)

		accessToken, _, _ := utils.CreateAuthToken(utils.UserAuthInfo{
			ID:    userIdStr,
			Email: user.Email,
		}, enum.AccessToken)

		session.SetToken(userIdStr, refreshToken)
		utils.SetCookie(c, accessToken)
		c.Redirect(http.StatusTemporaryRedirect, constants.FRONTEND_URL)
	}
}
