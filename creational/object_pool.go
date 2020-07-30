package creational

type PockerCard struct {
	King  string
	Queen string
	Jack  string
	Ten   string
	Nine  string
	Eight string
	Seven string
	Six   string
	Five  string
	Four  string
	Three string
	Two   string
	As    string
}

// NewStraight and NewFullhouse use the same PockerCard object as a parameter
func NewStraight(param *PockerCard) string {
	return param.Three + "," + param.Four + "," + param.Five + "," + param.Six + "," + param.Seven
}

func NewFullhouse(param *PockerCard) string {
	return param.King + "," + param.King + "," + param.King + "," + param.As + "," + param.As
}
