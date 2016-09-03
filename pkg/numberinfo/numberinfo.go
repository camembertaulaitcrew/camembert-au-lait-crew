package calcnumberinfo

import "github.com/moul/numberinfo"

type CALCNumberInfo struct {
	input  float64
	number *numberinfo.Float64Number
}

func New(number float64) *CALCNumberInfo {
	return &CALCNumberInfo{
		input:  number,
		number: numberinfo.Float64(number),
	}
}

func (n *CALCNumberInfo) All() map[string]interface{} {
	ret := map[string]interface{}{}
	ret["number"] = n.input
	ret["is-prime"] = n.number.IsPrime()
	return ret
}