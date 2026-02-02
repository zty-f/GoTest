package base

import (
	"fmt"
	"testing"
	"time"
)

type CBD struct {
	StartTs int64 `json:"start_ts"`
}

func TestNowYear(t *testing.T) {
	year := time.Now().Year() - 1
	fmt.Printf("%d\n", year)
	fmt.Println(int32(year))
	_ = "{\"start_ts\":1750003200,\"guide_exp_limit\":250,\"practice_exp_limit\":300,\"practice_exp_per_min\":1,\"guide_exp_once\":25,\"re_guide_exp_limit\":50,\"re_guide_exp_once\":5,\"lian_yi_lian_limit\":100,\"re_lian_yi_lian_limit\":50,\"lian_yi_lian_once\":10,\"re_lian_yi_lian_once\":5,\"follow_read_limit\":100,\"re_follow_read_limit\":50,\"follow_read_once\":10,\"re_follow_read_once\":5,\"max_rank_cert_template_uri\":\"10/sea/ec/dc/53fbb18524000db301ff04863151\"}"

	var cbd CBD
	x := &cbd
	fmt.Println(cbd)
	fmt.Println(x.StartTs)
}
