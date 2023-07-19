package changelog

import (
	"fmt"
	"github.com/convention-change/convention-change-log/convention"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"github.com/sinlov-go/sample-markdown/sample_mk"
	"strings"
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

	titleNodes           []sample_mk.Node
	historyFirstTagShort string
	historyFirstTitle    string
	historyFirstNodes    []sample_mk.Node
	historyFirstContent  string

	historyNodes []sample_mk.Node
}

func (c *changeLog) HistoryFirstTagShort() string {
	return c.historyFirstTagShort
}

func (c *changeLog) HistoryFirstTitle() string {
	return c.historyFirstTitle
}

func (c *changeLog) HistoryFirstNodes() []sample_mk.Node {
	return c.historyFirstNodes
}

func (c *changeLog) HistoryFirstContent() string {
	return c.historyFirstContent
}

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

	nodeStartIndex := 0
	searchCnt := 2
	firstTag := ""
	firstTitle := ""
	var firstTitleNode sample_mk.Node
	firstNodes := []sample_mk.Node{}
	for i, node := range nodes {
		if searchCnt == 0 {
			break
		}
		if node.Type() == sample_mk.NodeTypeHeader {
			header := node.(sample_mk.Header)
			if header.Level() == 2 {
				if searchCnt == 2 {
					nodeStartIndex = i
					firstHistoryTagStr := header.String()
					firstHistoryTagStr = strings.Replace(firstHistoryTagStr, "## ", "", 1)
					if strings.Index(firstHistoryTagStr, "[") == 0 {
						endIndex := strings.Index(firstHistoryTagStr, "]")
						firstTag = firstHistoryTagStr[1:endIndex]
					} else {
						spTagStr := strings.Split(firstHistoryTagStr, " ")
						if len(spTagStr) > 1 {
							firstTag = spTagStr[0]
						}
					}
					firstTitle = firstHistoryTagStr
					firstTitleNode = node
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

	changeLog.historyFirstTagShort = firstTag
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

	changeLog.historyNodes = nodes[nodeStartIndex:]

	return nil
}

type Reader interface {
	// HistoryFirstTagShort
	// return history first tag short not include convention.ConventionalChangeLogSpec TagPrefix
	HistoryFirstTagShort() string

	// HistoryFirstTitle
	// return history first title
	HistoryFirstTitle() string

	// HistoryFirstContent
	// return history first content without title
	HistoryFirstContent() string

	// HistoryFirstNodes
	// return history first nodes contains title
	HistoryFirstNodes() []sample_mk.Node

	// HistoryNodes
	// full history Nodes
	HistoryNodes() []sample_mk.Node
}
