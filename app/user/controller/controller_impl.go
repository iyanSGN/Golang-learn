package controller

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	dto "rearrange/app/user"
	"rearrange/app/user/repository"
	"rearrange/app/user/service"
	"rearrange/helpers"
	"rearrange/models"
	"rearrange/package/response"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nfnt/resize"
)

type controllerImpl struct {
	Service service.Service
}

func NewController(Service service.Service) Controller {
	return &controllerImpl{
		Service: Service,
	}
}

func (co *controllerImpl)GetAll(c echo.Context) error {
	result, err := co.Service.GetAll(c)
	if err != nil {
		return response.ErrorResponse(c,err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get All Admin", result)
}

func (co *controllerImpl)GetByID(c echo.Context) error {
	adminID := c.Param("id")

	id, err := strconv.ParseUint(adminID, 10, 64)
	if err != nil {
		return response.ErrorResponse(c,response.BuildError(response.ErrBadRequest, err))
	}

	result, err :=  co.Service.GetByID(c, uint(id))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get Admin by ID", result)


}


func CreateAdmin(c echo.Context) error {
	var newUser dto.PostUser
	if err := c.Bind(&newUser);

	err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message" : "bad request",
			"error" : err.Error(),
			"status_code" : http.StatusBadRequest,
		})
	}

	if err := helpers.HashedPassword(&newUser); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	var imageProfile []byte
	var formatImage	 string

	file, err := c.FormFile("image_profile")
	if err != nil {
		fmt.Println("no image uploaded")
	} else {
		src, err :=  file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message":     "Failed to open uploaded image",
				"error":       err.Error(),
				"status_code": http.StatusInternalServerError,
			})
		}
		defer src.Close()

		imageProfile, err = io.ReadAll(src)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message":     "Failed to read uploaded image",
				"error":       err.Error(),
				"status_code": http.StatusInternalServerError,
			})
		}

		newWidth := 600
		img, _, err := image.Decode(bytes.NewReader(imageProfile))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message":     "Failed to decode uploaded image",
				"error":       err.Error(),
				"status_code": http.StatusInternalServerError,
			})
		}

		resizedImg :=  resize.Resize(uint(newWidth), 0, img, resize.Lanczos3)
		var resizedImgBuffer bytes.Buffer

		switch filepath.Ext(file.Filename) {
		case ".jpg", ".jpeg":
			if err := jpeg.Encode(&resizedImgBuffer, resizedImg, nil);

			err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message":     "Failed to encode resized image",
					"error":       err.Error(),
					"status_code": http.StatusInternalServerError,
				})
			}

		case ".png" :
			if err := png.Encode(&resizedImgBuffer, resizedImg);
			err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message":     "Failed to encode resized png image",
					"error":       err.Error(),
					"status_code": http.StatusInternalServerError,
				})
			}

		default:
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message":     "Unknown image format",
				"status_code": http.StatusBadRequest,
			})
		}
		imageProfile = resizedImgBuffer.Bytes()
		randomString := GenerateRandomString(10)
		newImageName  := "profiles_pic" + randomString + ".jpg"
		formatImage = newImageName
	}

	userId, err := repository.CreateUser(newUser, imageProfile, formatImage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":     "Failed to create user",
			"error":       err.Error(),
			"status_code": http.StatusInternalServerError,
		})
	}

	go func () error  {
		if userId.FormatProfile == "" {
			user := userId.ID
			err = DefaultPicture(user)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":     "Default picture is set",
			"status_code": http.StatusCreated,
		})
	}()

	

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message":     "User created successfully",
			"status_code": http.StatusCreated,
		})

	}



func UpdateAdmin(c echo.Context) error {
	adminID := c.Param("id")
	id, err := strconv.ParseUint(adminID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updatedAdmin := models.MRegister{}
	if err := c.Bind(&updatedAdmin);
	err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if updatedAdmin.Password != "" {
		if err := helpers.HashPassword(&updatedAdmin);
		err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

	}

	err = repository.UpdateAdmin(uint(id),updatedAdmin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Admin has successfully updated",
		"status_code" :	http.StatusOK,
	})
}


func DeleteAdmin(c echo.Context) error {
	adminID := c.Param("id")
	id, err := strconv.Atoi(adminID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid user id"))
	}

	err = repository.DeleteAdmin(id)
	if err != nil {
	return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Selected user has been deleted",
		"status_code" : http.StatusOK,
	})
}





func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}


func DefaultPicture(id uint) error {
	userId := int(id)
	defaultImagePath := "picture/default.jpg"
	defaultImage, err := os.Open(defaultImagePath)
	if err != nil {
		return err
	}
	defer defaultImage.Close()

	defaultImageData, err := io.ReadAll(defaultImage)
	if err != nil {
		return err
	}

	randomString := GenerateRandomString(10)
	newImageName := "profiles_pic" + randomString + ".jpg"
	formatImage := newImageName

	updatedUser := dto.ImageUser{}

	if err := repository.DefaultPicture(userId, defaultImageData, updatedUser, formatImage); err != nil {
		return err
	}

	return nil
}