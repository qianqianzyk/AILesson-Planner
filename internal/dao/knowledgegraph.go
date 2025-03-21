package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"strings"
	"time"
)

func (d *Dao) FetchGraphDataByFileName(ctx context.Context, fileName, authorizationID string) (model.GraphData, error) {
	session := (*d.ne).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx,
		`MATCH (d:Document)
				WHERE d.fileName = $fileName AND d.authorization_id = $authorizationID
				CALL apoc.path.subgraphAll(d, {relationshipFilter: '>|<'}) YIELD nodes, relationships
				WITH d, nodes AS uniqueNodes, relationships AS allRelationships

				WITH d, uniqueNodes, allRelationships,
     				[node IN uniqueNodes WHERE node:Chunk AND node.fileName = d.fileName] AS chunkNodes

				WITH d, uniqueNodes, allRelationships, chunkNodes,
     				[node IN uniqueNodes 
         				WHERE node:Chunk AND node.fileName = d.fileName 
         				| {
             				element_id: elementId(node),
             				fileName: node.fileName,
             				content_offset: node.content_offset,
             				page_number: node.page_number,
             				length: node.length,
             				id: node.id,
             				position: node.position,
             				text: node.text
         				}
     				] AS chunks

				WITH d, uniqueNodes, allRelationships, chunks, chunkNodes,
     				[node IN uniqueNodes 
         				WHERE NOT node:Chunk 
           				AND node <> d 
           				AND ANY(chunk IN chunkNodes WHERE EXISTS((chunk)-[]-(node)))
						| {
             				element_id: elementId(node),
             				labels: labels(node),
             				properties: properties(node)
         				}
     				] AS otherNodes,
     				[rel IN allRelationships | {
         				element_id: elementId(rel),
         				type: type(rel),
						start_node_element_id: elementId(startNode(rel)),
         				end_node_element_id: elementId(endNode(rel))
     				}] AS relationshipsData
     
				RETURN
    				elementId(d) AS element_id,
    				d.fileName AS fileName,
    				d.fileSource AS fileSource,
    				d.fileType AS fileType,
    				d.fileSize AS fileSize,
    				d.is_cancelled AS isCancelled,
    				d.model AS model,
    				d.processingTime AS processingTime,
    				d.status AS status,
    				d.total_chunks AS totalChunks,
					d.processed_chunk AS processedChunk,
    				d.relationshipCount AS relationshipCount,
    				d.nodeCount AS nodeCount,
    				d.communityNodeCount AS communityNodeCount,
    				d.communityRelCount AS communityRelCount,
    				d.entityNodeCount AS entityNodeCount,
    				d.entityEntityRelCount AS entityEntityRelCount,
    				d.errorMessage AS errorMessage,
    				d.createdAt AS createdAt,
    				d.updatedAt AS updatedAt,
    				chunks,
    				otherNodes,
    				relationshipsData`,
		map[string]interface{}{"fileName": fileName, "authorizationID": authorizationID},
	)
	if err != nil {
		return model.GraphData{}, err
	}

	var graphData model.GraphData

	if result.Next(ctx) {
		record := result.Record()

		elementID, _ := record.Get("element_id")
		fileNameValue, _ := record.Get("fileName")
		fileSource, _ := record.Get("fileSource")
		fileType, _ := record.Get("fileType")
		fileSize, _ := record.Get("fileSize")
		isCancelled, _ := record.Get("isCancelled")
		modelValue, _ := record.Get("model")
		processingTime, _ := record.Get("processingTime")
		status, _ := record.Get("status")
		totalChunks, _ := record.Get("totalChunks")
		processedChunk, _ := record.Get("processedChunk")
		relationshipCount, _ := record.Get("relationshipCount")
		nodeCount, _ := record.Get("nodeCount")
		communityNodeCount, _ := record.Get("communityNodeCount")
		communityRelCount, _ := record.Get("communityRelCount")
		entityNodeCount, _ := record.Get("entityNodeCount")
		entityEntityRelCount, _ := record.Get("entityEntityRelCount")
		errorMessage, _ := record.Get("errorMessage")
		createdAt, _ := record.Get("createdAt")
		updatedAt, _ := record.Get("updatedAt")
		docProps := model.DocumentProperties{
			FileName:             fileNameValue.(string),
			FileSource:           fileSource.(string),
			FileType:             fileType.(string),
			FileSize:             int(fileSize.(int64)),
			IsCancelled:          isCancelled.(bool),
			Model:                modelValue.(string),
			ProcessingTime:       processingTime.(float64),
			Status:               status.(string),
			TotalChunks:          int(totalChunks.(int64)),
			ProcessedChunk:       int(processedChunk.(int64)),
			RelationshipCount:    int(relationshipCount.(int64)),
			NodeCount:            int(nodeCount.(int64)),
			CommunityNodeCount:   int(communityNodeCount.(int64)),
			CommunityRelCount:    int(communityRelCount.(int64)),
			EntityNodeCount:      int(entityNodeCount.(int64)),
			EntityEntityRelCount: int(entityEntityRelCount.(int64)),
			ErrorMessage:         errorMessage.(string),
			CreatedAt:            formatTime(createdAt),
			UpdatedAt:            formatTime(updatedAt),
		}
		graphData.Nodes = append(graphData.Nodes, model.Node{
			ElementID:  elementID.(string),
			Labels:     []string{"Document"},
			Properties: docProps,
		})

		chunks, _ := record.Get("chunks")
		for _, chunk := range chunks.([]interface{}) {
			chunkMap := chunk.(map[string]interface{})
			chunkProps := model.ChunkProperties{
				FileName:      chunkMap["fileName"].(string),
				ContentOffset: int(chunkMap["content_offset"].(int64)),
				PageNumber:    int(chunkMap["page_number"].(int64)),
				Length:        int(chunkMap["length"].(int64)),
				ID:            chunkMap["id"].(string),
				Position:      int(chunkMap["position"].(int64)),
				Text:          chunkMap["text"].(string),
			}
			graphData.Nodes = append(graphData.Nodes, model.Node{
				ElementID:  chunkMap["element_id"].(string),
				Labels:     []string{"Chunk"},
				Properties: chunkProps,
			})
		}

		otherNodes, _ := record.Get("otherNodes")
		for _, other := range otherNodes.([]interface{}) {
			otherMap := other.(map[string]interface{})
			labels := make([]string, 0)
			for _, label := range otherMap["labels"].([]interface{}) {
				labels = append(labels, label.(string))
			}

			props := otherMap["properties"].(map[string]interface{})

			id := ""
			if idVal, exists := props["id"]; exists && idVal != nil {
				id = idVal.(string)
			}
			otherProps := model.OtherNodeProperties{
				ID: id,
			}

			graphData.Nodes = append(graphData.Nodes, model.Node{
				ElementID:  otherMap["element_id"].(string),
				Labels:     labels,
				Properties: otherProps,
			})
		}

		relationships, _ := record.Get("relationshipsData")
		for _, relationship := range relationships.([]interface{}) {
			relationshipMap := relationship.(map[string]interface{})
			graphData.Relationships = append(graphData.Relationships, model.Relationship{
				ElementID:          relationshipMap["element_id"].(string),
				Type:               relationshipMap["type"].(string),
				StartNodeElementID: relationshipMap["start_node_element_id"].(string),
				EndNodeElementID:   relationshipMap["end_node_element_id"].(string),
			})
		}

		return graphData, nil
	}

	return model.GraphData{}, result.Err()
}

