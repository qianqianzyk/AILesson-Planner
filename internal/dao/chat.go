package dao

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"gorm.io/gorm"
	"strings"
)

func (d *Dao) CreateTopic(ctx context.Context, cov *model.ConversationSession) error {
	result := d.orm.WithContext(ctx).Model(&model.ConversationSession{}).Create(cov)
	return result.Error
}

func (d *Dao) UpdateTopic(ctx context.Context, cov *model.ConversationSession) error {
	result := d.orm.WithContext(ctx).Save(&cov)
	return result.Error
}

func (d *Dao) GetTopicByID(ctx context.Context, id uint) (*model.ConversationSession, error) {
	var cov *model.ConversationSession
	result := d.orm.WithContext(ctx).Model(&model.ConversationSession{}).Where("id = ?", id).First(&cov)
	if result.Error != nil {
		return nil, result.Error
	}
	return cov, nil
}

func (d *Dao) DelTopicByID(ctx context.Context, id uint, userID int64) error {
	err := d.orm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ? AND user_id = ?", id, userID).Delete(&model.ConversationSession{}).Error; err != nil {
			return err
		}
		if err := tx.Where("session_id = ?", id).Delete(&model.ConversationMessage{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (d *Dao) GetTopicList(ctx context.Context, userID int64) ([]model.ConversationSession, error) {
	var cov []model.ConversationSession
	result := d.orm.WithContext(ctx).Model(&model.ConversationSession{}).Where("user_id = ?", userID).Order("updated_at DESC").Find(&cov)
	return cov, result.Error
}

func (d *Dao) GetMessageByKey(ctx context.Context, key string) ([]string, error) {
	messages, err := d.rc.LrangeCtx(ctx, key, 0, -1)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (d *Dao) DelMessageByKey(ctx context.Context, key string) error {
	_, err := d.rc.DelCtx(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) StoreMessageInDB(messages []model.ConversationMessage) error {
	tx := d.orm.Begin()
	if err := tx.CreateInBatches(messages, 100).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (d *Dao) StoreMessageToRedis(ctx context.Context, key string, value string) error {
	err := d.rc.SetCtx(ctx, key, value)
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) GetRecentConversation(ctx context.Context, sessionID int, limit int) ([]model.ConversationMessage, error) {
	var messages []model.ConversationMessage
	err := d.orm.WithContext(ctx).Where("session_id = ?", sessionID).
		Order("created_at DESC").
		Limit(limit).
		Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (d *Dao) GetMessageList(ctx context.Context, sessionID, pageNum, pageSize int) ([]model.ConversationMessage, *int64, error) {
	var messages []model.ConversationMessage
	var sum int64
	err := d.orm.WithContext(ctx).Model(model.ConversationMessage{}).Where("session_id = ?", sessionID).Order("created_at DESC").Count(&sum).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&messages).Error
	return messages, &sum, err
}

func (d *Dao) SyncToElasticsearch(ctx context.Context, indexName string) error {
	var messages []model.ConversationMessage
	if err := d.orm.WithContext(ctx).Find(&messages).Error; err != nil {
		return err
	}

	var bulkBody []byte
	for _, message := range messages {
		data, _ := json.Marshal(message)
		bulkBody = append(bulkBody, []byte(fmt.Sprintf("{\"index\":{\"_index\":\"%s\",\"_id\":\"%d\"}}\n%s\n", indexName, message.ID, string(data)))...)
	}

	req := esapi.BulkRequest{
		Body:    bytes.NewReader(bulkBody),
		Refresh: "true",
	}

	res, err := req.Do(ctx, d.es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("bulk request failed with status: %s", res.Status())
	}

	return nil
}

func (d *Dao) SyncFileToElasticsearch(ctx context.Context, file *model.File, indexName string) error {
	data, err := json.Marshal(struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		UserID int    `json:"user_id"`
	}{
		ID:     int(file.ID),
		Name:   file.Name,
		UserID: file.UserID,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal file data: %v", err)
	}

	bulkBody := []byte(fmt.Sprintf("{\"index\":{\"_index\":\"%s\",\"_id\":\"%d\"}}\n", indexName, file.ID))
	bulkBody = append(bulkBody, data...)
	bulkBody = append(bulkBody, []byte("\n")...)

	req := esapi.BulkRequest{
		Body:    bytes.NewReader(bulkBody),
		Refresh: "true",
	}

	res, err := req.Do(ctx, d.es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("bulk request failed with status: %s", res.Status())
	}

	return nil
}

func (d *Dao) SyncMessageIndexToElasticsearch(ctx context.Context, messages []model.ConversationMessage) error {
	var bulkBody []byte
	for _, message := range messages {
		data, _ := json.Marshal(message)
		bulkBody = append(bulkBody, []byte(fmt.Sprintf("{\"index\":{\"_index\":\"conversation_messages\",\"_id\":\"%d\"}}\n%s\n", message.ID, string(data)))...)
	}

	req := esapi.BulkRequest{
		Body:    bytes.NewReader(bulkBody),
		Refresh: "true",
	}

	res, err := req.Do(ctx, d.es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("bulk request failed with status: %s", res.Status())
	}

	return nil
}

func (d *Dao) SearchConversations(ctx context.Context, query string, userID int) ([]model.ResponseConversationSession, error) {
	var result []model.ResponseConversationSession

	searchBody := fmt.Sprintf(`{
		"query": {
			"bool": {
				"must": [
					{ "match": { "message": "%s" } },
					{ "term": { "user_id": %d } }
				]
			}
		}
	}`, query, userID)

	req := esapi.SearchRequest{
		Index: []string{"conversation_messages"},
		Body:  strings.NewReader(searchBody),
	}

	res, err := req.Do(ctx, d.es)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("search request failed with status: %s", res.Status())
	}

	var esResponse struct {
		Hits struct {
			Hits []struct {
				Source model.ConversationMessage `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&esResponse); err != nil {
		return nil, err
	}

	groupedMessages := make(map[int][]model.ConversationMessage)
	for _, hit := range esResponse.Hits.Hits {
		message := hit.Source
		groupedMessages[message.SessionID] = append(groupedMessages[message.SessionID], message)
	}

	var sessions []model.ConversationSession
	if err := d.orm.WithContext(ctx).Where("user_id = ?", userID).Find(&sessions).Error; err != nil {
		return nil, err
	}

	sessionMap := make(map[int]model.ConversationSession)
	for _, session := range sessions {
		sessionMap[int(session.ID)] = session
	}

	for sessionID, messages := range groupedMessages {
		if session, exists := sessionMap[sessionID]; exists {
			result = append(result, model.ResponseConversationSession{
				ID:                  int(session.ID),
				UserID:              session.UserID,
				Title:               session.Title,
				CreatedAt:           session.CreatedAt,
				UpdatedAt:           session.UpdatedAt,
				ConversationMessage: messages,
			})
		}
	}

	return result, nil
}

func (d *Dao) SearchFilesByFilename(ctx context.Context, query string, userID int) ([]int, error) {
	searchQuery := fmt.Sprintf(`
	{
		"query": {
			"bool": {
				"must": [
					{
						"match": {
							"name": {
								"query": "%s",
								"fuzziness": "AUTO"
							}
						}
					},
					{
						"term": {
							"user_id": %d
						}
					}
				]
			}
		},
		"_source": ["id"]
	}`, query, userID)

	req := esapi.SearchRequest{
		Index: []string{"disk_files"},
		Body:  bytes.NewReader([]byte(searchQuery)),
	}

	res, err := req.Do(ctx, d.es)
	if err != nil {
		return nil, fmt.Errorf("failed to execute search query: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("search request failed with status: %s", res.Status())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse search result: %v", err)
	}

	var fileIDs []int
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		if fileID, ok := source["id"].(float64); ok {
			fileIDs = append(fileIDs, int(fileID))
		}
	}

	return fileIDs, nil
}
