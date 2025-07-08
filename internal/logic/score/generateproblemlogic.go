package score

import (
	"context"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateProblemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateProblemLogic {
	return &GenerateProblemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateProblemLogic) GenerateProblem(req *types.Empty) (resp *types.GenerateProblemResp, err error) {
	return &types.GenerateProblemResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: []interface{}{
			map[string]interface{}{
				"topic_type": 1,
				"question":   "甲午战争的主要战场在哪个地区？",
				"options": map[string]string{
					"A": "黄海",
					"B": "长江",
					"C": "南海",
					"D": "台湾海峡",
				},
				"correct_answer": "A",
				"explanation":    "甲午战争的主要战场是在黄海，特别是黄海海战成为关键性的战斗。",
			},
			map[string]interface{}{
				"topic_type": 2,
				"question":   "鸦片战争后，中国签订的条约中有哪些内容？",
				"options": map[string]string{
					"A": "割让香港给英国",
					"B": "开放广州、上海、宁波等港口",
					"C": "赔偿巨额战争赔款",
					"D": "获得对外直接贸易权利",
				},
				"correct_answer": []string{"A", "B", "C", "D"},
				"explanation":    "鸦片战争后，中国与英国签订了《南京条约》，其中包括割让香港、开放港口、赔偿战争赔款以及允许对外直接贸易。",
			},
			map[string]interface{}{
				"topic_type":     5,
				"question":       "鸦片战争的历史背景和影响是什么？",
				"correct_answer": "鸦片战争的历史背景包括中国实施禁烟政策和西方列强的经济扩张需求。战争结果是中国被迫开放市场，签订一系列不平等条约，标志着中国沦为半殖民地化的国家，社会经济遭到严重冲击。",
				"explanation":    "鸦片战争的影响深远，不仅加剧了中国的贫弱，还使得中国的对外贸易和主权受到了严重侵犯，成为中国近代史的转折点。",
			},
			map[string]interface{}{
				"topic_type":     5,
				"question":       "甲午战争的原因及其后果是什么？",
				"correct_answer": "甲午战争的原因包括中国内部的腐败和清朝政府的软弱无力，外部则是日本的崛起和其对朝鲜半岛的野心。战争后，中国的失败导致了《马关条约》的签订，割让领土和赔偿大量战争赔款，进一步加剧了中国的半殖民地化。",
				"explanation":    "甲午战争的后果是中国的国力进一步衰退，失去了对朝鲜的影响力，并为日本进一步的扩张奠定了基础。",
			},
			map[string]interface{}{
				"topic_type":     5,
				"question":       "土地革命时期中国共产党的土地政策如何影响社会发展？",
				"correct_answer": "土地革命时期，党通过分田分地的政策，减轻了农民的负担，增强了农民的支持。这一政策不仅有助于提升群众的社会阶级意识，也为后来的抗日战争和解放战争提供了重要的社会基础。",
				"explanation":    "土地革命的成功实施，解决了贫困农民的问题，推动了农村社会的变革，为中国共产党奠定了广泛的群众基础。",
			},
		},
	}, nil
}
