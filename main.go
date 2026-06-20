package main

import (
	"github.com/atotto/clipboard"
	webview "github.com/webview/webview_go"
)

const targetURL = "https://chat.mistral.ai/work"

// WKWebView blocks navigator.clipboard, so we bind a Go function and intercept Cmd+V.
const pasteJS = `
document.addEventListener('keydown', function(e) {
	if ((e.metaKey || e.ctrlKey) && e.key === 'v') {
		e.preventDefault();
		window.clipboardRead().then(function(text) {
			var el = document.activeElement;
			if (!el) return;
			if (el.isContentEditable) {
				document.execCommand('insertText', false, text);
				return;
			}
			if (el.tagName === 'INPUT' || el.tagName === 'TEXTAREA') {
				var s = el.selectionStart, end = el.selectionEnd;
				el.value = el.value.slice(0, s) + text + el.value.slice(end);
				el.selectionStart = el.selectionEnd = s + text.length;
				el.dispatchEvent(new Event('input', {bubbles: true}));
			}
		});
	}
}, true);
`

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Mistral Vibe")
	w.SetSize(1400, 900, webview.HintNone)

	w.Bind("clipboardRead", func() string {
		text, _ := clipboard.ReadAll()
		return text
	})

	w.Init(pasteJS)
	w.Navigate(targetURL)
	w.Run()
}
