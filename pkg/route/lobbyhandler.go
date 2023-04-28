package route

import (
	"backend/pkg/models"
	"backend/pkg/models/lobby"
	"backend/pkg/services"
)

func CreateLobby(context *models.Context, data *models.Data) error {
	return services.Create(context)
}

func JoinLobby(context *models.Context, data *models.Data) error {
	req := lobby.JoinRequest{}
	err := req.Decode(&data.Message)

	if err != nil {
		return err
	}

	return services.Join(context.ForLobby(&req.LobbyId), false)
}

func SetLobbyName(context *models.Context, data *models.Data) error {
	req := lobby.SetNameRequest{}
	err := req.Decode(&data.Message)

	if err != nil {
		return err
	}

	return services.NameChange(context.ForLobby(&req.LobbyId), &req.Text)
}

func GetLobby(context *models.Context, data *models.Data) error {
	req := lobby.GetRequest{}
	err := req.Decode(&data.Message)

	if err != nil {
		return err
	}

	return services.Get(context.ForLobby(&req.LobbyId))
}