func formatTime(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case dbtype.LocalDateTime:
		return v.Time().Format("2006-01-02 15:04:05")
	default:
		return ""
	}
}

func (d *Dao) UpdateNodeByElementID(ctx context.Context, nodeType, elementID string, updates map[string]interface{}) error {
	session := (*d.ne).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	setClauses := ""
	params := map[string]interface{}{
		"elementID": elementID,
	}

	i := 0
	for key, value := range updates {
		paramKey := fmt.Sprintf("param%d", i)
		setClauses += fmt.Sprintf("n.%s = $%s, ", key, paramKey)

		switch v := value.(type) {
		case string:
			params[paramKey] = v
		case int, int64, float64:
			params[paramKey] = v
		case bool:
			params[paramKey] = v
		default:
			return fmt.Errorf("unsupported property type for key %s", key)
		}

		i++
	}
	if len(setClauses) > 0 {
		setClauses = setClauses[:len(setClauses)-2]
	} else {
		return fmt.Errorf("no valid update fields provided")
	}

	query := fmt.Sprintf(`
		MATCH (n:%s) 
		WHERE elementId(n) = $elementID
		SET %s
		RETURN elementId(n) AS updatedElementId
	`, nodeType, setClauses)

	result, err := session.Run(ctx, query, params)
	if err != nil {
		return err
	}

	if result.Next(ctx) {
		return nil
	}
	return fmt.Errorf("node not found or update failed")
}

