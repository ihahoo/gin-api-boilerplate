package validate

import (
	"reflect"
	"regexp"
	"strconv"

	valid "github.com/asaskevich/govalidator"
	"github.com/ihahoo/go-api-lib/errors"
)

// IsEmpty 检查值是否是空
func IsEmpty(s interface{}) bool {
	v := reflect.ValueOf(s)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	default:
		return reflect.DeepEqual(s, reflect.Zero(v.Type()).Interface())
	}
}

// Required 检查是否为空
func Required(v interface{}, name string) (errors.E, bool) {
	if IsEmpty(v) {
		return errors.E{422, 400001, name + "不能为空"}, false
	}

	return errors.E{}, true
}

// Email 检查是否是Email
func Email(v string, isRequired bool, name string) (errors.E, bool) {
	if isRequired {
		if info, ok := Required(v, name); !ok {
			return info, false
		}
	}
	if len(v) > 0 {
		if !valid.IsEmail(v) {
			return errors.E{422, 400004, name + "格式不正确"}, false
		}
	}

	return errors.E{}, true
}

// MinLength 字符的最小长度
func MinLength(v string, length int, name string) (errors.E, bool) {
	if len(v) < length {
		return errors.E{422, 400005, name + "不能少于" + strconv.Itoa(length) + "个字符"}, false
	}

	return errors.E{}, true
}

// Mobile 手机号码
func Mobile(v string, isRequired bool, name string) (errors.E, bool) {
	if isRequired {
		if info, ok := Required(v, name); !ok {
			return info, false
		}
	}
	if len(v) > 0 {
		rxMobile := regexp.MustCompile("^(\\+?0?86\\-?)?1[345789]\\d{9}$")
		if !rxMobile.MatchString(v) {
			return errors.E{422, 400021, name + "错误"}, false
		}
	}

	return errors.E{}, true
}

// UUID uuid
func UUID(v string, isRequired bool, name string) (errors.E, bool) {
	if isRequired {
		if info, ok := Required(v, name); !ok {
			return info, false
		}
	}
	if len(v) > 0 {
		if !valid.IsUUID(v) {
			return errors.E{422, 400002, name + "格式错误"}, false
		}
	}

	return errors.E{}, true
}

// RealName 汉字姓名
func RealName(v string, isRequired bool, name string) (errors.E, bool) {
	if isRequired {
		if info, ok := Required(v, name); !ok {
			return info, false
		}
	}
	if len(v) > 0 {
		rxRealName := regexp.MustCompile("^[\u2E80-\uFE4F]{2,10}$")
		if !rxRealName.MatchString(v) {
			return errors.E{422, 400002, name + "格式错误"}, false
		}
	}

	return errors.E{}, true
}

// IDCard 身份证
func IDCard(v string, isRequired bool, name string) (errors.E, bool) {
	if isRequired {
		if info, ok := Required(v, name); !ok {
			return info, false
		}
	}
	if len(v) > 0 {
		rxIDCard := regexp.MustCompile("(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)")
		if !rxIDCard.MatchString(v) {
			return errors.E{422, 400002, name + "格式错误"}, false
		}
	}

	return errors.E{}, true
}

// MD5 MD5
func MD5(v string, isRequired bool, name string) (errors.E, bool) {
	if isRequired {
		if info, ok := Required(v, name); !ok {
			return info, false
		}
	}
	if len(v) > 0 {
		rxMD5 := regexp.MustCompile("^[a-f0-9]{32}$")
		if !rxMD5.MatchString(v) {
			return errors.E{422, 400002, name + "格式错误"}, false
		}
	}

	return errors.E{}, true
}
