package controllers

import (
	"errors"
	"example-api/database"
	"example-api/models"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/pquerna/ffjson/ffjson"
	"golang.org/x/exp/slices"
)

// @Description Create a client
// @Summary creates a client
// @Tags Clients v1
// @Accept json
// @Produce json
// @Param	client	body	models.ClientCreateQuery	true	"Client to add"
// @Success 200 {object} models.Client
// @Failure 415
// @Failure 500 {object} ResultError
// @Router /api/v1/clients [post]
func createClient(ctx *fiber.Ctx) (err error) {

	responseStatus := fiber.StatusOK
	responseBody := []byte{}

	supportedMIMETypes := []string{
		fiber.MIMEApplicationJSON,
	}

	defer func() {
		ctx.Status(responseStatus)
		ctx.Context().SetBody(responseBody)
	}()

	defer func() {
		if err != nil {
			resultErr := ResultError{}
			resultErr.Error.Description = err.Error()
			responseBody, err = ffjson.Marshal(resultErr)
			if err != nil {
				err = fmt.Errorf("unable to encode error result: %w", err)
				log.Error(err.Error())
			}
		}
	}()

	if !slices.Contains(supportedMIMETypes, ctx.GetReqHeaders()[fiber.HeaderContentType]) {
		responseStatus = fiber.StatusUnsupportedMediaType
		return
	}

	toCreateClientQuery := models.ClientCreateQuery{}
	ffjson.Unmarshal(ctx.Body(), &toCreateClientQuery)

	toCreateClient := models.Client{
		Name: toCreateClientQuery.Name,
		Tel: toCreateClientQuery.Tel,
		Zipcode: toCreateClientQuery.Zipcode,
		Address: toCreateClientQuery.Address,
	}

	createdClient, err := database.CreateClient(toCreateClient)
	if err != nil {
		err = fmt.Errorf("unable to create client: %w", err)
		responseStatus = fiber.StatusInternalServerError
		responseBody = []byte(err.Error())
		return
	}

	responseBody, err = ffjson.Marshal(&createdClient)
	if err != nil {
		err = fmt.Errorf("unable to encode response: %w", err)
		responseStatus = fiber.StatusInternalServerError
		responseBody = []byte(err.Error())
		return
	}

	return

}

// @Description Retrieve a client
// @Summary retrieves a client
// @Tags Clients v1
// @Produce json
// @Param	id	path	string	true	"Client ID"
// @Success 200 {object} models.Client
// @Success 400 {object} ResultError
// @Success 404 {object} ResultError
// @Failure 500 {object} ResultError
// @Router /api/v1/clients/{id} [get]
func getClient(ctx *fiber.Ctx) (err error) {

	responseStatus := fiber.StatusOK
	responseBody := []byte{}

	defer func() {
		ctx.Status(responseStatus)
		ctx.Context().SetBody(responseBody)
	}()

	defer func() {
		if err != nil {
			resultErr := ResultError{}
			resultErr.Error.Description = err.Error()
			responseBody, err = ffjson.Marshal(resultErr)
			if err != nil {
				err = fmt.Errorf("unable to encode error result: %w", err)
				log.Error(err.Error())
			}
		}
	}()

	clientIDToRetrieve := ctx.Params("id", "null")
	if clientIDToRetrieve == "null" {
		err = errors.New("must provide client id")
		responseStatus = fiber.StatusBadRequest
		responseBody = []byte(err.Error())
		return
	}

	parsedClientIdToRetrieve, err := uuid.Parse(clientIDToRetrieve)
	if err != nil {
		err = fmt.Errorf("unable to parse uuid: %w", err)
		responseStatus = fiber.StatusBadRequest
		responseBody = []byte(err.Error())
		return

	}

	retrievedClient, err := database.GetClient(parsedClientIdToRetrieve.String())
	if err != nil {

		if errors.Is(err, database.ErrClientNotFound) {
			responseStatus = fiber.StatusBadRequest
		} else {
			err = fmt.Errorf("unable to get client: %w", err)
			responseStatus = fiber.StatusInternalServerError

		}

		responseBody = []byte(err.Error())
		return
	}

	responseBody, err = ffjson.Marshal(&retrievedClient)
	if err != nil {
		err = fmt.Errorf("unable to encode response: %w", err)
		responseStatus = fiber.StatusInternalServerError
		responseBody = []byte(err.Error())
		return
	}

	return

}

