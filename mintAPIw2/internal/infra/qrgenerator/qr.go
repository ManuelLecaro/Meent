package qrgenerator

import (
	"context"
	"encoding/json"
	"fmt"
	"image/color"

	"github.com/skip2/go-qrcode"
)

type Generator struct {
	QR_DIR string
}

func  NewGenerator(dir string) *Generator {
	return &Generator{
		QR_DIR: dir,
	}
}

func (g *Generator) CreateQR(ctx context.Context, data map[string]interface{}, name string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error serializing map to JSON:", err)
		return err
	}

	return qrcode.WriteColorFile(
		string(jsonData),
		qrcode.Medium,
		256,
		color.RGBA{0, 0, 255, 255},
		color.RGBA{0, 255, 0, 255},
		g.nameGenerator(name),
	)
}

func (g *Generator) nameGenerator(name string) string {
	return fmt.Sprintf("%s/%s", g.QR_DIR, name)
}
