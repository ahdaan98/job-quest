package interfaces

import models "chat/pkg/utils"

type ChatUseCase interface {
	MessageConsumer()
	GetFriendChat(string, string, models.Pagination) ([]models.Message, error)
	GetGroupMessages(groupID string, limit, offset int) ([]models.Message, error)
}