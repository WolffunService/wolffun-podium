package lib

import "context"

// PodiumInterface defines the interface to be implemented
type PodiumInterface interface {
	GetTop(ctx context.Context, leaderboard string, page, pageSize int) (*MemberList, error)
	GetTopPercent(ctx context.Context, leaderboard string, percentage int) (*MemberList, error)
	UpdateScore(ctx context.Context, leaderboard, memberID string, score, scoreTTL int) (*MemberList, error)
	IncrementScore(ctx context.Context, leaderboard, memberID string, increment, scoreTTL int) (*MemberList, error)
	UpdateScores(ctx context.Context, leaderboards []string, memberID string, score, scoreTTL int) (*ScoreList, error)
	RemoveMemberFromLeaderboard(ctx context.Context, leaderboard, member string) (*Response, error)
	GetMember(ctx context.Context, leaderboard, memberID string) (*Member, error)
	GetMembers(ctx context.Context, leaderboard string, memberIDs []string) (*MemberList, error)
	Healthcheck(ctx context.Context) (string, error)
	DeleteLeaderboard(ctx context.Context, leaderboard string) (*Response, error)
}
