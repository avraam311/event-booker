package server

import (
	"net/http"

	
	"github.com/avraam311/event-booker/internal/api/handlers/events"
	"github.com/avraam311/event-booker/internal/middlewares"
	
	"github.com/wb-go/wbf/ginext"
)

func NewRouter(ginMode string, handlerEv *events.Handler) *ginext.Engine {
	e := ginext.New(ginMode)

	e.Use(middlewares.CORSMiddleware())
	e.Use(ginext.Logger())
	e.Use(ginext.Recovery())

	api := e.Group("/event-booker/api")
	{
		api.POST("/events", handlerEv.CreateEvent)
		api.POST("/events/:id/book", handlerEv.BookSeat)
		api.POST("/events/confirm/:id", handlerEv.Confirm)
		api.GET("/events/:id", handlerEv.GetEventInfo)
	}

	return e
}

func NewServer(addr string, router *ginext.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}