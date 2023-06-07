package ipfs

import (
	"context"
	"encoding/json"
	"fmt"
	"mintapi/internal/core/domain"
	"mintapi/internal/core/port/resource"
	"mintapi/internal/infra/config"
	"mintapi/internal/infra/qrgenerator"
	"os"
	"path/filepath"

	shell "github.com/ipfs/go-ipfs-api"
)

const ExternalURLData = "https://www.youtube.com/watch?v=dQw4w9WgXcQ&pp=ygUIcmlja3JvbGw%3D"

type Trait struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
}

type Metadata struct {
	ID          uint64  `json:"_"`
	Description string  `json:"description"`
	ExternalURL string  `json:"external_url"`
	Image       string  `json:"image_url"`
	Name        string  `json:"name"`
	Attributes  []Trait `json:"attributes"`
}

type IPFSService struct {
	shell       *shell.Shell
	qrGenerator resource.QRGenerator
}

func NewIPFSService() resource.Minter {
	return &IPFSService{
		shell:       shell.NewShell(config.GetIPFSURL()),
		qrGenerator: qrgenerator.NewGenerator(config.GetQRDirectory()),
	}
}

func (s *IPFSService) UploadFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Subir el archivo al nodo de IPFS
	cid, err := s.shell.Add(file)
	if err != nil {
		return "", err
	}

	// Construir la URL del archivo utilizando el CID
	url := fmt.Sprintf("https://ipfs.io/ipfs/%s", cid)

	return url, nil
}

func (s *IPFSService) UploadToIPFS(ctx context.Context, data *domain.Event) (string, error) {
	dataForQR := map[string]interface{}{
		"event": data.Name,
	}

	if err := s.qrGenerator.CreateQR(ctx, dataForQR, fmt.Sprintf("%d.png", data.Id)); err != nil {
		return "", err
	}

	dirPath := fmt.Sprintf("%s/%s", config.GetQRDirectory(), fmt.Sprintf("%d.png", data.Id))

	qrURI, err := s.UploadFile(dirPath)
	if err != nil {
		return "", err
	}

	metadata := Metadata{
		ID:          data.Id,
		Description: "TicketQR Data",
		ExternalURL: ExternalURLData,
		Image:       qrURI,
		Name:        data.Name,
		Attributes:  []Trait{{TraitType: "ID", Value: data.Id}},
	}

	pathFileMD, err := saveStructAsJSON(metadata, config.GetQRDirectory())
	if err != nil {
		return "", err
	}

	return s.UploadFile(pathFileMD)
}

func saveStructAsJSON(data Metadata, directoryPath string) (string, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("ticket_metadata_%d.json", data.ID)

	filePath := filepath.Join(directoryPath, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
