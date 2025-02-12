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

type Entrada struct {
	Id string `json:"id"`
	// Fecha de la entrada
	Fecha string `json:"fecha"`

	Vendedor *Empresa `json:"vendedor"`

	Producto *Producto `json:"producto"`

	AContenedor *Contenedor `json:"a_contenedor"`
	// Información adicional sobre la entrada
	Info string `json:"info,omitempty"`
}
