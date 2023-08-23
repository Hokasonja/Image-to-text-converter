package file

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

// Create 텍스트 파일 생성 함수
func Create(filepath string, strings []string) error {
	prefix := "./public/text/"
	suffix := ".txt"

	filename, err := removePathNExtension(filepath)
	if err != nil {
		return err
	}

	file, err := os.Create(prefix + filename + suffix)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, s := range strings {
		cleanString := removeANSICodes(s)
		_, err = file.WriteString(cleanString)
		if err != nil {
			return err
		}
	}

	return nil
}

// 파일 경로와 파일 확장자를 제거한 파일명 반환
func removePathNExtension(filepath string) (string, error) {
	filepaths := strings.Split(filepath, "/")

	if len(filepaths) < 1 {
		return "", errors.New("파일경로가 올바르지 않습니다")
	}

	filename := filepaths[len(filepaths)-1]
	filenames := strings.Split(filename, ".")
	if len(filenames) < 2 {
		return "", errors.New("파일명을 제대로 입력해 주세요")
	}

	return strings.Join(filenames[:len(filenames)-1], "."), nil
}

// 색을 표현하기위한 이스케이프 코드 제거
func removeANSICodes(input string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(input, "")
}
