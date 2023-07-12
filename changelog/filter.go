package changelog

import (
	"container/list"
	"fmt"
	"github.com/elliotchance/orderedmap/v2"
	"github.com/sinlov-go/convention-change-log/convention"
	"sync"
)

var filterLock = &sync.RWMutex{}

func filter(commits []convention.Commit, logSpec convention.ConventionalChangeLogSpec) map[string][]convention.Commit {
	if len(commits) == 0 {
		return nil
	}
	scopes := logSpec.Types
	if len(scopes) == 0 {
		return nil
	}

	filteredCommits := make(map[string][]convention.Commit)
	for _, scope := range scopes {
		if scope.Hidden {
			continue
		}
		filteredCommits[scope.Type] = make([]convention.Commit, 0, defaultLen)
	}

	for _, commit := range commits {

		for _, scope := range scopes {
			if scope.Hidden {
				continue
			}
			if commit.Type == scope.Type {
				filteredCommits[scope.Type] = append(filteredCommits[scope.Type], commit)
			}
		}
	}

	// remove len 0
	for k, v := range filteredCommits {
		if len(v) == 0 {
			filterLock.Lock()
			delete(filteredCommits, k)
			filterLock.Unlock()
		}
	}

	return filteredCommits
}

func SortCommitsByLogSpec(commits map[string][]convention.Commit, logSpec convention.ConventionalChangeLogSpec) (*orderedmap.OrderedMap[string, []convention.Commit], error) {
	if len(commits) == 0 {
		return nil, fmt.Errorf("commits can not be empty")
	}

	if len(logSpec.Types) == 0 {
		return nil, fmt.Errorf("logSpec.Types can not be empty")
	}

	// Sort
	keySort := list.New()
	keyEl := 0
	for _, typeDef := range logSpec.Types {
		if typeDef.Hidden {
			continue
		}
		if typeDef.Sort <= keyEl {
			keySort.PushFront(typeDef.Type)
		} else {
			keyEl = typeDef.Sort
			keySort.PushBack(typeDef.Type)
		}
	}
	if keySort.Len() == 0 {
		return nil, fmt.Errorf("all ConventionalChangeLogSpec type is hidden, please check")
	}
	sortedCommits := orderedmap.NewOrderedMap[string, []convention.Commit]()
	for e := keySort.Front(); e != nil; e = e.Next() {
		k := e.Value.(string)
		sortCommit, ok := commits[k]
		if ok {
			sortedCommits.Set(k, sortCommit)
		}
	}
	return sortedCommits, nil
}
