package events

import (
	"context"
	"time"

	"github.com/wb-go/wbf/zlog"
)

func (s *Service) CleanExpiredBooks(ctx context.Context) {
	books, err := s.repo.GetAllBooks(ctx)
	if err != nil {
		zlog.Logger.Warn().Err(err).Msg("failed to get all books")
	}

	for _, b := range books {
		cancelTime := b.CreatedAt.Add(time.Hour)
		now := time.Now()
		if cancelTime.Equal(now) || cancelTime.After(now) {
			err := s.repo.DeleteBook(ctx, b.ID)
			if err != nil {
				zlog.Logger.Warn().Err(err).Msg("failed to delete book")
			}
		}
	}
}
