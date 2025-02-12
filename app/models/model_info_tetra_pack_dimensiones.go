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

// Dimensiones del tetra pack
type InfoTetraPackDimensiones struct {
	// Alto del tetra pack en cm
	Alto float64 `json:"alto,omitempty"`
	// Ancho del tetra pack en cm
	Ancho float64 `json:"ancho,omitempty"`
	// Largo del tetra pack en cm
	Largo float64 `json:"largo,omitempty"`
	// Capacidad del tetra pack en litros
	Capacidad float64 `json:"capacidad,omitempty"`
}
