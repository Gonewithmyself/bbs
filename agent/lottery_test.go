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
	t.Error("")
}
