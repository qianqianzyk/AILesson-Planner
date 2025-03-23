// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5

package handler

import (
	"net/http"

	app "github.com/qianqianzyk/AILesson-Planner/internal/handler/app"
	chat "github.com/qianqianzyk/AILesson-Planner/internal/handler/chat"
	class "github.com/qianqianzyk/AILesson-Planner/internal/handler/class"
	course "github.com/qianqianzyk/AILesson-Planner/internal/handler/course"
	disk "github.com/qianqianzyk/AILesson-Planner/internal/handler/disk"
	email "github.com/qianqianzyk/AILesson-Planner/internal/handler/email"
	es "github.com/qianqianzyk/AILesson-Planner/internal/handler/es"
	graph "github.com/qianqianzyk/AILesson-Planner/internal/handler/graph"
	score "github.com/qianqianzyk/AILesson-Planner/internal/handler/score"
	share "github.com/qianqianzyk/AILesson-Planner/internal/handler/share"
	student "github.com/qianqianzyk/AILesson-Planner/internal/handler/student"
	user "github.com/qianqianzyk/AILesson-Planner/internal/handler/user"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: app.PingHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/answer",
				Handler: chat.GetChatAnswerHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/history",
				Handler: chat.GetChatHistoryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search",
				Handler: chat.SearchMessagesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/topic",
				Handler: chat.CreateChatTopicHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/topic",
				Handler: chat.UpdateChatTopicHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/topic",
				Handler: chat.DeleteChatTopicHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/topic",
				Handler: chat.GetChatListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/ws/connect",
				Handler: chat.WsConnectHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/chat"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/export_tplan",
				Handler: chat.ExportLessonPlanHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/tplan",
				Handler: chat.GenerateLessonPlanHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/tplan",
				Handler: chat.UpdateLessonPlanHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/tplan",
				Handler: chat.GetLessonPlanListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/tplan",
				Handler: chat.DeleteLessonPlanHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/upload_file",
				Handler: chat.UploadTPlanFileHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/lesson"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/allList",
				Handler: class.GetAllClassListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/chapter",
				Handler: class.GetStudentChapterScoreHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/gpa",
				Handler: class.GetStudentGPAHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: class.GetClassListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/score",
				Handler: class.GetStudentScoresHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/stu",
				Handler: class.GetStudentListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/class"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/classList",
				Handler: course.GetClassListByCourseHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: course.CreateCourseHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/del",
				Handler: course.DeleteCourseHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: course.GetCourseListHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/put",
				Handler: course.UpdateCourseHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/table",
				Handler: course.GetCourseTableHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/course"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/cert",
				Handler: disk.GetUploadCertHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/check",
				Handler: disk.CheckFileMD5Handler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/collect",
				Handler: disk.CollectFileHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/collect",
				Handler: disk.GetCollectFileHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/directory",
				Handler: disk.GetDiskDirectoryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/directory",
				Handler: disk.CreateDiskDirectoryHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/directory",
				Handler: disk.PutDiskDirectoryHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/directory",
				Handler: disk.DeleteDiskFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file",
				Handler: disk.UploadDiskFileHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/file",
				Handler: disk.GetFileInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/fileType",
				Handler: disk.GetFilesByTypeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/find",
				Handler: disk.SearchFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/merge",
				Handler: disk.MergeFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/move",
				Handler: disk.MoveFileHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/recycle",
				Handler: disk.GetRecycleDiskFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/recycle",
				Handler: disk.RecoverDiskFileHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/recycle",
				Handler: disk.CompleteDelDiskFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/share",
				Handler: disk.ShareLinkHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/share",
				Handler: disk.GetShareLinkHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/shareDir",
				Handler: disk.GetShareDirectoryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/store",
				Handler: disk.StoreResourceHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/disk"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/code",
				Handler: email.SendCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/verify",
				Handler: email.VerifyEmailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/email"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/index",
				Handler: es.SyncIndexByHandHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/es"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/material",
				Handler: graph.GetMaterialGraphHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/material",
				Handler: graph.UpdateGraphNodeHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/material",
				Handler: graph.DeleteGraphNodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/material",
				Handler: graph.CreateGraphNodeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/material_list",
				Handler: graph.GetMaterialGraphListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/relationship",
				Handler: graph.CreateGraphNodeRelationShipHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/relationship",
				Handler: graph.UpdateGraphNodeRelationShipHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/relationship",
				Handler: graph.DeleteGraphNodeRelationShipHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/graph"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: score.CreateScoresHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: score.DeleteScoresHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/export",
				Handler: score.ExportScoresHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/import",
				Handler: score.ImportScoresHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/number",
				Handler: score.GetCountNumberHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/perf",
				Handler: score.GetClassPerformanceHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/scores",
				Handler: score.GetStudentTranscriptsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/student",
				Handler: score.GetStudentScoreHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/transcripts",
				Handler: score.GetStudentTranscriptHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/transcripts",
				Handler: score.SendTranscriptsHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: score.UpdateScoresHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/score"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/resource",
				Handler: share.CreateShareResourceHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/resource",
				Handler: share.DeleteShareResourceHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/resource",
				Handler: share.GetShareResourceListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search_resource",
				Handler: share.SearchShareResourceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/share_file",
				Handler: share.UploadShareFileHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/share_file",
				Handler: share.DeleteShareFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/share_resource",
				Handler: share.StoreShareResourceHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/share"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/info",
				Handler: student.CreateStudentInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/info",
				Handler: student.UpdateStudentInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: student.GetStudentInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/info",
				Handler: student.DelStudentInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/student"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPut,
				Path:    "/findpwd",
				Handler: user.FindPasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/reg",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: user.UploadFileHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/personal",
				Handler: user.GetPersonalInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/personal",
				Handler: user.PutPersonalInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/pwd",
				Handler: user.UpdatePasswordHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/user"),
	)
}
