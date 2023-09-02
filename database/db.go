package database

import (
	"example-api/models"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

var (
	clients      []models.Client = []models.Client{}
	clientsMutex *sync.RWMutex   = &sync.RWMutex{}
)

func GetClient(id string) (retrievedClient models.Client, err error) {

	clientsMutex.RLock()
	defer clientsMutex.RUnlock()

	idx := slices.IndexFunc(clients, func(c models.Client) bool {
		return c.Id == id
	})

	if idx == -1 {
		err = ErrClientNotFound
		return
	}

	retrievedClient = clients[idx]

	return
}

func GetClients() (retrievedClients []models.Client, err error) {

	clientsMutex.RLock()
	defer clientsMutex.RUnlock()

	retrievedClients = clients
	return
}

func UpdateClient(id string, desired models.Client) (updatedClient models.Client, err error) {

	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	idx := slices.IndexFunc(clients, func(c models.Client) bool {
		return c.Id == id
	})

	if idx == -1 {
		err = ErrClientNotFound
		return
	}

	clients[idx] = desired
	updatedClient = clients[idx]

	return
}

func DeleteClient(id string) (err error) {

	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	idx := slices.IndexFunc(clients, func(c models.Client) bool {
		return c.Id == id
	})

	if idx == -1 {
		err = ErrClientNotFound
		return
	}

	clients = append(clients[idx:], clients[:idx+1]...)

	return
}

func CreateClient(client models.Client) (createdClient models.Client, err error) {

	client.Id = uuid.New().String()

	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	clients = append(clients, client)
	createdClient = client

	return

}
