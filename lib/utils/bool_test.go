package util

import "testing"

func TestXOR(t *testing.T) {

	type args struct {
		src    bool
		target bool
	}

	tests := []struct {
		name  string
		args  args
		wantV bool
	}{
		{
			name: "true * true",
			args: args{
				true,
				true,
			},
			wantV: false,
		},
		{
			name: "true * false",
			args: args{
				true,
				false,
			},
			wantV: true,
		},
		{
			name: "false * true",
			args: args{
				false,
				true,
			},
			wantV: true,
		},
		{
			name: "false * false",
			args: args{
				src:    false,
				target: false,
			},
			wantV: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := XOR(tt.args.src, tt.args.target); gotV != tt.wantV {
				t.Errorf("XOR() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}
