package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/fabienogli/stoik/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSlugifier_RetrieveURL(t *testing.T) {
	type fields struct {
		retrieverSender func(t *testing.T) retrieverSender
	}
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
	}{
		{
			name: "url retriever err",
			fields: fields{
				retrieverSender: func(t *testing.T) retrieverSender {
					m := mocks.NewRetrieverSender(t)
					m.On("RetrieveURL", mock.Anything, "mySlug").Return("", fmt.Errorf("my err: %w", assert.AnError))
					return m
				},
			},
			args: args{
				slug: "mySlug",
			},
			want:    "",
			wantErr: assert.AnError,
		},
		{
			name: "no error",
			fields: fields{
				retrieverSender: func(t *testing.T) retrieverSender {
					m := mocks.NewRetrieverSender(t)
					m.On("RetrieveURL", mock.Anything, "mySlug").Return("https://google.com", nil)
					return m
				},
			},
			args: args{
				slug: "mySlug",
			},
			want:    "https://google.com",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Slugifier{
				retrieverSender: tt.fields.retrieverSender(t),
			}
			got, err := s.RetrieveURL(context.Background(), tt.args.slug)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSlugifier_CreateSlug(t *testing.T) {
	type fields struct {
		retrieverSender func(t *testing.T) retrieverSender
		hasher          func(t *testing.T) hasher
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
	}{
		{
			name: "error while creating Slug",
			fields: fields{
				retrieverSender: func(t *testing.T) retrieverSender {
					m := mocks.NewRetrieverSender(t)
					m.On("CreateSlug", mock.Anything, "https://google.com", "mySlug").Return(fmt.Errorf("my err: %w", assert.AnError))
					return m
				},
				hasher: func(t *testing.T) hasher {
					m := mocks.NewHasher(t)
					m.On("GenerateShortLink", "https://google.com").Return("mySlug")
					return m
				},
			},
			args: args{
				url: "https://google.com",
			},
			want:    "",
			wantErr: assert.AnError,
		},
		{
			name: "nominal case",
			fields: fields{
				retrieverSender: func(t *testing.T) retrieverSender {
					m := mocks.NewRetrieverSender(t)
					m.On("CreateSlug", mock.Anything, "https://google.com", "mySlug").Return(nil)
					return m
				},
				hasher: func(t *testing.T) hasher {
					m := mocks.NewHasher(t)
					m.On("GenerateShortLink", "https://google.com").Return("mySlug")
					return m
				},
			},
			args: args{
				url: "https://google.com",
			},
			want:    "mySlug",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Slugifier{
				retrieverSender: tt.fields.retrieverSender(t),
				hasher:          tt.fields.hasher(t),
			}
			got, err := s.CreateSlug(context.Background(), tt.args.url)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
