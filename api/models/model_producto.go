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

type Producto struct {
	Id uint64 `json:"id"`
	// Nombre del producto
	Nombre string `json:"nombre,omitempty"`
	// Tipo de producto
	Tipo string `json:"tipo,omitempty"`

	Cantidad *ProductoCantidad `json:"cantidad"`
	// Información sobre el producto, como está hecho, añadas, etc.
	Info OneOfProductoInfoItems `json:"info,omitempty"`
}
