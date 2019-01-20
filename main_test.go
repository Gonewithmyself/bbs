package main

import (
	_ "bbs/routers"
	"bbs/spider"
	"testing"
)

func Test_main(t *testing.T) {

	t.Parallel()
	spider.TransWords()
	t.Error("123")
}
