package score

import (
	"context"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentWrongProblemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentWrongProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentWrongProblemLogic {
	return &GetStudentWrongProblemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentWrongProblemLogic) GetStudentWrongProblem(req *types.Empty) (resp *types.GetStudentWrongProblemResp, err error) {
	return &types.GetStudentWrongProblemResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: map[string]interface{}{
			"questions": []map[string]interface{}{
				{
					"question": "鸦片战争爆发的主要原因是？",
					"options": map[string]string{
						"A": "中国实行闭关锁国政策",
						"B": "英国希望打开中国市场，获取更多利益",
						"C": "中国大量输出茶叶、丝绸等商品，导致贸易逆差",
						"D": "清政府腐败无能",
					},
					"correct_answer": []string{"A", "B"},
					"wrong_answer":   []string{"C", "D"},
					"topic_type":     2,
				},
				{
					"question": "鸦片战争期间，中国签订的第一个不平等条约是？",
					"options": map[string]string{
						"A": "南京条约",
						"B": "北京条约",
						"C": "天津条约",
						"D": "瑷珲条约",
					},
					"correct_answer": "A",
					"wrong_answer":   "B",
					"topic_type":     1,
				},
				{
					"question": "甲午战争爆发的直接导火索是什么？",
					"options": map[string]string{
						"A": "朝鲜东学党起义",
						"B": "中日关于台湾的争端",
						"C": "清政府与英国的贸易摩擦",
						"D": "清军袭击日本舰队",
					},
					"correct_answer": "A",
					"wrong_answer":   "D",
					"topic_type":     1,
				},
				{
					"question": "甲午战争的失败直接导致了什么后果？",
					"options": map[string]string{
						"A": "《南京条约》的签订",
						"B": "《马关条约》的签订",
						"C": "中国割让澳门给葡萄牙",
						"D": "清政府加强了自强运动",
					},
					"correct_answer": "B",
					"wrong_answer":   "D",
					"topic_type":     1,
				},
				{
					"question": "甲午战争后签订的《马关条约》内容包括哪些？",
					"options": map[string]string{
						"A": "割让台湾及附属岛屿给日本",
						"B": "允许日本在中国开设工厂",
						"C": "中国需赔款 2 亿两白银",
						"D": "中国承认朝鲜独立",
					},
					"correct_answer": []string{"A", "B", "C", "D"},
					"wrong_answer":   []string{"A", "C"},
					"topic_type":     2,
				},
				{
					"question":       "简述土地革命的背景、经过及影响。",
					"correct_answer": "土地革命是中国共产党在1927年至1949年间进行的一系列农民革命活动，目的是通过没收地主的土地，分配给贫苦农民，来推动社会变革。土地革命的背景包括对地主阶级的反抗及农民贫困问题，经过包括各地农民起义与政策执行，影响则表现为农村土地分配的变化，并为后续的中国共产党政权建立提供了社会基础。",
					"wrong_answer":   "土地革命主要发生在1920年代末期，目的是通过改变农村的经济结构，推动工业化进程。",
					"topic_type":     5,
				},
			},
		},
	}, nil
}
