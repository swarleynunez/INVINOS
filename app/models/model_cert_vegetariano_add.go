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

// Certificado de producto vegetariano
type CertVegetarianoAdd struct {
	// Tipo de certificado
	Tipo string `json:"tipo,omitempty"`
	// Nombre del certificado
	Name string `json:"name,omitempty"`

	Certificado *CertTemplateAdd `json:"certificado,omitempty"`
}
