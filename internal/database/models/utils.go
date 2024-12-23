package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"strings"
	"unicode"
)

type FieldUpdate struct {
	Field string
	Value any
}

func updateFields(db *sql.DB, table string, updates []FieldUpdate, condition string) error {
	// slice to hold set clauses for SQL command
	clauses := make([]string, len(updates))
	args := make([]interface{}, len(updates))

	// populate clauses and arguments
	for i, update := range updates {
		clauses[i] = fmt.Sprintf("%s = $%d", camelToSnakeCase(update.Field), i+1)
		args[i] = update.Value
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, strings.Join(clauses, ", "), condition)

	_, err := db.Query(query, args...)

	if err != nil {
		return err
	}

	return nil
}

func camelToSnakeCase(camel string) string {
	var result string
	var words []string
	v := make([]rune, 0, len(camel)+10)

	for _, r := range camel {
		if unicode.IsUpper(r) {
			if len(v) > 0 {
				words = append(words, string(v))
				v = v[:0]
			}
		}
		v = append(v, unicode.ToLower(r))
	}

	if len(v) > 0 {
		words = append(words, string(v))
	}

	result = strings.Join(words, "_")

	return result
}
