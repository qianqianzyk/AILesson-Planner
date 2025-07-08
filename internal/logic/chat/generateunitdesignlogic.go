package chat

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateUnitDesignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateUnitDesignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateUnitDesignLogic {
	return &GenerateUnitDesignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateUnitDesignLogic) GenerateUnitDesign(req *types.GenerateUnitDesignReq) (resp *types.GenerateUnitDesignResp, err error) {
	sessionID := req.SessionID
	chatAnswer := "# 中国近代史第一章教学案例\n\n​**​专业学科​**​：中国近代史  \n​**​所需课时​**​：8课时\n\n---\n\n## 第一节 鸦片战争前后的中国与世界\n\n### 一、中国封建社会的衰弱\n1. ​**​经济基础​**​：\n    - 小农经济占主导地位\n    - 土地兼并严重，农民负担加重\n2. ​**​政治制度​**​：\n    - 君主专制制度僵化\n    - 官僚体系腐败现象加剧\n3. ​**​社会矛盾​**​：\n    - 民族矛盾与阶级矛盾交织\n    - 白莲教起义等农民反抗频发\n\n### 二、世界资本主义的发展与殖民扩张\n- ​**​工业革命影响​**​：\n    - 英国完成工业革命，急需海外市场\n    - 生产力飞跃推动殖民扩张\n- ​**​殖民体系形成​**​：\n    - 列强争夺原料产地和商品倾销地\n    - 印度、东南亚已沦为殖民地\n\n### 三、鸦片战争的爆发\n1. ​**​直接导火索​**​：\n    - 1839年林则徐虎门销烟\n    - 英国国会通过对华战争拨款\n2. ​**​战争进程​**​：\n    - 1840年英军封锁珠江口\n    - 清军节节败退至南京条约签订\n3. ​**​关键战役​**​：\n    - 穿鼻海战\n    - 镇江保卫战\n\n---\n\n## 第二节 西方列强对中国的侵略\n\n### 一、军事侵略\n- ​**​典型事件​**​：\n   - 第二次鸦片战争（1856-1860）\n   - 八国联军侵华（1900）\n- ​**​侵略方式​**​：\n   - 舰炮外交\n   - 强占军事要地\n\n### 二、政治控制\n1. ​**​条约体系​**​：\n    - 《南京条约》开割地赔款先例\n    - 《天津条约》扩大领事裁判权\n2. ​**​代理人统治​**​：\n    - 扶持清政府作为统治工具\n    - 海关总税务司由英国人掌控\n\n### 三、经济掠夺\n- ​**​资源掠夺​**​：\n   - 掠夺茶叶、生丝等原料\n   - 控制铁路、矿山开采权\n- ​**​资本输出​**​：\n   - 开设汇丰等外资银行\n   - 借款附带政治条件\n\n### 四、文化渗透\n1. ​**​传教活动​**​：\n    - 建立教堂和教会学校\n    - 教案频发（如天津教案）\n2. ​**​意识形态输出​**​：\n    - 宣扬\"西方文明优越论\"\n    - 培养亲西方知识分子群体\n\n---\n\n## 第三节 反抗外国武装侵略的斗争\n\n### 一、抵御外来侵略的斗争历程\n- ​**​三元里抗英​**​（1841）：\n   - 民间自发组织抗英斗争\n   - 击毙英军少校毕霞\n- ​**​黑旗军抗法​**​（1883-1885）：\n   - 刘永福部取得纸桥大捷\n   - 延缓法国吞并越南进程\n\n### 二、义和团运动与列强瓜分中国图谋的破产\n1. ​**​运动特点​**​：\n    - \"扶清灭洋\"口号\n    - 波及直隶、山东等省份\n2. ​**​历史影响​**​：\n    - 粉碎列强立即瓜分中国计划\n    - 迫使列强采取\"以华治华\"策略\n\n---\n\n## 第四节 反侵略战争的失败与民族意识的觉醒\n\n### 一、反侵略战争的失败及其原因\n- ​**​根本原因​**​：\n   - 社会制度腐败（封建专制）\n   - 经济技术落后\n- ​**​直接表现​**​：\n   - 军事指挥体系混乱\n   - 武器装备代差明显\n\n### 二、民族意识的觉醒\n1. ​**​思想启蒙​**​：\n    - 魏源《海国图志》\"师夷长技以制夷\"\n    - 严复翻译《天演论》引入进化论\n2. ​**​实践探索​**​：\n    - 洋务运动（1861-1894）\n    - 戊戌变法（1898）\n\n---\n\n## 教学特色\n- ​**​时空对比表​**​：18世纪中英社会发展对比\n- ​**​史料研读​**​：林则徐《四洲志》选段分析\n- ​**​思辨讨论​**​：\"鸦片战争是否具有历史必然性？\"\n- ​**​数字技术​**​：3D还原圆明园被毁前后对比"

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	messages := []model.ConversationMessage{
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "user",
			Message:     "请按照要求为我生成一份大单元教学设计案例",
			MessageType: 1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "ai",
			Message:     chatAnswer,
			MessageType: 4,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	_, err = service.SaveMessage(sessionID, messages)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrSyncToMySQL, err)
	}

	time.Sleep(5 * time.Second)

	return &types.GenerateUnitDesignResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.ChatAnswer{Message: chatAnswer},
	}, nil
}
