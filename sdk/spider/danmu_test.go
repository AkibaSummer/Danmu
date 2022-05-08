package spider

import (
	"github.com/AkibaSummer/Danmu/sdk/utils"
	"testing"
)

func TestDanmuSpider_Init(t *testing.T) {
	utils.Init()
	_ = NewDanmuSpider(21144080)
}
