package goutils

import (
	"bytes"
	"html/template"

	chromav2html "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	mermaid "go.abhg.dev/goldmark/mermaid"
)

// MarkdownToTemplateHTML 将 Markdown 转换为 template.HTML。
// 启用 goldmark 全部内置扩展特性，包括：
//   - GFM: 表格(Table)、删除线(Strikethrough)、自动链接(Linkify)、任务列表(TaskList)
//   - 语法高亮: 代码块语法高亮，使用 manni 主题，显示行号
//   - Mermaid 图表: 支持渲染 Mermaid 图表
//   - 定义列表(DefinitionList): PHP Markdown Extra 风格
//   - 脚注(Footnote): 支持多文档 ID 前缀防冲突
//   - 排版优化(Typographer): 智能引号、破折号、省略号等
//   - CJK 优化: 东亚换行处理、转义空格支持强调语法
//   - 自动标题 ID: 为标题生成锚点
//   - 自定义属性: 标题支持 {#id .class} 语法
//   - 硬换行: 将换行符渲染为 <br>
//
// enableUnsafe 控制是否允许原始 HTML 和潜在危险链接：
//   - false: 安全模式，原始 HTML 会被转义（推荐用于用户输入）
//   - true:  信任模式，保留原始 HTML（仅用于可信内容）
func MarkdownToTemplateHTML(source string, enableUnsafe bool) (template.HTML, error) {
	// 构建扩展列表（全部启用）
	extensions := []goldmark.Extender{
		extension.GFM,
		highlighting.NewHighlighting(
			highlighting.WithStyle("manni"),
			highlighting.WithFormatOptions(
				chromav2html.WithLineNumbers(true),
			),
		),
		&mermaid.Extender{},
		extension.Footnote,
		extension.DefinitionList,
		extension.NewCJK(
			extension.WithEastAsianLineBreaks(extension.EastAsianLineBreaksSimple),
			extension.WithEscapedSpace(),
		),
		extension.Typographer,
	}

	// 构建解析器选项
	parserOpts := []parser.Option{
		parser.WithAutoHeadingID(), // 自动标题 ID
		parser.WithAttribute(),     // 自定义属性支持
	}

	// 构建渲染器选项
	rendererOpts := []renderer.Option{
		html.WithHardWraps(), // 硬换行：\n 渲染为 <br>
	}
	if enableUnsafe {
		rendererOpts = append(rendererOpts, html.WithUnsafe())
	}

	// 创建 goldmark 实例
	md := goldmark.New(
		goldmark.WithExtensions(extensions...),
		goldmark.WithParserOptions(parserOpts...),
		goldmark.WithRendererOptions(rendererOpts...),
	)

	// 执行转换
	var buf bytes.Buffer
	if err := md.Convert([]byte(source), &buf); err != nil {
		return "", err
	}

	return template.HTML(buf.String()), nil
}
