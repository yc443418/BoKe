package utils

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

const (
	defaultSeparator   = "-"
	defaultRandomChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 去除了容易混淆的字符
)

// Options SKU生成配置选项
type Options struct {
	Prefix       string   // SKU前缀（必填）
	Attributes   []string // 属性代码（如颜色、尺寸）
	UseDate      bool     // 是否包含日期部分
	DateLayout   string   // 日期格式（默认"20060102"）
	RandomLength int      // 随机码长度（默认4）
	Separator    string   // 分隔符（默认"-"）
}

// GenerateSKU 生成带校验的SKU编号
// 格式: [前缀]-[属性]-[日期]-[随机码]（各部分根据配置可选）
func GenerateSKU(opt Options) (string, error) {
	// 参数校验
	if err := validateOptions(&opt); err != nil {
		return "", err
	}

	var parts []string

	// 添加前缀
	parts = append(parts, strings.ToUpper(opt.Prefix))

	// 处理属性部分
	if len(opt.Attributes) > 0 {
		attrPart := strings.Join(opt.Attributes, opt.Separator)
		parts = append(parts, strings.ToUpper(attrPart))
	}

	// 处理日期部分
	if opt.UseDate {
		layout := opt.DateLayout
		if layout == "" {
			layout = "20060102" // 默认年月日格式
		}
		parts = append(parts, time.Now().Format(layout))
	}

	// 生成随机码
	if opt.RandomLength > 0 {
		random, err := generateRandomString(opt.RandomLength)
		if err != nil {
			return "", err
		}
		parts = append(parts, random)
	}

	return strings.Join(parts, opt.Separator), nil
}

// validateOptions 校验配置参数
func validateOptions(opt *Options) error {
	if opt.Prefix == "" {
		return errors.New("prefix is required")
	}

	if opt.RandomLength < 0 {
		return errors.New("random length cannot be negative")
	}

	// 设置默认值
	if opt.Separator == "" {
		opt.Separator = defaultSeparator
	}

	return nil
}

// generateRandomString 生成指定长度的随机字符串
func generateRandomString(length int) (string, error) {
	if length <= 0 {
		return "", nil
	}

	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(length)

	for i := 0; i < length; i++ {
		sb.WriteByte(defaultRandomChars[rand.Intn(len(defaultRandomChars))])
	}

	return sb.String(), nil
}
