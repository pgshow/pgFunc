package spider

import (
	"github.com/eddycjy/fake-useragent"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"regexp"
	"strings"
)

// GetRandomUA 随机获取 userAgent
func GetRandomUA() string {
	return browser.Computer()
}

// HtmlMinify html代码压缩，去掉两个以上连续的空格
func HtmlMinify(str string) (result string, err error) {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)

	result, err = m.String("text/html", str)
	if err == nil {
		result = strings.Replace(result, "\n", "", -1) // 换行符

		reg, _ := regexp.Compile(`\s{2,}`) // 两个及两个以上的空白
		result = reg.ReplaceAllString(result, "")
	}

	return
}
