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
		Path: path,
		spec: spec,
	}
	err := parse(&reader)
	if err != nil {
		return &reader, err
	}
	return &reader, nil
}

type changeLog struct {
	Path string

	spec convention.ConventionalChangeLogSpec

	titleNodes      []sample_mk.Node
	historyTagShort string
	historyNodes    []sample_mk.Node
}

func (c *changeLog) HistoryNodes() []sample_mk.Node {
	return c.historyNodes
}

func (c *changeLog) HistoryTagShort() string {
	return c.historyTagShort
}

func parse(changeLog *changeLog) error {
	readLine, err := filepath_plus.ReadFileAsLines(changeLog.Path)
	if err != nil {
		return err
	}

	nodes := sample_mk.Parse(readLine)

	if len(nodes) == 0 {
		return fmt.Errorf("can not find any sample markdown node by path: %s", changeLog.Path)
	}

	historyNodeStartIndex := 0
	historyTag := ""
	for i, node := range nodes {
		if node.Type() == sample_mk.NodeTypeHeader {
			header := node.(sample_mk.Header)
			if header.Level() > 1 {
				historyNodeStartIndex = i
				firstHistoryTagStr := header.String()
				firstHistoryTagStr = strings.Replace(firstHistoryTagStr, "## ", "", 1)
				if strings.Index(firstHistoryTagStr, "[") == 0 {
					endIndex := strings.Index(firstHistoryTagStr, "]")
					historyTag = firstHistoryTagStr[1:endIndex]
				} else {
					spTagStr := strings.Split(firstHistoryTagStr, " ")
					if len(spTagStr) > 1 {
						historyTag = spTagStr[0]
					}
				}
				break
			}
		}
	}
	changeLog.titleNodes = nodes[:historyNodeStartIndex]
	changeLog.historyTagShort = historyTag
	changeLog.historyNodes = nodes[historyNodeStartIndex:]

	return nil
}

type Reader interface {
	HistoryTagShort() string

	HistoryNodes() []sample_mk.Node
}
