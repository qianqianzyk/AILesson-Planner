package chat

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"time"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateUnitHomeworkDesignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateUnitHomeworkDesignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateUnitHomeworkDesignLogic {
	return &GenerateUnitHomeworkDesignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateUnitHomeworkDesignLogic) GenerateUnitHomeworkDesign(req *types.GenerateUnitHomeworkDesignReq) (resp *types.GenerateUnitHomeworkDesignResp, err error) {
	sessionID := req.SessionID

	chatAnswer := "# 中国近代史第一章单元小测设计\n\n## 基本信息\n- ​**​学科​**​: 中国近代史  \n- ​**​阶段​**​: 大学  \n- ​**​教材​**​: 《中国近现代史纲要》（2023年版）  \n- ​**​册别​**​: 全册  \n- ​**​课题​**​: 第一章单元小测  \n\n## 题型配置\n- ​**​题型​**​: 单选题  \n- ​**​数量​**​: 10  \n- ​**​难度​**​: 中等偏上（含史料分析）  \n\n## 考点分布\n1. 世界资本主义的发展与殖民扩张  \n2. 中英贸易结构差异  \n3. 《南京条约》核心条款  \n4. 鸦片战争的多维度影响  \n\n---\n\n## 题目列表\n\n### 1. 19世纪初，中英贸易中中国长期保持顺差的关键商品是？\n​**​选项​**​  \nA. 鸦片与棉花  \nB. 茶叶、生丝与瓷器  \nC. 机械制品与军火  \nD. 香料与白银  \n​**​答案​**​: B\n\n### 2. 《南京条约》中直接冲击清政府财政体系的条款是？\n​**​选项​**​  \nA. 割让香港岛  \nB. 五口通商  \nC. 2100万银元赔款  \nD. 领事裁判权  \n​**​答案​**​: C\n\n### 3. 英国发动鸦片战争的直接导火索是？\n​**​选项​**​  \nA. 马戛尔尼使团礼仪争端  \nB. 林则徐虎门销烟  \nC. 东印度公司破产  \nD. 中俄《尼布楚条约》修订  \n​**​答案​**​: B\n\n### 4. 19世纪殖民扩张的新特点是？\n​**​选项​**​  \nA. 宗教传播为主  \nB. 建立原料-市场双向体系  \nC. 奴隶贸易盛行  \nD. 军事据点式占领  \n​**​答案​**​: B\n\n### 5. 以下最能体现协定关税危害的史料是？\n​**​选项​**​  \nA. 虎门炮台防御图  \nB. 1843年中英海关税率对比表  \nC. 《海国图志》手稿  \nD. 十三行贸易账簿  \n​**​答案​**​: B\n\n### 6. 鸦片战争后中国社会性质的根本变化体现在？\n​**​选项​**​  \nA. 自然经济完全解体  \nB. 开始沦为半殖民地半封建社会  \nC. 民族资本主义兴起  \nD. 科举制度废除  \n​**​答案​**​: B\n\n### 7. 英国选择广州作为突破口的原因不包括？\n​**​选项​**​  \nA. 传统贸易口岸  \nB. 长江流域腹地广阔  \nC. 军事防御薄弱  \nD. 距离印度殖民地近  \n​**​答案​**​: B\n\n### 8. 下列条约内容体现经济侵略特征的是？\n​**​选项​**​  \nA. 设立租界  \nB. 片面最惠国待遇  \nC. 公使驻京  \nD. 军舰巡查权  \n​**​答案​**​: B\n\n### 9. 工业革命对殖民扩张的影响是？\n​**​选项​**​  \nA. 降低原料需求  \nB. 催生商品输出需求  \nC. 减少军事冲突  \nD. 促进文化融合  \n​**​答案​**​: B\n\n### 10. 鸦片战争后，中国自然经济开始解体的主要表现是？\n​**​选项​**​  \nA. 农民放弃耕种  \nB. 手工棉纺织业衰败  \nC. 土地兼并加剧  \nD. 人口大量减少  \n​**​答案​**​: B\n\n---\n\n## 命题特色\n\n- ​**​史料实证​**​：第5题结合税率对比表考查条约影响\n- ​**​时空观念​**​：第7题分析地理因素在战争中的作用\n- ​**​因果解释​**​：第10题考查经济结构变化的内在逻辑"

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	messages := []model.ConversationMessage{
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "user",
			Message:     "请按照要求为我生成一份单元作业设计案例",
			MessageType: 1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "ai",
			Message:     chatAnswer,
			MessageType: 6,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	_, err = service.SaveMessage(sessionID, messages)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrSyncToMySQL, err)
	}

	time.Sleep(5 * time.Second)

	return &types.GenerateUnitHomeworkDesignResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.ChatAnswer{Message: chatAnswer},
	}, nil
}
