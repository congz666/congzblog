package markdown

import (
	"html/template"

	"github.com/russross/blackfriday"
)

// MarkdownToHTML 将 markdown 文档转换成 template.HTML
func MarkdownToHTML(input string) template.HTML {
	return HTML(input)
}

// HTML 将 markdown 文档转换成 html 文件
func HTML(input string) template.HTML {

	// 因为默认字符串会被模板转译，所以返回一个 template.HTML
	// 就可以叫 HTML 原样输出了

	htmlContent := markdown([]byte(input))

	return template.HTML(htmlContent)
}

// 自定义的 blackfriday ，为了支持生成锚点
// 下面的主要内容是从 https://github.com/shurcooL/github_flavored_markdown 扒出来的
func markdown(text []byte) []byte {
	renderer := &renderer{Html: blackfriday.HtmlRenderer(markdownHTMLFlags, "", "").(*blackfriday.Html)}

	return blackfriday.MarkdownOptions(text, renderer, blackfriday.Options{
		Extensions: markdownExtensions})
}

// 定义 HTML 渲染器的配置选项
const markdownHTMLFlags = 0 |
	blackfriday.HTML_USE_XHTML |
	blackfriday.HTML_USE_SMARTYPANTS |
	blackfriday.HTML_SMARTYPANTS_FRACTIONS |
	blackfriday.HTML_SMARTYPANTS_DASHES |
	blackfriday.HTML_SMARTYPANTS_LATEX_DASHES |
	blackfriday.HTML_FOOTNOTE_RETURN_LINKS

// 定义 markdown 扩展，其实就是复制的 commonExtensions
const markdownExtensions = 0 |
	blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
	blackfriday.EXTENSION_TABLES |
	blackfriday.EXTENSION_FENCED_CODE |
	blackfriday.EXTENSION_AUTOLINK |
	blackfriday.EXTENSION_STRIKETHROUGH |
	blackfriday.EXTENSION_SPACE_HEADERS |
	blackfriday.EXTENSION_HEADER_IDS |
	blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
	blackfriday.EXTENSION_DEFINITION_LISTS

type renderer struct {
	*blackfriday.Html
}
