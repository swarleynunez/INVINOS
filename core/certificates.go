package core

import (
	"context"
	"errors"
	"github.com/swarleynunez/INVINOS/app/models"
)

var (
	ErrUnknownCertType = errors.New("unknown certificate type")
)

func CheckForCompositionCertificates(ctx context.Context, info *models.OneOfProductoInfoItems) error {

	if len(info.InfoVino.Composicion) > 0 { // Case: "InfoVino"
		for compk, comp := range info.InfoVino.Composicion {
			if comp.Uva != nil {
				err := checkInfoComponenteForCertificates(ctx, &comp)
				if err != nil {
					return err
				}
				info.InfoVino.Composicion[compk] = comp
			}
		}
	} else if len(info.InfoMosto.Composicion) > 0 { // Case: "InfoMosto"
		for compk, comp := range info.InfoMosto.Composicion {
			if comp.Uva != nil {
				err := checkInfoComponenteForCertificates(ctx, &comp)
				if err != nil {
					return err
				}
				info.InfoMosto.Composicion[compk] = comp
			}
		}
	} else if len(info.InfoOtros.Composicion) > 0 { // Case: "InfoOtros"
		for compk, comp := range info.InfoOtros.Composicion {
			if comp.Uva != nil {
				err := checkInfoComponenteForCertificates(ctx, &comp)
				if err != nil {
					return err
				}
				info.InfoOtros.Composicion[compk] = comp
			}
		}
	}

	return nil
}

func CheckForProductCertificates(ctx context.Context, info *models.OneOfProductoInfoItems) error {

	if len(info.InfoVino.Certificados) > 0 { // Case: "InfoVino"
		for certsk, certs := range info.InfoVino.Certificados {
			ci, err := checkCertificadoItem(ctx, certs)
			if err != nil {
				return err
			}
			info.InfoVino.Certificados[certsk] = ci.(models.AnyOfInfoVinoCertificadosItems)
		}
	} else if len(info.InfoUva.Certificados) > 0 { // Case: "InfoUva"
		for certsk, certs := range info.InfoUva.Certificados {
			ci, err := checkCertificadoItem(ctx, certs)
			if err != nil {
				return err
			}
			info.InfoUva.Certificados[certsk] = ci.(models.AnyOfInfoUvaCertificadosItems)
		}
	} else if len(info.InfoMosto.Certificados) > 0 { // Case: "InfoMosto"
		for certsk, certs := range info.InfoMosto.Certificados {
			ci, err := checkCertificadoItem(ctx, certs)
			if err != nil {
				return err
			}
			info.InfoMosto.Certificados[certsk] = ci.(models.AnyOfInfoMostoCertificadosItems)
		}
	}

	return nil
}

