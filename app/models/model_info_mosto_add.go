/*
 * API for INVINOS blockchain broker
 *
 * Esta es la API para el brocker de INVINOS con el que hay que interactuar para realizar transacciones en la blockchain.
 *
 * API version: 1.0
 * Contact: francisco.delicado@uclm.es
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type InfoMostoAdd struct {
	// Tipo de mosto según proceso de elaboración
	Tipo string `json:"tipo,omitempty"`
	// Tipo de vino según color de vino
	Color string `json:"color,omitempty"`

	Quimica *InfoMostoAddQuimica `json:"quimica,omitempty"`
	// Uvas que componen el mosto
	Composicion []InfoComponente `json:"composicion,omitempty"`
	// Certificados del mosto
	Certificados []AnyOfInfoMostoAddCertificadosItems `json:"certificados,omitempty"`
}
