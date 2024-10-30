package core

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/swarleynunez/INVINOS/core/bindings"
	"github.com/swarleynunez/INVINOS/core/types"
	"github.com/swarleynunez/INVINOS/core/utils"
	"math/big"
)

// ////////////
// Checkers //
// ////////////
func existProductType(einst *bindings.EntityInfo, id string) (r bool) {

	r, err := einst.ExistProductType(&bind.CallOpts{}, id)
	utils.CheckError(err, utils.WarningMode)

	return
}

func existCompany(einst *bindings.EntityInfo, id string) (r bool) {

	r, err := einst.ExistCompany(&bind.CallOpts{}, id)
	utils.CheckError(err, utils.WarningMode)

	return
}

func existContainer(einst *bindings.EntityInfo, id string) (r bool) {

	r, err := einst.ExistContainer(&bind.CallOpts{}, id)
	utils.CheckError(err, utils.WarningMode)

	return
}

func isProductAvailable(tinst *bindings.Traceability, prodID uint64) (r bool) {

	r, err := tinst.IsProductAvailable(&bind.CallOpts{}, big.NewInt(int64(prodID)))
	utils.CheckError(err, utils.WarningMode)

	return
}

func isLotNumberAvailable(tinst *bindings.Traceability, lotn string) (r bool) {

	r, err := tinst.IsLotNumberAvailable(&bind.CallOpts{}, lotn)
	utils.CheckError(err, utils.WarningMode)

	return
}

// ///////////
// Getters //
// ///////////
func GetProduct(tinst *bindings.Traceability, prodID uint64) (p types.Product) {

	p, err := tinst.Products(&bind.CallOpts{}, big.NewInt(int64(prodID)))
	utils.CheckError(err, utils.WarningMode)

	return
}

func GetProductQuantity(tinst *bindings.Traceability, prodID uint64) uint64 {

	qty, err := tinst.GetProductQuantity(&bind.CallOpts{}, big.NewInt(int64(prodID)))
	utils.CheckError(err, utils.WarningMode)

	return qty.Uint64()
}

func GetProductFromLotNumber(tinst *bindings.Traceability, lotn string) uint64 {

	prodID, err := tinst.LotNumbers(&bind.CallOpts{}, lotn)
	utils.CheckError(err, utils.WarningMode)

	return prodID.Uint64()
}

func GetTransitionTypeName(ainst *bindings.Auth, tranID uint64) string {

	tt, err := ainst.TransitionTypes(&bind.CallOpts{}, big.NewInt(int64(tranID)))
	utils.CheckError(err, utils.WarningMode)

	return tt.Info
}
