package handler

import (
	"MyGram/helper"
	"MyGram/model"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// User Endpoint
func (h HttpServer) UserRegister(c *gin.Context) {
	contentType := helper.GetContentType(c)
	user := model.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	if user.Age < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Belum cukup umur!",
			"status":  http.StatusBadRequest,
		})
		return
	}

	if len(user.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password harus lebih dari 6 karakter!",
			"status":  http.StatusBadRequest,
		})
		return
	}

	_, err := mail.ParseAddress(user.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email tidak valid!",
			"status":  http.StatusBadRequest,
		})
		return
	}

	res, err := h.app.UserRegister(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h HttpServer) UserLogin(c *gin.Context) {
	contentType := helper.GetContentType(c)
	user := model.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	if len(user.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password harus lebih dari 6 karakter!",
			"status":  http.StatusBadRequest,
		})
	}

	passwordClient := user.Password

	res, err := h.app.UserLogin(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "invalid email or password",
		})
		return
	}

	isValid := helper.ComparePass([]byte(res.Password), []byte(passwordClient))
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "invalid email or password",
		})
		return
	}

	token := helper.GenerateToken(uint(res.ID), res.Username)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Photo Endpoint
func (h HttpServer) PhotoGetAll(c *gin.Context) {

	res, err := h.app.PhotoGetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h HttpServer) PhotoGet(c *gin.Context) {

	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var photo model.Photo

	photo.ID = id

	res, err := h.app.PhotoGet(photo)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h HttpServer) PhotoCreate(c *gin.Context) {
	var newPhoto model.Photo

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))

	if err := c.ShouldBindJSON(&newPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	newPhoto.UserID = int(userID)

	res, err := h.app.PhotoCreate(newPhoto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h HttpServer) PhotoUpdate(c *gin.Context) {
	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	var photo model.Photo

	photo.ID = id

	photo, err = h.app.PhotoAuthorization(photo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}

	if photo.UserID != userID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "you're not allowed to access this",
		})

		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	photo.ID = id

	res, err := h.app.PhotoUpdate(photo)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h HttpServer) PhotoDelete(c *gin.Context) {
	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	var photo model.Photo

	photo.ID = id

	photo, err = h.app.PhotoAuthorization(photo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}

	if photo.UserID != userID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "you're not allowed to access this",
		})

		return
	}

	photo.ID = id

	err = h.app.PhotoDelete(photo)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Photo deleted successfully",
	})
}

// Comment Endpoint
func (h HttpServer) CommentGetAll(c *gin.Context) {

	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var comment model.Comment

	comment.ID = id

	res, err := h.app.CommentGetAll(comment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h HttpServer) CommentGet(c *gin.Context) {

	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var comment model.Comment

	comment.ID = id

	res, err := h.app.CommentGet(comment)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h HttpServer) CommentCreate(c *gin.Context) {

	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var newComment model.Comment

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))

	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	newComment.UserID = int(userID)
	newComment.PhotoID = id

	res, err := h.app.CommentCreate(newComment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h HttpServer) CommentUpdate(c *gin.Context) {
	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	var comment model.Comment

	comment.ID = id

	comment, err = h.app.CommentAuthorization(comment)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}

	if comment.UserID != userID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "you're not allowed to access this",
		})

		return
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	comment.ID = id

	res, err := h.app.CommentUpdate(comment)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h HttpServer) CommentDelete(c *gin.Context) {
	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	var comment model.Comment

	comment.ID = id

	comment, err = h.app.CommentAuthorization(comment)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}

	if comment.UserID != userID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "you're not allowed to access this",
		})

		return
	}

	comment.ID = id

	err = h.app.CommentDelete(comment)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Comment deleted successfully",
	})
}

// SocialMedia Endpoint
func (h HttpServer) SocialMediaGetAll(c *gin.Context) {

	res, err := h.app.SocialMediaGetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h HttpServer) SocialMediaGet(c *gin.Context) {

	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var photo model.SocialMedia

	photo.ID = id

	res, err := h.app.SocialMediaGet(photo)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h HttpServer) SocialMediaCreate(c *gin.Context) {
	var newSocialMedia model.SocialMedia

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))

	if err := c.ShouldBindJSON(&newSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	newSocialMedia.UserID = int(userID)

	res, err := h.app.SocialMediaCreate(newSocialMedia)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h HttpServer) SocialMediaUpdate(c *gin.Context) {
	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	var photo model.SocialMedia

	photo.ID = id

	photo, err = h.app.SocialMediaAuthorization(photo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}

	if photo.UserID != userID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "you're not allowed to access this",
		})

		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	photo.ID = id

	res, err := h.app.SocialMediaUpdate(photo)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h HttpServer) SocialMediaDelete(c *gin.Context) {
	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	var photo model.SocialMedia

	photo.ID = id

	photo, err = h.app.SocialMediaAuthorization(photo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}

	if photo.UserID != userID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "you're not allowed to access this",
		})

		return
	}

	photo.ID = id

	err = h.app.SocialMediaDelete(photo)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Social Media deleted successfully",
	})
}
