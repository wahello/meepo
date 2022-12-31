package well_known_option

import "github.com/PeerXu/meepo/pkg/internal/option"

const OPTION_KCP_KEY = "kcpKey"

var WithKcpKey, GetKcpKey = option.New[string](OPTION_KCP_KEY)
