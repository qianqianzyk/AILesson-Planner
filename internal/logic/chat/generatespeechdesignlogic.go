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

type GenerateSpeechDesignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateSpeechDesignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateSpeechDesignLogic {
	return &GenerateSpeechDesignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateSpeechDesignLogic) GenerateSpeechDesign(req *types.GenerateSpeechDesignReq) (resp *types.GenerateSpeechDesignResp, err error) {
	sessionID := req.SessionID

	chatAnswer := "# 说课稿设计\n\n## 教材信息\n\n**教材名称**：《中国近现代史纲要》（2023年版）\n\n**专业学科**：中国近代史\n\n**对应章节**：第一章第三节\n\n**衔接关系**：\n\n- 前承：清朝闭关政策  \n- 后启：第二次鸦片战争  \n- 横向关联：同时期英国议会改革（1832年）\n\n***\n\n## 教学基本信息\n\n**课题名称**：探究鸦片战争的起因和影响\n\n**授课对象**：计算机科学与技术2303/2304班\n\n**总课时**：1.5课时（75分钟）\n\n**说课场景**：同行教学评审交流\n\n***\n\n## 设计框架\n\n### 一、开头导入\n\n**数字档案对比法** \n**展示两组数据**：  \n\n1. 1830年中英贸易逆差（中国顺差600万两白银）  \n\n2. 1845年英国输华商品总值（比1836年下降53%） \n\n   **设问**：\n\n   - 为什么工业强国商品在中国滞销？ \n   - 贸易数据突变与战争决策有何关联？\n\n***\n\n### 二、教材分析\n\n**教材定位**  \n\n1. **知识结构**： \n    - 战争起因三维度：经济（鸦片贸易）、政治（朝贡体系）、技术（军事代差） \n    - 影响双面性：主权丧失与技术输入  \n2. **学术更新**： \n    - 新增大英档案馆解密文件引用 \n    - 采用GDP购买力平价法重估战争损失\n\n***\n\n### 三、学情分析\n\n**学生特征**  \n\n- **优势**： \n   - 计算机专业学生擅数据建模（可开展贸易数据可视化项目） \n   - 熟悉网络资源检索（引导使用大英博物馆数字档案）  \n- **挑战**： \n   - 历史时空观念较弱（需强化19世纪全球化进程认知） \n   - 政策背景理解困难（需解析东印度公司运作机制）\n\n***\n\n### 四、教学目标\n\n**三维目标** \n**知识目标**  \n\n- 解构战争爆发的复合型原因  \n- 辨析条约体系的连锁反应 \n  **能力目标**  \n- 掌握多源史料比对技术  \n- 培养地缘政治分析思维 \n  **价值目标**  \n- 理解技术落后与制度僵化的辩证关系  \n- 树立科技报国的责任意识\n\n***\n\n### 五、教学重难点\n\n**重点突破**  \n\n- **多维起因分析**： \n  经济（鸦片白银）、外交（律劳卑事件）、文化（天朝观念）  \n- **影响分层**： \n  直接（五口通商）vs间接（买办阶级形成）  \n\n**难点化解**  \n\n- **半殖民地化概念**： \n  通过海关主权旁移案例（赫德体系）具象化  \n- **长时段影响**： \n  使用GIS技术演示通商口岸辐射效应\n\n***\n\n### 六、教法学法\n\n**教学策略**  \n\n- **DBQ教学法**： \n  提供东印度公司财报、林则徐奏折等原始文献  \n- **数字人文工具**： \n  运用Python进行贸易数据回归分析  \n\n**学法指导**  \n\n- **证据链构建**： \n  训练从《南京条约》文本提取关键条款  \n- **跨学科迁移**： \n  用TCP/IP协议类比条约体系叠加效应\n\n***\n\n### 七、教学过程\n\n**四段式设计** \n**1. 情境建构（15分钟）**  \n\n- 播放BBC纪录片《鸦片战争》冲突片段  \n- 发放角色卡（清朝官员/英国商人/中国茶农）  \n\n**2. 实证探究（30分钟）**  \n\n- 小组任务： \n  ▫️ 计算机组：用Matplotlib绘制中英军费增长曲线 \n  ▫️ 文献组：解析《巴麦尊致中国皇帝书》语义  \n\n**3. 仿真推演（20分钟）**  \n\n- 战争决策模拟系统： \n  ▫️ 输入变量：白银外流量/舰船吨位比 \n  ▫️ 输出结果：战争爆发概率模型  \n\n**4. 迁移拓展（10分钟）**  \n\n- 当代关联：对比芯片战争中的技术封锁  \n- 布置E考据作业：爬取鸦片战争研究论文高频词\n\n***\n\n### 八、特色创新\n\n**数字人文融合**：将Python数据分析应用于军费比对  \n\n**计算思维迁移**：用决策树模型解析战争必然性  \n\n**虚实结合**：AR重现虎门炮台防御体系漏洞  \n\n---\n\n### 九、教学反思\n\n- 需注意技术工具使用时长控制（防止喧宾夺主） \n- 加强历史语境还原（避免现代思维代入）\n\n"

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	messages := []model.ConversationMessage{
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "user",
			Message:     "请按照要求为我生成一份说课稿设计案例",
			MessageType: 1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "ai",
			Message:     chatAnswer,
			MessageType: 7,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	_, err = service.SaveMessage(sessionID, messages)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrSyncToMySQL, err)
	}

	time.Sleep(5 * time.Second)

	return &types.GenerateSpeechDesignResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.ChatAnswer{Message: chatAnswer},
	}, nil
}
