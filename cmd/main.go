package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	file = flag.String("file", "", "file to read dsn secret from (overrides -dsn)")
	dsn  = flag.String("dsn", "", "dsn secret raw string")
	dir  = flag.String("dir", "", "migration directory")
)

func init() {
	flag.Parse()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := validate()
	if err != nil {
		log.Fatalf("validation failed: %v", err)
	}

	dsn, err := getDSN()
	if err != nil {
		log.Fatalf("failed to get dsn: %v", err)
	}

	if err := run(ctx, dsn, *dir); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("done")
}

func validate() error {
	var errors []string
	if *file == "" && *dsn == "" {
		errors = append(errors, "file or dsn must be set")
	}
	if *dir == "" {
		errors = append(errors, "dir must be set")
	}

	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, ", "))
	}

	return nil
}

func getDSN() (string, error) {
	if *file != "" {
		content, err := os.ReadFile(*file)
		if err != nil {
			return "", fmt.Errorf("failed to read file: %v", err)
		}
		return string(content), nil
	}

	return *dsn, nil
}

func run(ctx context.Context, dsn string, dir string) error {
	// Create a new migration instance
	m, err := migrate.New(fmt.Sprintf("file://%s", dir), dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