func (d *Dao) DeleteNodeByElementID(ctx context.Context, elementID, nodeType, filename, authorizationID string) error {
	session := (*d.ne).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	if nodeType == "Document" {
		query := `
			MATCH (d:Document) 
			WHERE elementId(d) = $elementID AND d.fileName IS NOT NULL AND d.authorization_id = $authorizationID
			CALL apoc.path.subgraphAll(d, {relationshipFilter: '>|<'}) YIELD nodes, relationships
			WITH d, nodes AS uniqueNodes, relationships AS allRelationships

			WITH d, uniqueNodes, allRelationships,
				[node IN uniqueNodes WHERE node:Chunk AND node.fileName = d.fileName] AS chunkNodes

			WITH d, uniqueNodes, allRelationships, chunkNodes,
				[node IN uniqueNodes 
					WHERE node:Chunk AND node.fileName = d.fileName 
				| {
					element_id: elementId(node),
					fileName: node.fileName,
					content_offset: node.content_offset,
					page_number: node.page_number,
					length: node.length,
					id: node.id,
					position: node.position,
					text: node.text
				}
			] AS chunks

			WITH d, uniqueNodes, allRelationships, chunks, chunkNodes,
				[node IN uniqueNodes 
					WHERE NOT node:Chunk 
					AND node <> d 
					AND ANY(chunk IN chunkNodes WHERE EXISTS((chunk)-[]-(node)))
					| {
					element_id: elementId(node),
					labels: labels(node),
					properties: properties(node)
				}
			] AS otherNodes,
			[rel IN allRelationships | {
				element_id: elementId(rel),
				type: type(rel),
				start_node_element_id: elementId(startNode(rel)),
				end_node_element_id: elementId(endNode(rel))
			}] AS relationshipsData

			UNWIND chunks AS chunk
			MATCH (c:Chunk) WHERE c.id = chunk.id
			DETACH DELETE c

			UNWIND otherNodes AS node
			MATCH (n) WHERE elementId(n) = node.element_id
			DETACH DELETE n

			UNWIND relationshipsData AS rel
			MATCH ()-[r]->()
			WHERE elementId(r) = rel.element_id
			DELETE r

			DETACH DELETE d
			RETURN COUNT(d) AS deletedCount
		`
		params := map[string]interface{}{
			"elementID":       elementID,
			"authorizationID": authorizationID,
		}

		result, err := session.Run(ctx, query, params)
		if err != nil {
			return fmt.Errorf("failed to execute delete query: %w", err)
		}
		if result.Next(ctx) {
			count, _ := result.Record().Get("deletedCount")
			if count.(int64) > 0 {
				return nil
			}
		}
	}

	query := `
		MATCH (n) 
		WHERE elementId(n) = $elementID
		DETACH DELETE n
		RETURN COUNT(n) AS deletedCount
	`

	params := map[string]interface{}{
		"elementID": elementID,
	}

	result, err := session.Run(ctx, query, params)
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}

	if result.Next(ctx) {
		count, _ := result.Record().Get("deletedCount")
		if count.(int64) > 0 {
			updateDocQuery := `
				MATCH (d:Document)
				WHERE d.fileName = $fileName AND d.authorization_id = $authorizationID
				CALL apoc.path.subgraphAll(d, {relationshipFilter: '>|<'}) 
				YIELD nodes, relationships
				WITH d, nodes, relationships, 
     				[node IN nodes WHERE node:Chunk AND node.fileName = d.fileName] AS chunkNodes
				WITH d,
     				size(relationships) AS totalRelationships,
     				chunkNodes,
     				size(chunkNodes) AS totalChunks,
     				size([node IN nodes 
         				WHERE NOT node:Chunk 
           				AND node <> d 
           				AND ANY(chunk IN chunkNodes WHERE EXISTS((chunk)-[]-(node)))
     				]) AS entityNodeCount
					SET d.nodeCount = entityNodeCount + totalChunks,
    					d.relationshipCount = totalRelationships,
						d.total_chunks = totalChunks,
						d.entityNodeCount = entityNodeCount,
    					d.updatedAt = $updatedAt
					RETURN COUNT(d) AS updatedCount
			`
			updatedAt := time.Now().Format("2006-01-02 15:04:05")
			updateParams := map[string]interface{}{
				"fileName":        filename,
				"authorizationID": authorizationID,
				"updatedAt":       updatedAt,
			}

			updateResult, err := session.Run(ctx, updateDocQuery, updateParams)
			if err != nil {
				return fmt.Errorf("failed to execute update query: %w", err)
			}

			if updateResult.Next(ctx) {
				updatedCount, _ := updateResult.Record().Get("updatedCount")
				if updatedCount.(int64) > 0 {
					return nil
				}
			}
		}
	}

	return fmt.Errorf("node not found or delete failed")
}

