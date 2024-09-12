package interfaces

import "github.com/ahdaan67/JobQuest/pkg/utils/models"

type ChatClient interface {
	GetChat(userID string, req models.ChatRequest) ([]models.TempMessage, error)
	GetGroupChat(userID string, groupID string, req models.ChatRequest) ([]models.TempMessage, error)
}