package meta

// 交易类型
const (
	Transfer int = iota // 0: 转账交易
	Register            // 1: 注册账户
	Publish             // 2: 发布合约
	Invoke              // 3: 调用合约
)

type Transaction struct {
	From      string            `json:"from"`
	To        string            `json:"to"`
	Dest      string            `json:"dest"`
	Contract  string            `json:"contract"`
	Method    string            `json:"method"`
	Args      map[string]string `json:"args"`
	Data      TransactionData   `json:"data"`
	Value     int               `json:"value"`
	Id        []byte            `json:"id"`
	Timestamp string            `json:"timestamp"`
	Hash      []byte            `json:"hash"`
	PublicKey string            `json:"public_key"`
	Sign      []byte            `json:"sign"`
	Type      int               `json:"type"`
}

type TransactionData struct {
	Read map[string]string
	Set  map[string]string
	Code string
}

type PostTran struct {
	From       string `json:"from"`
	To         string `json:"to"`
	Dest       string `json:"dest"`
	Contract   string `json:"contract"`
	Method     string `json:"method"`
	Args       string `json:"args"`
	Value      int    `json:"value"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Sign       string `json:"sign"`
	Type       int    `json:"type"`
}

type Block struct {
	Height     int
	Timestamp  string
	PrevHash   []byte
	MerkleRoot []byte
	Signature  []byte
	Hash       []byte
	TX         []Transaction
	StateRoot  []byte
	eventRoot  []byte
}

type BlockHeader struct {
	Height     int
	Timestamp  string
	Hash       []byte
	PrevHash   []byte
	MerkleRoot []byte
}

//跨链注册信息
//一个典型的信息：
// {"Id":"ssbc","Relayers":[{"Id":"","PublicKey":null,"IP":"127.0.0.1","Port":"63042"}],"Servers":[{"Id":"","PublicKey":null,"IP":"127.0.0.1","Port":"63042"}]}
type RegisterInformation struct {
	Id       string //链名
	Relayers []Node //中继节点
	Servers  []Node //服务节点
}
