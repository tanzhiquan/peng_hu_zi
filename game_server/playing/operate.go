package playing

import (
	"peng_hu_zi/game_server/card"
	"fmt"
)

type OperateType int

const (
	OperateGetInitCards	OperateType = iota + 1
	OperateDispatchCard

	OperateEnterRoom
	OperateLeaveRoom

	OperateDropCard
	OperateChiCard
	OperatePengCard
	OperateSaoCard
	OperatePaoCard
	OperateTiLongCard
	OperateHu
)

type Operate struct {//玩家操作
	Op			OperateType
	Operator	*Player				//操作者
	Data		interface{}
	Result		chan interface{}
}

func (op *Operate) String() string {
	if op == nil {
		return "{operator=nil, op=nil}"
	}
	return fmt.Sprintf("{operator=%v, op=%v}", op.Operator, op.Op)
}

func newOperate(op OperateType, operator *Player, data interface{}) *Operate{
	return &Operate{
		Op:	op,
		Data: data,
		Operator: operator,
		Result: make(chan interface{}),
	}
}
/*
type OperateGetInitCardsData struct {
	PlayingCards *card.PlayingCards
}
func NewOperateGetInitCards(operator *Player, data *OperateGetInitCardsData) *Operate {
	return newOperate(OperateGetInitCards, operator, data)
}

type OperateDispatchCardData struct {
	Card 	*card.Card
}
func NewOperateDispatchCard(operator *Player, data *OperateDispatchCardData) *Operate {
	return newOperate(OperateDispatchCard, operator, data)
}
*/


type OperateEnterRoomData struct {
}
type OperateEnterRoomResult struct {
	Ok	bool
}
func NewOperateEnterRoom(operator *Player, data *OperateEnterRoomData) *Operate {
	return newOperate(OperateEnterRoom, operator, data)
}

type OperateLeaveRoomData struct {
}
type OperateLeaveRoomResult struct {
	Ok	bool
}
func NewOperateLeaveRoom(operator *Player, data *OperateLeaveRoomData) *Operate {
	return newOperate(OperateLeaveRoom, operator, data)
}


type OperateDropCardData struct{
	Card *card.Card
}
type OperateDropCardResult struct {
	OK bool
}
func NewOperateDropCard(operator *Player, data *OperateDropCardData) *Operate {
	return newOperate(OperateDropCard, operator, data)
}

type OperateChiCardData struct {
	Card 	*card.Card
	Group	*card.Cards
}
type OperateChiCardResultData struct {
	ok bool
}
func NewOperateChiCard(operator *Player, data *OperateChiCardData) *Operate {
	return newOperate(OperateChiCard, operator, data)
}

type OperatePengCardData struct {
	Card *card.Card
}
func NewOperatePengCard(operator *Player, data *OperatePengCardData) *Operate {
	return newOperate(OperatePengCard, operator, data)
}

type OperateSaoCardData struct {
	Card *card.Card
}
func NewOperateSaoCard(operator *Player, data *OperateSaoCardData) *Operate {
	return newOperate(OperateSaoCard, operator, data)
}

type OperatePaoCardData struct {
	Card *card.Card
}
func NewOperatePaoCard(operator *Player, data *OperatePaoCardData) *Operate {
	return newOperate(OperatePaoCard, operator, data)
}

type OperateTiLongCardData struct {
	Card *card.Card
}
func NewOperateTiLongCard(operator *Player, data *OperateTiLongCardData) *Operate {
	return newOperate(OperateTiLongCard, operator, data)
}

type OperateHuData struct {
	HuPlayer		*Player			// 胡牌的玩家
	FromPlayer		*Player			// 牌来源于哪个玩家
	Desc			string			// 胡的描述
	PlayerScore 	[]*PlayerScore	// 玩家的分数
}
func NewOperateHu(operator *Player, data *OperateHuData) *Operate {
	return newOperate(OperateHu, operator, data)
}