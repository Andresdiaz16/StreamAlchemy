package obs

type IdentifyOp struct {
	RpcVersion int `json:"rpcVersion"`
}

type AuthIdentifyOp struct {
	RpcVersion int `json:"rpcVersion"`
}

///TODO: it will require to send SceneItemEnabled on other endpoints so be aware

type RequestData struct {
	SceneName        string `json:"sceneName,omitempty"`
	SourceName       string `json:"sourceName,omitempty"`
	SceneItemId      int    `json:"sceneItemId,omitempty"`
	SceneItemEnabled bool   `json:"sceneItemEnabled"`
}

type RequestOp struct {
	RequestType string      `json:"requestType"`
	RequestId   string      `json:"requestId"`
	RequestData RequestData `json:"requestData,omitempty"`
}

/*  names for properties directly taken from: 
	https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#message-types-opcodes
*/
type OBSRequest struct {
	Op int       `json:"op"`
	D  RequestOp `json:"d"`
}

/*  names for properties directly taken from: 
	https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#message-types-opcodes
*/
type OBSIdentify struct {
	Op int        `json:"op"`
	D  IdentifyOp `json:"d"`
}
