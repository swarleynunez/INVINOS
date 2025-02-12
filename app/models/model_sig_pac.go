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

// Información SIGPAC sobre la parcela de cultivo
type SigPac struct {
	// Provincia de la parcela
	Provincia float64 `json:"provincia"`
	// Municipio de la parcela
	Municipio float64 `json:"municipio"`
	// Numeración que se le da al polígono
	Poligono float64 `json:"poligono"`
	// Numeración que se le da a la parcela
	Parcela float64 `json:"parcela"`
	// Numeración que se le da al recinto
	Recinto float64 `json:"recinto"`
}
