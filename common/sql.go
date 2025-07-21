package common

import (
	"database/sql"
	"time"
)

func DbStr(str string) sql.NullString {
	return sql.NullString{
		String: str,
		Valid:  true,
	}
}

func DbInt64(v int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: v,
		Valid: true,
	}
}

func DbInt64Nullable(v *int64) sql.NullInt64 {
	var val int64
	if v != nil {
		val = *v
	}
	return sql.NullInt64{
		Int64: val,
		Valid: v != nil,
	}
}

func DbTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

func FromDbStr(dbStr sql.NullString) string {
	if dbStr.Valid {
		return dbStr.String
	}
	return ""
}

func FromDbInt64(dbInt sql.NullInt64) int64 {
	if dbInt.Valid {
		return dbInt.Int64
	}
	return 0
}

func FromNullableDbInt64(dbInt sql.NullInt64) *int64 {
	if dbInt.Valid {
		return &dbInt.Int64
	}
	return nil
}

func LikeStr(str string) string {
	return "%" + str + "%"
}

func IfNilDefault[T any](v *T) T {
	if v == nil {
		var result T
		return result
	}
	return *v
}

func IfNil[T any](v *T, defaultVal T) T {
	if v == nil {
		return defaultVal
	}
	return *v
}
