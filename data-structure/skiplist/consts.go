package go_skipList

import "errors"

const (
	maxLevel  = 64
	skipListP = 0.25
)

var (
	errNoKey                = errors.New("not exist mem")
	errInvalidScoreMemArray = errors.New("invalid score member array")
	errInvalidRank          = errors.New("invalid rank param")
	errInvalidScoreType     = errors.New("score type only support int&float")
)
