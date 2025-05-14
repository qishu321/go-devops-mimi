package tools

import (
	"fmt"
	"strings"
	"time"

	"github.com/mozillazg/go-pinyin"
	"github.com/robfig/cron/v3"
)

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

func ConvertToPinYin(src string) (dst string) {
	args := pinyin.NewArgs()
	args.Fallback = func(r rune, args pinyin.Args) []string {
		return []string{string(r)}
	}

	for _, singleResult := range pinyin.Pinyin(src, args) {
		for _, result := range singleResult {
			dst = dst + result
		}
	}
	return
}

// buildScriptWithEnv 会生成：
// declare -A env_task_s
// env_task_s=(
//
//	["KEY"]="VALUE"
//	...
//
// )
//
// <原始脚本内容>
// exit
func BuildScriptWithEnv(envs map[string]string, content string) string {
	var sb strings.Builder
	// 1. 声明关联数组
	sb.WriteString("declare -A env_task_s\n")
	sb.WriteString("env_task_s=(\n")
	for k, v := range envs {
		// ["ENV"]="prod"
		sb.WriteString(fmt.Sprintf(`  [%q]=%q`+"\n", k, v))
	}
	sb.WriteString(")\n\n")
	// 2. 原脚本
	sb.WriteString(content)
	sb.WriteString("\n")
	// // 3. 退出
	sb.WriteString("exit\n")
	return sb.String()
}

// mustParseTime 解析 RFC3339 时间
func MustParseTime(ts string) time.Time {
	t, _ := time.Parse(time.RFC3339, ts)
	return t
}

func ParseTime(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, fmt.Errorf("空时间字符串")
	}

	// 1) RFC3339
	if t, err := time.Parse(time.RFC3339, value); err == nil {
		return t, nil
	}

	// 2) "2006-01-02 15:04:05"
	const layout2 = "2006-01-02 15:04:05"
	if t, err := time.ParseInLocation(layout2, value, time.Local); err == nil {
		return t, nil
	}

	// 3) "2006-01-02"
	const layout3 = "2006-01-02"
	if t, err := time.ParseInLocation(layout3, value, time.Local); err == nil {
		return t, nil
	}

	return time.Time{}, fmt.Errorf("无法解析时间: %s", value)
}

// ValidateCronExpr 检验给定的 Cron 表达式是否有效。
// 支持：
//   - 六字段格式（分 时 日 月 周）
//   - 描述符：@yearly、@monthly、@every 等
func ValidateCronExpr(expr string) error {
	// 构造一个支持秒和描述符的解析器
	parser := cron.NewParser(
		cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)
	// 尝试解析表达式
	if _, err := parser.Parse(expr); err != nil {
		return fmt.Errorf("无效的 Cron 表达式 [%s]：%w", expr, err)
	}
	return nil
}
