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

// Información sobre el descargadero
type InfoDescargadero struct {
	Tipo string `json:"tipo,omitempty"`

	Dimensiones *InfoDescargaderoDimensiones `json:"dimensiones,omitempty"`
	// Material del que está hecho el descargadero
	Material string `json:"material,omitempty"`
}
