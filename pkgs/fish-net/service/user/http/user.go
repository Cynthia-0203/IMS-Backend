package http

import (
	"fishnet/domain"
	"fishnet/service/user/pack"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var req pack.CreateUserRequest
	print(1)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"msg": "bad request",
		})
		return
	}
	print(2)
	if len(req.Username) < 3 {
		c.JSON(400, gin.H{
			"msg": "bad request",
		})
		return
	}
	print(3)
	users, err := _userUsecase.QueryUser(nil, &req.Username, nil, 1, 0)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "bad request",
		})
		return
	}
	if len(users) > 0 {
		c.JSON(400, gin.H{
			"msg": "user already exists",
		})
		return
	}
	user := &domain.User{
		Username: req.Username,
		Nickname: req.Username,
		Icon:     "https://pics.com/avatar.png",
	}
	err = _userUsecase.CreateUser([]*domain.User{user})
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "ok",
		"data": gin.H{
			"id": user.ID,
		},
	})
}

func DeleteUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "bad request",
		})
		return
	}
	err = _userUsecase.DeleteUser(userId)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "ok",
	})
}

func UpdateUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "bad request",
		})
		return
	}
	var req *pack.UpdateUserRequest
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, gin.H{
			"msg": "bad request",
		})
		return
	}
	isExists, err := _userUsecase.CheckUserExist(&userId, nil)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "internal server error",
		})
		return
	}
	if !isExists {
		c.JSON(404, gin.H{
			"msg": "not found",
		})
		return
	}
	err = _userUsecase.UpdateUser(userId, req.Nickname, req.Icon)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "ok",
	})
}

func QueryUser(c *gin.Context) {
	var req pack.QueryUserRequest
	if userIDStr := c.Param("userId"); userIDStr != "" {
		userId, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"msg": "bad request",
			})
			return
		}
		users, err := _userUsecase.QueryUser(&userId, nil, nil, 1, 0)
		if err != nil {
			c.JSON(500, gin.H{
				"msg": "internal server error",
			})
			return
		}
		if len(users) == 0 {
			c.JSON(404, gin.H{
				"msg": "not found",
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": "ok",
			"data": gin.H{
				"list":  pack.Users(users),
				"total": len(users),
			},
		})
		return
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{
			"msg": "bad request",
		})
		return
	}
	users, err := _userUsecase.QueryUser(&req.UserID, &req.Username, &req.Nickname, req.Limit, req.Offset)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "ok",
		"data": gin.H{
			"list":  pack.Users(users),
			"total": len(users),
		},
	})
}
