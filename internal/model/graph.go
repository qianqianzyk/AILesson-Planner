package model

type GraphData struct {
	Nodes         []Node         `json:"nodes"`
	Relationships []Relationship `json:"relationships"`
}

type ChunkProperties struct {
	FileName      string `json:"fileName"`
	ContentOffset int    `json:"content_offset"`
	PageNumber    int    `json:"page_number"`
	Length        int    `json:"length"`
	ID            string `json:"id"`
	Position      int    `json:"position"`
	Text          string `json:"text"`
}

type DocumentProperties struct {
	ChunkNodeCount       int     `json:"chunkNodeCount"`
	ChunkRelCount        int     `json:"chunkRelCount"`
	CommunityNodeCount   int     `json:"communityNodeCount"`
	CommunityRelCount    int     `json:"communityRelCount"`
	CreatedAt            string  `json:"createdAt"`
	EntityEntityRelCount int     `json:"entityEntityRelCount"`
	EntityNodeCount      int     `json:"entityNodeCount"`
	ErrorMessage         string  `json:"errorMessage"`
	FileName             string  `json:"fileName"`
	FileSize             int     `json:"fileSize"`
	FileSource           string  `json:"fileSource"`
	FileType             string  `json:"fileType"`
	IsCancelled          bool    `json:"is_cancelled"`
	Model                string  `json:"model"`
	NodeCount            int     `json:"nodeCount"`
	ProcessedChunk       int     `json:"processed_chunk"`
	ProcessingTime       float64 `json:"processingTime"`
	RelationshipCount    int     `json:"relationshipCount"`
	Status               string  `json:"status"`
	TotalChunks          int     `json:"total_chunks"`
	UpdatedAt            string  `json:"updatedAt"`
}

type OtherNodeProperties struct {
	ID string `json:"id"`
}

type Node struct {
	ElementID  string      `json:"element_id"`
	Labels     []string    `json:"labels"`
	Properties interface{} `json:"properties"`
}

type Relationship struct {
	ElementID          string `json:"element_id"`
	Type               string `json:"type"`
	StartNodeElementID string `json:"start_node_element_id"`
	EndNodeElementID   string `json:"end_node_element_id"`
}
