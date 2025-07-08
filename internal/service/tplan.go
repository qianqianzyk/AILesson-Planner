package service

import (
	"bytes"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/unidoc/unioffice/document"
	"io"
	"regexp"
	"strings"
)

func CreateTPlan(tPlan *model.TPlan) error {
	err := d.CreateTPlan(ctx, tPlan)
	return err
}

func GetTPlanByMessageID(messageID int) (*model.TPlan, error) {
	tPlan, err := d.GetTPlanByMessageID(ctx, messageID)
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

func DeleteTPlanByID(messageID, userID int) error {
	tPlan, err := GetTPlanByMessageID(messageID)
	if err != nil {
		return err
	}
	deletedUrls := FormatUrls(tPlan.TemplateFile, tPlan.ResourceFile, tPlan.TextBookImg, tPlan.TPlanUrl)
	err = d.DeleteTPlanByMessageID(ctx, messageID, userID)
	if err != nil {
		return err
	}
	DeleteFiles(deletedUrls)
	return nil
}

func GenerateWordDoc(tPlan *model.TPlan, message string) (string, error) {
	doc := document.New()

	//para := doc.AddParagraph()
	//para.SetStyle("Heading1")
	//para.AddRun().AddText(fmt.Sprintf("%s 教案设计", tPlan.TextBookName))
	//
	//para = doc.AddParagraph()
	//run := para.AddRun()
	//run.Properties().SetBold(true)
	//run.AddText("专业学科: ")
	//run = para.AddRun()
	//run.AddText(fmt.Sprintf("%s", tPlan.Subject))
	//
	//para = doc.AddParagraph()
	//run = para.AddRun()
	//run.Properties().SetBold(true)
	//run.AddText("课题名称: ")
	//run = para.AddRun()
	//run.AddText(fmt.Sprintf("%s", tPlan.TopicName))
	//
	//para = doc.AddParagraph()
	//run = para.AddRun()
	//run.Properties().SetBold(true)
	//run.AddText("总课时: ")
	//run = para.AddRun()
	//run.AddText(fmt.Sprintf("%s", tPlan.TopicHours))

	cleanedContent := cleanHTML(message)
	lines := strings.Split(cleanedContent, "\n")

	var para document.Paragraph
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
				para.SetStyle("Heading1")
			case 2:
				para = doc.AddParagraph()
				para.SetStyle("Heading2")
			case 3:
				para = doc.AddParagraph()
				para.SetStyle("Heading3")
			case 4:
				para = doc.AddParagraph()
				para.SetStyle("Heading4")
			case 5:
				para = doc.AddParagraph()
				para.SetStyle("Heading5")
			case 6:
				para = doc.AddParagraph()
				para.SetStyle("Heading6")
			}

			boldRegex := regexp.MustCompile(`\*\*(.*?)\*\*`)
			locs := boldRegex.FindAllStringIndex(line, -1)

			last := 0
			for _, loc := range locs {
				if loc[0] > last {
					para.AddRun().AddText(line[last:loc[0]])
				}
				boldText := line[loc[0]+2 : loc[1]-2]
				run := para.AddRun()
				run.Properties().SetBold(true)
				run.AddText(boldText)
				last = loc[1]
			}
			if last < len(line) {
				para.AddRun().AddText(line[last:])
			}
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
					para = doc.AddParagraph()
					boldRegex := regexp.MustCompile(`\*\*(.*?)\*\*`)
					locs := boldRegex.FindAllStringIndex(line, -1)

					last := 0
					for _, loc := range locs {
						if loc[0] > last {
							para.AddRun().AddText(line[last:loc[0]])
						}
						boldText := line[loc[0]+2 : loc[1]-2]
						run := para.AddRun()
						run.Properties().SetBold(true)
						run.AddText(boldText)
						last = loc[1]
					}
					if last < len(line) {
						para.AddRun().AddText(line[last:])
					}
				}
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
