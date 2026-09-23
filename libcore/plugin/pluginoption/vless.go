package pluginoption

import (
	"encoding/json"

	"github.com/sagernet/sing-box/option"
)

type VLESSOutboundOptions struct {
	option.VLESSOutboundOptions
	Encryption string          `json:"encryption,omitempty" examples:"none"`
	Finalmask  json.RawMessage `json:"finalmask,omitempty"`
}
