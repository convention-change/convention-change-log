package changelog

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/convention-change/convention-change-log/convention"
	"github.com/elliotchance/orderedmap/v2"
	"github.com/sinlov-go/go-common-lib/pkg/date"
	"github.com/sinlov-go/sample-markdown/sample_mk"
)

const (
	firstLevel  = 1
	secondLevel = 2
	thirdLevel  = 3
)

var (
	NotErrCommitsLenZero = fmt.Errorf("commits len is zero")
)

// GenerateMarkdownNodes
//
// gitHttpInfo: git repo info by convention.GitRepositoryHttpInfo
//
// changelogDesc: changelog desc by ConventionalChangeLogDesc
//
// commit: commit list by []convention.Commit
//
// logSpec: log spec by convention.ConventionalChangeLogSpec
func GenerateMarkdownNodes(
	gitHttpInfo convention.GitRepositoryHttpInfo,
	changelogDesc ConventionalChangeLogDesc,
	commits []convention.Commit,
	logSpec convention.ConventionalChangeLogSpec,
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
		// type header + 2 version header
		nodes := make([]sample_mk.Node, 0, 3)

		// Adding title
		versionHeader := generateVersionHeaderValue(gitHttpInfo, logSpec, changelogDesc)
		nodes = append([]sample_mk.Node{
			sample_mk.NewHeader(firstLevel, logSpec.Header),
			sample_mk.NewBasicItem(fmt.Sprintf(titleDesc, changelogDesc.ToolsKitName, changelogDesc.ToolsKitURL)),
			sample_mk.NewHeader(secondLevel, versionHeader),
		}, nodes...)

		return nil, NotErrCommitsLenZero
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

	breakingChanges := make([]convention.BreakingChanges, 0)
	if sortedCommits.Len() > 0 {
		//var markDownNodes map[string][]sample_mk.Node
		for el := sortedCommits.Front(); el != nil; el = el.Next() {
			elCommits := el.Value

			if len(elCommits) > 0 {
				for _, commit := range elCommits {
					if commit.BreakingChanges.Describe != "" {
						breakingChanges = append(breakingChanges, commit.BreakingChanges)
					}
				}
			}

			markdownNodes := convertToListMarkdownNodes(elCommits, gitHttpInfo, logSpec)
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

	// add BREAKING CHANGE:
	if len(breakingChanges) > 0 {
		bkNode := []sample_mk.Node{
			sample_mk.NewHeader(thirdLevel, convention.MarkdownBreakingChangesToken),
		}
		for _, breakingChange := range breakingChanges {
			bkNode = append(bkNode, sample_mk.NewListItem(breakingChange.Describe))
		}
		nodes = append(bkNode, nodes...)
	}

	// Adding title
	versionHeader := generateVersionHeaderValue(gitHttpInfo, logSpec, changelogDesc)
	nodes = append([]sample_mk.Node{
		sample_mk.NewHeader(firstLevel, logSpec.Header),
		sample_mk.NewBasicItem(fmt.Sprintf(titleDesc, changelogDesc.ToolsKitName, changelogDesc.ToolsKitURL)),
		sample_mk.NewHeader(secondLevel, versionHeader),
	}, nodes...)

	return nodes, nil
}

func convertToListMarkdownNodes(commits []convention.Commit, gitHttpInfo convention.GitRepositoryHttpInfo, spec convention.ConventionalChangeLogSpec) []sample_mk.Node {
	result := make([]sample_mk.Node, 0, len(commits))

	for _, commit := range commits {
		commitRaw := commit.String()
		if commit.IssueInfo.IssueReferencesId > 0 {
			issueInfo := commit.IssueInfo
			if gitHttpInfo.Host != "" {
				issueTemplate := new(convention.IssueRenderTemplate)
				issueTemplate.Scheme = gitHttpInfo.Scheme
				issueTemplate.Host = gitHttpInfo.Host
				issueTemplate.Owner = gitHttpInfo.Owner
				issueTemplate.Repository = gitHttpInfo.Repository
				issueTemplate.Id = strconv.FormatUint(issueInfo.IssueReferencesId, 10)
				render, err := convention.RaymondRender(spec.IssueUrlFormat, issueTemplate)
				if err != nil {
					fmt.Printf("convertToListMarkdownNodes spec.IssueUrlFormat %s err: %v\n", spec.IssueUrlFormat, err)
				} else {
					commitRaw = fmt.Sprintf("%s, %s [%s%s](%s)",
						commitRaw, issueInfo.IssueReference, issueInfo.IssuePrefix, issueTemplate.Id, render)
				}
			} else {
				commitRaw = fmt.Sprintf("%s, %s %s%d",
					commitRaw, issueInfo.IssueReference, issueInfo.IssuePrefix, issueInfo.IssueReferencesId)
			}
		}
		result = append(result, sample_mk.NewListItem(commitRaw))
	}

	return result
}

// generateVersionHeaderValue
// if generate compareUrl error will return sample
func generateVersionHeaderValue(
	gitRepoInfo convention.GitRepositoryHttpInfo,
	spec convention.ConventionalChangeLogSpec,
	changelogDesc ConventionalChangeLogDesc,
) string {
	versionTxt := changelogDesc.Version
	if spec.TagPrefix != "" {
		versionTxt = strings.Replace(changelogDesc.Version, spec.TagPrefix, "", 1)
	}

	if changelogDesc.PreviousTag == "" {
		return fmt.Sprintf("%s (%s)", versionTxt, date.FormatDateByDefault(changelogDesc.When, changelogDesc.Location))
	}

	compareRender := new(convention.CompareRenderTemplate)
	compareRender.Scheme = gitRepoInfo.Scheme
	compareRender.Host = gitRepoInfo.Host
	compareRender.Owner = gitRepoInfo.Owner
	compareRender.Repository = gitRepoInfo.Repository
	compareRender.PreviousTag = changelogDesc.PreviousTag
	compareRender.CurrentTag = changelogDesc.Version

	compareUrl, err := convention.RaymondRender(spec.CompareUrlFormat, compareRender)
	if err != nil {
		return fmt.Sprintf("%s (%s)", versionTxt, date.FormatDateByDefault(changelogDesc.When, changelogDesc.Location))
	}
	return fmt.Sprintf("[%s](%s) (%s)", versionTxt, compareUrl, date.FormatDateByDefault(changelogDesc.When, changelogDesc.Location))
}
