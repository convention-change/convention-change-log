package changelog

import (
	"fmt"
	"strings"
	"time"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/sinlov-go/convention-change-log/convention"
	"github.com/sinlov-go/go-common-lib/pkg/date"
	"github.com/sinlov-go/sample-markdown/sample_mk"
)

const (
	firstLevel  = 1
	secondLevel = 2
	thirdLevel  = 3
)

func GenerateMarkdownNodes(
	commits []convention.Commit,
	logSpec convention.ConventionalChangeLogSpec,
	changelogDesc ConventionalChangeLogDesc,
) ([]sample_mk.Node, error) {

	if changelogDesc.Version == "" {
		return nil, fmt.Errorf("changelogDesc.Version can not be empty")
	}

	if changelogDesc.When.IsZero() {
		return nil, fmt.Errorf("changelogDesc.When can not be Zero")
	}

	if changelogDesc.ToolsKitName == "" {
		return nil, fmt.Errorf("changelogDesc.ToolsKitName can not be empty")
	}
	if changelogDesc.ToolsKitURL == "" {
		return nil, fmt.Errorf("changelogDesc.ToolsKitURL can not be empty")
	}
	if len(commits) == 0 {
		return nil, fmt.Errorf("commits can not be empty")
	}

	if changelogDesc.Location == nil {
		changelogDesc.Location = time.Local
	}

	filteredCommits := filter(commits, logSpec)

	sortedCommits, errSort := SortCommitsByLogSpec(filteredCommits, logSpec)
	if errSort != nil {
		return nil, errSort
	}

	nodesLen := 0
	markDownNodes := orderedmap.NewOrderedMap[string, []sample_mk.Node]()
	//markDownNodes := make(map[string][]sample_mk.Node, len(filteredCommits))
	if sortedCommits.Len() > 0 {
		//var markDownNodes map[string][]sample_mk.Node
		for el := sortedCommits.Front(); el != nil; el = el.Next() {
			markdownNodes := convertToListMarkdownNodes(el.Value)
			markDownNodes.Set(el.Key, markdownNodes)
			nodesLen += len(markdownNodes)
		}
	}

	// type header + 2 version header
	nodes := make([]sample_mk.Node, 0, nodesLen+3)

	// Adding each type
	for el := markDownNodes.Front(); el != nil; el = el.Next() {
		//fmt.Println(el.Key, el.Value)
		sectionFromType := convention.ParseSectionFromType(logSpec, el.Key)
		nodes = append(nodes, sample_mk.NewHeader(thirdLevel, sectionFromType))
		nodes = append(nodes, el.Value...)
	}

	// Adding title
	versionHeader := generateVersionHeaderValue(logSpec, changelogDesc)
	nodes = append([]sample_mk.Node{
		sample_mk.NewHeader(firstLevel, title),
		sample_mk.NewBasicItem(fmt.Sprintf(titleDesc, changelogDesc.ToolsKitName, changelogDesc.ToolsKitURL)),
		sample_mk.NewHeader(secondLevel, versionHeader),
	}, nodes...)

	return nodes, nil
}

func convertToListMarkdownNodes(commits []convention.Commit) []sample_mk.Node {
	result := make([]sample_mk.Node, 0, len(commits))

	for _, commit := range commits {
		result = append(result, sample_mk.NewListItem(commit.String()))
	}

	return result
}

func generateVersionHeaderValue(
	spec convention.ConventionalChangeLogSpec,
	changelogDesc ConventionalChangeLogDesc,
) string {
	versionTxt := changelogDesc.Version
	if spec.TagPrefix != "" {
		versionTxt = strings.Replace(changelogDesc.Version, spec.TagPrefix, "", 1)
	}
	if changelogDesc.VersionNotesUrl == "" {
		return fmt.Sprintf("%s (%s)", versionTxt, date.FormatDateByDefault(changelogDesc.When, changelogDesc.Location))
	}
	return fmt.Sprintf("[%s](%s) (%s)", versionTxt, changelogDesc.VersionNotesUrl, date.FormatDateByDefault(changelogDesc.When, changelogDesc.Location))
}