func (d *Dao) GetDocumentList(ctx context.Context, authorizationID string, graphType int) ([]model.DocumentProperties, error) {
	session := (*d.ne).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
		MATCH (d:Document)
		WHERE d.fileName IS NOT NULL AND d.authorization_id = $authorizationID
	`
	params := map[string]interface{}{
		"authorizationID": authorizationID,
	}
	if graphType == 2 {
		query += ` AND d.model = "手动创建"`
	} else if graphType == 1 {
		query += ` AND d.model <> "手动创建"`
	}
	query += `
		RETURN
			d.fileName AS fileName,
			d.fileSource AS fileSource,
			d.fileType AS fileType,
			d.fileSize AS fileSize,
			d.is_cancelled AS isCancelled,
			d.model AS model,
			d.processingTime AS processingTime,
			d.status AS status,
			d.total_chunks AS totalChunks,
			d.processed_chunk AS processedChunk,
			d.relationshipCount AS relationshipCount,
			d.nodeCount AS nodeCount,
			d.communityNodeCount AS communityNodeCount,
			d.communityRelCount AS communityRelCount,
			d.entityNodeCount AS entityNodeCount,
			d.entityEntityRelCount AS entityEntityRelCount,
			d.errorMessage AS errorMessage,
			d.createdAt AS createdAt,
			d.updatedAt AS updatedAt
		ORDER BY d.createdAt DESC
	`

	result, err := session.Run(ctx, query, params)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var documents []model.DocumentProperties
	for result.Next(ctx) {
		record := result.Record()

		fileNameValue, _ := record.Get("fileName")
		fileSource, _ := record.Get("fileSource")
		fileType, _ := record.Get("fileType")
		fileSize, _ := record.Get("fileSize")
		isCancelled, _ := record.Get("isCancelled")
		modelValue, _ := record.Get("model")
		processingTime, _ := record.Get("processingTime")
		status, _ := record.Get("status")
		totalChunks, _ := record.Get("totalChunks")
		processedChunk, _ := record.Get("processedChunk")
		relationshipCount, _ := record.Get("relationshipCount")
		nodeCount, _ := record.Get("nodeCount")
		communityNodeCount, _ := record.Get("communityNodeCount")
		communityRelCount, _ := record.Get("communityRelCount")
		entityNodeCount, _ := record.Get("entityNodeCount")
		entityEntityRelCount, _ := record.Get("entityEntityRelCount")
		errorMessage, _ := record.Get("errorMessage")
		createdAt, _ := record.Get("createdAt")
		updatedAt, _ := record.Get("updatedAt")
		doc := model.DocumentProperties{
			FileName:             fileNameValue.(string),
			FileSource:           fileSource.(string),
			FileType:             fileType.(string),
			FileSize:             int(fileSize.(int64)),
			IsCancelled:          isCancelled.(bool),
			Model:                modelValue.(string),
			ProcessingTime:       processingTime.(float64),
			Status:               status.(string),
			TotalChunks:          int(totalChunks.(int64)),
			ProcessedChunk:       int(processedChunk.(int64)),
			RelationshipCount:    int(relationshipCount.(int64)),
			NodeCount:            int(nodeCount.(int64)),
			CommunityNodeCount:   int(communityNodeCount.(int64)),
			CommunityRelCount:    int(communityRelCount.(int64)),
			EntityNodeCount:      int(entityNodeCount.(int64)),
			EntityEntityRelCount: int(entityEntityRelCount.(int64)),
			ErrorMessage:         errorMessage.(string),
			CreatedAt:            formatTime(createdAt),
			UpdatedAt:            formatTime(updatedAt),
		}
		documents = append(documents, doc)
	}

	return documents, nil
}

func (d *Dao) CreateNode(ctx context.Context, label string, properties map[string]interface{}) (string, error) {
	session := (*d.ne).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	if label == "" || len(properties) == 0 {
		return "", fmt.Errorf("label 和 properties 不能为空")
	}

	var keys []string
	var params []string
	cypherParams := map[string]interface{}{}

	for key, value := range properties {
		paramName := "param_" + key
		keys = append(keys, key)
		params = append(params, key+": $"+paramName)

		switch v := value.(type) {
		case string:
			cypherParams[paramName] = v
		case json.Number:
			if intValue, err := v.Int64(); err == nil {
				cypherParams[paramName] = intValue
			} else if floatValue, err := v.Float64(); err == nil {
				cypherParams[paramName] = floatValue
			} else {
				return "", fmt.Errorf("invalid json.Number for key %s", key)
			}
		case int, int64, float64:
			cypherParams[paramName] = v
		case bool:
			cypherParams[paramName] = v
		default:
			return "", fmt.Errorf("unsupported property type for key %s", key)
		}
	}

	query := fmt.Sprintf("CREATE (n:%s {%s}) RETURN elementId(n) AS id", label, strings.Join(params, ", "))

	result, err := session.Run(ctx, query, cypherParams)
	if err != nil {
		return "", err
	}

	if result.Next(ctx) {
		id, _ := result.Record().Get("id")
		return id.(string), nil
	}

	return "", fmt.Errorf("创建节点失败")
}

func (d *Dao) CreateGraphNodeRelationShip(ctx context.Context, startNodeElementID, endNodeElementID, relationshipType, filename string) (string, error) {
	session := (*d.ne).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	relationshipQuery := `
		MATCH (startNode) WHERE elementId(startNode) = $startNodeElementID
		MATCH (endNode) WHERE elementId(endNode) = $endNodeElementID
		MERGE (startNode)-[r:` + relationshipType + `]->(endNode)
		RETURN elementId(r) AS relationshipElementID
	`

	updateDocQuery := `
		MATCH (d:Document)
		WHERE d.fileName = $fileName
		CALL apoc.path.subgraphAll(d, {relationshipFilter: '>|<'}) 
		YIELD nodes, relationships
		WITH d, nodes, relationships, 
			 [node IN nodes WHERE node:Chunk AND node.fileName = d.fileName] AS chunkNodes
		WITH d,
			 size(relationships) AS totalRelationships,
			 chunkNodes,
			 size(chunkNodes) AS totalChunks,
			 size([node IN nodes 
					 WHERE NOT node:Chunk 
					 AND node <> d 
					 AND ANY(chunk IN chunkNodes WHERE EXISTS((chunk)-[]-(node)))
				]) AS entityNodeCount
		SET d.nodeCount = entityNodeCount + totalChunks,
			d.relationshipCount = totalRelationships,
			d.total_chunks = totalChunks,
			d.entityNodeCount = entityNodeCount,
			d.updatedAt = $updatedAt
		RETURN COUNT(d) AS updatedCount
	`

	updatedAt := time.Now().Format("2006-01-02 15:04:05")

	relationshipElementID, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, relationshipQuery, map[string]any{
			"startNodeElementID": startNodeElementID,
			"endNodeElementID":   endNodeElementID,
		})
		if err != nil {
			return "", fmt.Errorf("failed to create relationship: %w", err)
		}

		if !result.Next(ctx) {
			return "", fmt.Errorf("relationship not created")
		}
		relID, _ := result.Record().Get("relationshipElementID")

		updateResult, err := tx.Run(ctx, updateDocQuery, map[string]any{
			"fileName":  filename,
			"updatedAt": updatedAt,
		})
		if err != nil {
			return "", fmt.Errorf("failed to execute update query: %w", err)
		}

		if updateResult.Next(ctx) {
			updatedCount, _ := updateResult.Record().Get("updatedCount")
			if updatedCount.(int64) > 0 {
				return relID.(string), nil
			}
		}

		return relID.(string), nil
	})

	if err != nil {
		return "", err
	}
	return relationshipElementID.(string), nil
}

func (d *Dao) UpdateGraphNodeRelationShip(ctx context.Context, elementID, relationshipType, filename string) error {
	session := (*d.ne).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	updateRelationshipQuery := `
		MATCH ()-[r]->() 
		WHERE elementId(r) = $elementID
		SET r.type = $relationshipType
	`

	updateDocQuery := `
		MATCH (d:Document)
		WHERE d.fileName = $fileName
		CALL apoc.path.subgraphAll(d, {relationshipFilter: '>|<'}) 
		YIELD nodes, relationships
		WITH d, nodes, relationships, 
			 [node IN nodes WHERE node:Chunk AND node.fileName = d.fileName] AS chunkNodes
		WITH d,
			 size(relationships) AS totalRelationships,
			 chunkNodes,
			 size(chunkNodes) AS totalChunks,
			 size([node IN nodes 
					 WHERE NOT node:Chunk 
					 AND node <> d 
					 AND ANY(chunk IN chunkNodes WHERE EXISTS((chunk)-[]-(node)))
				]) AS entityNodeCount
		SET d.nodeCount = entityNodeCount + totalChunks,
			d.relationshipCount = totalRelationships,
			d.total_chunks = totalChunks,
			d.entityNodeCount = entityNodeCount,
			d.updatedAt = $updatedAt
	`

	updatedAt := time.Now().Format("2006-01-02 15:04:05")

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, updateRelationshipQuery, map[string]any{
			"elementID":        elementID,
			"relationshipType": relationshipType,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to update relationship: %w", err)
		}

		if _, err := result.Consume(ctx); err != nil {
			return nil, fmt.Errorf("relationship not found")
		}

		updateResult, err := tx.Run(ctx, updateDocQuery, map[string]any{
			"fileName":  filename,
			"updatedAt": updatedAt,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to execute update query: %w", err)
		}

		if _, err := updateResult.Consume(ctx); err != nil {
			return nil, fmt.Errorf("failed to update document statistics")
		}

		return nil, nil
	})

	return err
}

func (d *Dao) DeleteGraphNodeRelationShip(ctx context.Context, elementID, filename string) error {
	session := (*d.ne).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	deleteRelationshipQuery := `
		MATCH ()-[r]->() 
		WHERE elementId(r) = $elementID
		DELETE r
	`

	updateDocQuery := `
		MATCH (d:Document)
		WHERE d.fileName = $fileName
		CALL apoc.path.subgraphAll(d, {relationshipFilter: '>|<'}) 
		YIELD nodes, relationships
		WITH d, nodes, relationships, 
			 [node IN nodes WHERE node:Chunk AND node.fileName = d.fileName] AS chunkNodes
		WITH d,
			 size(relationships) AS totalRelationships,
			 chunkNodes,
			 size(chunkNodes) AS totalChunks,
			 size([node IN nodes 
					 WHERE NOT node:Chunk 
					 AND node <> d 
					 AND ANY(chunk IN chunkNodes WHERE EXISTS((chunk)-[]-(node)))
				]) AS entityNodeCount
		SET d.nodeCount = entityNodeCount + totalChunks,
			d.relationshipCount = totalRelationships,
			d.total_chunks = totalChunks,
			d.entityNodeCount = entityNodeCount,
			d.updatedAt = $updatedAt
	`

	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, deleteRelationshipQuery, map[string]any{
			"elementID": elementID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to delete relationship: %w", err)
		}

		if _, err := result.Consume(ctx); err != nil {
			return nil, fmt.Errorf("relationship not found or already deleted")
		}

		updateResult, err := tx.Run(ctx, updateDocQuery, map[string]any{
			"fileName":  filename,
			"updatedAt": updatedAt,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to execute update query: %w", err)
		}

		if _, err := updateResult.Consume(ctx); err != nil {
			return nil, fmt.Errorf("failed to update document statistics")
		}

		return nil, nil
	})

	return err
}
