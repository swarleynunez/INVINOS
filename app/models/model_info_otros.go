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

type InfoOtros struct {
	// Descripción del producto
	Descripcion string `json:"descripcion,omitempty"`
	// Variedades de uva que componen el producto
	Composicion []InfoComponente `json:"composicion_otros,omitempty"`
}
