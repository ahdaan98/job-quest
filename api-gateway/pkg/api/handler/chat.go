package handler

import (
	interfaces "github.com/ahdaan67/JobQuest/pkg/client/interface"
	"github.com/ahdaan67/JobQuest/pkg/helper"
	"github.com/ahdaan67/JobQuest/pkg/utils/models"
	"github.com/ahdaan67/JobQuest/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var User = make(map[string]*websocket.Conn)

type ChatHandler struct {
	GRPC_Client interfaces.ChatClient
	helper      *helper.Helper
}

func NewChatHandler(chatClient interfaces.ChatClient, helper *helper.Helper) *ChatHandler {
	return &ChatHandler{
		GRPC_Client: chatClient,
		helper:      helper,
	}
}

// WebSocket
func (ch *ChatHandler) EmployerMessage(c *gin.Context) {
	fmt.Println("++== call hit in message funtion")
	tokenString := c.Request.Header.Get("Authorization")

	splitToken := strings.Split(tokenString, " ")
	if tokenString == "" {
		errs := response.ClientResponse(http.StatusUnauthorized, "Missing Authorization header", nil, "")
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	splitToken[1] = strings.TrimSpace(splitToken[1])
	userID, err := ch.helper.ValidateToken(splitToken[1])

	if err != nil {
		errs := response.ClientResponse(http.StatusUnauthorized, "Invalid token", nil, err.Error())
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Websocket Connection Issue", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	defer func() {
		delete(User, strconv.Itoa(int(userID.Id)))
		conn.Close()
	}()

	user := strconv.Itoa(int(userID.Id))
	User[user] = conn

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}

		ch.helper.SendMessageToUser(User, msg, user)
	}
}

// GetChat handles the HTTP request to retrieve chat details.
//
// @Summary Retrieve chat details
// @Description Retrieves chat details based on the provided request
// @Tags Chat
// @Accept json
// @Produce json
// @Param body body models.ChatRequest true "Chat request details"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "Successfully retrieved chat details"
// @Failure 400 {object} response.Response{} "Details not in correct format" or "User ID not found in JWT claims" or "Failed to get chat details"
// @Router /employer/chats [post]
func (ch *ChatHandler) GetChat(c *gin.Context) {
	var chatRequest models.ChatRequest
	if err := c.ShouldBindJSON(&chatRequest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIDInterface, exists := c.Get("id")
	if !exists {
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	// Change type assertion to uint
	userID, ok := userIDInterface.(uint)
	if !ok {
		errs := response.ClientResponse(http.StatusInternalServerError, "User ID type assertion failed", nil, "")
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	userIDStr := strconv.Itoa(int(userID))

	result, err := ch.GRPC_Client.GetChat(userIDStr, chatRequest)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to get chat details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	errs := response.ClientResponse(http.StatusOK, "Successfully retrieved chat details", result, nil)
	c.JSON(http.StatusOK, errs)
}

// GroupMessage handles WebSocket group chat messages.
//
// @Summary Process WebSocket group chat messages
// @Description Processes WebSocket messages for group chat based on the provided group ID
// @Tags Chat
// @Accept json
// @Produce json
// @Param groupID path string true "Group ID"
// @Security ApiKeyAuth
// @Success 200 {string} string "WebSocket connection established"
// @Failure 400 {object} response.Response{} "Missing Authorization header" or "Invalid token" or "Websocket Connection Issue" or "Error reading WebSocket message" or "Details not in correct format"
// @Router /group/:groupID/chat [get]
func (ch *ChatHandler) GroupMessage(c *gin.Context) {

	tokenString := c.Request.Header.Get("Authorization")

	splitToken := strings.Split(tokenString, " ")
	if tokenString == "" {
		errs := response.ClientResponse(http.StatusUnauthorized, "Missing Authorization header", nil, "")
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	splitToken[1] = strings.TrimSpace(splitToken[1])
	userID, err := ch.helper.ValidateToken(splitToken[1])

	if err != nil {
		errs := response.ClientResponse(http.StatusUnauthorized, "Invalid token", nil, err.Error())
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	groupID := c.Param("groupID")

	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Websocket Connection Issue", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	defer func() {
		groupKey := groupID + "_" + strconv.Itoa(int(userID.Id))
		delete(User, groupKey)
		conn.Close()
	}()

	user := strconv.Itoa(int(userID.Id))
	groupKey := groupID + "_" + user
	User[groupKey] = conn

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}

		ch.helper.SendMessageToGroup(User, msg, groupID, user)
	}
}