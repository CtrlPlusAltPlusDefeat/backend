package handlers

import (
	apigateway "backend/pkg/aws-helpers/api-gateway"
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/services"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

// DisconnectHandler we want to delete the connection information from db here and return an ok response
func DisconnectHandler(context context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigateway.Response, error) {
	log.Printf("DisconnectHandler requestId: %s, connectionId:%s \n\r", req.RequestContext.RequestID, req.RequestContext.ConnectionID)

	con := models.NewContext(context, &req.RequestContext.ConnectionID, &req.RequestContext.DomainName, &req.RequestContext.Stage)
	data, _ := models.NewData(req.Body)

	connection, err := db.Connection.Get(&req.RequestContext.ConnectionID)
	if err != nil {
		return apigateway.Response{}, err
	}
	con.ForSession(&connection.SessionId)

	lobbies, err := db.LobbyPlayer.GetLobbiesBySessionId(&connection.SessionId)

	for _, l := range lobbies {
		if l.IsOnline {
			con.ForLobby(&models.Lobby{LobbyId: l.LobbyId})
			//error should already be outputted higher up
			_ = services.LeaveLobby(con, data)
		}
	}

	if err != nil {
		return apigateway.Response{}, err
	}

	err = db.Connection.Remove(&req.RequestContext.ConnectionID)

	if err != nil {
		return apigateway.Response{}, err
	}

	return apigateway.OkResponse(), nil
}
