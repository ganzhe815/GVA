package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/entity"
	"server/model/request"
	"server/model/response"
	"server/utils"
)

func (b *BaseApi) Login(c *gin.Context) {
	var body request.Login
	err := c.ShouldBindJSON(&body)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(body, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &entity.SysUser{Username: body.Username, Password: body.Password}
	user, err := userService.Login(u)
	if err != nil {
		global.LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
		return
	}
	if user.Enable != 1 {
		global.LOG.Error("登陆失败! 用户被禁止登录!")
		response.FailWithMessage("用户被禁止登录", c)
		return
	}

}

// TokenNext 登录以后签发jwt
/*func (b *BaseApi) TokenNext(c *gin.Context, user entity.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.CONFIG.Server.UseMultipoint {
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}*/
