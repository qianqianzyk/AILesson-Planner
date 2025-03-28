package utils

import (
	"errors"
	"github.com/qianqianzyk/AILesson-Planner/internal/logs"
)

var (
	ErrServer                      = NewError(200500, logs.LevelError, "系统异常，请稍后重试!")
	ErrParam                       = NewError(200501, logs.LevelInfo, "参数错误")
	ErrUsernameOrPassword          = NewError(200502, logs.LevelInfo, "账号或密码不符合格式")
	ErrUserExist                   = NewError(200503, logs.LevelInfo, "用户或邮箱已存在")
	ErrCreateUser                  = NewError(200504, logs.LevelInfo, "用户创建失败，请联系管理员")
	ErrUserNotExist                = NewError(200505, logs.LevelInfo, "该用户不存在")
	ErrLogin                       = NewError(200506, logs.LevelInfo, "邮箱或密码错误")
	ErrGenToken                    = NewError(200507, logs.LevelInfo, "token生成错误")
	ErrSendCodeLimited             = NewError(200508, logs.LevelInfo, "操作频繁，请等待60秒后再尝试发送")
	ErrVerifyEmail                 = NewError(200509, logs.LevelInfo, "邮箱验证失败，请稍后重试")
	ErrVerifyCode                  = NewError(200510, logs.LevelInfo, "邮箱验证码错误，请重试")
	ErrUpdateUser                  = NewError(200511, logs.LevelInfo, "更新失败，请重试")
	ErrOldPassword                 = NewError(200512, logs.LevelInfo, "原密码错误，请重试")
	ErrUserID                      = NewError(200513, logs.LevelInfo, "获取用户ID失败")
	ErrCreateTopic                 = NewError(200514, logs.LevelInfo, "创建会话失败")
	ErrGetTopic                    = NewError(200515, logs.LevelInfo, "获取会话失败")
	ErrUpdateTopic                 = NewError(200516, logs.LevelInfo, "修改会话失败")
	ErrGetTopicList                = NewError(200517, logs.LevelInfo, "获取会话列表失败")
	ErrUpgradeWs                   = NewError(200518, logs.LevelInfo, "websocket 升级失败")
	ErrReadMessage                 = NewError(200519, logs.LevelInfo, "websocket 消息接收失败")
	ErrMessageFormatted            = NewError(200520, logs.LevelInfo, "websocket 消息格式错误")
	ErrSendMessage                 = NewError(200521, logs.LevelInfo, "websocket 消息发送失败")
	ErrSaveToRedis                 = NewError(200522, logs.LevelInfo, "无法保存消息到redis")
	ErrSyncToMySQL                 = NewError(200523, logs.LevelInfo, "无法同步消息到mysql")
	ErrGenAnswer                   = NewError(200524, logs.LevelInfo, "回答生成失败")
	ErrGetMessageList              = NewError(200525, logs.LevelInfo, "获取会话历史记录失败")
	ErrDelTopic                    = NewError(200526, logs.LevelInfo, "删除会话失败")
	ErrAuthUser                    = NewError(200527, logs.LevelInfo, "权限不足")
	ErrSyncIndex                   = NewError(200528, logs.LevelInfo, "索引同步失败")
	ErrSearch                      = NewError(200529, logs.LevelInfo, "搜索失败")
	ErrAvatarLimited               = NewError(200530, logs.LevelInfo, "头像大小不能超过5MB")
	ErrFileNotImage                = NewError(200531, logs.LevelInfo, "上传的文件不是图片")
	ErrAvatarUpload                = NewError(200532, logs.LevelInfo, "上传头像失败")
	ErrFileExisted                 = NewError(200533, logs.LevelInfo, "该文件/目录已经存在")
	ErrCreateFile                  = NewError(200534, logs.LevelInfo, "创建文件/目录失败，请重试")
	ErrGetFile                     = NewError(200535, logs.LevelInfo, "获取文件/目录失败，请重试")
	ErrUpdateFile                  = NewError(200536, logs.LevelInfo, "更新文件/目录失败，请重试")
	ErrFileUpload                  = NewError(200537, logs.LevelInfo, "上传文件失败")
	ErrFileStats                   = NewError(200538, logs.LevelInfo, "统计文件信息失败")
	ErrDelFile                     = NewError(200539, logs.LevelInfo, "删除文件/目录失败，请重试")
	ErrRecoverFile                 = NewError(200540, logs.LevelInfo, "恢复文件/目录失败，请重试")
	ErrGenCode                     = NewError(200541, logs.LevelInfo, "生成提取码失败，请重试")
	ErrFormattedCode               = NewError(200542, logs.LevelInfo, "提取码格式错误，请重试")
	ErrGenLink                     = NewError(200543, logs.LevelInfo, "生成链接失败，请重试")
	ErrFile                        = NewError(200544, logs.LevelInfo, "文件非法，请重试")
	ErrGetLink                     = NewError(200545, logs.LevelInfo, "该链接不存在")
	ErrExpireLink                  = NewError(200546, logs.LevelInfo, "该链接已过期")
	ErrLinkCode                    = NewError(200547, logs.LevelInfo, "提取码错误，请重试")
	ErrCollectFile                 = NewError(200548, logs.LevelInfo, "收藏失败，请重试")
	ErrMoveFile                    = NewError(200549, logs.LevelInfo, "移动失败，请重试")
	ErrSearchFile                  = NewError(200550, logs.LevelInfo, "搜索失败，请重试")
	ErrCheckFileMD5                = NewError(200551, logs.LevelInfo, "检查文件MD5失败")
	ErrGenFileCert                 = NewError(200552, logs.LevelInfo, "生成上传凭证失败")
	ErrMergeFile                   = NewError(200553, logs.LevelInfo, "文件合并失败")
	ErrCreateCourse                = NewError(200554, logs.LevelInfo, "课程创建失败")
	ErrFoundStudentID              = NewError(200555, logs.LevelInfo, "Excel中缺失必填列，学号列")
	ErrImportScores                = NewError(200556, logs.LevelInfo, "一键导入失败，请重试")
	ErrGetCourse                   = NewError(200557, logs.LevelInfo, "无法获取课程信息")
	ErrGetStudentScore             = NewError(200558, logs.LevelInfo, "无法获取学生成绩")
	ErrExportScores                = NewError(200559, logs.LevelInfo, "一键导出失败，请重试")
	ErrUpdateScores                = NewError(200560, logs.LevelInfo, "更新学生信息成绩失败")
	ErrDeleteScores                = NewError(200561, logs.LevelInfo, "删除学生成绩失败")
	ErrUpsertScores                = NewError(200562, logs.LevelInfo, "添加学生信息成绩失败")
	ErrDeleteCourse                = NewError(200563, logs.LevelInfo, "删除课程失败")
	ErrGetClassList                = NewError(200564, logs.LevelInfo, "获取班级列表失败")
	ErrGetStudentList              = NewError(200565, logs.LevelInfo, "获取学生列表失败")
	ErrGetStudentGPA               = NewError(200566, logs.LevelInfo, "获取学生绩点排名失败")
	ErrGetChapterScores            = NewError(200567, logs.LevelInfo, "获取学生章节成绩失败")
	ErrUpdateCourse                = NewError(200568, logs.LevelInfo, "更新课程信息失败")
	ErrStudentExist                = NewError(200569, logs.LevelInfo, "该学号已存在")
	ErrDeleteStudents              = NewError(200570, logs.LevelInfo, "删除学生信息失败")
	ErrGetGraph                    = NewError(200571, logs.LevelInfo, "获取知识图谱失败，请重试")
	ErrUpdateGraphNode             = NewError(200572, logs.LevelInfo, "更新知识图谱节点失败，请重试")
	ErrDeleteGraphNode             = NewError(200573, logs.LevelInfo, "删除知识图谱节点失败，请重试")
	ErrGraphNodeProperty           = NewError(200574, logs.LevelInfo, "该知识图谱节点配置信息非法，请重试")
	ErrCreateGraphNode             = NewError(200575, logs.LevelInfo, "创建知识图谱节点失败，请重试")
	ErrCreateGraphNodeRelationship = NewError(200576, logs.LevelInfo, "新增知识图谱节点关系失败，请重试")
	ErrUpdateGraphNodeRelationship = NewError(200577, logs.LevelInfo, "更新知识图谱节点关系失败，请重试")
	ErrDeleteGraphNodeRelationship = NewError(200578, logs.LevelInfo, "删除知识图谱节点关系失败，请重试")
	ErrCreateExperiencePost        = NewError(200579, logs.LevelInfo, "创建经验帖失败，请重试")
	ErrDeleteExperiencePost        = NewError(200580, logs.LevelInfo, "删除经验帖失败，请重试")
	ErrGetExperiencePost           = NewError(200581, logs.LevelInfo, "获取经验贴失败，请重试")
	ErrGenTPlan                    = NewError(200582, logs.LevelInfo, "生成教案失败，请重试")
	ErrUpdateTPlan                 = NewError(200583, logs.LevelInfo, "更新教案失败，请重试")
	ErrGetTPlan                    = NewError(200584, logs.LevelInfo, "获取教案失败，请重试")
	ErrDeleteTPlan                 = NewError(200585, logs.LevelInfo, "删除教案失败，请重试")
	ErrExportTPlan                 = NewError(200586, logs.LevelInfo, "导出教案失败，请重试")

	ErrTimeLimited = errors.New("请等待60秒后再尝试发送邮件")
)
