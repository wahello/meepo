package well_known_option

import "github.com/PeerXu/meepo/pkg/internal/option"

const OPTION_KCP_SNDWND = "kcpSndwnd"

var WithKcpSndwnd, GetKcpSndwnd = option.New[int](OPTION_KCP_SNDWND)
