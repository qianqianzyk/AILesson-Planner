package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"net/http"
	"strconv"
)

const (
	ModelName       = "qwen-plus"
	ResultFormatted = "text"
	ResultNormalGen = "stop"
	ContextLimit    = 6
)

func CreateTopic(cov model.ConversationSession) error {
	err := d.CreateTopic(ctx, &cov)
	return err
}

func UpdateTopic(cov model.ConversationSession) error {
	err := d.UpdateTopic(ctx, &cov)
	return err
}

func GetTopicByID(id uint) (*model.ConversationSession, error) {
	cov, err := d.GetTopicByID(ctx, id)
	return cov, err
}

func DelTopicByID(id uint, userID int64) error {
	err := d.DelTopicByID(ctx, id, userID)
	return err
}

func GetTopicList(userID int64) ([]model.ConversationSession, error) {
	cov, err := d.GetTopicList(ctx, userID)
	return cov, err
}

func GetMessageList(sessionID, pageNum, pageSize int) ([]model.ConversationMessage, *int64, error) {
	messages, totalSize, err := d.GetMessageList(ctx, sessionID, pageNum, pageSize)
	return messages, totalSize, err
}

func SyncMessagesToMySQL(userID int64) error {
	key := "chat:messages:" + strconv.FormatInt(userID, 10)
	messages, err := d.GetMessageByKey(ctx, key)
	if err != nil {
		return err
	}

	var chatMessages []model.ConversationMessage
	for _, msg := range messages {
		var chatMessage model.ConversationMessage
		err := json.Unmarshal([]byte(msg), &chatMessage)
		if err != nil {
			return err
		}
		chatMessages = append(chatMessages, chatMessage)
	}

	err = d.StoreMessageInDB(chatMessages)
	if err != nil {
		return err
	}

	err = d.DelMessageByKey(ctx, key)
	if err != nil {
		return err
	}

	return nil
}

func SaveMessageToRedis(userID int64, message model.ConversationMessage) error {
	key := "chat:messages:" + strconv.FormatInt(userID, 10)
	msgBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = d.StoreMessageToRedis(ctx, key, string(msgBytes))
	if err != nil {
		return err
	}

	return nil
}

func SaveMessageToMySQL(messages []model.ConversationMessage) error {
	err := d.StoreMessageInDB(messages)
	if err != nil {
		return err
	}
	return nil
}

func GetAnswerTextByTongyi(endpoint string, apiKey string, question string, sessionID int) (string, error) {
	messages, err := d.GetRecentConversation(ctx, sessionID, ContextLimit)
	if err != nil {
		return "", err
	}

	messagesConverted := convertToMessages(messages)
	messagesConverted = append(messagesConverted, model.Message{
		Role:    "user",
		Content: question,
	})

	requestBody := model.RequestBody{
		Model: ModelName,
		Input: model.Input{
			Messages: messagesConverted,
		},
		Parameters: model.Parameters{
			ResultFormat: ResultFormatted,
		},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return "", errors.New("failed to call Tongyi API, status code: " + resp.Status)
	}

	var responsePayload model.ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&responsePayload)
	if err != nil {
		return "", err
	}
	if responsePayload.Output.FinishReason == ResultNormalGen {
		return responsePayload.Output.Text, nil
	}

	return "", err
}

func SyncMessageIndexToES(messages []model.ConversationMessage) error {
	err := d.SyncMessageIndexToElasticsearch(ctx, messages)
	if err != nil {
		return err
	}
	return nil
}

//func StartCronSync(interval time.Duration) {
//	ticker := time.NewTicker(interval)
//	go func() {
//		for range ticker.C {
//			if err := d.SyncToElasticsearch(ctx); err != nil {
//				log.Printf("sync failed: %v", err)
//			}
//		}
//	}()
//}

func convertToMessages(conversationHistory []model.ConversationMessage) []model.Message {
	var messages []model.Message
	for i := len(conversationHistory) - 1; i >= 0; i-- {
		msg := conversationHistory[i]
		var role string
		if msg.Role == "user" {
			role = "user"
		} else if msg.Role == "ai" {
			role = "assistant"
		}

		messages = append(messages, model.Message{
			Role:    role,
			Content: msg.Message,
		})
	}

	return messages
}
