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
	"encoding/json"
	"github.com/swarleynunez/INVINOS/api/models"
	"github.com/swarleynunez/INVINOS/core"
	"github.com/swarleynunez/INVINOS/core/utils"
	"net/http"
)

func AddProductPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

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
