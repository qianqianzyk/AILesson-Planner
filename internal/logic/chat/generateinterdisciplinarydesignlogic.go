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

type GenerateInterdisciplinaryDesignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateInterdisciplinaryDesignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateInterdisciplinaryDesignLogic {
	return &GenerateInterdisciplinaryDesignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateInterdisciplinaryDesignLogic) GenerateInterdisciplinaryDesign(req *types.GenerateInterdisciplinaryDesignReq) (resp *types.GenerateInterdisciplinaryDesignResp, err error) {
	sessionID := req.SessionID

	chatAnswer := "# 中国近代史跨学科教学设计\n\n## 基本信息\n- ​**​主学课​**​: 中国近代史  \n- ​**​阶段​**​: 大学  \n- ​**​册别​**​: 全册  \n- ​**​教材​**​: 《中国近现代史纲要》（2023年版）  \n- ​**​课题​**​: 鸦片战争前后的中国与世界  \n- ​**​涉及学科​**​: 历史、军事、社科  \n- ​**​课时安排​**​: 2课时（90分钟）\n\n---\n\n## 教学目标\n### 知识目标\n1. 理解鸦片战争前清朝的朝贡体系与闭关政策\n2. 分析英国工业革命后的全球扩张需求\n3. 掌握《南京条约》等不平等条约的核心内容\n4. 比较中西方军事技术差距的具体表现\n\n### 能力目标\n1. 培养多维度历史比较分析能力\n2. 提升军事战略思维与地缘政治分析能力\n3. 掌握社会科学研究方法论\n\n### 价值目标\n1. 认识现代化进程中的国家主权意识\n2. 培养国际视野下的历史反思能力\n\n---\n\n## 教学活动设计\n\n### 第一环节：历史维度探究（30分钟）\n1. ​**​时间轴对比分析​**​\n    - 分组整理1750-1840年间中英两国在以下领域的对比数据：​\n      - GDP总量与人均值  \n      - 对外贸易构成  \n      - 军事支出占比\n2. ​**​文献分析工作坊​**​\n    - 精读《粤海关志》与《东印度公司档案》节选\n    - 制作中英贸易认知差异对比表\n\n​**​教学资源​**​:\n - 《中国近代经济史统计资料》\n - 大英博物馆数字档案\n\n---\n\n### 第二环节：军事维度探究（25分钟）\n1. ​**​武器装备实景模拟​**​\n    - 通过3D模型对比分析：\n      - 清军鸟枪与英军燧发枪射速比  \n      - 虎门炮台防御体系漏洞\n2. ​**​战略沙盘推演​**​\n    - 分组扮演中英指挥官\n    - 复盘珠江口战役决策过程\n\n​**​教学资源​**​:\n - 军事科学院《中西火器发展史》\n - 数字战场还原系统\n\n---\n\n### 第三环节：社科维度探究（25分钟）\n1. ​**​社会结构对比研究​**​\n    - 制作1840年前后：\n      - 中国士绅阶层分布图  \n      - 英国资产阶级力量增长曲线\n2. ​**​国际法理分析​**​\n    - 对比《威斯特伐利亚体系》与《朝贡体系》\n    - 撰写条约体系转型分析报告\n\n​**​教学资源​**​:\n - 《国际法历史案例集》\n - 清华大学社会科学数据库\n\n---\n\n## 跨学科整合\n- ​**​历史军事联动​**​：分析马戛尔尼使团军礼争议背后的制度差异  \n- ​**​军事社科交汇​**​：探讨条约口岸开放对传统社会结构的冲击  \n- ​**​历史社科融合​**​：研究白银外流导致的社会治理危机  \n\n---\n\n## 评估方式\n1. 小组提交《多维对比分析报告》\n2. 课堂战略推演表现评估\n3. 设计1份包含图表的历史测试题\n\n---\n\n## 延伸学习\n1. 参观鸦片战争博物馆虚拟展\n2. 阅读茅海建《天朝的崩溃》第3章"

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	messages := []model.ConversationMessage{
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "user",
			Message:     "请按照要求为我生成一份跨学科设计案例",
			MessageType: 1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "ai",
			Message:     chatAnswer,
			MessageType: 5,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	_, err = service.SaveMessage(sessionID, messages)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrSyncToMySQL, err)
	}

	time.Sleep(5 * time.Second)

	return &types.GenerateInterdisciplinaryDesignResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.ChatAnswer{
			Message: chatAnswer,
		},
	}, nil
}
