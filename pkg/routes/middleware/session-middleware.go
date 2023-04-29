package middleware

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/routes"
)

func SessionMiddleware(next routes.Handler) routes.Handler {
	return func(context *models.Context, data *models.Data) error {
		res, err := db.Connection.Get(context.ConnectionId())

		if err != nil {
			return err
		}

		return next(context.ForSession(&res.SessionId), data)
	}
}
