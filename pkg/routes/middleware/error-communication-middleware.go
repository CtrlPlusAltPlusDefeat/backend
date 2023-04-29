package middleware

import (
	"backend/pkg/models"
	"backend/pkg/routes"
	"backend/pkg/ws"
)

func ErrorCommunicateMiddleware(next routes.Handler) routes.Handler {
	return func(context *models.Context, data *models.Data) error {
		err := next(context, data)

		if err != nil {
			res := models.ErrorResponse{Error: "Something went wrong handling this request."}

			err = ws.Send(context, data.Route(), res)

			return err
		}

		return nil
	}
}
