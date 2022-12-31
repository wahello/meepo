package well_known_option

import "github.com/PeerXu/meepo/pkg/internal/option"

const OPTION_MUX_NOCOMP = "muxNocomp"

var WithMuxNocomp, GetMuxNocomp = option.New[bool](OPTION_MUX_NOCOMP)
