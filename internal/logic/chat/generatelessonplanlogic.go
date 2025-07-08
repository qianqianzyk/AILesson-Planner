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

type GenerateLessonPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateLessonPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateLessonPlanLogic {
	return &GenerateLessonPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateLessonPlanLogic) GenerateLessonPlan(req *types.GenerateLessonPlanReq) (resp *types.GenerateLessonPlanResp, err error) {
	sessionID := req.SessionID
	subject := req.Subject
	textBookName := req.TextBookName
	topicName := req.TopicName
	topicHours := req.TopicHours
	templateFile := req.TemplateFile
	resourceFile := req.ResourceFile
	textBookImg := req.TextBookImg
	description := req.Description

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	answer := "# 探究鸦片战争的起因和影响 教案设计\n\n**课题名称**：探究鸦片战争的起因和影响\n\n​**​专业学科​**​：中国近代史\n\n​**​总课时​**​：3.0 课时（每课时45分钟）\n\n------\n\n## **一、教学目标**\n\n1. **知识目标**：\n    - 理解鸦片战争的历史背景、直接导火索及根本原因。\n    - 掌握《南京条约》等不平等条约的核心内容及其对中国近代化的深远影响。\n    - 分析鸦片战争对中国社会、经济、外交的多维度冲击。\n2. **能力目标**：\n    - 培养学生通过史料对比（如中英贸易数据、清廷奏折与英国外交档案）形成独立历史观点。\n    - 提升批判性思维，区分直接原因（如林则徐禁烟）与结构性矛盾（如朝贡体系与殖民扩张的冲突）。\n3. **情感目标**：\n    - 激发学生对近代中国命运的历史共情，认识落后挨打的深层逻辑。\n    - 引导学生反思历史事件对国家主权意识的塑造作用。\n\n------\n\n## **二、教学内容与课时分配**\n\n### **第一课时：鸦片战争的历史背景与起因**\n\n1. **知识要点**：\n    - 19世纪全球贸易格局中的中英地位差异（茶叶、丝绸与鸦片三角贸易）。\n    - 清朝闭关政策与英国工业革命后扩张需求的碰撞。\n    - 林则徐虎门销烟事件的双重性：民族气节与外交失策。\n2. **师生互动**：\n    - **角色模拟**：分组扮演清廷官员、英国东印度公司代表，辩论“白银外流与鸦片贸易合法性”。\n    - **史料分析**：对比《道光朝实录》与英国议会辩论记录，讨论双方决策逻辑差异。\n3. **教学工具**：\n    - 动态PPT展示1830-1840年中英贸易数据对比图。\n    - 短视频（3分钟）还原虎门销烟现场场景。\n\n------\n\n### **第二课时：鸦片战争的进程与结果**\n\n1. **知识要点**：\n    - 关键战役分析：穿鼻海战、镇江保卫战中的战术对比（传统水师vs蒸汽舰队）。\n    - 《南京条约》核心条款解析：五口通商、协定关税、治外法权的制度性危害。\n2. **师生互动**：\n    - **地图推演**：使用动态地图软件标注英军进攻路线，学生分组总结清军战略失误。\n    - **条约谈判模拟**：中英代表就赔款数额、香港割让条款进行“外交磋商”。\n3. **深度讨论**：\n    - 为何林则徐被称为“开眼看世界第一人”，却未能阻止战争爆发？\n    - 对比同时期日本黑船事件，分析中日应对西方冲击的路径差异。\n\n------\n\n### **第三课时：鸦片战争的深远影响与历史反思**\n\n1. **知识要点**：\n    - 朝贡体系崩塌与“条约体系”强加的国际秩序重构。\n    - 经济殖民化萌芽：手工纺织业解体与买办阶级兴起。\n    - 思想启蒙契机：魏源《海国图志》与洋务运动的先声作用。\n2. **师生互动**：\n    - **辩论赛**：正方“鸦片战争加速中国近代化” vs 反方“战争彻底破坏发展自主权”。\n    - **历史档案研读**：分组分析马戛尔尼使团报告与鸦片战争后英国商人信件，透视认知变化。\n3. **批判性思考**：\n    - 从全球史视角重新定位鸦片战争：是区域性冲突还是帝国主义全球扩张的必然？\n    - 当代香港问题与《南京条约》的历史勾连及其现实启示。\n\n------\n\n## **三、评估方式**\n\n1. **形成性评估**（40%）：\n    - 课堂角色扮演表现（语言准确性、历史逻辑性）。\n    - 小组辩论中的论点构建与史料运用能力。\n2. **总结性评估**（60%）：\n    - 期末论文选题（示例）：\n     《从茶叶到鸦片：重估19世纪中英经济博弈的深层逻辑》\n     《虎门销烟的双重镜像：民族主义叙事与全球毒品贸易史的对话》\n\n------\n\n## **四、教学资源拓展**\n\n1. **核心文献**：\n    - 茅海建《天朝的崩溃：鸦片战争再研究》（制度分析范本）。\n    - 费正清《剑桥中国晚清史》第三章（全球史视角）。\n2. **数字资源**：\n    - 大英博物馆数字化鸦片战争档案（https://www.britishmuseum.org/）。\n    - 香港历史博物馆“鸦片战争”虚拟展厅。\n\n------\n\n## **五、教学反思预设**\n\n1. 针对学生可能存在的“历史决定论”误区，设计辨析环节：如果没有鸦片贸易，中英冲突是否仍不可避免？\n2. 警惕民族主义情绪过度影响理性分析，通过比较1840年英国议会投票记录（271:262反对开战）展现历史偶然性。\n\n------\n\n此教案通过**问题导向学习**和**多模态史料介入**，突破传统战争史教学的单一叙事模式，着力培养历史思维的复杂性与跨学科整合能力。\n\n"
	//answer, err := service.GenerateLessonPlan(l.svcCtx.Config.AI.TPlanEndpoint, textBookName, subject, topicHours, topicName, templateFile, resourceFile, textBookImg, description)
	//if err != nil {
	//	return nil, utils.AbortWithException(utils.ErrGenAnswer, err)
	//}

	messages := []model.ConversationMessage{
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "user",
			Message:     description,
			MessageType: 1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "ai",
			Message:     answer,
			MessageType: 3,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	messageID, err := service.SaveMessage(sessionID, messages)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrSyncToMySQL, err)
	}

	tPlan := model.TPlan{
		UserID:       int(userID),
		MessageID:    messageID,
		Subject:      subject,
		TextBookName: textBookName,
		TopicHours:   topicHours,
		TopicName:    topicName,
		TemplateFile: templateFile,
		ResourceFile: resourceFile,
		TextBookImg:  textBookImg,
		Description:  description,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = service.CreateTPlan(&tPlan)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGenTPlan, err)
	}

	time.Sleep(4 * time.Second)

	return &types.GenerateLessonPlanResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.ChatAnswer{
			Message: answer,
		},
	}, nil
}
