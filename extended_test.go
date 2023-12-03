package decimal

import (
	"testing"
)

func TestNewFrom10k(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name string
		args args
		want Decimal
	}{
		{
			name: "6.52",
			args: args{value: 65200},
			want: NewFromFloat(6.52),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFrom10k(tt.args.value); !got.Equal(tt.want) {
				t.Errorf("NewFrom10k() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal_Int10k(t *testing.T) {
	tests := []struct {
		name  string
		value Decimal
		want  int64
	}{
		{
			name:  "6.52",
			value: NewFromFloat(6.52),
			want:  65200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.value.Int10k(); got != tt.want {
				t.Errorf("Decimal.Int10k() = %v, want %v", got, tt.want)
			}
		})
	}
}
