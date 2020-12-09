package v1

import (
	"net/http"
	"os"

	"github.com/ahmdrz/goinsta/v2"
	"github.com/arthurhung/instagram-tidyup/models"
	"github.com/arthurhung/instagram-tidyup/pkg/app"
	"github.com/arthurhung/instagram-tidyup/pkg/e"
	"github.com/arthurhung/instagram-tidyup/pkg/logging"
	"github.com/arthurhung/instagram-tidyup/service/instidyup"
	"github.com/gin-gonic/gin"
)

// LoginForm ...
type LoginForm struct {
	Username string `form:"username" json:"username" valid:"Required;MaxSize(100)"`
	Password string `form:"password" json:"password" valid:"Required;MaxSize(100)"`
}

type loginResp struct {
	Token          string `json:"token"`
	FollowingCount int    `json:"followingCount"`
}

// LoginIG ...
func LoginIG(c *gin.Context) {
	var (
		insta *goinsta.Instagram
		appG  = app.Gin{C: c}
		form  LoginForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	logging.Info(form.Username, form.Password)
	exist, err := models.CheckUserExist(form.Username)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, err)
		return
	}
	if exist == false {
		err := models.CreateUser(form.Username)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER_FAIL, nil)
			return
		}
	}

	encodedAccount := os.Getenv("IG")
	if encodedAccount != "" {
		insta, _ = instidyup.ImportFromBase64String(encodedAccount)

	} else {
		insta = instidyup.GetInsta(form.Username, form.Password)
		instidyup.StoreSession(insta)
	}

	if insta == nil {
		appG.Response(http.StatusBadRequest, e.ERROR_LOGIN_FAIL, nil)
	}
	following := insta.Account.Following()

	var followingUsers []goinsta.User

	for following.Next() {
		for _, user := range following.Users {
			user.Mute("all")
			followingUsers = append(followingUsers, user)

		}
	}

	appG.Response(http.StatusOK, e.SUCCESS, loginResp{
		Token:          insta.Account.Username,
		FollowingCount: len(followingUsers),
	})
}
