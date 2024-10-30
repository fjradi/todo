package internal

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"reflect"
)

type ApiError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return fe.Error()
}

func parseValidationError(err error) any {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), msgForTag(fe)}
		}
		return out
	}

	return err.Error()
}

func parseSqlError(err error) (httpStatusCode int, message any) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			httpStatusCode = http.StatusBadRequest
			switch pgErr.ConstraintName {
			case "todo_uk":
				message = "duplicate todo name"
			default:
				message = pgErr.Message
			}
			return
		}
	}

	httpStatusCode = http.StatusInternalServerError
	message = err.Error()
	return
}

func parseStringToUUID(id string) (uuid pgtype.UUID, err error) {
	err = uuid.Scan(id)

	return
}

func registerTag(fld reflect.StructField) string {
	name := fld.Tag.Get("form")
	if name == "" {
		name = fld.Tag.Get("json")
	}
	if name == "" {
		name = fld.Tag.Get("uri")
	}
	return name
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(registerTag)
	}
}
