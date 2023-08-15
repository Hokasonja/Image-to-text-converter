package file

import (
	"testing"
)

func TestUpload(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				filename: "test1.png",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Upload("../public/images/" + tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Fatalf("wantErr %t, gotErr %t\n", tt.wantErr, err != nil)
			}
		})
	}
}
