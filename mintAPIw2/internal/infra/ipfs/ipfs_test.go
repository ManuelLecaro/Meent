package ipfs

import (
	"context"
	"mintapi/internal/core/domain"
	"mintapi/internal/infra/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPFSService_UploadFile(t *testing.T) {
	config.LoadConfiguration()

	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should work correctly",
			args: args{
				filePath: "../../../files/test/test.json",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewIPFSService().(*IPFSService)

			got, err := s.UploadFile(tt.args.filePath)

			if (err != nil) != tt.wantErr {
				t.Errorf("IPFSService.UploadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("IPFSService.UploadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPFSService_UploadToIPFS(t *testing.T) {
	config.LoadConfiguration()
	type args struct {
		ctx  context.Context
		data *domain.Event
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should work correctly",
			args: args{
				ctx: context.Background(),
				data: &domain.Event{
					Id:   110111,
					Name: "test_concert",
				},
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewIPFSService().(*IPFSService)

			got, err := s.UploadToIPFS(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("IPFSService.UploadToIPFS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, len(got), 67)
		})
	}
}
