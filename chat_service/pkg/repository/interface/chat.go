package interfaces

import models "chat/pkg/utils"

type ChatRepository interface {
	StoreFriendsChat(models.MessageReq) error
	GetFriendChat(string, string, models.Pagination) ([]models.Message, error)
	UpdateReadAsMessage(string, string) error
	GetGroupMessages(groupID string, limit, offset int) ([]models.Message, error)
}