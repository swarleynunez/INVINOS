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

// Información sobre un componente de un producto. Se especifica el % de dicho componente y la uva que lo componen
type InfoComponente struct {
	// Porcentaje del componente en el producto
	Porcentaje float64 `json:"porcentaje,omitempty"`

	Uva *InfoUva `json:"uva,omitempty"`
}
