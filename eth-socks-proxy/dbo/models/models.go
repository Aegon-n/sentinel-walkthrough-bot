package models

type Nodes struct {
	NodesList   []List `json:"list"`
}
type Location struct {
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Country   string  `json:"country"`
}
type NetSpeed struct {
	Download float64 `json:"download"`
	Upload   float64 `json:"upload"`
}
type TONNode struct {
	Location       Location `json:"location"`
	NetSpeed       NetSpeed `json:"netSpeed"`
	APIPort        int      `json:"APIPort"`
	PricePerGB     int      `json:"pricePerGB"`
	Description    string   `json:"description"`
	RatingPoints   int      `json:"ratingPoints"`
	RatingCount    int      `json:"ratingCount"`
	AccountAddress string   `json:"accountAddress"`
	IP             string   `json:"IP"`
	EncMethod      string   `json:"encMethod"`
	NodeType       string   `json:"nodeType"`
	Version        string   `json:"version"`
	TxHash         string   `json:"txHash"`
}

type KV struct {
	Key   string
	Value string
}

type User struct {
	Node             string
	TelegramUsername string
	ChatID					 string
	Token						 string					
}

type TxReceipt struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	Timestamp        string   `json:"timestamp"`
	GasPrice         string   `json:"gasPrice"`
	GasUser          string   `json:"gasUsed"`
	LogIndex         string   `json:"logIndex"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

type TXDetails struct {
	Result struct {
		BlockHash        string `json:"blockHash"`
		BlockNumber      string `json:"blockNumber"`
		From             string `json:"from"`
		Gas              string `json:"gas"`
		GasPrice         string `json:"gasPrice"`
		Hash             string `json:"hash"`
		Input            string `json:"input"`
		Nonce            string `json:"nonce"`
		To               string `json:"to"`
		TransactionIndex string `json:"transactionIndex"`
		Value            string `json:"value"`
		V                string `json:"v"`
		R                string `json:"r"`
		S                string `json:"s"`
	} `json:"result"`
}

type TMTxn struct {
	Tx struct {
		Value struct {
			Msg []TMMsg `json:"msg"`
			Fee struct {
				Amount []TMAmount `json:"amount"`
			} `json:"fee"`
		} `json:"value"`
	} `json:"tx"`
}

type TMMsg struct {
	Type  string `json:"type"`
	Value struct {
		Coins   []TMAmount `json:"Coins"`
		From    string     `json:"From"`
		To      string     `json:"To"`
		Address string     `json:"address"`
	} `json:"value"`
}

type TMAmount struct {
	Denom  string
	Amount string
}

type TxReceiptList struct {
	Results []TxReceipt `json:"result"`
	Status  string      `json:"status"`
}

type GeoLocation struct {
	//As string `json:"as"`
	City        string `json:"city_name"`
	Country     string `json:"country_name"`
	CountryCode string `json:"country_code"`
	//Isp string `json:"isp"`
	//Lat float64 `json:"lat"`
	//Lon float64 `json:"lon"`
	//Query string `json:"query"`
	//Region string `json:"region"`
	RegionName string `json:"region_name"`
	//Status string `json:"status"`
	//TimeZone string `json:"timezone"`
	//Zip string `json:"zip"`
}



type ExpiredUsers struct {
	Key   string
	Value string
}

type UserResp struct {
	Message  string `json:"messages"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type InlineButtonOptions struct {
	Label string
	URL   string
}

type ButtonHelper struct {
	Type               string
	Labels             []string
	InlineKeyboardOpts []InlineButtonOptions
}

type Load struct {
	CPU    float64 `json:"cpu"`
	Memory int     `json:"memory"`
}

type VpnResponse struct {
	Node        Node   `json:"node"`
	SessionName string `json:"session_name"`
	Success     bool   `json:"success"`
}
type Vpn struct {
	Username			string	`json: "username"`
	Password 			string	`json: "password"`
	TelegramLink	string 	`json:"telegram_link"`
}
type BestServer struct {
	Latency float64 `json:"latency"`
	Host    string  `json:"host"`
}

type Node struct {
	Vpn      Vpn      `json:"vpn"`
	NetSpeed NodeNetSpeed `json:"net_speed"`
	Location Location `json:"location"`
}

type NodeNetSpeed struct {
	Download   float64    `json:"download"`
	BestServer BestServer `json:"best_server"`
	Upload     float64    `json:"upload"`
}

type List struct {
	AccountAddr       string   `json:"account_addr"`
	IP                string   `json:"ip"`
	Latency           float64  `json:"latency"`
	VpnType           string   `json:"vpn_type"`
	Location          Location `json:"location"`
	NetSpeed          NetSpeed `json:"net_speed"`
	EncMethod         string   `json:"enc_method"`
	Version           string   `json:"version"`
	ActiveConnections int      `json:"active_connections"`
	Load              Load     `json:"load"`
	Rating            float64  `json:"rating,omitempty"`
	PricePerGB        float64  `json:"price_per_GB"`
	Moniker           string   `json:"moniker,omitempty"`
	Description       string   `json:"description,omitempty"`
}

type SocksResponse struct {
	Success bool   `json:"success"`
	List    []List `json:"list"`
}
type MasterResponce struct {
	Success bool   `json:"success"`
	IP      string `json:"ip"`
	Port    int    `json:"port"`
	Token   string `json:"token"`
	VpnAddr string `json:"vpn_addr"`
	Message string `json:"message"`
}
type VpnUsage struct {
	Usage   Usage `json:"usage"`
}
type Usage struct {
	Down float64 `json:"down"`
	Up   float64 `json:"up"`
}

type LimitResponse struct {
	ClientList []string `json:"client_list"`
}