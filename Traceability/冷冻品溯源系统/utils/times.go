package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type HTime struct {
	time.Time
}

var (
	formatTime = "2006-01-02 15:04:05"
)

func (t HTime) MarshalJSON() ([]byte, error) { // MarshalJSON 方法只能被 HTime 类型的对象调用
	formatted := fmt.Sprintf("\"%s\"", t.Format(formatTime))
	return []byte(formatted), nil
}

func (t *HTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+formatTime+`"`, string(data), time.Local)
	//第一个参数是时间格式，第二个参数是 JSON 字符串，第三个参数是时区（这里使用本地时区 time.Local）
	*t = HTime{Time: now}
	//*t 表示对指针 t 指向的原始对象进行解引用操作，即获取原始的 HTime 对象
	//将新创建的 HTime 对象（包含解析后的时间值 now）赋值给指针 t 指向的原始对象
	return
}

// Value 方法：用于将自定义类型转换为数据库可以理解的值，通常在插入或更新数据库时使用
func (t HTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 方法：用于将数据库查询结果扫描到自定义类型中，通常在查询数据库时使用
func (t *HTime) Scan(v interface{}) error {
	//使用类型断言 v.(time.Time) 检查传入的值是否为 time.Time 类型
	value, ok := v.(time.Time)
	if ok {
		*t = HTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
