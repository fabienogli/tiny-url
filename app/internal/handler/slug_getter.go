package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/fabienogli/stoik/internal/domain"
	"github.com/gin-gonic/gin"
)

type slugRetriever interface {
	RetrieveURL(ctx context.Context, slug string) (string, error)
}

type SlugGetter struct {
	slugRetriever slugRetriever
}

func NewSlugGetter(slugRetriever slugRetriever) *SlugGetter {
	return &SlugGetter{
		slugRetriever: slugRetriever,
	}
}

func (s *SlugGetter) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")
		if slug == "" {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no slug given"))
			return
		}
		url, err := s.slugRetriever.RetrieveURL(c.Request.Context(), slug)
		if errors.Is(err, domain.ErrNoResult) {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Redirect(http.StatusMovedPermanently, url)
		return
	}
}