// @Description Retrieve all clients
// @Summary retrieves all clients
// @Tags Clients v1
// @Produce json
// @Success 200 {object} []models.Client
// @Success 404 {object} ResultError
// @Failure 500 {object} ResultError
// @Router /api/v1/clients [get]
func getClients(ctx *fiber.Ctx) (err error) {

	responseStatus := fiber.StatusOK
	responseBody := []byte{}

	defer func() {
		ctx.Status(responseStatus)
		ctx.Context().SetBody(responseBody)
	}()

	defer func() {
		if err != nil {
			resultErr := ResultError{}
			resultErr.Error.Description = err.Error()
			responseBody, err = ffjson.Marshal(resultErr)
			if err != nil {
				err = fmt.Errorf("unable to encode error result: %w", err)
				log.Error(err.Error())
			}
		}
	}()

	retrievedClients, err := database.GetClients()
	if err != nil {

		err = fmt.Errorf("unable to get client: %w", err)
		responseStatus = fiber.StatusInternalServerError
		responseBody = []byte(err.Error())
		return
	}

	responseBody, err = ffjson.Marshal(&retrievedClients)
	if err != nil {
		err = fmt.Errorf("unable to encode response: %w", err)
		responseStatus = fiber.StatusInternalServerError
		responseBody = []byte(err.Error())
		return
	}

	return

}

// @Description Update a client
// @Summary updates a client
// @Tags Clients v1
// @Accept json
// @Produce json
// @Param	id	path	string	true	"Client ID"
// @Param	client	body	models.ClientUpdateQuery	true	"Desired client"
// @Success 200 {object} models.Client
// @Success 400 {object} ResultError
// @failure 404 {object} ResultError
// @Failure 415
// @Failure 500 {object} ResultError
// @Router /api/v1/clients/{id} [put]
func replaceClient(ctx *fiber.Ctx) (err error) {

	responseStatus := fiber.StatusOK
	responseBody := []byte{}

	supportedMIMETypes := []string{
		fiber.MIMEApplicationJSON,
	}

	defer func() {
		ctx.Status(responseStatus)
		ctx.Context().SetBody(responseBody)
	}()

	defer func() {
		if err != nil {
			resultErr := ResultError{}
			resultErr.Error.Description = err.Error()
			responseBody, err = ffjson.Marshal(resultErr)
			if err != nil {
				err = fmt.Errorf("unable to encode error result: %w", err)
				log.Error(err.Error())
			}
		}
	}()

	if !slices.Contains(supportedMIMETypes, ctx.GetReqHeaders()[fiber.HeaderContentType]) {
		responseStatus = fiber.StatusUnsupportedMediaType
		return
	}

	clientIDToUpdate := ctx.Params("id", "null")
	if clientIDToUpdate == "null" {
		err = errors.New("must provide client id")
		responseStatus = fiber.StatusBadRequest
		responseBody = []byte(err.Error())
		return
	}

	parsedclientIDToUpdate, err := uuid.Parse(clientIDToUpdate)
	if err != nil {
		err = fmt.Errorf("unable to parse uuid: %w", err)
		responseStatus = fiber.StatusBadRequest
		responseBody = []byte(err.Error())
		return

	}

	desiredClientQuery := models.ClientUpdateQuery{}
	ffjson.Unmarshal(ctx.Body(), &desiredClientQuery)

	desiredClient := models.Client{
		Name: desiredClientQuery.Name,
		Tel: desiredClientQuery.Tel,
		Zipcode: desiredClientQuery.Zipcode,
		Address: desiredClientQuery.Address,
	}

	updatedClient, err := database.UpdateClient(parsedclientIDToUpdate.String(), desiredClient)
	if err != nil {
		if errors.Is(err, database.ErrClientNotFound) {
			responseStatus = fiber.StatusBadRequest
		} else {
			err = fmt.Errorf("unable to replace client: %w", err)
			responseStatus = fiber.StatusInternalServerError
		}

		responseBody = []byte(err.Error())
		return
	}

	responseBody, err = ffjson.Marshal(&updatedClient)
	if err != nil {
		err = fmt.Errorf("unable to encode response: %w", err)
		responseStatus = fiber.StatusInternalServerError
		responseBody = []byte(err.Error())
		return
	}

	return

}

