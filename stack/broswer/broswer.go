package broswer

// Browser 浏览器的前进后退功能
type Browser struct {
	backStack    []string // 后退栈
	forwardStack []string // 前进栈
}

// Goback 后退
func (b *Browser) Goback() string {
	backLen := len(b.backStack)
	b.forwardStack = append(b.forwardStack, b.backStack[backLen-1])
	pageName := b.backStack[backLen-2]
	b.backStack = b.backStack[0 : backLen-1]
	return pageName
}

// Forward 前进
func (b *Browser) Forward() string {
	forwardLen := len(b.forwardStack)
	if forwardLen == 0 {
		return b.backStack[len(b.backStack)-1]
	}
	pageName := b.forwardStack[forwardLen-1]
	b.backStack = append(b.backStack, pageName)
	b.forwardStack = b.forwardStack[0 : forwardLen-1]
	return pageName
}

// Push 页面进栈
func (b *Browser) Push(pageName string) string {
	b.backStack = append(b.backStack, pageName)
	b.forwardStack = b.forwardStack[0:0]
	return pageName
}
