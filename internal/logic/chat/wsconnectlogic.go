package chat

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// ClientMessage 用户发送的消息结构体
type ClientMessage struct {
	SessionID int64  `json:"session_id"` // 对话的 Session ID
	Message   string `json:"message"`    // 用户的问题
}

// ServerMessage 服务端发送的消息结构体
type ServerMessage struct {
	SessionID int64  `json:"session_id"` // 对话的 Session ID
	Message   string `json:"message"`    // AI的回答
}

type WsConnectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsConnectLogic {
	return &WsConnectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsConnectLogic) WsConnect(w http.ResponseWriter, r *http.Request) error {
	// 将 HTTP 连接升级为 WebSocket 连接
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		return utils.AbortWithException(utils.ErrUpgradeWs, err)
	}
	defer conn.Close()

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return utils.AbortWithException(utils.ErrUserID, err)
	}

	// 将 WebSocket 连接添加到 WebSocket 管理器
	l.svcCtx.WebsocketManager.AddConnection(userID, conn)
	defer func() {
		err := service.SyncMessagesToMySQL(userID)
		if err != nil {
			utils.LogError(&utils.ErrSyncToMySQL, err)
		}
		l.svcCtx.WebsocketManager.RemoveConnection(userID, conn)
	}()

	// 处理 WebSocket 消息
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				return nil
			} else {
				return utils.AbortWithException(utils.ErrReadMessage, err)
			}
		}

		// 解析用户发送的消息
		var clientMsg ClientMessage
		if err := json.Unmarshal(message, &clientMsg); err != nil {
			return utils.AbortWithException(utils.ErrMessageFormatted, err)
		}

		// 调用 AI 服务获取答案
		aiReply := "Hello World!"

		// 将 AI 的回答发送回客户端
		serverMsg := ServerMessage{
			SessionID: clientMsg.SessionID,
			Message:   aiReply,
		}
		serverMsgBytes, _ := json.Marshal(serverMsg)
		if err := conn.WriteMessage(websocket.TextMessage, serverMsgBytes); err != nil {
			return utils.AbortWithException(utils.ErrSendMessage, err)
		}

		// 将消息存储到数据库中
		chatMessage := model.ConversationMessage{
			SessionID: int(clientMsg.SessionID),
			Role:      "user",
			Message:   clientMsg.Message,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		err = service.SaveMessageToRedis(userID, chatMessage)
		if err != nil {
			return utils.AbortWithException(utils.ErrSaveToRedis, err)
		}

		serverMessage := model.ConversationMessage{
			SessionID: int(serverMsg.SessionID),
			Role:      "ai",
			Message:   aiReply,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		err = service.SaveMessageToRedis(userID, serverMessage)
		if err != nil {
			return utils.AbortWithException(utils.ErrSaveToRedis, err)
		}

		l.svcCtx.WebsocketManager.IncrementMessageCount(userID)

		if l.svcCtx.WebsocketManager.GetMessageCount(userID) >= 10 {
			err = service.SyncMessagesToMySQL(userID)
			if err != nil {
				return utils.AbortWithException(utils.ErrSyncToMySQL, err)
			}
			l.svcCtx.WebsocketManager.ResetMessageCount(userID)
		}
	}
}
