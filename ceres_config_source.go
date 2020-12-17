package CeresConfig

import (
	"crypto/md5"
	"fmt"
	"time"
)

type Source interface {
	Read() (*DataSet, error)
	Write(*DataSet) error
	IsChanged() <-chan struct{}
	Watch()
	String() string
	UnWatch()
}

func (c *DataSet) Sum() string {
	h := md5.New()
	h.Write(c.Data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

type DataSet struct {
	Data      []byte
	Checksum  string
	Format    string
	Source    string
	Timestamp time.Time
}
