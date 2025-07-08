package score

import (
	"context"
	"errors"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportScoresLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportScoresLogic {
	return &ImportScoresLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportScoresLogic) ImportScores(w http.ResponseWriter, r *http.Request, req *types.ImportScoresReq) (resp *types.ImportScoresResp, err error) {
	courseID := req.CourseID

	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrParam, err)
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			zap.L().Warn("文件关闭错误", zap.Error(err))
		}
	}(file)

	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if len(rows) < 2 {
		return nil, utils.AbortWithException(utils.ErrServer, errors.New("excel 文件内容为空"))
	}

	headers := make(map[string]int)
	for i, col := range rows[0] {
		headers[col] = i
	}

	idxStudentID, ok := headers["学号"]
	if !ok {
		return nil, utils.AbortWithException(utils.ErrFoundStudentID, err)
	}

	idxName, hasName := headers["姓名"]
	idxCollege, hasCollege := headers["学院"]
	idxClass, hasClass := headers["班级"]
	idxMajor, hasMajor := headers["专业"]
	idxRegularScore, hasRegularScore := headers["平时成绩"]
	idxFinalScore, hasFinalScore := headers["期末成绩"]
	idxTotalScore, hasTotalScore := headers["最终成绩"]

	err = service.ImportStudentScores(rows, courseID, idxStudentID, idxName, idxCollege, idxClass, idxMajor, idxRegularScore, idxFinalScore, idxTotalScore, hasName, hasCollege, hasClass, hasMajor, hasRegularScore, hasFinalScore, hasTotalScore)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrImportScores, err)
	}

	return &types.ImportScoresResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
