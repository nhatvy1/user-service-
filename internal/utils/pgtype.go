package utils

import "github.com/jackc/pgx/v5/pgtype"

func OptionalBool(v *bool) pgtype.Bool {
	if v == nil {
		return pgtype.Bool{Valid: false}
	}
	return pgtype.Bool{Bool: *v, Valid: true}
}

func OptionalText(v *string) pgtype.Text {
	if v == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *v, Valid: true}
}

func NullableBool(v *bool) pgtype.Bool {
	return pgtype.Bool{Bool: *v, Valid: true}
}
