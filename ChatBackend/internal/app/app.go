package app

import (
	"chat/internal/config"
	"chat/internal/handler"
	"chat/internal/lib/slogx"
	"chat/internal/repository"
	"chat/internal/service"
	"log/slog"
	"os"
	"time"

	"github.com/gin-contrib/cors"
)

type App struct {
	cfg *config.Config
	log *slog.Logger
}

func NewApp(cfg *config.Config, log *slog.Logger) *App {
	return &App{cfg: cfg, log: log}
}

func (a *App) MustRun() {
	// Инициализация базы данных
	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     a.cfg.Host,
		Port:     a.cfg.Port,
		Username: a.cfg.Username,
		DBName:   a.cfg.DBName,
		SSLMode:  a.cfg.SSLMode,
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		a.log.Error("Failed to initialize database", slogx.Error(err))
		os.Exit(1)
	}

	// Инициализация репозитория
	authPostgres := repository.NewAuthPostgres(db)
	controlPostgres := repository.NewControlPostgres(db)
	chatPostgres := repository.NewChatPostgres(db)
	audioPostgres := repository.NewAudioPostgres(db)
	videoPostgres := repository.NewVideoPostgres(db)

	authService := service.NewAuthService(authPostgres)
	middlewareService := service.NewMiddlewareService()
	controlService := service.NewControlService(controlPostgres)
	chatService := service.NewChatService(chatPostgres, a.cfg.LLM)
	audioService := service.NewAudioService(audioPostgres)
	videoService := service.NewVideoService(videoPostgres)

	// Инициализация обработчиков
	authHandler := handler.NewAuthHandler(authService, a.log)
	middlewareHandler := handler.NewMiddlewareHandler(middlewareService, a.log)
	controlHandler := handler.NewControlHandler(controlService, a.log)
	chatHandler := handler.NewChatHandler(chatService, a.log)
	audioHandler := handler.NewAudioHandler(audioService, a.log)
	videoHandler := handler.NewVideoHandler(videoService, a.log)

	// Инициализация роутера
	handler := NewHandler(
		middlewareHandler,
		authHandler,
		controlHandler,
		chatHandler,
		audioHandler,
		videoHandler,
	)
	router := handler.InitRoutes()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "HEAD"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	srv := new(Server)
	a.log.Info("Run HTTP server", slog.String("address", a.cfg.Address))
	if err = srv.Run(a.cfg.Address, router); err != nil {
		a.log.Error("Failed to run HTTP server", slogx.Error(err))
		os.Exit(1)
	}
}
