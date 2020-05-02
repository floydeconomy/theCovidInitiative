package x

import (
	"errors"
	"github.com/floydeconomy/arisaedo-go/common"
	"sync/atomic"
)

// Case represents structure of COVID-19 cases
type Case struct {
	// body
	header *Header

	// caches
	cache struct {
		identifier atomic.Value
	}
}

// Cases represents a variety of cases
type Cases []Case

//// NewCase creates a new cases
//func Compose(header *Header, body *Body) *Case {
//	c := &Case{header: CopyHeader(header)}
//	return c
//}

func (c *Case) Header() *Header { return c.header }


type Nonce [8]byte

// Header represents the entire COVID-19 case/death related data
type Header struct {
	// body
	body Body

	// cache
	cache struct {
		identifier atomic.Value
		country    atomic.Value
		province   atomic.Value
	}
}

// Body represents the cases/death related to COVID-19
// todo: implement signature and nonce fields
type Body struct {
	// IPFS Identifiers
	CountryID  common.Identifier `json:"country"`
	ProvinceID common.Identifier `json:"province"`

	// Case
	Time      uint64 `json:"time"`
	Confirmed uint64 `json:"confirmed"`
	Death     uint64 `json:"death"`
	Recovered uint64 `json:"recovered"`
	Active    uint64 `json:"active"`
}

func (h *Header) Confirmed() uint64             { return h.body.Confirmed }
func (h *Header) Death() uint64                 { return h.body.Death }
func (h *Header) Recovered() uint64             { return h.body.Recovered }
func (h *Header) Active() uint64                { return h.body.Active }
func (h *Header) Time() uint64                  { return h.body.Time }
func (h *Header) CountryID() common.Identifier  { return h.body.CountryID }
func (h *Header) ProvinceID() common.Identifier { return h.body.ProvinceID }
func (h *Header) Body() Body                    { return h.body }

// SanityCheck checks whether the case is valid
func (h *Header) SanityCheck() error {
	if h.Time() <= 0 {
		return errors.New("invalid time")
	}

	// todo: fix this
	if h.CountryID() == "" {
		return errors.New("country id doesn't exists")
	}

	// ret: all passed
	return nil
}
