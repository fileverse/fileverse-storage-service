package upload

import (
	"context"
	"fmt"
	"mime/multipart"
	"storage/src/pkg/logger"
	"storage/src/services/dbservice"
	"storage/src/services/ipfs/pinata"

	"github.com/spf13/viper"
)

func CanUpload(ctx context.Context, size int64) (bool, error) {
	log := logger.GetContextLogger(ctx)
	if ctx.Value("isAuthenticated") == false {
		return false, fmt.Errorf("unauthorized to upload file")
	}

	contract := ctx.Value("contract").(string)
	limit := dbservice.GetLimit(ctx, contract)

	if limit.StorageUse+size > limit.StorageLimit {
		log.Debug("storage limit exceeded", "limit", limit.StorageLimit, "use", limit.StorageUse, "size", size)
		return false, fmt.Errorf("storage limit exceeded")
	}

	return true, nil
}

func UploadFile(ctx context.Context, file *multipart.FileHeader, tags []string) (resp UploadFileResp, err error) {
	log := logger.GetContextLogger(ctx)

	// step 1: upload file to ipfs
	pinataResp, err := pinata.UploadFileToIPFS(ctx, file)
	if err != nil {
		log.Error("failed to upload file to ipfs", "error", err)
		return
	}

	contract := ctx.Value("contract").(string)

	err = dbservice.UpdateLimit(ctx, contract, pinataResp.PinSize)
	if err != nil {
		log.Error("failed to update limit", "error", err)
	}

	fileverseGateway := viper.GetString("fileverse_gateway")
	fileObj := dbservice.Files{
		InvokerAddress:  ctx.Value("invoker").(string),
		ContractAddress: contract,
		GatewayURL:      fileverseGateway + pinataResp.IpfsHash,
		ChainID:         ctx.Value("chainId").(string),
		IpfsHash:        pinataResp.IpfsHash,
		FileSize:        pinataResp.PinSize,
		Tags:            tags,
		Namespace:       ctx.Value("namespace").(string),
	}

	err = fileObj.Create(ctx)
	if err != nil {
		log.Error("failed to save file metadata", "error", err)
		return
	}

	resp = UploadFileResp{
		IpfsUrl:         fileverseGateway + pinataResp.IpfsHash,
		IpfsHash:        pinataResp.IpfsHash,
		IpfsStorage:     pinataResp.IpfsStorage,
		FileSize:        pinataResp.PinSize,
		Mimetype:        file.Header.Get("Content-Type"),
		ContractAddress: contract,
	}
	return
}
