package validate

import (
	"testing"

	"github.com/ihahoo/go-api-lib/errors"
	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	s1 := ""
	s2 := 0
	s3 := "abc"
	assert.Equal(t, IsEmpty(s1), true, "s1为空")
	assert.Equal(t, IsEmpty(s2), true, "s2为空")
	assert.Equal(t, IsEmpty(s3), false, "s3不为空")
}

func TestRequired(t *testing.T) {
	test, ok := Required("", "名字")
	assert.Equal(t, ok, false, "名字为空")
	assert.Equal(t, test, errors.E{Status: 422, Code: 400001, Msg: "名字不能为空"}, "当空时正确返回数据")

	test2, ok2 := Required(22, "年龄")
	assert.Equal(t, ok2, true, "年龄不为空")
	assert.Equal(t, test2, errors.E{}, "年龄有值, 不为空")
}

func TestEmail(t *testing.T) {
	test1, ok1 := Email("test123@abc", false, "Email")
	assert.Equal(t, ok1, false, "Email格式应是不正确")
	assert.Equal(t, test1, errors.E{422, 400004, "Email格式不正确"}, "Email格式错误时要返回正确数据")

	test2, ok2 := Email("", true, "Email")
	assert.Equal(t, ok2, false, "Email为空")
	assert.Equal(t, test2, errors.E{422, 400001, "Email不能为空"}, "Email应该为空")

	test3, ok3 := Email("test123@abc.com", true, "Email")
	assert.Equal(t, ok3, true, "Email格式应该正确")
	assert.Equal(t, test3, errors.E{}, "Email格式正确, 不应有错误信息返回")

	test4, ok4 := Email("", false, "Email")
	assert.Equal(t, ok4, true, "Email格式应该正确")
	assert.Equal(t, test4, errors.E{}, "Email格式正确, 不应有错误信息返回")
}

func TestMinLength(t *testing.T) {
	test1, ok1 := MinLength("abc", 4, "MinLength")
	assert.Equal(t, ok1, false, "MinLength test1应该返回错误")
	assert.Equal(t, test1, errors.E{422, 400005, "MinLength不能少于4个字符"}, "MinLength test1应该返回正确的错误数据")

	test2, ok2 := MinLength("abcd", 4, "MinLength")
	assert.Equal(t, ok2, true, "MinLength test2应该返回正确")
	assert.Equal(t, test2, errors.E{}, "MinLength test2正确，不应有错误信息返回")
}

func TestMobile(t *testing.T) {
	test1, ok1 := Mobile("83818888888", true, "Mobile")
	assert.Equal(t, ok1, false, "Mobile test1应该返回错误")
	assert.Equal(t, test1, errors.E{422, 400021, "Mobile错误"}, "Mobile test1应该返回正确的错误数据")

	test2, ok2 := Mobile("13818888888", true, "Mobile")
	assert.Equal(t, ok2, true, "Mobile test2应该返回正确")
	assert.Equal(t, test2, errors.E{}, "Mobile test2正确，不应有错误信息返回")
}

func TestUUID(t *testing.T) {
	test1, ok1 := UUID("abc-ffi98888888", true, "UUID")
	assert.Equal(t, ok1, false, "UUID test1应该返回错误")
	assert.Equal(t, test1, errors.E{422, 400002, "UUID格式错误"}, "UUID test1应该返回正确的错误数据")

	test2, ok2 := UUID("b2d3410d-6fb7-4e20-8c2d-4dc33bb64e5d", true, "UUID")
	assert.Equal(t, ok2, true, "UUID test2应该返回正确")
	assert.Equal(t, test2, errors.E{}, "UUID test2正确，不应有错误信息返回")
}

func TestRealName(t *testing.T) {
	test1, ok1 := RealName("王小2", true, "RealName")
	assert.Equal(t, ok1, false, "RealName test1应该返回错误")
	assert.Equal(t, test1, errors.E{422, 400002, "RealName格式错误"}, "RealName test1应该返回正确的错误数据")

	test2, ok2 := RealName("王小二", true, "UUID")
	assert.Equal(t, ok2, true, "RealName test2应该返回正确")
	assert.Equal(t, test2, errors.E{}, "RealName test2正确，不应有错误信息返回")
}

func TestIDCard(t *testing.T) {
	test1, ok1 := IDCard("31011019830911143a", true, "IDCard")
	assert.Equal(t, ok1, false, "IDCard test1应该返回错误")
	assert.Equal(t, test1, errors.E{422, 400002, "IDCard格式错误"}, "IDCard test1应该返回正确的错误数据")

	test2, ok2 := IDCard("310110198309111431", true, "IDCard")
	assert.Equal(t, ok2, true, "IDCard test2应该返回正确")
	assert.Equal(t, test2, errors.E{}, "IDCard test2正确，不应有错误信息返回")

	test3, _ := IDCard("31011019830911143X", false, "IDCard")
	assert.Equal(t, test3, errors.E{}, "IDCard test3正确，不应有错误信息返回")
}

func TestMD5(t *testing.T) {
	test1, ok1 := MD5("fc5e038d38a57032085441e7fe7010b", true, "MD5")
	assert.Equal(t, ok1, false, "MD5 test1应该返回错误")
	assert.Equal(t, test1, errors.E{422, 400002, "MD5格式错误"}, "MD5 test1应该返回正确的错误数据")

	test2, ok2 := MD5("fc5e038d38a57032085441e7fe7010b0", true, "MD5")
	assert.Equal(t, ok2, true, "MD5 test2应该返回正确")
	assert.Equal(t, test2, errors.E{}, "MD5 test2正确，不应有错误信息返回")
}
