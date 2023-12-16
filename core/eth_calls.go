package core

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/swarleynunez/INVINOS/core/types"
	"github.com/swarleynunez/INVINOS/core/utils"
	"math/big"
)

// ////////////
// Checkers //
// ////////////
func existProductType(id string) (r bool) {

	r, err := _einst.ExistProductType(&bind.CallOpts{}, id)
	utils.CheckError(err, utils.WarningMode)

	return
}

func existCompany(id string) (r bool) {

	r, err := _einst.ExistCompany(&bind.CallOpts{}, id)
	utils.CheckError(err, utils.WarningMode)

	return
}

func existContainer(id string) (r bool) {

	r, err := _einst.ExistContainer(&bind.CallOpts{}, id)
	utils.CheckError(err, utils.WarningMode)

	return
}

func isProductAvailable(prodID uint64) (r bool) {

	r, err := _tinst.IsProductAvailable(&bind.CallOpts{}, big.NewInt(int64(prodID)))
	utils.CheckError(err, utils.WarningMode)

	return
}

func isLotNumberAvailable(lotn string) (r bool) {

	r, err := _tinst.IsLotNumberAvailable(&bind.CallOpts{}, lotn)
	utils.CheckError(err, utils.WarningMode)

	return
}

// ///////////
// Getters //
// ///////////
func GetProduct(prodID uint64) (p types.Product) {

	p, err := _tinst.Products(&bind.CallOpts{}, big.NewInt(int64(prodID)))
	utils.CheckError(err, utils.WarningMode)

	return
}

func GetProductQuantity(prodID uint64) uint64 {

	qty, err := _tinst.GetProductQuantity(&bind.CallOpts{}, big.NewInt(int64(prodID)))
	utils.CheckError(err, utils.WarningMode)

	return qty.Uint64()
}

func GetProductFromLotNumber(lotn string) uint64 {

	prodID, err := _tinst.LotNumbers(&bind.CallOpts{}, lotn)
	utils.CheckError(err, utils.WarningMode)

	return prodID.Uint64()
}

func GetTransitionTypeName(tranID uint64) string {

	tt, err := _ainst.TransitionTypes(&bind.CallOpts{}, big.NewInt(int64(tranID)))
	utils.CheckError(err, utils.WarningMode)

	return tt.Info
}
