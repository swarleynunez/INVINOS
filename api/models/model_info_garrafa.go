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

// Información sobre la garrafa
type InfoGarrafa struct {
	Tipo string `json:"tipo,omitempty"`

	Dimensiones *InfoGarrafaDimensiones `json:"dimensiones,omitempty"`
	// Material del que está hecha la garrafa
	Material string `json:"material,omitempty"`

	Etiqueta *Etiqueta `json:"etiqueta,omitempty"`
}
