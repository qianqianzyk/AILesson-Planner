package service

import "github.com/qianqianzyk/AILesson-Planner/internal/model"

func FetchGraphDataByFileName(fileName, authorizationID string) (interface{}, error) {
	graph, err := d.FetchGraphDataByFileName(ctx, fileName, authorizationID)
	return graph, err
}

func UpdateNodeByElementID(nodeType, elementID, authorizationID string, updates map[string]interface{}) error {
	err := d.UpdateNodeByElementID(ctx, nodeType, elementID, authorizationID, updates)
	return err
}

func DeleteNodeByElementID(elementID, nodeType, filename, authorizationID string) error {
	err := d.DeleteNodeByElementID(ctx, elementID, nodeType, filename, authorizationID)
	return err
}

func GetDocumentList(authorizationID string, graphType int) ([]model.DocumentProperties, error) {
	documents, err := d.GetDocumentList(ctx, authorizationID, graphType)
	return documents, err
}

func CreateNode(label string, properties map[string]interface{}) (string, error) {
	elementID, err := d.CreateNode(ctx, label, properties)
	return elementID, err
}

func CreateGraphNodeRelationShip(startNodeElementID, endNodeElementID, relationshipType, filename string) (string, error) {
	elementID, err := d.CreateGraphNodeRelationShip(ctx, startNodeElementID, endNodeElementID, relationshipType, filename)
	return elementID, err
}

func UpdateGraphNodeRelationShip(elementID, relationshipType, filename string) error {
	err := d.UpdateGraphNodeRelationShip(ctx, elementID, relationshipType, filename)
	return err
}

func DeleteGraphNodeRelationShip(elementID, filename string) error {
	err := d.DeleteGraphNodeRelationShip(ctx, elementID, filename)
	return err
}
