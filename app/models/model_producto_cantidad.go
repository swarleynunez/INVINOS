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

// Cantidad del producto. Se especifica cantidad y unidades
type ProductoCantidad struct {
	// Cantidad del producto
	Valor uint64 `json:"valor,omitempty"`
	// Unidades de medida
	Unidades string `json:"unidades,omitempty"`
}
