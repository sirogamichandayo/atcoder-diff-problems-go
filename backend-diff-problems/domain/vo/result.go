package vo

import "fmt"

type Result string

const (
	Ce  Result = "CE"
	Mle Result = "MLE"
	Tle Result = "TLE"
	Re  Result = "RE"
	Ole Result = "OLE"
	Ie  Result = "IE"
	Wa  Result = "WA"
	Ac  Result = "AC"
	Wj  Result = "WJ"
	Wr  Result = "WR"
)

func (r Result) IsAc() bool {
	return r == Ac
}

func (r Result) Valid() error {
	switch r {
	case Ce, Mle, Tle, Re, Ole, Ie, Wa, Ac, Wj, Wr:
		return nil
	default:
		return fmt.Errorf("invalid result : " + string(r))
	}
}

func ParseResult(s string) (r Result, err error) {
	r = Result(s)
	err = r.Valid()
	return
}
