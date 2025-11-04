package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/avraam311/event-booker/internal/api/handlers"
	"github.com/avraam311/event-booker/internal/models"
	"github.com/avraam311/event-booker/internal/repository/events"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) CreateEvent(c *ginext.Context) {
	var ev models.EventDTO

	if err := json.NewDecoder(c.Request.Body).Decode(&ev); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to decode request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	if err := h.validator.Struct(ev); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to validate request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("validation error: %s", err.Error()))
		return
	}

	id, err := h.service.CreateEvent(c.Request.Context(), &ev)
	if err != nil {
		zlog.Logger.Error().Err(err).Interface("event", ev).Msg("failed to create event")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.Created(c.Writer, id)
}

func (h *Handler) BookSeat(c *ginext.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to convert param id into int")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid id: %s", err.Error()))
		return
	}
	id := uint(idInt)

	var book models.BookDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&book); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to decode request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	if err := h.validator.Struct(book); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to validate request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("validation error: %s", err.Error()))
		return
	}

	bookID, err := h.service.BookSeat(c.Request.Context(), id, &book)
	if err != nil {
		if errors.Is(err, events.ErrEventNotFound) {
			zlog.Logger.Warn().Err(err).Msg("event not found")
			handlers.Fail(c.Writer, http.StatusNotFound, fmt.Errorf("event not found"))
			return
		} else if errors.Is(err, events.ErrNoSeatsOrEventNotFound) {
			zlog.Logger.Warn().Err(err).Msg("no seats left or event not found")
			handlers.Fail(c.Writer, http.StatusNotFound, fmt.Errorf("no seats left or event not found"))
			return
		}

		zlog.Logger.Error().Err(err).Interface("book", book).Msg("failed to book seat")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.Created(c.Writer, bookID)
}

func (h *Handler) Confirm(c *ginext.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to convert param id into int")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid id: %s", err.Error()))
		return
	}
	id := uint(idInt)

	err = h.service.Confirm(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, events.ErrBookNotFound) {
			zlog.Logger.Warn().Err(err).Msg("book not found")
			handlers.Fail(c.Writer, http.StatusNotFound, fmt.Errorf("book not found"))
			return
		}

		zlog.Logger.Error().Err(err).Interface("book", id).Msg("failed to confirm book")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.OK(c.Writer, "book confirmed")
}
