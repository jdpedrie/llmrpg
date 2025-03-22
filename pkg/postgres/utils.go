package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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