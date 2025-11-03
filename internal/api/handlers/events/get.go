package events

import (
	// "errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/avraam311/event-booker/internal/api/handlers"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) GetEventInfo(c *ginext.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		zlog.Logger.Warn().Err(err).Msg("id is not proper unsigned integer or empty parameter")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("non-empty and proper id required"))
		return
	}
	id := uint(idInt)

	ev, err := h.service.GetEventInfo(c.Request.Context(), id)
	if err != nil {
		// if errors.Is(err, events.ErrEventNotFound) {
		// 	zlog.Logger.Warn().Err(err).Msg("event not found")
		// 	handlers.Fail(c.Writer, http.StatusNotFound, fmt.Errorf("event not found"))
		// 	return
		// }

		zlog.Logger.Error().Err(err).Msg("failed to get event")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.OK(c.Writer, ev)
}