func checkInfoComponenteForCertificates(ctx context.Context, comp *models.InfoComponente) error {

	for ck := range comp.Uva.Certificados {
		// Case: "DOP" certificate
		err := checkDOPCertificate(ctx, &comp.Uva.Certificados[ck].CertDop)
		if err != nil {
			return err
		}

		// Case: "IGP" certificate
		err = checkIGPCertificate(ctx, &comp.Uva.Certificados[ck].CertIgp)
		if err != nil {
			return err
		}

		// Case: "Otros" certificate
		err = checkOtrosCertificate(ctx, &comp.Uva.Certificados[ck].CertOtros)
		if err != nil {
			return err
		}

		// Case: "Ecologico" certifier
		err = checkEcologicoCertificates(ctx, &comp.Uva.Certificados[ck].CertEcologico)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkCertificadoItem(ctx context.Context, cii interface{}) (interface{}, error) {

	switch cii.(type) {
	// Case: "InfoVino" certificate item
	case models.AnyOfInfoVinoCertificadosItems:
		// Cast interface to certificate item type
		ci := cii.(models.AnyOfInfoVinoCertificadosItems)

		// Case: "DOP" certificate
		err := checkDOPCertificate(ctx, &ci.CertDop)
		if err != nil {
			return nil, err
		}

		// Case: "IGP" certificate
		err = checkIGPCertificate(ctx, &ci.CertIgp)
		if err != nil {
			return nil, err
		}

		// Case: "Otros" certificate
		err = checkOtrosCertificate(ctx, &ci.CertOtros)
		if err != nil {
			return nil, err
		}

		// Case: "Ecologico" certifier
		err = checkEcologicoCertificates(ctx, &ci.CertEcologico)
		if err != nil {
			return nil, err
		}

		// Case: "Vegano" certificates
		err = checkVeganoCertificate(ctx, &ci.CertVegano)
		if err != nil {
			return nil, err
		}

		// Case: "Vegetariano" certificates
		err = checkVegetarianoCertificate(ctx, &ci.CertVegetariano)
		if err != nil {
			return nil, err
		}

		// Case: "Sulfitos" certificates
		err = checkSulfitosCertificate(ctx, &ci.CertSulfitos)
		if err != nil {
			return nil, err
		}

		return ci, nil

	// Case: "InfoUva" certificate item
	case models.AnyOfInfoUvaCertificadosItems:
		// Cast interface to certificate item type
		ci := cii.(models.AnyOfInfoUvaCertificadosItems)

		// Case: "DOP" certificate
		err := checkDOPCertificate(ctx, &ci.CertDop)
		if err != nil {
			return nil, err
		}

		// Case: "IGP" certificate
		err = checkIGPCertificate(ctx, &ci.CertIgp)
		if err != nil {
			return nil, err
		}

		// Case: "Otros" certificate
		err = checkOtrosCertificate(ctx, &ci.CertOtros)
		if err != nil {
			return nil, err
		}

		// Case: "Ecologico" certifier
		err = checkEcologicoCertificates(ctx, &ci.CertEcologico)
		if err != nil {
			return nil, err
		}

		return ci, nil

	// Case: "InfoMosto" certificate item
	case models.AnyOfInfoMostoCertificadosItems:
		// Cast interface to certificate item type
		ci := cii.(models.AnyOfInfoMostoCertificadosItems)

		// Case: "DOP" certificate
		err := checkDOPCertificate(ctx, &ci.CertDop)
		if err != nil {
			return nil, err
		}

		// Case: "IGP" certificate
		err = checkIGPCertificate(ctx, &ci.CertIgp)
		if err != nil {
			return nil, err
		}

		// Case: "Otros" certificate
		err = checkOtrosCertificate(ctx, &ci.CertOtros)
		if err != nil {
			return nil, err
		}

		// Case: "Ecologico" certifier
		err = checkEcologicoCertificates(ctx, &ci.CertEcologico)
		if err != nil {
			return nil, err
		}

		// Case: "Vegano" certificates
		err = checkEcologicoCertificates(ctx, &ci.CertEcologico)
		if err != nil {
			return nil, err
		}

		// Case: "Vegetariano" certificates
		err = checkVegetarianoCertificate(ctx, &ci.CertVegetariano)
		if err != nil {
			return nil, err
		}

		return ci, nil

	default:
		return nil, ErrUnknownCertType
	}
}

func checkDOPCertificate(ctx context.Context, cert *models.CertDop) error {

	if cert.Certificado != nil {
		for fk, f := range cert.Certificado.Files {
			if f.Name != "" {
				cid, err := uploadIPFSFile(ctx, f.Name)
				if err != nil {
					return err
				}

				// Link CID to the file
				cert.Certificado.Files[fk].CID = cid
			}
		}
	}

	return nil
}

func checkIGPCertificate(ctx context.Context, cert *models.CertIgp) error {

	if cert.Certificado != nil {
		for fk, f := range cert.Certificado.Files {
			if f.Name != "" {
				cid, err := uploadIPFSFile(ctx, f.Name)
				if err != nil {
					return err
				}

				// Link CID to the file
				cert.Certificado.Files[fk].CID = cid
			}
		}
	}

	return nil
}

func checkOtrosCertificate(ctx context.Context, cert *models.CertOtros) error {

	if cert.Certificado != nil {
		for fk, f := range cert.Certificado.Files {
			if f.Name != "" {
				cid, err := uploadIPFSFile(ctx, f.Name)
				if err != nil {
					return err
				}

				// Link CID to the file
				cert.Certificado.Files[fk].CID = cid
			}
		}
	}

	return nil
}

func checkEcologicoCertificates(ctx context.Context, cert *models.CertEcologico) error {

	// Certifier
	if cert.Certificadora != nil {
		for fk, f := range cert.Certificadora.Files {
			if f.Name != "" {
				cid, err := uploadIPFSFile(ctx, f.Name)
				if err != nil {
					return err
				}

				// Link CID to the file
				cert.Certificadora.Files[fk].CID = cid
			}
		}
	}

	// Certificates
	for ek, e := range cert.Extra {
		// Case: "Ecologico-SistProduccionBiodinamico" certificate
		if e.SistProduccionBiodinamico.Certificado != nil {
			for fk, f := range e.SistProduccionBiodinamico.Certificado.Files {
				if f.Name != "" {
					cid, err := uploadIPFSFile(ctx, f.Name)
					if err != nil {
						return err
					}

					// Link CID to the file
					cert.Extra[ek].SistProduccionBiodinamico.Certificado.Files[fk].CID = cid
				}
			}
		}

		// Case: "Ecologico-SistProduccionNatural" certificate
		if e.SistProduccionNatural.Certificado != nil {
			for fk, f := range e.SistProduccionNatural.Certificado.Files {
				if f.Name != "" {
					cid, err := uploadIPFSFile(ctx, f.Name)
					if err != nil {
						return err
					}

					// Link CID to the file
					cert.Extra[ek].SistProduccionNatural.Certificado.Files[fk].CID = cid
				}
			}
		}
	}

	return nil
}

func checkVeganoCertificate(ctx context.Context, cert *models.CertVegano) error {

	if cert.Certificado != nil {
		for fk, f := range cert.Certificado.Files {
			if f.Name != "" {
				cid, err := uploadIPFSFile(ctx, f.Name)
				if err != nil {
					return err
				}

				// Link CID to the file
				cert.Certificado.Files[fk].CID = cid
			}
		}
	}

	return nil
}

func checkVegetarianoCertificate(ctx context.Context, cert *models.CertVegetariano) error {

	if cert.Certificado != nil {
		for fk, f := range cert.Certificado.Files {
			if f.Name != "" {
				cid, err := uploadIPFSFile(ctx, f.Name)
				if err != nil {
					return err
				}

				// Link CID to the file
				cert.Certificado.Files[fk].CID = cid
			}
		}
	}

	return nil
}

func checkSulfitosCertificate(ctx context.Context, cert *models.CertSulfitos) error {

	if cert.ConSulfitos.Certificado != nil {
		for fk, f := range cert.ConSulfitos.Certificado.Files {
			if f.Name != "" {
				cid, err := uploadIPFSFile(ctx, f.Name)
				if err != nil {
					return err
				}

				// Link CID to the file
				cert.ConSulfitos.Certificado.Files[fk].CID = cid
			}
		}
	} else if cert.SinSulfitos.Certificado != nil {
		for fk, f := range cert.SinSulfitos.Certificado.Files {
			if f.Name != "" {
				cid, err := uploadIPFSFile(ctx, f.Name)
				if err != nil {
					return err
				}

				// Link CID to the file
				cert.SinSulfitos.Certificado.Files[fk].CID = cid
			}
		}
	}

	return nil
}
