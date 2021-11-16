package test

import (
	"testing"

	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/service"
	"github.com/stretchr/testify/assert"
)

func Test_Translate(t *testing.T) {
	service := service.ProvideMeowTranslatorService()

	type args struct {
		body string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Returns meow when body is letter",
			args: args{body: "A"},
			want: "Translated to meow language: meow (translated from leather language: A)",
		},
		{
			name: "Returns purr when body is digit",
			args: args{body: "1"},
			want: "Translated to meow language: purr (translated from leather language: 1)",
		},
		{
			name: "Returns sniff when body is not letter and not digit",
			args: args{body: "$"},
			want: "Translated to meow language: sniff (translated from leather language: $)",
		},
		{
			name: "Returns meow purr sniff when body contains letter, digit and another one",
			args: args{body: "A1$"},
			want: "Translated to meow language: meow purr sniff (translated from leather language: A1$)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := service.Translate(tt.args.body)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Benchmark_Translate(b *testing.B) {
	service := service.ProvideMeowTranslatorService()

	for i := 0; i < b.N; i++ {
		service.Translate("")
	}
}
