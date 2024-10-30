package core

import (
	"errors"
	"github.com/swarleynunez/INVINOS/core/bindings"
	"math/big"
)

const (
	AddProductAction        = "add_product"
	AddCompanyAction        = "add_company"
	AddContainerAction      = "add_container"
	ProductEntryAction      = "product_entry"
	ProductProcessingAction = "product_processing"
	ProductPartitionAction  = "product_partition"
	ProductOutputAction     = "product_output"
)

func AddProductTypeTxn(einst *bindings.EntityInfo, id, info string) error {

	// Checking zone
	if !existProductType(einst, id) {
		// Send transaction
		_, err := einst.CreateProductType(ethTransactor(), id, info)
		return err
	} else {
		return errors.New(AddProductAction + ": transaction not sent")
	}
}

func AddCompanyTxn(einst *bindings.EntityInfo, id, info string) error {

	// Checking zone
	if !existCompany(einst, id) {
		// Send transaction
		_, err := einst.CreateCompany(ethTransactor(), id, info)
		return err
	} else {
		return errors.New(AddCompanyAction + ": transaction not sent")
	}
}

func AddContainerTxn(einst *bindings.EntityInfo, id, info string) error {

	// Checking zone
	if !existContainer(einst, id) {
		// Send transaction
		_, err := einst.CreateContainer(ethTransactor(), id, info)
		return err
	} else {
		return errors.New(AddContainerAction + ": transaction not sent")
	}
}

func ProductEntryTxn(tinst *bindings.Traceability, einst *bindings.EntityInfo, qty uint64, prodTypeID, coID, ctrID, info string) error {

	// Checking zone
	if existProductType(einst, prodTypeID) && existCompany(einst, coID) && existContainer(einst, ctrID) {
		// Send transaction
		_, err := tinst.ProductEntry(ethTransactor(), big.NewInt(int64(qty)), prodTypeID, coID, ctrID, info)
		return err
	} else {
		return errors.New(ProductEntryAction + ": transaction not sent")
	}
}

func ProductProcessingTxn(tinst *bindings.Traceability, einst *bindings.EntityInfo, prodID, lostQty uint64, prodTypeID, ctrID, info string) error {

	// Checking zone
	if isProductAvailable(tinst, prodID) &&
		lostQty <= GetProductQuantity(tinst, prodID) &&
		existProductType(einst, prodTypeID) &&
		existContainer(einst, ctrID) {
		// Send transaction
		_, err := tinst.ProductProcessing(ethTransactor(), big.NewInt(int64(prodID)), big.NewInt(int64(lostQty)), prodTypeID, ctrID, info)
		return err
	} else {
		return errors.New(ProductProcessingAction + ": transaction not sent")
	}
}

func ProductPartitionTxn(tinst *bindings.Traceability, einst *bindings.EntityInfo, prodID, qty uint64, ctrID, info string) error {

	// Checking zone
	if isProductAvailable(tinst, prodID) && qty < GetProductQuantity(tinst, prodID) && existContainer(einst, ctrID) {
		// Send transaction
		_, err := tinst.ProductPartition(ethTransactor(), big.NewInt(int64(prodID)), big.NewInt(int64(qty)), ctrID, info)
		return err
	} else {
		return errors.New(ProductPartitionAction + ": transaction not sent")
	}
}

func ProductOutputTxn(tinst *bindings.Traceability, einst *bindings.EntityInfo, prodID uint64, lotn, coID, ctrID, info string) error {

	// Checking zone
	if isProductAvailable(tinst, prodID) && isLotNumberAvailable(tinst, lotn) && existCompany(einst, coID) && existContainer(einst, ctrID) {
		// Send transaction
		_, err := tinst.ProductOutput(ethTransactor(), big.NewInt(int64(prodID)), lotn, coID, ctrID, info)
		return err
	} else {
		return errors.New(ProductOutputAction + ": transaction not sent")
	}
}
