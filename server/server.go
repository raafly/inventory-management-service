package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/raafly/invetory-management/config"
	"github.com/raafly/invetory-management/db"
	"github.com/raafly/invetory-management/listing"
)

type Server struct {
	App *fiber.App
	conf *config.AppConfig
}

func NewServer() *Server {
	app := fiber.New()
	conf, err := config.NewAppConfig()
	if err != nil {
		log.Fatalf("failed initialize config %s", err)
	}

	return &Server{
		App: app,
		conf: conf,
	}
}

func (s *Server) Run() {
	db := db.NewDB(s.conf)


	listing.NewAuthRoutes(s.App, db)
	listing.NewCategoriesRoutes(s.App, db)
	listing.NewItemRoutes(s.App, db)
	listing.NewHistoryRoutes(s.App, db)

	s.GracefulShutdown()
}

func (s *Server) GracefulShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func ()  {
		if err := s.App.Listen(":" + s.conf.Fiber.Port); err != nil {
			log.Fatalf("error when listening to :%s, %s", s.conf.Fiber.Port, err)
		}
	} ()

	log.Printf("server is running on :%s", s.conf.Fiber.Port)

	<-stop

	log.Println("server gracefully shutdown")

	if err := s.App.Shutdown(); err != nil {
		log.Fatalf("error when shutting down the server, %s", err)
	}

	log.Println("process clean up...")
}