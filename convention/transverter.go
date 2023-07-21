package convention

import (
	"fmt"
	"github.com/sinlov-go/go-git-tools/git"
)

func ConvertGitCommits2ConventionCommits(commits []git.Commit, spec ConventionalChangeLogSpec, gitHttpInfo GitRepositoryHttpInfo) ([]Commit, error) {
	if len(commits) == 0 {
		return nil, fmt.Errorf("commits is empty")
	}

	var result []Commit
	for _, c := range commits {
		commit, err := NewCommitWithLogSpec(c, spec, gitHttpInfo)
		if err != nil {
			return nil, err
		}
		result = append(result, commit)
	}
	return result, nil
}
