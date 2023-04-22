package ws

import "backend/pkg/db"

func Connect(id string) error {
	return db.Connection.GetClient().Add(id)

}
