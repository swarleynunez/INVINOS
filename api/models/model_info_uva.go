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

type InfoUva struct {
	// Variedad de uva
	Variedad string `json:"variedad,omitempty"`
	// Añada de la uva
	Anada float64 `json:"añada,omitempty"`
	// Tipo de vino según color de vino
	Color string `json:"color_uva,omitempty"`
	// Tipo de conducción de la uva
	Conduccion string `json:"conduccion,omitempty"`
	// Tipo de vendimia de la uva
	Vendimia string `json:"vendimia,omitempty"`
	// Certificados de la uva
	Certificados []AnyOfInfoUvaCertificadosItems `json:"certificados_uva,omitempty"`
}
