package file

import (
	"testing"
)

func TestResize(t *testing.T) {
	type args struct {
		filename string
	}

	type size struct {
		height int
		width  int
	}

	tests := []struct {
		name    string
		args    args
		want    size
		wantErr bool
	}{
		{
			name:    "resize test",
			args:    args{filename: "../../public/images/test2.png"},
			want:    size{width: 100, height: 100},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := Upload(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Fatalf("wantErr %t, gotErr %t\n", tt.wantErr, err != nil)
			}

			newImg := Resize(img)
			got := size{
				height: newImg.Bounds().Dy(),
				width:  newImg.Bounds().Dx(),
			}
			if tt.want.height != got.height || tt.want.width != got.width {
				t.Fatalf("want size (%d, %d), got (%d, %d)\n", tt.want.height, tt.want.width, got.height, got.width)
			}
		})
	}
}
