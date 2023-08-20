package output

import "testing"

func TestOutput(t *testing.T) {
	type args struct {
		multiline []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "output test1",
			args: args{
				multiline: []string{"hello", "my name", "hokasonja"},
			},
			want: "hello\nmy name\nhokasonja\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Output(tt.args.multiline)
			if tt.want != got {
				t.Fatalf("want %s, got %s\n", tt.want, got)
			}
		})
	}
}
