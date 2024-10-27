package pkg_kit

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"os"
	"regexp"
	"strings"
)

const (
	PkgVersionRegexp     = `^(\W*)(\"version\")(\W*\:\W*)(\".*\")(\W*\,\W*)$`
	PkgVersionGroupIndex = 4
)

// ReplaceJsonVersionByLine
// replace json version by line
func ReplaceJsonVersionByLine(path string, versionNoPrefix string) error {
	return ReplaceFileLineByLine(path, PkgVersionRegexp, PkgVersionGroupIndex, fmt.Sprintf("\"%s\"", versionNoPrefix))
}

// ReplaceFileLineByLine
//
// read file as lines, this method will read all line, so if file is too big, will be slow
func ReplaceFileLineByLine(path string, reg string, index int, replace string) error {
	compLine, err := regexp.Compile(reg)
	if err != nil {
		return err
	}
	if !filepath_plus.PathExistsFast(path) {
		return fmt.Errorf("read path %s not exists", path)
	}
	if filepath_plus.PathIsDir(path) {
		return fmt.Errorf("read path %s is dir", path)
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("AlterFile get file info at path: %v, err: %v", path, err)
	}
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("read path %s error %s", path, err)
	}
	defer func(file *os.File) {
		errFileClose := file.Close()
		if errFileClose != nil {
			fmt.Printf("read close file err: %v\n", errFileClose)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	var readLine []string
	for scanner.Scan() {
		lineText := scanner.Text()
		findGroup := compLine.FindStringSubmatch(lineText)
		if len(findGroup) < index+1 {
			readLine = append(readLine, lineText)
			continue
		}
		findGroup[index] = replace
		readLine = append(readLine, strings.Join(findGroup[1:], ""))
	}
	joinData := strings.Join(readLine, "\n")

	err = os.WriteFile(path, []byte(joinData), fileInfo.Mode())
	if err != nil {
		return fmt.Errorf("ReplaceFileLineByLine write data at path: %v, err: %v", path, err)
	}

	return nil
}

// ReplaceVersion
// this way is not good, but it's work, will change json order
func ReplaceVersion(path string, version string) error {
	pkgBytes, err := filepath_plus.ReadFileAsByte(path)
	if err != nil {
		return err
	}
	var pkgJsonMap map[string]interface{}
	err = json.Unmarshal(pkgBytes, &pkgJsonMap)
	if err != nil {
		return fmt.Errorf("ReplaceVersion json.Unmarshal error %s", err)
	}
	pkgJsonMap["version"] = version
	err = filepath_plus.WriteFileAsJsonBeauty(path, pkgJsonMap, true)
	if err != nil {
		return fmt.Errorf("ReplaceVersion filepath_plus.WriteFileAsJsonBeauty error %s", err)
	}

	return nil
}
