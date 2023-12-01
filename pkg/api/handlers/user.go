package handlers

import (
	"net/http"
	"strconv"

	"github.com/Anandhu4456/go-Ecommerce/pkg/helper"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userusecase services.UserUsecase
}

// Constructor function
func NewUserHandler(userUsecase services.UserUsecase) *UserHandler {
	return &UserHandler{
		userusecase: userUsecase,
	}
}

func (uH *UserHandler) AddAddress(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var address models.AddAddress
	if err := c.BindJSON(&address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := uH.userusecase.AddAddress(userId, address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't add address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added address", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (uH *UserHandler) ChangePassword(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't find user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var changePass models.ChangePassword

	if err := c.BindJSON(&changePass); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := uH.userusecase.ChangePassword(userId, changePass.OldPassword, changePass.NewPassword, changePass.RePassword); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't change the password", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "password changed successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (uH *UserHandler) EditUser(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var userData models.EditUser
	if err := c.BindJSON(&userData); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := uH.userusecase.EditUser(userId, userData); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't change the user details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully changed user details", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (uH *UserHandler) GetAddresses(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	addresses, err := uH.userusecase.GetAddresses(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully got all addresses", addresses, nil)
	c.JSON(http.StatusOK, successRes)
}

func (uH *UserHandler) GetCart(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "limit number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	products, err := uH.userusecase.GetCart(userId, page, limit)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't retrieve cart products", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully got all products in cart", products, nil)
	c.JSON(http.StatusOK, successRes)
}

func (uH *UserHandler) GetUserDetails(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userDetails, err := uH.userusecase.GetUserDetails(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully got user details", userDetails, nil)
	c.JSON(http.StatusOK, successRes)
}

func (uH *UserHandler) Login(c *gin.Context) {
	var user models.UserLogin
	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userToken, err := uH.userusecase.Login(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user couldn't login", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "user successfully logged in", userToken, nil)
	// c.SetCookie("Authorization",userToken.Token,3600,"/","yoursstore.online",true,false)
	c.SetCookie("Authorization", userToken.Token, 3600, "", "", true, false)
	c.JSON(http.StatusOK, successRes)
}

func (uH *UserHandler) RemoveFromCart(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't find user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	cartId, err := uH.userusecase.GetCartID(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get cart id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	invId, err := strconv.Atoi(c.Query("inventory_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := uH.userusecase.RemoveFromCart(cartId, invId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "remove from cart failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully removed from cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (uH *UserHandler) SignUp(c *gin.Context) {
	var user models.UserDetails

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userToken, err := uH.userusecase.SignUp(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't signup user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully signed up", userToken, nil)
	c.JSON(http.StatusOK, successRes)
}

func (uH *UserHandler) UpdateQuantityAdd(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	cartId, err := uH.userusecase.GetCartID(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get cart id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	invId, err := strconv.Atoi(c.Query("inventory_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := uH.userusecase.UpdateQuantityAdd(cartId, invId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't update quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added quantity", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
