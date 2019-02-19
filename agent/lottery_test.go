package agent

import "testing"

func TestGetHeader(t *testing.T) {
	t.Error(LoadHeader("conf.ini"))
}

func TestLottery(t *testing.T) {
	GetLotteryInfo()
	t.Error("")
}

func TestLotteryAnalyse(t *testing.T) {
	AnalyseLottery()
	"01 14 18 20 22 26 12"
	"01 14 18 20 22 26 09"
	"03 05 06 10 13 19 15"
	"02 05 06 21 27 28 16"
	"02 04 11 12 21 23 13"
	t.Error("")
}
