package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"net/http"
	"strings"
	"sync"
)

func GenerateProblem(courseID, chapter int, tags, endpoint, apiKey string, problemType, problemNum int) ([]*model.Problem, error) {
	course, err := d.GetCourseByID(ctx, courseID)
	if err != nil {
		return nil, fmt.Errorf("查询课程名称失败: %w", err)
	}

	courseName := course.CourseName
	prompt := buildProblemPrompt(courseName, tags, problemType)

	if problemNum == 0 {
		problemNum = 5
	}

	//response, err := callTongyiAPI(endpoint, apiKey, prompt, "请你在接下来的回答中按照要求返回json格式的字符串")
	//if err != nil {
	//	return nil, fmt.Errorf("调用通义千问生成题目失败: %w", err)
	//}
	//
	//problem, err := parseResponse(response, courseID, chapter)
	//if err != nil {
	//	return nil, fmt.Errorf("解析AI返回的题目信息失败: %w", err)
	//}
	//
	//err = d.CreateProblem(ctx, problem)
	//if err != nil {
	//	return nil, fmt.Errorf("保存题目到数据库失败: %w", err)
	//}

	var wg sync.WaitGroup
	problemsChan := make(chan *model.Problem, problemNum)

	for i := 0; i < problemNum; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			response, err := callTongyiAPI(endpoint, apiKey, prompt, "请你在接下来的回答中按照要求返回json格式的字符串")
			if err != nil {
				problemsChan <- nil
				return
			}

			problem, err := parseResponse(response, courseID, chapter)
			if err != nil {
				problemsChan <- nil
				return
			}

			problemsChan <- problem
		}()
	}

	wg.Wait()
	close(problemsChan)

	var problems []*model.Problem
	for problem := range problemsChan {
		if problem != nil {
			problems = append(problems, problem)
		}
	}

	err = d.CreateProblemsBatch(ctx, problems)
	if err != nil {
		return nil, fmt.Errorf("批量创建题目记录失败: %w", err)
	}

	return problems, nil
}

func buildProblemPrompt(courseName, tags string, problemType int) string {
	problemTypeDesc := map[int]string{
		1: "单选题",
		2: "多选题",
		3: "判断题",
		4: "填空题",
		5: "简答题",
	}

	typeDesc := "随机生成"
	if problemType > 0 && problemType <= 5 {
		typeDesc = problemTypeDesc[problemType]
	}

	tagsDesc := "如果未提供，请根据题目内容生成合适的标签"
	if tags != "" {
		tagsDesc = "题目标签：" + tags
	}

	return fmt.Sprintf(`
		你是一个专业的出题助手，可以根据课程内容、题目标签和题目类型生成高质量的试题。
		请根据以下要求生成习题：
		- 课程名称：%s
		- %s
		- 题目类型：%s
		请严格输出 JSON 格式：
		{
			"question": "题目内容",
			"options": ["选项A", "选项B", "选项C", "选项D"], // 仅适用于选择题
			"answer": "正确答案",
			"score": 5.0,
			"tags": "生成的知识点标签",
			"problem_type": 1  // 1=单选, 2=多选, 3=判断, 4=填空, 5=简答
		}`, courseName, tagsDesc, typeDesc)
}

func callTongyiAPI(endpoint, apiKey, prompt, question string) (string, error) {
	messages := []model.Message{
		{Role: "system", Content: prompt},
		{Role: "user", Content: question},
	}

	requestBody := model.RequestBody{
		Model: ModelName,
		Input: model.Input{
			Messages: messages,
		},
		Parameters: model.Parameters{
			ResultFormat: ResultFormatted,
		},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("调用通义 API 失败, 状态码: %d", resp.StatusCode)
	}

	var responsePayload model.ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&responsePayload)
	if err != nil {
		return "", err
	}

	if responsePayload.Output.FinishReason == ResultNormalGen {
		return responsePayload.Output.Text, nil
	}

	return "", errors.New("AI 生成失败")
}

func parseResponse(response string, courseID, chapter int) (*model.Problem, error) {
	var generatedProblem struct {
		Question    string   `json:"question"`
		Options     []string `json:"options"`
		Answer      string   `json:"answer"`
		Score       float64  `json:"score"`
		Tags        string   `json:"tags"`
		ProblemType int      `json:"problem_type"`
	}

	err := json.Unmarshal([]byte(response), &generatedProblem)
	if err != nil {
		return nil, fmt.Errorf("解析AI返回的题目信息失败: %w", err)
	}

	if generatedProblem.ProblemType < 1 || generatedProblem.ProblemType > 5 {
		return nil, errors.New("AI 生成的 problem_type 无效")
	}

	optionsStr := strings.Join(generatedProblem.Options, "|||")

	return &model.Problem{
		CourseID:    courseID,
		Chapter:     chapter,
		Tags:        generatedProblem.Tags,
		Question:    generatedProblem.Question,
		Options:     optionsStr,
		Answer:      generatedProblem.Answer,
		Score:       generatedProblem.Score,
		ProblemType: generatedProblem.ProblemType,
	}, nil
}
