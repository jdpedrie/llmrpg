package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewNullString creates a new sql.NullString from a string
func NewNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// StringFromNullString safely extracts a string from a sql.NullString
func StringFromNullString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

// NewText creates a new pgtype.Text from a string
func NewText(s string) pgtype.Text {
	if s == "" {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{
		String: s,
		Valid:  true,
	}
}

// StringFromText safely extracts a string from a pgtype.Text
func StringFromText(s pgtype.Text) string {
	if s.Valid {
		return s.String
	}
	return ""
}

// NewNullUUID creates a new sql.NullString from a UUID
func NewNullUUID(id uuid.UUID) sql.NullString {
	if id == uuid.Nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: id.String(),
		Valid:  true,
	}
}

// UUIDFromNullString safely converts a sql.NullString to a UUID
func UUIDFromNullString(s sql.NullString) (uuid.UUID, error) {
	if !s.Valid {
		return uuid.Nil, nil
	}
	return uuid.Parse(s.String)
}

// TimeFromNullTime safely extracts a time.Time pointer from a sql.NullTime
func TimeFromNullTime(t sql.NullTime) *sql.NullTime {
	if !t.Valid {
		return nil
	}
	return &t
}

// TimeFromTimestamptz safely extracts a time.Time pointer from a pgtype.Timestamptz
func TimeFromTimestamptz(t pgtype.Timestamptz) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

// TimestamptzFromTime creates a pgtype.Timestamptz from a time.Time pointer
func TimestamptzFromTime(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{Valid: false}
	}
	return pgtype.Timestamptz{
		Time:  *t,
		Valid: true,
	}
}

// NewFromEnv creates a new PostgreSQL connection pool from environment variables
func NewFromEnv() (*pgxpool.Pool, error) {
	// Use standard PostgreSQL environment variables like PGHOST, PGPORT, etc.
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		var connVars []string
		
		// Build connection string from individual environment variables
		if host := os.Getenv("PGHOST"); host != "" {
			connVars = append(connVars, fmt.Sprintf("host=%s", host))
		}
		
		if port := os.Getenv("PGPORT"); port != "" {
			connVars = append(connVars, fmt.Sprintf("port=%s", port))
		}
		
		if user := os.Getenv("PGUSER"); user != "" {
			connVars = append(connVars, fmt.Sprintf("user=%s", user))
		}
		
		if password := os.Getenv("PGPASSWORD"); password != "" {
			connVars = append(connVars, fmt.Sprintf("password=%s", password))
		}
		
		if dbname := os.Getenv("PGDATABASE"); dbname != "" {
			connVars = append(connVars, fmt.Sprintf("dbname=%s", dbname))
		} else {
			connVars = append(connVars, "dbname=llmrpg")
		}
		
		// Default settings for SSL mode and connection timeout
		connVars = append(connVars, "sslmode=disable")
		connVars = append(connVars, "connect_timeout=10")
		
		connString = strings.Join(connVars, " ")
	}
	
	// Create the connection pool
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("error parsing postgres connection string: %w", err)
	}
	
	// Set pool configuration
	config.MaxConns = 10
	
	// Create the pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("error connecting to postgres: %w", err)
	}
	
	// Ping the database to verify connection
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("error pinging postgres: %w", err)
	}
	
	return pool, nil
}