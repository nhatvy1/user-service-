package utils

import "github.com/jackc/pgx/v5/pgtype"

func NullableText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false} // NULL trong PostgreSQL
	}
	return pgtype.Text{
		String: *s,   // giá trị thực
		Valid:  true, // không null
	}
}
