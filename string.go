// 字符串相关方法封装

package goutils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/antlabs/strsim"
)

// RemoveAllWhitespace 删除字符串中所有的空白符
func RemoveAllWhitespace(s string) string {
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, " ", "", -1)
	return s
}

// ReverseString 翻转字符串
func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// RemoveDuplicateWhitespace 删除字符串中重复的空白字符为单个空白字符
// trim: 是否去掉首位空白
func RemoveDuplicateWhitespace(s string, trim bool) string {
	ws, err := regexp.Compile(`\s+`)
	if err != nil {
		return s
	}
	s = ws.ReplaceAllString(s, " ")
	if trim {
		s = strings.TrimSpace(s)
	}
	return s
}

// StrSimilarity 提供不同的算法检测文本相似度
func StrSimilarity(s1, s2 string, algorithm string) float64 {
	// strsim: 完全相同返回 1，完全不同返回 0
	switch strings.ToLower(algorithm) {
	case "levenshtein":
		return strsim.Compare(s1, s2)
	case "dice":
		return strsim.Compare(s1, s2, strsim.DiceCoefficient())
	case "jaro":
		return strsim.Compare(s1, s2, strsim.Jaro())
	case "hamming":
		return strsim.Compare(s1, s2, strsim.Hamming())
	}
	return strsim.Compare(s1, s2, strsim.Hamming())
}

// YiWanString 将数字转换为 亿/万
func YiWanString(num float64) string {
	YI := float64(100000000.0)
	WAN := float64(10000.0)
	yi := num / YI
	if math.Abs(yi) >= 1 {
		return fmt.Sprintf("%.2f 亿", yi)
	}
	wan := num / WAN
	if math.Abs(wan) >= 1 {
		return fmt.Sprintf("%.2f 万", wan)
	}
	return fmt.Sprint(num)
}

// SplitStringFields 将传入字符串分割为slice
func SplitStringFields(s string) []string {
	return regexp.MustCompile("[\\/\\:\\,\\;\\.\\s\\-\\|\\#\\$\\%\\&\\+\\=\\?]+").Split(s, -1)
}

// MD5 string md5
func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

// RemoveMarkdownTags 移除markdown文本中的标记返回纯文本内容
func RemoveMarkdownTags(markdownText string) string {
	// 移除内联CSS样式
	markdownText = regexp.MustCompile("<style>[\\s\\S]*?</style>").ReplaceAllString(markdownText, "")

	// 移除内联JavaScript代码
	markdownText = regexp.MustCompile("<script>[\\s\\S]*?</script>").ReplaceAllString(markdownText, "")

	// 移除HTML标记
	markdownText = regexp.MustCompile(`<[^>]+>`).ReplaceAllString(markdownText, "")

	// 移除分隔线
	markdownText = regexp.MustCompile(`-{3,}\n`).ReplaceAllString(markdownText, "")

	// 移除粗体
	markdownText = regexp.MustCompile(`\*\*(.*?)\*\*`).ReplaceAllString(markdownText, "$1")

	// 移除斜体
	markdownText = regexp.MustCompile(`\*(.*?)\*`).ReplaceAllString(markdownText, "$1")

	// 移除标题
	markdownText = regexp.MustCompile(`#+\s+(.*?)\n`).ReplaceAllString(markdownText, "$1\n")

	// 移除链接/图片
	markdownText = regexp.MustCompile(`\[(.*?)\]\(.*?\)`).ReplaceAllString(markdownText, "$1")

	// 移除代码块
	markdownText = regexp.MustCompile("```.*?```").ReplaceAllString(markdownText, "")

	// 移除列表
	markdownText = regexp.MustCompile(`[-*]\s+(.*?)\n`).ReplaceAllString(markdownText, "$1\n")

	// 移除删除线
	markdownText = regexp.MustCompile(`~~(.*?)~~`).ReplaceAllString(markdownText, "$1")

	// 移除引用
	markdownText = regexp.MustCompile(`>\s*`).ReplaceAllString(markdownText, "")
	return markdownText
}

// IsAlphaNumericEnding 检查字符串是否以字母和数字结尾
func IsAlphaNumericEnding(s string) bool {
	// 使用正则表达式检查是否以字母或数字结尾
	re := regexp.MustCompile(`[a-zA-Z0-9]$`)
	return re.MatchString(s)
}
