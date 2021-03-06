package wotsp

import (
	"fmt"
)

// Mode constants specify internal parameters according to the given mode of
// operation. The available parameter sets include w = 4 and w = 16. The
// default, which is used when no explicit mode is chosen, is w = 16. This
// allows the default Mode to be selected by specifying wotsp.Mode(0).
//
// See RFC 8391 for details on the different parameter sets.
type Mode int

const (
	// W16 indicates the parameter set of WOTSP where w = 16. W16 is the default
	// mode.
	//
	// Passing W16 to Opts opts is equivalent to passing Mode(0), or not setting
	// the Mode at all.
	W16 Mode = iota

	// W4 indicates the parameter set of WOTSP where w = 4.
	W4
)

// params construct a modeParams instance based on the operating Mode, or an
// error if the mode is not valid.
func (m Mode) params() (p params, err error) {
	switch m {
	case W4:
		p.w = 4
		p.logW = 2
		p.l1 = 128
		p.l2 = 5
	case W16:
		p.w = 16
		p.logW = 4
		p.l1 = 64
		p.l2 = 3
	default:
		err = fmt.Errorf("invalid mode %s, must be either wotsp.W4 or wotsp.W16", m)
		return
	}

	p.l = p.l1 + p.l2
	return
}

// String implements fmt.Stringer.
func (m Mode) String() string {
	switch m {
	case W4:
		return "W4"
	case W16:
		return "W16"
	default:
		return fmt.Sprintf("<invalid %d>", m)
	}
}
