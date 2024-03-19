package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateShortLink(t *testing.T) {
	type args struct {
		initialLink string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nominal case",
			args: args{
				initialLink: "https://medium.com/equify-tech/the-three-fundamental-stages-of-an-engineering-career-54dac732fc74",
			},
			want: "MTAwMTMx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sha256LinkGenerator{}
			got := s.GenerateShortLink(tt.args.initialLink)
			assert.Equal(t, tt.want, got)
		})
	}
}
