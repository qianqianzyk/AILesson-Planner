package service

import (
	"bytes"
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/unidoc/unioffice/v2/document"
	"io"
	"regexp"
	"strings"
)

func CreateTPlan(tPlan *model.TPlan) error {
	err := d.CreateTPlan(ctx, tPlan)
	return err
}

func GetTPlanByID(id int) (*model.TPlan, error) {
	tPlan, err := d.GetTPlanByID(ctx, id)
	return tPlan, err
}

func ExtractUrlsFromString(input string) []string {
	if input == "" {
		return []string{}
	}
	urls := strings.Split(input, ",")
	for i, url := range urls {
		urls[i] = strings.TrimSpace(url)
	}
	return urls
}

func FormatUrls(templateFile, resourceFile, textBookImg, tPlanUrl string) []string {
	resourceFileUrls := ExtractUrlsFromString(resourceFile)
	textBookImgUrls := ExtractUrlsFromString(textBookImg)
	urls := append(resourceFileUrls, textBookImgUrls...)
	urls = append(urls, templateFile)
	urls = append(urls, tPlanUrl)
	return urls
}

func ExtractDelUrls(previousUrls, nowUrls []string) []string {
	deletedUrls := []string{}
	for _, prevUrl := range previousUrls {
		found := false
		for _, newUrl := range nowUrls {
			if prevUrl == newUrl {
				found = true
				break
			}
		}
		if !found {
			deletedUrls = append(deletedUrls, prevUrl)
		}
	}
	return deletedUrls
}

func DeleteFiles(deletedUrls []string) {
	for _, delUrl := range deletedUrls {
		DeleteObjectByUrlAsync(delUrl)
	}
}

func UpdateTPlan(tPlan *model.TPlan) error {
	err := d.UpdateTPlan(ctx, tPlan)
	return err
}

func GetTPlanList(userID int) ([]model.TPlan, error) {
	tPlans, err := d.GetTPlanList(ctx, userID)
	return tPlans, err
}

func DeleteTPlanByID(id, userID int) error {
	tPlan, err := GetTPlanByID(id)
	if err != nil {
		return err
	}
	deletedUrls := FormatUrls(tPlan.TemplateFile, tPlan.ResourceFile, tPlan.TextBookImg, tPlan.TPlanUrl)
	err = d.DeleteTPlanByID(ctx, id, userID)
	if err != nil {
		return err
	}
	DeleteFiles(deletedUrls)
	return nil
}

func GenerateWordDoc(tPlan *model.TPlan) (string, error) {
	doc := document.New()

	para := doc.AddParagraph()
	para.SetStyle("Heading1")
	para.AddRun().AddText(fmt.Sprintf("%s 教案", tPlan.Subject))

	para = doc.AddParagraph()
	run := para.AddRun()
	run.Properties().SetBold(true)
	run.AddText("教材名称: ")
	run = para.AddRun()
	run.AddText(fmt.Sprintf("%s", tPlan.TextBookName))

	para = doc.AddParagraph()
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.AddText("课题名称: ")
	run = para.AddRun()
	run.AddText(fmt.Sprintf("%s", tPlan.TopicName))

	para = doc.AddParagraph()
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.AddText("总课时: ")
	run = para.AddRun()
	run.AddText(fmt.Sprintf("%s", tPlan.TopicHours))

	cleanedContent := cleanHTML(tPlan.TPlanContent)
	lines := strings.Split(cleanedContent, "\n")

	processed := false
	for _, line := range lines {
		if !processed && (strings.HasPrefix(line, "#") || strings.HasPrefix(line, "##") ||
			strings.HasPrefix(line, "###") || strings.HasPrefix(line, "####") ||
			strings.HasPrefix(line, "#####") || strings.HasPrefix(line, "######")) {
			processed = true
		}
		if !processed {
			continue
		}
		if strings.HasPrefix(line, "#") {
			currentHeadingLevel := strings.Count(line, "#")
			line = strings.TrimSpace(strings.TrimPrefix(line, strings.Repeat("#", currentHeadingLevel)))

			switch currentHeadingLevel {
			case 1:
				para = doc.AddParagraph()
				para.SetStyle("Heading2")
			case 2:
				para = doc.AddParagraph()
				para.SetStyle("Heading3")
			case 3:
				para = doc.AddParagraph()
				para.SetStyle("Heading4")
			case 4:
				para = doc.AddParagraph()
				para.SetStyle("Heading5")
			case 5:
				para = doc.AddParagraph()
				para.SetStyle("Heading6")
			case 6:
				para = doc.AddParagraph()
				para.SetStyle("Heading7")
			}

			para.AddRun().AddText(line)
		} else {
			re := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
			matches := re.FindAllStringSubmatch(line, -1)
			if len(matches) > 0 {
				para = doc.AddParagraph()
				lastIndex := 0

				for _, match := range matches {
					linkText := match[1]
					linkURL := match[2]

					para.AddRun().AddText(line[lastIndex:strings.Index(line, match[0])])

					hl := para.AddHyperLink()
					hl.SetTarget(linkURL)
					run := hl.AddRun()
					run.Properties().SetStyle("Hyperlink")
					run.AddText(linkText)

					hl.SetToolTip("hover to see this")

					lastIndex = strings.Index(line, match[0]) + len(match[0])
				}

				if lastIndex < len(line) {
					para.AddRun().AddText(line[lastIndex:])
				}
			} else {
				doc.AddParagraph().AddRun().AddText(line)
			}
		}
	}

	var buf bytes.Buffer
	if err := doc.Save(&buf); err != nil {
		return "", err
	}

	objectKey := GenerateObjectKey("ai", ".docx")
	objectUrl, err := PutObject(objectKey, io.NopCloser(&buf), int64(buf.Len()), "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	if err != nil {
		return "", err
	}

	if tPlan.TPlanUrl != "" {
		DeleteObjectByUrlAsync(tPlan.TPlanUrl)
	}
	err = d.UpdateTPlanUrl(ctx, int(tPlan.ID), objectUrl)
	if err != nil {
		return "", err
	}
	return objectUrl, nil
}

func cleanHTML(input string) string {
	re := regexp.MustCompile(`<.*?>`)
	cleanedContent := re.ReplaceAllString(input, "\n")
	cleanedContent = strings.ReplaceAll(cleanedContent, "\n\n", "\n")
	return cleanedContent
}
