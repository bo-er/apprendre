## 使用Go来替换文本


替换掉字符串中包含`%s`,` `,`%v`,`:`的字符，使用`|`号隔开
```go
var trimRegexPattern = regexp.MustCompile(`(\%s)|(\%v)|\s|\:`)

func TrimAllTokens(s string) string {
	return trimRegexPattern.ReplaceAllString(s, "")
}
```