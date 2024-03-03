package core

import (
	"context"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/swarleynunez/INVINOS/core/utils"
	"os"
)

func connectToIPFS() (ipfsc *rpc.HttpApi) {

	ipfsc, err := rpc.NewLocalApi()
	utils.CheckError(err, utils.FatalMode)

	return
}

func uploadIPFSFile(ctx context.Context, filename string) (string, error) {

	// Get tmp file from fs
	filepath := utils.FormatPath(utils.GetEnv("TMP_DIR"), filename)
	node, err := getUnixfsNode(filepath)
	if err != nil {
		return "", err
	}

	// Upload file to the IPFS network
	p, err := _ipfsc.Unixfs().Add(ctx, node)
	if err != nil {
		return "", err
	}

	// Delete tmp file from fs
	_ = os.Remove(filepath)

	return p.RootCid().String(), nil
}

func DownloadIPFSFile(ctx context.Context, strCid, filepath string) error {

	// Decode string CID into a CID object
	c, err := cid.Decode(strCid)
	if err != nil {
		return err
	}

	// Get file from the IPFS network
	node, err := _ipfsc.Unixfs().Get(ctx, path.FromCid(c))
	if err != nil {
		return err
	}

	// Write IPFS file to fs
	err = files.WriteTo(node, filepath)
	if err != nil {
		return err
	}

	return nil
}

func getUnixfsNode(path string) (files.Node, error) {

	st, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	f, err := files.NewSerialFile(path, false, st)
	if err != nil {
		return nil, err
	}

	return f, nil
}