// @Description Patch a client
// @Summary patches a client
// @Tags Clients v1
// @Accept json
// @Produce json
// @Param	id	path	string	true	"Client ID"
// @Param	client	body	models.ClientUpdateQuery	true	"Client fields to update"
// @Success 200 {object} models.Client
// @Failure 400 {object} ResultError
// @Failure 404 {object} ResultError
// @Failure 415
// @Failure 500 {object} ResultError
// @Router /api/v1/clients/{id} [patch]
func updateClient(ctx *fiber.Ctx) (err error) {

	responseStatus := fiber.StatusOK
	responseBody := []byte{}

	supportedMIMETypes := []string{
		fiber.MIMEApplicationJSON,
	}

	defer func() {
		ctx.Status(responseStatus)
		ctx.Context().SetBody(responseBody)
	}()

	defer func() {
		if err != nil {
			resultErr := ResultError{}
			resultErr.Error.Description = err.Error()
			responseBody, err = ffjson.Marshal(resultErr)
			if err != nil {
				err = fmt.Errorf("unable to encode error result: %w", err)
				log.Error(err.Error())
			}
		}
	}()

	if !slices.Contains(supportedMIMETypes, ctx.GetReqHeaders()[fiber.HeaderContentType]) {
		responseStatus = fiber.StatusUnsupportedMediaType
		return
	}

	clientIDToUpdate := ctx.Params("id", "null")
	if clientIDToUpdate == "null" {
		err = errors.New("must provide client id")
		responseStatus = fiber.StatusBadRequest
		responseBody = []byte(err.Error())
		return
	}

	parsedclientIDToUpdate, err := uuid.Parse(clientIDToUpdate)
	if err != nil {
		err = fmt.Errorf("unable to parse uuid: %w", err)
		responseStatus = fiber.StatusBadRequest
		responseBody = []byte(err.Error())
		return

	}

	retrievedClient, err := database.GetClient(parsedclientIDToUpdate.String())
	if err != nil {

		if errors.Is(err, database.ErrClientNotFound) {
			responseStatus = fiber.StatusBadRequest
		} else {
			err = fmt.Errorf("unable to get client: %w", err)
			responseStatus = fiber.StatusInternalServerError

		}

		responseBody = []byte(err.Error())
		return
	}

	desiredClientQuery := models.ClientUpdateQuery{}
	ffjson.Unmarshal(ctx.Body(), &desiredClientQuery)

	desiredClient := models.Client{
		Name: desiredClientQuery.Name,
		Tel: desiredClientQuery.Tel,
		Zipcode: desiredClientQuery.Zipcode,
		Address: desiredClientQuery.Address,
	}

	// updates retrievedClient with desiredClient's non-zero fields
	v := reflect.ValueOf(models.Client{})
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Field(i).Type().Name()
		field := reflect.ValueOf(&desiredClient).FieldByName(fieldName)

		if !field.IsZero() {
			toSetField := reflect.ValueOf(&retrievedClient).FieldByName(fieldName)

			if toSetField.CanSet() {
				toSetField.Set(field)
			}
		}

	}

	updatedClient, err := database.UpdateClient(parsedclientIDToUpdate.String(), retrievedClient)
	if err != nil {
		if errors.Is(err, database.ErrClientNotFound) {
			responseStatus = fiber.StatusBadRequest
		} else {
			err = fmt.Errorf("unable to update client: %w", err)
			responseStatus = fiber.StatusInternalServerError
		}

		responseBody = []byte(err.Error())
		return
	}

	responseBody, err = ffjson.Marshal(&updatedClient)
	if err != nil {
		err = fmt.Errorf("unable to encode response: %w", err)
		responseStatus = fiber.StatusInternalServerError
		responseBody = []byte(err.Error())
		return
	}

	return

}

// @Description Delete a client
// @Summary deletes a client
// @Tags Clients v1
// @Produce json
// @Param	id	path	string	true	"Client ID"
// @Success 204
// @Failure 400 {object} ResultError
// @Failure 404 {object} ResultError
// @Failure 500 {object} ResultError
// @Router /api/v1/clients/{id} [delete]
func deleteClient(ctx *fiber.Ctx) (err error) {

	responseStatus := fiber.StatusNoContent
	responseBody := []byte{}

	defer func() {
		ctx.Status(responseStatus)
		ctx.Context().SetBody(responseBody)
	}()

	defer func() {
		if err != nil {
			resultErr := ResultError{}
			resultErr.Error.Description = err.Error()
			responseBody, err = ffjson.Marshal(resultErr)
			if err != nil {
				err = fmt.Errorf("unable to encode error result: %w", err)
				log.Error(err.Error())
			}
		}
	}()

	clientIDToRetrieve := ctx.Params("id", "null")
	if clientIDToRetrieve == "null" {
		err = errors.New("must provide client id")
		responseStatus = fiber.StatusBadRequest
		responseBody = []byte(err.Error())
		return
	}

	parsedClientIdToRetrieve, err := uuid.Parse(clientIDToRetrieve)
	if err != nil {
		err = fmt.Errorf("unable to parse uuid: %w", err)
		responseStatus = fiber.StatusBadRequest
		responseBody = []byte(err.Error())
		return

	}

	err = database.DeleteClient(parsedClientIdToRetrieve.String())
	if err != nil {

		if errors.Is(err, database.ErrClientNotFound) {
			responseStatus = fiber.StatusBadRequest
		} else {
			err = fmt.Errorf("unable to delete client: %w", err)
			responseStatus = fiber.StatusInternalServerError

		}

		responseBody = []byte(err.Error())
		return
	}

	return

}
