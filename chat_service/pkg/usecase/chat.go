package usecase

import (
	"chat/pkg/helper"
	"chat/pkg/pb/auth"
	interfaces "chat/pkg/repository/interface"
	services "chat/pkg/usecase/interface"
	models "chat/pkg/utils"
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
)

type ChatUseCase struct {
	chatRepository interfaces.ChatRepository
	authClient     auth.AuthServiceClient
}

func NewChatUseCase(repository interfaces.ChatRepository, authclient auth.AuthServiceClient) services.ChatUseCase {
	return &ChatUseCase{
		chatRepository: repository,
		authClient:     authclient,
	}
}

func (c *ChatUseCase) MessageConsumer() {
	// cfg, err := config.LoadConfig()
	// if err != nil {
	// 	fmt.Println("Error loading config:", err)
	// 	return
	// }

	configs := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, configs)
	if err != nil {
		fmt.Println("Error creating Kafka consumer:", err)
		return
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("Error creating partition consumer:", err)
		return
	}
	defer partitionConsumer.Close()
	fmt.Println("Kafka consumer started")
	for {
		select {
		case message := <-partitionConsumer.Messages():
			msg, err := c.UnmarshelChatMessage(message.Value)
			fmt.Println("message usecase", message.Value)
			if err != nil {
				fmt.Println("Error unmarshalling message:", err)
				continue
			}
			fmt.Println("Received message:", msg)
			err = c.chatRepository.StoreFriendsChat(*msg)
			if err != nil {
				fmt.Println("Error storing message in repository:", err)
				continue
			}
		case err := <-partitionConsumer.Errors():
			fmt.Println("Kafka consumer error:", err)
		}
	}
}

func (c *ChatUseCase) UnmarshelChatMessage(data []byte) (*models.MessageReq, error) {
	var message models.MessageReq
	err := json.Unmarshal(data, &message)
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}

	message.Timestamp = time.Now()
	return &message, nil
}

func (c *ChatUseCase) GetFriendChat(userID, friendID string, pagination models.Pagination) ([]models.Message, error) {
	var err error
	pagination.OffSet, err = helper.Pagination(pagination.Limit, pagination.OffSet)
	if err != nil {
		return nil, err
	}
	_ = c.chatRepository.UpdateReadAsMessage(userID, friendID)
	return c.chatRepository.GetFriendChat(userID, friendID, pagination)
}

func (uc *ChatUseCase) GetGroupMessages(groupID string, limit, offset int) ([]models.Message, error) {
	// Retrieve messages from the repository
	return uc.chatRepository.GetGroupMessages(groupID, limit, offset)
}