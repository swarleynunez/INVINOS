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

// Dimensiones del descargadero
type InfoDescargaderoDimensiones struct {
	// Alto del descargadero en cm
	Alto float64 `json:"alto,omitempty"`
	// Ancho del descargadero en cm
	Ancho float64 `json:"ancho,omitempty"`
	// Largo del descargadero en cm
	Largo float64 `json:"largo,omitempty"`
	// Capacidad del descargadero en litros
	Capacidad float64 `json:"capacidad,omitempty"`
}
