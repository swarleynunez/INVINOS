package core

import (
	"errors"
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

func AddProductTypeTxn(id, info string) error {

	// Checking zone
	if !existProductType(id) {
		// Send transaction
		_, err := _einst.CreateProductType(ethTransactor(), id, info)
		return err
	} else {
		return errors.New(AddProductAction + ": transaction not sent")
	}
}

func AddCompanyTxn(id, info string) error {

	// Checking zone
	if !existCompany(id) {
		// Send transaction
		_, err := _einst.CreateCompany(ethTransactor(), id, info)
		return err
	} else {
		return errors.New(AddCompanyAction + ": transaction not sent")
	}
}

func AddContainerTxn(id, info string) error {

	// Checking zone
	if !existContainer(id) {
		// Send transaction
		_, err := _einst.CreateContainer(ethTransactor(), id, info)
		return err
	} else {
		return errors.New(AddContainerAction + ": transaction not sent")
	}
}

func ProductEntryTxn(qty uint64, prodTypeID, coID, ctrID, info string) error {

	// Checking zone
	if existProductType(prodTypeID) && existCompany(coID) && existContainer(ctrID) {
		// Send transaction
		_, err := _tinst.ProductEntry(ethTransactor(), big.NewInt(int64(qty)), prodTypeID, coID, ctrID, info)
		return err
	} else {
		return errors.New(ProductEntryAction + ": transaction not sent")
	}
}

func ProductProcessingTxn(prodID, lostQty uint64, prodTypeID, ctrID, info string) error {

	// Checking zone
	if isProductAvailable(prodID) &&
		lostQty <= GetProductQuantity(prodID) &&
		existProductType(prodTypeID) &&
		existContainer(ctrID) {
		// Send transaction
		_, err := _tinst.ProductProcessing(ethTransactor(), big.NewInt(int64(prodID)), big.NewInt(int64(lostQty)), prodTypeID, ctrID, info)
		return err
	} else {
		return errors.New(ProductProcessingAction + ": transaction not sent")
	}
}

func ProductPartitionTxn(prodID, qty uint64, ctrID, info string) error {

	// Checking zone
	if isProductAvailable(prodID) && qty < GetProductQuantity(prodID) && existContainer(ctrID) {
		// Send transaction
		_, err := _tinst.ProductPartition(ethTransactor(), big.NewInt(int64(prodID)), big.NewInt(int64(qty)), ctrID, info)
		return err
	} else {
		return errors.New(ProductPartitionAction + ": transaction not sent")
	}
}

func ProductOutputTxn(prodID uint64, lotn, coID, ctrID, info string) error {

	// Checking zone
	if isProductAvailable(prodID) && isLotNumberAvailable(lotn) && existCompany(coID) && existContainer(ctrID) {
		// Send transaction
		_, err := _tinst.ProductOutput(ethTransactor(), big.NewInt(int64(prodID)), lotn, coID, ctrID, info)
		return err
	} else {
		return errors.New(ProductOutputAction + ": transaction not sent")
	}
}
