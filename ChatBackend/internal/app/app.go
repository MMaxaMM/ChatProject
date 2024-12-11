package app

import (
	"chat/internal/config"
	"chat/internal/handler"
	"chat/internal/lib/slogx"
	minioclient "chat/internal/minio-client"
	"chat/internal/repository"
	"chat/internal/service"
	"log/slog"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
)

type App struct {
	cfg *config.Config
	log *slog.Logger
}

func NewApp(cfg *config.Config, log *slog.Logger) *App {
	return &App{cfg: cfg, log: log}
}

func (a *App) MustRun() {
	// Подготовка окружения
	tempDir := os.TempDir()
	if err := os.MkdirAll(tempDir, 1777); err != nil {
		a.log.Error("Failed to create temporary directory", slogx.Error(err))
		os.Exit(1)
	}

	// Инициализация базы данных
	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     a.cfg.Database.Host,
		Port:     a.cfg.Database.Port,
		Username: a.cfg.Database.Username,
		DBName:   a.cfg.Database.DBName,
		SSLMode:  a.cfg.Database.SSLMode,
		Password: a.cfg.Database.Password,
	})
	if err != nil {
		a.log.Error("Failed to initialize database", slogx.Error(err))
		os.Exit(1)
	}

	// Инициализация репозитория
	rep := repository.NewPostgresRepository(db, a.cfg.Filestorage)

	// Инициализация сервисов
	if err = minioclient.Connect(a.cfg.Minio); err != nil {
		a.log.Error("Failed to connect to minio", slogx.Error(err))
		os.Exit(1)
	}
	service := service.NewService(a.cfg, rep)

	// Инициализация обработчиков
	handler := handler.NewHandler(service, a.log)
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

	srv := NewServer(a.cfg.HTTPServer)
	a.log.Info("Run HTTP server", slog.String("address", a.cfg.HTTPServer.Address))
	if err = srv.Run(router); err != nil {
		a.log.Error("Failed to run HTTP server", slogx.Error(err))
		os.Exit(1)
	}
}
