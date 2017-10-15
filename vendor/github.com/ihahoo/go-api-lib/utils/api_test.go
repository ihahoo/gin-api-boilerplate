package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestResTime(t *testing.T) {
	testTime, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	assert.Equal(t, ResTime(testTime), "2012-11-01T22:08:41Z", "时间格式不匹配")
}

func TestReqDate(t *testing.T) {
	testTime, _ := time.ParseInLocation("2006-01-02", "2012-11-01", time.Local)
	assert.Equal(t, ReqDate("2012-11-01"), testTime, "时间格式不匹配")
}
