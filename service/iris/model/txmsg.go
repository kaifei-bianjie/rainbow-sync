package iris

import (
	"github.com/irisnet/rainbow-sync/service/iris/constant"
	"github.com/irisnet/irishub/app/v1/gov"
	"strconv"
)

type Param struct {
	Subspace string `bson:"subspace"`
	Key      string `bson:"key"`
	Value    string `bson:"value"`
}

type Params []Param

type DocTxMsgSubmitProposal struct {
	Title          string `bson:"title"`          //  Title of the proposal
	Description    string `bson:"description"`    //  Description of the proposal
	Proposer       string `bson:"proposer"`       //  Address of the proposer
	InitialDeposit Coins  `bson:"initialDeposit"` //  Initial deposit paid by sender. Must be strictly positive.
	ProposalType   string `bson:"proposalType"`   //  Initial deposit paid by sender. Must be strictly positive.
	Params         Params `bson:"params"`
}

type DocTxMsgSubmitTokenAdditionProposal struct {
	DocTxMsgSubmitProposal
	Symbol          string `bson:"symbol"`
	CanonicalSymbol string `bson:"canonical_symbol"`
	Name            string `bson:"name"`
	Decimal         uint8  `bson:"decimal"`
	MinUnitAlias    string `bson:"min_unit_alias"`
	InitialSupply   uint64 `bson:"initial_supply"`
}

func (doctx *DocTxMsgSubmitTokenAdditionProposal) Type() string {
	return constant.Iris_TxTypeSubmitTokenAdditionProposal
}

func (doctx *DocTxMsgSubmitTokenAdditionProposal) BuildMsg(txMsg interface{}) {
	msg := txMsg.(MsgSubmitTokenAdditionProposal)
	doctx.Title = msg.Title
	doctx.Description = msg.Description
	doctx.ProposalType = msg.ProposalType.String()
	doctx.Proposer = msg.Proposer.String()
	doctx.Params = loadParams(msg.Params)
	doctx.InitialDeposit = loadInitialDeposit(msg.InitialDeposit)
	doctx.Symbol = msg.Symbol
	doctx.MinUnitAlias = msg.MinUnitAlias
	doctx.CanonicalSymbol = msg.CanonicalSymbol
	doctx.Name = msg.Name
	doctx.Decimal = msg.Decimal
	doctx.InitialSupply = msg.InitialSupply
}

func loadParams(params []gov.Param) (result []Param) {
	for _, val := range params {
		result = append(result, Param{Subspace: val.Subspace, Value: val.Value, Key: val.Key})
	}
	return
}

func loadInitialDeposit(coins SdkCoins) (result Coins) {
	for _, val := range coins {
		amt, _ := strconv.ParseFloat(val.Amount.String(), 64)
		result = append(result, &Coin{Amount: amt, Denom: val.Denom})
	}
	return
}

type DocTxMsgSetMemoRegexp struct {
	Owner      string `bson:"owner"`
	MemoRegexp string `bson:"memo_regexp"`
}

func (doctx *DocTxMsgSetMemoRegexp) Type() string {
	return constant.Iris_TxTypeSetMemoRegexp
}

func (doctx *DocTxMsgSetMemoRegexp) BuildMsg(txMsg interface{}) {
	msg := txMsg.(MsgSetMemoRegexp)
	doctx.MemoRegexp = msg.MemoRegexp
	doctx.Owner = msg.Owner.String()
}

type DocTxMsgRequestRand struct {
	Consumer      string `bson:"consumer"`       // request address
	BlockInterval uint64 `bson:"block-interval"` // block interval after which the requested random number will be generated
}

func (doctx *DocTxMsgRequestRand) Type() string {
	return constant.Iris_TxTypeRequestRand
}

func (doctx *DocTxMsgRequestRand) BuildMsg(txMsg interface{}) {
	msg := txMsg.(MsgRequestRand)
	doctx.Consumer = msg.Consumer.String()
	doctx.BlockInterval = msg.BlockInterval
}
