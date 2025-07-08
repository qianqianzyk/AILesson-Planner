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

type GenerateOutlineDesignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateOutlineDesignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateOutlineDesignLogic {
	return &GenerateOutlineDesignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateOutlineDesignLogic) GenerateOutlineDesign(req *types.GenerateOutlineDesignReq) (resp *types.GenerateOutlineDesignResp, err error) {
	sessionID := req.SessionID
	chatAnswer := "# **这是依据您的主题生成的内容大纲，觉得怎么样？**\n\n# 中国近代史重点复习\n\n## 第一章：中国封建社会的衰落\n\n### 1.1 政治原因\n\n#### 1.1.1 封建中央集权制度\n\n- 封建的中央集权制度达到顶峰，皇权对人民的束缚达到最大，阶级矛盾尖锐。\n\n#### 1.1.2 文化原因\n\n- 明朝实行八股取试，清朝大行文字狱，摧残了知识分子，严重阻碍了文化的发展和科技进步。儒家思想顽固保守，夜郎自大。\n\n### 1.2 经济原因\n\n#### 1.2.1 自然经济\n\n- 自给自足的自然经济为基础的封建经济不再起积极作用，反而成为阻碍经济发展的重要因素。\n\n#### 1.2.2 外交原因\n\n- 明朝自郑和以来鲜有主动的对外交往，清朝则走向闭关锁国的道路，阻碍了文化的交流发展。\n\n## 第二章：近代中国社会的性质与特征\n\n### 2.1 社会性质\n\n#### 2.1.1 半殖民地半封建社会\n\n- 资本-帝国主义侵略势力逐步操纵了中国的财政和经济命脉，控制了中国的政治，日益成为支配中国的决定性力量。\n\n#### 2.1.2 封建势力\n\n- 中国的封建势力日益衰败并同外国侵略势力勾结，成为资本-帝国主义压迫、奴役中国人民的社会基础和统治支柱。\n\n### 2.2 基本特征\n\n#### 2.2.1 经济基础\n\n- 中国自然经济的基础虽然遭到破坏，但是封建剥削制度的根基即封建地主的土地所有制依然在广大地区内保持着，成为中国发展进步的严重障碍。\n\n#### 2.2.2 民族资本主义\n\n- 中国新兴的民族资本主义经济虽然已经产生，并在政治、文化生活中起了一定的作用，但是在帝国主义和封建主义的压迫下，它的发展很缓慢，力量很软弱，而且它的大部分与外国资本-帝国主义和本国封建主义都有或多或少的联系。\n\n## 第三章：近代中国的社会主要矛盾与历史任务\n\n### 3.1 社会主要矛盾\n\n#### 3.1.1 帝国主义与中华民族的矛盾\n\n- 帝国主义与中华民族的矛盾是近代中国社会的最主要矛盾。\n\n#### 3.1.2 封建主义与人民大众的矛盾\n\n- 封建主义与人民大众的矛盾是近代中国社会的另一主要矛盾。\n\n### 3.2 两大历史任务\n\n#### 3.2.1 推翻半殖民地半封建社会制度\n\n- 必须推翻帝国主义、封建主义联合统治的半殖民地半封建的社会制度，争得民族独立和人民解放。\n\n#### 3.2.2 实现国家富强和人民富裕\n\n- 必须改变中国经济技术落后的面貌，实现国家的富强和人民的富裕。\n\n## 第四章：近代西方列强对中国的侵略活动\n\n### 4.1 军事侵略\n\n#### 4.1.1 发动侵略战争\n\n- 发动侵略战争，屠杀中国人民。\n\n#### 4.1.2 侵占中国领土\n\n- 侵占中国领土，划分势力范围。\n\n### 4.2 政治控制\n\n#### 4.2.1 控制内政、外交\n\n- 控制中国的内政、外交。\n\n#### 4.2.2 镇压反抗\n\n- 镇压中国人民的反抗，扶植、收买代理人。\n\n### 4.3 经济掠夺\n\n#### 4.3.1 控制通商口岸\n\n- 控制中国的通商口岸，剥夺中国的关税自主权。\n\n#### 4.3.2 实行商品倾销和资本输出\n\n- 实行商品倾销和资本输出，操纵中国的经济命脉。\n\n### 4.4 文化渗透\n\n#### 4.4.1 宗教外衣\n\n- 披着宗教外衣，进行侵略活动。\n\n#### 4.4.2 制造舆论\n\n- 为侵略中国制造舆论。\n\n## 第五章：近代中国反对外国侵略战争的意义\n\n### 5.1 教育和振奋民族精神\n\n#### 5.1.1 提高民族觉醒意识\n\n- 教育了中国人民，振奋了中华民族的民族精神，鼓舞了人民反帝反封建的斗志，大大提高了中国人民的民族觉醒意识。\n\n#### 5.1.2 打击帝国主义野心\n\n- 使中国人民免于奴役，沉重打击了帝国主义侵华的野心，粉碎了他们瓜分中国和把中国变成完全殖民地的图谋。\n\n### 5.2 增强民族认同感和凝聚力\n\n#### 5.2.1 民族利益休戚与共\n\n- 增强了中华民族整体民族利益休戚与共的民族认同感和凝聚力，成为中华民族自立自强并永远立于世界民族之林的根本所在。\n\n## 第六章：近代中国反侵略战争失败的原因\n\n### 6.1 社会制度的腐败\n\n#### 6.1.1 政治腐败\n\n- 清政府统治腐败，自给自足的自然经济不能满足社会发展的需要。\n\n#### 6.1.2 军事落后\n\n- 军队训练、管理、指挥落后，装备落后。\n\n### 6.2 思想保守\n\n#### 6.2.1 儒家思想\n\n- 儒家思想顽固保守，夜郎自大，闭关锁国。\n\n## 第七章：民族意识的觉醒历程\n\n### 7.1 洋务运动\n\n#### 7.1.1 主要代表\n\n- 主要代表：奕訢、曾国藩、左宗棠、张之洞、李鸿章。\n\n#### 7.1.2 指导思想\n\n- 指导思想：中体西用。\n\n#### 7.1.3 主要内容\n\n- 主要内容：军事、经济、教育方面的改革。\n\n#### 7.1.4 历史作用\n\n- 历史作用：开启了中国近代工业化的开始，促进了中国东南沿海、长江沿岸城市的发展。\n\n#### 7.1.5 失败原因\n\n- 失败原因：封建性、对外国依赖性、企业管理腐朽性。\n\n### 7.2 戊戌维新运动\n\n#### 7.2.1 主要代表\n\n- 主要代表：康有为、梁启超。\n\n#### 7.2.2 意义\n\n- 意义：爱国救亡运动、资产阶级性质的政治改革运动、思想启蒙运动。\n\n#### 7.2.3 失败原因\n\n- 失败原因：反对派势力强大，维新派自身的局限性。\n\n### 7.3 辛亥革命\n\n#### 7.3.1 发展历程\n\n- 发展历程：武昌起义、南京临时政府成立、清帝退位。\n\n#### 7.3.2 历史意义\n\n- 历史意义：结束封建君主专制制度，建立资产阶级共和国，激发民族觉醒。\n\n#### 7.3.3 失败原因\n\n- 失败原因：没有提出彻底的反帝反封建革命纲领，不能依靠和充分发动群众，不能建立一个坚强的革命政党。\n\n### 7.4 新文化运动\n\n#### 7.4.1 兴起\n\n- 兴起：1915年，陈独秀、胡适等领导，以《新青年》为主要阵地。\n\n#### 7.4.2 意义\n\n- 意义：启蒙运动、思想解放运动，为马克思主义在中国的传播创造了有利条件。\n\n### 7.5 五四运动\n\n#### 7.5.1 爆发\n\n- 爆发：1919年5月4日，巴黎和会上中国的外交失败。\n\n#### 7.5.2 意义\n\n- 意义：打击帝国主义和封建主义，中国新民主主义革命的开端，为中国共产党的成立作了思想上的准备。\n\n## 第八章：中国共产党的成立及其意义\n\n### 8.1 成立\n\n#### 8.1.1 一大内容\n\n- 一大内容：通过党纲，确定党的名称，确定奋斗目标，选举中央局书记。\n\n#### 8.1.2 二大内容\n\n- 二大内容：制定党的最高纲领和最低纲领。\n\n#### 8.1.3 建党精神\n\n- 建党精神：坚持真理、坚守理想，践行初心、担当使命，不怕牺牲、英勇斗争，对党忠诚、不负人民。\n\n### 8.2 意义\n\n#### 8.2.1 改变发展方向\n\n- 意义：深刻改变了近代以后中华民族发展的方向和进程，深刻改变了中国人民和中华民族的前途和命运。\n\n#### 8.2.2 改变世界格局\n\n- 意义：深刻改变了世界发展的趋势和格局。\n\n## 第九章：中国革命的新面貌\n\n### 9.1 反帝反封建纲领\n\n#### 9.1.1 提出纲领\n\n- 第一次提出了反帝反封建的民主革命纲领，为中国人民指出了明确的斗争目标。\n\n#### 9.1.2 发动群众\n\n- 发动工农群众开展革命斗争。\n\n### 9.2 国共合作\n\n#### 9.2.1 合作意义\n\n- 国共合作有利于中国革命的发展，也有利于两党的发展。\n\n#### 9.2.2 合作失败\n\n- 合作失败的原因和教训。\n\n### 9.3 大革命\n\n#### 9.3.1 失败原因\n\n- 失败的原因和教训。"

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	messages := []model.ConversationMessage{
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "user",
			Message:     "主题：中国近代史重点复习\n请按照要求为我生成一份PPT",
			MessageType: 1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "ai",
			Message:     chatAnswer,
			MessageType: 8,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	_, err = service.SaveMessage(sessionID, messages)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrSyncToMySQL, err)
	}

	time.Sleep(4 * time.Second)

	return &types.GenerateOutlineDesignResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.ChatAnswer{Message: chatAnswer},
	}, nil
}
