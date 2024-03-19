package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type slugCreator interface {
	CreateSlug(ctx context.Context, url string) (string, error)
}

type SlugPoster struct {
	slugCreator  slugCreator
	serverDomain string
}

func NewSlugPoster(slugCreator slugCreator, serverDomain string) *SlugPoster {
	return &SlugPoster{
		slugCreator:  slugCreator,
		serverDomain: serverDomain,
	}
}

func (s *SlugPoster) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query createQuery
		err := c.ShouldBindJSON(&query)
		if err != nil {
			//we do not care for error since it is for an example
			data, _ := json.Marshal(createQuery{
				Url: "example-url",
			})
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("malformed request: should be %s", data))
			return
		}
		if query.Url == "" {
			//we do not care for error since it is for an example
			data, _ := json.Marshal(createQuery{
				Url: "example-url",
			})
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("malformed request: should be %s", data))
			return
		}
		slug, err := s.slugCreator.CreateSlug(c.Request.Context(), query.Url)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusCreated, response{
			Url:     query.Url,
			TinyURL: fmt.Sprintf("https://%s/%s", s.serverDomain, slug),
		})
	}
}
