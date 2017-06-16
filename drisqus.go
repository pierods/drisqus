// Package drisqus is a high-level wrapper over Disqus' API
package drisqus

import (
	"github.com/pierods/gisqus"
)

// Drisqus is the entry point of the library
type Drisqus struct {
	gisqus gisqus.Gisqus
}

// NewDrisqus returns an instance of Drisqus, given an instance of gisqus.Gisqus
func NewDrisqus(g gisqus.Gisqus) Drisqus {
	return Drisqus{
		gisqus: g,
	}
}
