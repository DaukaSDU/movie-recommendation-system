package main

import (
	"github.com/DaukaSDU/movie-app"
	"github.com/DaukaSDU/movie-app/pkg/handler"
	"github.com/DaukaSDU/movie-app/pkg/repository"
	"github.com/DaukaSDU/movie-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	// Set up CORS middleware with appropriate options
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},  // Allow frontend domain
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"}, // Allow specific methods
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Set log formatter to JSON for structured logs
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Load configuration files
	if err := initConfigs(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// Initialize PostgreSQL database connection
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("error initializing DB: %s", err.Error())
	}

	// Set up repositories, services, and handlers
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// Initialize the server
	srv := new(movie.Server)

	// Wrap the handlers with CORS middleware
	corsHandler := corsOptions.Handler(handlers.InitRoutes())

	// Run the server with the CORS-enabled routes
	if err := srv.Run(viper.GetString("port"), corsHandler); err != nil {
		logrus.Fatalf("error occurred while running HTTP server: %s", err.Error())
	}
}

// initConfigs initializes the configuration for the application
func initConfigs() error {
	viper.AddConfigPath("configs") // Path to configuration files
	viper.SetConfigName("config")  // Config file name (config.yaml or config.json)
	return viper.ReadInConfig()    // Read configuration file
}
