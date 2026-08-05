package command

import (
	"fmt"

	"github.com/bar-counter/slog"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/sinlov-go/go-git-tools/git"
)

// DefaultBranchCheckPatterns returns the default branch check patterns.
func DefaultBranchCheckPatterns() []string {
	return []string{"main"}
}

// ResolveBranchCheckPatterns returns patterns if non-empty, otherwise returns default patterns.
func ResolveBranchCheckPatterns(patterns []string) []string {
	if len(patterns) == 0 {
		return DefaultBranchCheckPatterns()
	}
	return patterns
}

// MatchBranchPatterns checks if branchName matches any of the patterns using doublestar glob matching.
func MatchBranchPatterns(branchName string, patterns []string) bool {
	for _, pattern := range patterns {
		ok, err := doublestar.Match(pattern, branchName)
		if err != nil {
			slog.Warnf("branch check pattern %q error: %v", pattern, err)
			continue
		}
		if ok {
			return true
		}
	}
	return false
}

// CheckBranchAtRoot opens the git repo at rootPath and checks if the current branch matches patterns.
// Returns (branchName, matched, error).
func CheckBranchAtRoot(rootPath string, patterns []string) (string, bool, error) {
	repo, err := git.NewRepositoryRemoteByPath("origin", rootPath)
	if err != nil {
		return "", false, fmt.Errorf("load git repository for branch check error: %s", err)
	}
	branchName, err := repo.HeadBranchName()
	if err != nil {
		return "", false, fmt.Errorf("get HEAD branch name error: %s", err)
	}
	return branchName, MatchBranchPatterns(branchName, patterns), nil
}
