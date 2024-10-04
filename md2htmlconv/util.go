package main

import (
	"fmt"
	"log"
	"regexp" // 정규식을 사용하기 우히ㅏㄴ 패키지
	"runtime"
	"strconv" // 문자열과 기본 데이터 타입 간의 변환을 지원
	"strings"

	"os/exec"
)

func generateHTMLHead(title string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
</head>
<body>
`, title)
}

func wrapHTMLBody() string {
	return `</body>
</html>`
}

// h1 ~ h6 까지 적용 (개발 완료)
func convertHeaders(mdText string) string {
	// 해당 정규식의 경우에는 # Heading 1, # Another heading 등이 매칭이 됨.
	re := regexp.MustCompile(`(?m)^(#{1,6}) (.*)$`)
	// 뒤의 함수는 match된 문자열을 받아서 새로운 문자열로 반환하는 함수
	return re.ReplaceAllStringFunc(mdText, func(match string) string {
		// 헤더 레벨 계산
		// 첫 번재 공백을 기준으로 하여 문자열을 2개의 부분으로 나눔. ex) ## example
		splitTextArray := strings.SplitN(match, " ", 2)

		// #의 개수 및 텍스트 추출
		headingLevel := len(splitTextArray[0])
		headingText := splitTextArray[1]

		tag := "h" + strconv.Itoa(headingLevel)

		// Itoa : 정수형 문자를 문자열로 변경
		return "<" + tag + ">" + headingText + "</" + tag + ">"
	})
}

// 텍스트 포맷 (개발 완료)
func convertTextFormat(mdText string) string {
	re := regexp.MustCompile(`\*\*(.*?)\*\*`) // bold text
	mdText = re.ReplaceAllString(mdText, "<strong>$1</strong>")

	re = regexp.MustCompile(`\_(.*?)\_`)
	mdText = re.ReplaceAllString(mdText, "<em>$1</em>")

	return mdText
}

func Convert(mdText string) string {
	mdText = convertHeaders(mdText)
	mdText = convertTextFormat(mdText)

	return mdText
}

func openBrowser(filePath string) {
	var err error

	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", filePath).Start()
	case "linux":
		err = exec.Command("/usr/bin/google-chrome", filePath).Start()
	case "windows":
		err = exec.Command("start", filePath).Start()
	default:
		log.Fatalf("unsupported platform")
	}

	if err != nil {
		log.Fatalf("Error opening browser: %v", err)
	}
}