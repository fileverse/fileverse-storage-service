package upload

type UploadFileResp struct {
	IpfsUrl         string `json:"ipfsUrl,omitempty"`
	IpfsHash        string `json:"ipfsHash,omitempty"`
	IpfsStorage     string `json:"ipfsStorage,omitempty"`
	CachedUrl       string `json:"cachedUrl,omitempty"`
	FileSize        int64  `json:"fileSize,omitempty"`
	Mimetype        string `json:"mimetype,omitempty"`
	FileId          string `json:"fileId,omitempty"`
	ContractAddress string `json:"contractAddress,omitempty"`
}
