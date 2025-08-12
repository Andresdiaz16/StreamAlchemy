package obs

// TODO: Add the authentication field
// right now it assumes that the websocket plugins doesnt have authentication

/*  names for properties directly taken from: 
	https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#message-types-opcodes
*/
type ConnectionResponse struct {
	D  D   `json:"d"`
	Op int `json:"op"`
}

/*  names for properties directly taken from: 
	https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#message-types-opcodes
*/
type ObsResponse struct {
	D  D   `json:"d"`
	Op int `json:"op"`
}
