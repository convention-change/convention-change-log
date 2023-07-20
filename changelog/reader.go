package changelog

import (
	"fmt"
	"github.com/convention-change/convention-change-log/convention"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"github.com/sinlov-go/sample-markdown/sample_mk"
	"regexp"
	"strings"
)

const (
	regHistoryMarkdownTagHeadLine     = "^(#+)\\W(.*)\\W(\\(.*\\))$"
	regHistoryMarkdownTagLinkHeadLine = "^(#+)\\W(\\[.*\\])(\\()(.*)(\\))\\W(\\()(.*)(\\))$"
	regHistoryMarkdownTagTitle        = "^(#+)\\W(.*)$"
)

// NewReader
// path: string reader load ChangeLog.md path
//
// spec: convention.ConventionalChangeLogSpec please use
// convention.DefaultConventionalChangeLogSpec or convention.LoadConventionalChangeLogSpecByData
func NewReader(path string, spec convention.ConventionalChangeLogSpec) (Reader, error) {

	reader := changeLog{
		path: path,
		spec: spec,
	}
	err := parse(&reader)
	if err != nil {
		return &reader, err
	}
	return &reader, nil
}

type changeLog struct {
	path string

	spec convention.ConventionalChangeLogSpec

	titleNodes            []sample_mk.Node
	historyFirstTagShort  string
	historyFirstTag       string
	historyFirstTitle     string
	historyFirstNodes     []sample_mk.Node
	historyFirstContent   string
	historyFirstChangeUrl string

	historyNodes []sample_mk.Node
}

// HistoryFirstTagShort
// return history first tag short not include convention.ConventionalChangeLogSpec TagPrefix
func (c *changeLog) HistoryFirstTagShort() string {
	return c.historyFirstTagShort
}

// HistoryFirstTag
// return history first tag this will append convention.ConventionalChangeLogSpec TagPrefix
func (c *changeLog) HistoryFirstTag() string {
	return c.historyFirstTag
}

// HistoryFirstTitle
// return history first title, title not contain ## or ###
func (c *changeLog) HistoryFirstTitle() string {
	return c.historyFirstTitle
}

// HistoryFirstContent
// return history first content without title
func (c *changeLog) HistoryFirstContent() string {
	return c.historyFirstContent
}

// HistoryFirstChangeUrl
// return history first change url like https://github.com/convention-change/convention-change-log/compare/v1.0.0...v1.1.0
func (c *changeLog) HistoryFirstChangeUrl() string {
	return c.historyFirstChangeUrl
}

// HistoryFirstNodes
// return history first nodes contains title
func (c *changeLog) HistoryFirstNodes() []sample_mk.Node {
	return c.historyFirstNodes
}

// HistoryNodes
// full history Nodes
func (c *changeLog) HistoryNodes() []sample_mk.Node {
	return c.historyNodes
}

func parse(changeLog *changeLog) error {
	readLine, err := filepath_plus.ReadFileAsLines(changeLog.path)
	if err != nil {
		return err
	}

	nodes := sample_mk.Parse(readLine)

	if len(nodes) == 0 {
		return fmt.Errorf("can not find any sample markdown node at path: %s", changeLog.path)
	}

	compileMarkdownTagHeadLine, _ := regexp.Compile(regHistoryMarkdownTagHeadLine)
	compileMarkdownTagLinkHeadLine, _ := regexp.Compile(regHistoryMarkdownTagLinkHeadLine)
	compileMarkdownTagTitle, _ := regexp.Compile(regHistoryMarkdownTagTitle)

	nodeStartIndex := 0
	searchCnt := 2
	firstTagShort := ""
	firstTitle := ""
	firstChangeUrl := ""
	var firstTitleNode sample_mk.Node
	firstNodes := []sample_mk.Node{}
	for i, node := range nodes {
		if searchCnt == 0 {
			break
		}
		if node.Type() == sample_mk.NodeTypeHeader {
			header := node.(sample_mk.Header)
			firstHistoryTitleStr := header.String()
			if compileMarkdownTagHeadLine.MatchString(firstHistoryTitleStr) {
				if searchCnt == 2 {
					findStringSub := compileMarkdownTagHeadLine.FindStringSubmatch(firstHistoryTitleStr)
					if len(findStringSub) > 3 {
						firstHistoryVersionStr := findStringSub[2]
						if strings.Index(firstHistoryVersionStr, "[") == 0 {
							endIndex := strings.Index(firstHistoryVersionStr, "]")
							firstTagShort = firstHistoryVersionStr[1:endIndex]

							markdownLinkSub := compileMarkdownTagLinkHeadLine.FindStringSubmatch(firstHistoryTitleStr)
							if len(markdownLinkSub) > 8 {
								firstChangeUrl = markdownLinkSub[4]
							}
						} else {
							firstTagShort = firstHistoryVersionStr
						}
						titleSub := compileMarkdownTagTitle.FindStringSubmatch(firstHistoryTitleStr)
						firstTitle = titleSub[2]
						firstTitleNode = node
						nodeStartIndex = i
					}
				}
				searchCnt--
				continue
			}
		}
		if searchCnt == 1 {
			firstNodes = append(firstNodes, node)
		}
	}
	changeLog.titleNodes = nodes[:nodeStartIndex]

	changeLog.historyFirstTagShort = firstTagShort
	changeLog.historyFirstTag = fmt.Sprintf("%s%s", changeLog.spec.TagPrefix, firstTagShort)
	changeLog.historyFirstTitle = firstTitle
	historyFirstContent := sample_mk.GenerateText(firstNodes)
	changeLog.historyFirstContent = historyFirstContent
	if firstTitleNode != nil {
		changeLog.historyFirstNodes = append([]sample_mk.Node{
			firstTitleNode,
		}, firstNodes...)
	} else {
		changeLog.historyFirstNodes = firstNodes
	}
	changeLog.historyFirstChangeUrl = firstChangeUrl

	changeLog.historyNodes = nodes[nodeStartIndex:]

	return nil
}

type Reader interface {
	// HistoryFirstTagShort
	// return history first tag short not include convention.ConventionalChangeLogSpec TagPrefix
	HistoryFirstTagShort() string

	// HistoryFirstTag
	// return history first tag this will append convention.ConventionalChangeLogSpec TagPrefix
	HistoryFirstTag() string

	// HistoryFirstTitle
	// return history first title, title not contain ## or ###
	HistoryFirstTitle() string

	// HistoryFirstContent
	// return history first content without title
	HistoryFirstContent() string

	// HistoryFirstChangeUrl
	// return history first change url like https://github.com/convention-change/convention-change-log/compare/v1.0.0...v1.1.0
	HistoryFirstChangeUrl() string

	// HistoryFirstNodes
	// return history first nodes contains title
	HistoryFirstNodes() []sample_mk.Node

	// HistoryNodes
	// full history Nodes
	HistoryNodes() []sample_mk.Node
}
