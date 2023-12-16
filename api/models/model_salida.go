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

type Salida struct {
	Id string `json:"id"`
	// Fecha de la venta
	Fecha string `json:"fecha"`

	Producto *ProductoVendido `json:"producto"`

	Comprador *Empresa `json:"comprador"`

	AContenedor *Contenedor `json:"a_contenedor,omitempty"`
	// Información adicional sobre la venta
	Info string `json:"info,omitempty"`
}
