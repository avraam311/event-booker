package events

import (
	"context"
	"strconv"
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
		now := time.Now().UTC()
		if b.Book == "not confirmed" && (cancelTime.Equal(now) || cancelTime.Before(now)) {
			err := s.repo.DeleteBook(ctx, b.ID, b.EventID)
			if err != nil {
				zlog.Logger.Warn().Err(err).Str("book id", strconv.Itoa(int(b.ID))).Msg("failed to delete book")
				continue
			}
			zlog.Logger.Info().Str("book id", strconv.Itoa(int(b.ID))).Msg("book deleted")
		}
	}
}
