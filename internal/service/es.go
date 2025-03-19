package service

import "github.com/qianqianzyk/AILesson-Planner/internal/model"

func SyncIndex(indexName string) error {
	err := d.SyncToElasticsearch(ctx, indexName)
	return err
}

func SearchMessages(query string, userID int) ([]model.ResponseConversationSession, error) {
	result, err := d.SearchConversations(ctx, query, userID)
	return result, err
}

func SearchFiles(query string, userID int) ([]int, error) {
	result, err := d.SearchFilesByFilename(ctx, query, userID)
	return result, err
}
