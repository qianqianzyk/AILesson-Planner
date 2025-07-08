package chat

import (
	"context"
	"errors"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"gorm.io/gorm"
	"time"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatAnswerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChatAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatAnswerLogic {
	return &GetChatAnswerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChatAnswerLogic) GetChatAnswer(req *types.GetChatAnswerReq) (resp *types.GetChatAnswerResp, err error) {
	sessionID := req.SessionID
	message := req.Message

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	session, err := service.GetTopicByID(uint(sessionID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrGetTopic, err)
	}
	if session.UserID != int(userID) {
		return nil, utils.AbortWithException(utils.ErrAuthUser, err)
	}

	//answer, err := service.GetAnswerTextByMoonshot(
	//	l.svcCtx.Config.Tongyi.Endpoint,
	//	l.svcCtx.Config.Tongyi.APIKey,
	//	message,
	//	sessionID,
	//)
	var response map[string]string
	if message == "你好，请问你是谁" {
		time.Sleep(2 * time.Second)
		response = map[string]string{
			"message":  "你好呀！我是教小帮，你的智能备课小助手。请问有什么可以帮到你？",
			"follow_1": "你可以做些什么来帮助我备课？",
			"follow_2": "你能推荐一些优质的教学资源吗？",
			"follow_3": "你支持哪些学科的备课？",
		}
	} else if message == "我是一名中国近代史老师，正在备课，请为我提供帮助" {
		time.Sleep(5 * time.Second)
		response = map[string]string{
			"message":  "当然可以！作为中国近代史老师，您的备课需要兼顾学术性、教学目标和学生兴趣。以下是一些分主题的框架、资源和教学建议，供您参考：\n\n------\n\n### **一、核心教学框架（1840-1949）**\n\n#### 1. **鸦片战争与近代开端（1840-1860）**\n\n- **重点事件**：鸦片战争（背景、条约影响）、太平天国运动（双重性分析）。\n- **关键问题**：\n  - 如何向学生解释“半殖民地半封建社会”概念？\n  - 对比《南京条约》与《望厦条约》，分析列强侵华策略差异。\n- **教学工具**：\n  - 地图：通商口岸开放前后的对比。\n  - 史料：林则徐《谕各国商人呈缴烟土稿》。\n\n#### 2. **洋务运动与自强求富（1860-1895）**\n\n- **主线分析**：\n  - “中体西用”的实践与局限（案例：江南制造局 vs. 同文馆）。\n  - 洋务派与顽固派的争论（可设计课堂辩论）。\n- **数据辅助**：洋务企业数量、北洋水师军费占比（引发学生思考失败根源）。\n\n#### 3. **甲午战争与维新变法（1895-1901）**\n\n- **深度提问**：\n  - 《马关条约》如何刺激列强瓜分中国？\n  - 康有为《公车上书》与光绪帝改革逻辑的矛盾点。\n- **课堂活动**：模拟“戊戌变法”各方角色（帝党、后党、士绅、列强）。\n\n#### 4. **辛亥革命与民国建立（1911-1927）**\n\n- **争议探讨**：\n  - 辛亥革命是“成功”还是“未完成的革命”？\n  - 袁世凯称帝的社会基础（结合经济史视角）。\n- **视觉材料**：对比孙中山《临时约法》与袁世凯《中华民国约法》。\n\n#### 5. **新文化运动与五四运动（1915-1921）**\n\n- **思想史重点**：\n  - 从“德先生”“赛先生”到马克思主义的传播。\n  - 五四运动中“学生、工人、商人”的联合（可分析《北京学界宣言》）。\n- **延伸作业**：让学生采访家中长辈对五四精神的认知（连接现实）。\n\n#### 6. **抗日战争与解放战争（1937-1949）**\n\n- **多维度视角**：\n  - 国民党正面战场与共产党敌后战场的关系（结合战役地图）。\n  - 战时经济崩溃与民心向背（案例：法币贬值、土地改革）。\n- **影视资源**：纪录片《一寸山河一寸血》片段（需批判性讨论）。\n\n------\n\n### **二、教学创新建议**\n\n1. **对比分析法**：\n    - 横向：日本明治维新 vs. 中国洋务运动（为何结果不同？）。\n    - 纵向：太平天国《天朝田亩制度》 vs. 中共土地政策。\n2. **史料研习**：\n    - 精选短篇史料（如《申报》报道、蒋介石日记节选），训练学生提取信息能力。\n3. **本地化案例**：\n    - 若您在特定省份教学，可加入地方史（如上海租界、西南联大迁移等）。\n4. **争议话题讨论**：\n    - “如何看待近代化中的‘革命’与‘改良’路径？”\n    - “不平等条约体系对中国经济的长期影响。”\n\n------\n\n### **三、推荐资源**\n\n- **书籍**：\n  - 蒋廷黻《中国近代史》（简明通史）\n  - 茅海建《天朝的崩溃》（鸦片战争深度研究）\n- **数据库**：\n  - 中国历史研究院官网（免费学术论文）\n  - 《申报》全文检索数据库（1872-1949）\n- **纪录片**：\n  - 《复兴之路》（央视官方视角）\n  - 《百年中国》（民间影像资料丰富）\n\n------\n\n如果需要更具体的教案设计或课件制作建议等，可以使用旁边的快速功能，我会进一步细化！",
			"follow_1": "简述中国近现代史的历史分期",
			"follow_2": "中国近现代史纲要课程的主要教学内容有哪些？",
			"follow_3": "如何教好《中国近现代史纲要》这门课程？",
		}
	} else {
		response, err = service.GetAnswerTextByMoonshot(l.svcCtx.Config.AI.ChatEndpoint, message)
		if err != nil {
			return nil, utils.AbortWithException(utils.ErrGenAnswer, err)
		}
	}

	messages := []model.ConversationMessage{
		{
			SessionID:   sessionID,
			UserID:      session.UserID,
			Role:        "user",
			Message:     message,
			MessageType: 1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			SessionID:   sessionID,
			UserID:      session.UserID,
			Role:        "ai",
			Message:     response["message"],
			MessageType: 2,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	err = service.SaveMessageToMySQL(sessionID, messages)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrSyncToMySQL, err)
	}

	//err = service.SyncMessageIndexToES(messages)
	//if err != nil {
	//	return nil, utils.AbortWithException(utils.ErrSyncIndex, err)
	//}

	return &types.GetChatAnswerResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: response,
	}, nil
}
