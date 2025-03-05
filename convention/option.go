package convention

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sinlov-go/go-common-lib/pkg/date"
	"github.com/sinlov-go/go-git-tools/git"
)

const (
	leftScope                              = "("
	rightScope                             = ")"
	MarkdownBreakingChangesToken           = "BREAKING CHANGE: "
	MarkdownBreakingChangesSynonymousToken = "BREAKING-CHANGE: "
)

var (
	ErrEmptyCommit = errors.New("empty commit")

	headerRegex = regexp.MustCompile(`(?P<type>[a-zA-Z\-._]+)(?P<scope>\([a-zA-Z\-._]+\))?(?P<attention>!)?:\s(?P<description>.+)`)
)

type OptionFn func(*Commit) error

func GetRawHeader(gitCommit git.Commit) OptionFn {
	return func(c *Commit) error {
		if gitCommit.Message == "" {
			return ErrEmptyCommit
		}

		message := strings.TrimSpace(gitCommit.Message)
		messages := strings.Split(message, "\n")

		c.RawHeader = messages[0]

		return nil
	}
}

func GetTypeAndScope(gitCommit git.Commit) OptionFn {
	return func(c *Commit) error {
		if !headerRegex.MatchString(c.RawHeader) {
			c.Type = MiscType

			return nil
		}

		headerSubMatches := headerRegex.FindStringSubmatch(c.RawHeader)
		c.Type = strings.ToLower(headerSubMatches[1])
		c.Scope = strings.ToLower(headerSubMatches[2])
		c.Scope = strings.TrimLeft(c.Scope, leftScope)
		c.Scope = strings.TrimRight(c.Scope, rightScope)

		return nil
	}
}

func AddAuthorDate(gitCommit git.Commit) OptionFn {
	return func(c *Commit) error {
		c.RawHeader = fmt.Sprintf("%s (%s)", c.RawHeader, date.FormatDateByDefault(gitCommit.Author.When, time.Local))

		return nil
	}
}

func GetBreakChangesAndIssue(gitCommit git.Commit, spec ConventionalChangeLogSpec) OptionFn {
	return func(c *Commit) error {
		if gitCommit.Message == "" {
			return nil
		}
		message := strings.TrimSpace(gitCommit.Message)
		messages := strings.Split(message, "\n")
		breakingChangesDesc := ""
		issuePrefix := ""
		issueReference := ""
		var issueNum uint64
		for _, line := range messages {
			if strings.Index(line, MarkdownBreakingChangesToken) == 0 {
				breakingChangesDesc = strings.Replace(line, MarkdownBreakingChangesToken, "", 1)
				continue
			}
			if strings.Index(line, MarkdownBreakingChangesSynonymousToken) == 0 {
				breakingChangesDesc = strings.Replace(line, MarkdownBreakingChangesSynonymousToken, "", 1)
				continue
			}
			lineSplitSpace := strings.SplitN(line, " ", 2)
			if len(lineSplitSpace) > 1 {
				issueNumCheck := lineSplitSpace[1]
				if len(spec.IssuePrefixes) > 0 {
					for _, prefix := range spec.IssuePrefixes {
						if strings.Index(issueNumCheck, prefix) == 0 {
							issueNumStr := strings.Replace(issueNumCheck, prefix, "", 1)
							num, err := strconv.Atoi(issueNumStr)
							if err != nil {
								continue
							}
							issueReference = lineSplitSpace[0]
							issuePrefix = prefix
							issueNum = uint64(num)
							break
						}
					}
				}
			}
		}
		if breakingChangesDesc != "" {
			bc := BreakingChanges{
				Describe: breakingChangesDesc,
			}
			c.BreakingChanges = bc
		}
		if issueNum > 0 {
			ii := IssueInfo{
				IssueReference:    issueReference,
				IssuePrefix:       issuePrefix,
				IssueReferencesId: issueNum,
			}
			c.IssueInfo = ii
		}

		return nil
	}
}
