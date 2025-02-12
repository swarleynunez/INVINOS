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

// Dimensiones de la lata
type InfoLataDimensiones struct {
	// Alto de la lata en cm
	Alto float64 `json:"alto,omitempty"`
	// Ancho de la lata en cm
	Ancho float64 `json:"ancho,omitempty"`
	// Largo de la lata en cm
	Largo float64 `json:"largo,omitempty"`
	// Capacidad de la lata en litros
	Capacidad float64 `json:"capacidad,omitempty"`
}
