/*
 * API for INVINOS blockchain broker
 *
 * Esta es la API para el brocker de INVINOS con el que hay que interactuar para realizar transacciones en la blockchain.
 *
 * API version: 1.0
 * Contact: francisco.delicado@uclm.es
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package handlers

import (
	"context"
	"encoding/json"
	"github.com/swarleynunez/INVINOS/api/models"
	"github.com/swarleynunez/INVINOS/core"
	"github.com/swarleynunez/INVINOS/core/utils"
	"net/http"
)

func AddProductTypePost(w http.ResponseWriter, r *http.Request) {

	// Get POST parameters
	var tp models.TipoProducto
	err := json.NewDecoder(r.Body).Decode(&tp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.CheckError(err, utils.WarningMode)
		return
	}

	// Parameter checking zone
	if tp.Id == "" {
		w.WriteHeader(http.StatusBadRequest)
		utils.CheckError(ErrMissingParameters, utils.WarningMode)
		return
	}

	// Manage POST certificate files (composition certificates)
	err = core.CheckForCompositionCertificates(context.Background(), &tp.Info)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.CheckError(err, utils.WarningMode)
		return
	}

	// Manage POST certificate files (product certificates)
	err = core.CheckForProductCertificates(context.Background(), &tp.Info)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.CheckError(err, utils.WarningMode)
		return
	}

	// Blockchain interaction
	err = core.AddProductTypeTxn(tp.Id, utils.MarshalJSON(tp))
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		utils.CheckError(err, utils.WarningMode)
		return
	}

	// Successful request
	w.WriteHeader(http.StatusOK)
}

func AddCompanyPost(w http.ResponseWriter, r *http.Request) {

	// Get POST parameters
	var emp models.Empresa
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.CheckError(err, utils.WarningMode)
		return
	}

	// Parameter checking zone
	if emp.Id == "" || emp.Nombre == "" || emp.Cif == "" {
		w.WriteHeader(http.StatusBadRequest)
		utils.CheckError(ErrMissingParameters, utils.WarningMode)
		return
	}

	// Blockchain interaction
	err = core.AddCompanyTxn(emp.Id, utils.MarshalJSON(emp))
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		utils.CheckError(err, utils.WarningMode)
		return
	}

	// Successful request
	w.WriteHeader(http.StatusOK)
}

func AddContainerPost(w http.ResponseWriter, r *http.Request) {

	// Get POST parameters
	var ctr models.Contenedor
	err := json.NewDecoder(r.Body).Decode(&ctr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.CheckError(err, utils.WarningMode)
		return
	}

	// Parameter checking zone
	if ctr.Id == "" {
		w.WriteHeader(http.StatusBadRequest)
		utils.CheckError(ErrMissingParameters, utils.WarningMode)
		return
	}

	// Blockchain interaction
	err = core.AddContainerTxn(ctr.Id, utils.MarshalJSON(ctr))
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		utils.CheckError(err, utils.WarningMode)
		return
	}

	// Successful request
	w.WriteHeader(http.StatusOK)
}
