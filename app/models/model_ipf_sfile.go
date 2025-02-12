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

// Archivo previamente almacenado en el servidor
type File struct {
	// Nombre y extensión del archivo
	Name string `json:"name,omitempty"`
	// Hash IPFS del archivo
	CID string `json:"cid,omitempty"`
}
