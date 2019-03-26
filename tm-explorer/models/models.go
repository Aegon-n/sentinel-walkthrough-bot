package models

type Responce struct {
	JsonRPC 	string 					`json: "jsonrpc"`
	ID			string					`json: "id"`
	Result   map[string]interface{} 	`json: "result"`
}
