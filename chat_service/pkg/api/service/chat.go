package service

import (
	pb "chat/pkg/pb/chat"
	interfaces "chat/pkg/usecase/interface"
	models "chat/pkg/utils"
	"context"
	"fmt"
	"strconv"
	"time"
)

type ChatServer struct {
	chatUseCase interfaces.ChatUseCase
	pb.UnimplementedChatServiceServer
}

func NewChatServer(UseCaseChat interfaces.ChatUseCase) pb.ChatServiceServer {
	return &ChatServer{
		chatUseCase: UseCaseChat,
	}
}

func (c *ChatServer) GetFriendChat(ctx context.Context, req *pb.GetFriendChatRequest) (*pb.GetFriendChatResponse, error) {
	fmt.Println("line service1")
	ind, _ := time.LoadLocation("Asia/Kolkata")
	result, err := c.chatUseCase.GetFriendChat(req.UserID, req.FriendID, models.Pagination{Limit: req.Limit, OffSet: req.OffSet})
	if err != nil {
		return nil, err
	}

	var finalResult []*pb.Message
	for _, val := range result {
		finalResult = append(finalResult, &pb.Message{
			MessageID:   val.ID,
			SenderId:    val.SenderID,
			RecipientId: val.RecipientID,
			Content:     val.Content,
			Timestamp:   val.Timestamp.In(ind).String(),
		})
	}
	return &pb.GetFriendChatResponse{FriendChat: finalResult}, nil
}

func (c *ChatServer) GetGroupChat(ctx context.Context, req *pb.GetGroupChatRequest) (*pb.GetGroupChatResponse, error) {
	fmt.Println("Handling GetGroupChat request")

	// Convert Limit and OffSet from string to int
	limit, err := strconv.Atoi(req.Limit)
	if err != nil {
		return nil, fmt.Errorf("invalid limit value: %v", err)
	}

	offset, err := strconv.Atoi(req.OffSet)
	if err != nil {
		return nil, fmt.Errorf("invalid offset value: %v", err)
	}

	// Call the use case to retrieve group messages
	groupID := req.GroupID
	result, err := c.chatUseCase.GetGroupMessages(groupID, limit, offset)
	if err != nil {
		return nil, err
	}

	// Convert the result to the protobuf format
	var finalResult []*pb.Message
	for _, val := range result {
		finalResult = append(finalResult, &pb.Message{
			MessageID:   val.ID,
			SenderId:    val.SenderID,
			RecipientId: val.RecipientID,
			Content:     val.Content,
			Timestamp:   val.Timestamp.String(), // Convert timestamp to string as needed
		})
	}
	return &pb.GetGroupChatResponse{GroupChat: finalResult}, nil
}