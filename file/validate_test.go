package file

import "testing"

func TestValidate(t *testing.T) {
	type args struct {
		filename string
		want     string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				filename: "test1.png",
				want:     "test1.png",
			},
		},
		{
			name: "test2",
			args: args{
				filename: "test2.jpg",
				want:     "test2.jpg",
			},
		},
		{
			name: "test3",
			args: args{
				filename: "test3.gif",
				want:     "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Validate(tt.args.filename)
			if got != tt.args.want {
				t.Fatalf("want %s, got %s\n", tt.args.want, got)
			}
		})
	}
}
