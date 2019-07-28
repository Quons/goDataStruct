package main

import "testing"

func TestChar(t *testing.T) {
	a := `[{"Id":1,"CourseId":2,"AwardId":1,"AwardName":"金币","AwardType":1,"Number":50,"Index":1,"Image":"","Status":1,"AddTime":"2019-06-20T21:02:27+08:00","UpdateTime":"0001-01-01T00:00:00Z"},{"Id":2,"CourseId":2,"AwardId":1,"AwardName":"金币","AwardType":1,"Number":80,"Index":2,"Image":"","Status":1,"AddTime":"2019-06-20T21:04:34+08:00","UpdateTime":"0001-01-01T00:00:00Z"},{"Id":3,"CourseId":2,"AwardId":1,"AwardName":"课程礼包","AwardType":3,"Number":1,"Index":3,"Image":"https://appd.knowbox.cn/codebox/course/gift/l1gift.png","Status":1,"AddTime":"2019-07-01T14:17:37+08:00","UpdateTime":"0001-01-01T00:00:00Z"}]`
	t.Log(a)

}
