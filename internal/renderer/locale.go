package renderer

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func warnIfBadLocale() {
	if runtime.GOOS == "windows" {
		warnWindowsTerminal()
		return
	}

	lang := os.Getenv("LANG")
	lc := os.Getenv("LC_ALL")
	lcctype := os.Getenv("LC_CTYPE")
	combined := strings.ToLower(lang + " " + lc + " " + lcctype)
	if strings.TrimSpace(combined) == "" {
		return
	}
	if strings.Contains(combined, "utf-8") || strings.Contains(combined, "utf8") {
		return
	}
	msg := "faman: هشدار — locale شما UTF-8 نیست؛ ممکن است حروف فارسی خراب دیده شوند."
	hint := "  پیشنهاد: export LANG=en_US.UTF-8   |  FAMAN_PLAIN=1 faman <cmd>"
	printWarn(msg, hint)
}

// On Windows, LANG is often unset — do not treat that as an error.
// Remind once-style tips only when classic conhost / non-UTF8 is likely.
func warnWindowsTerminal() {
	// Windows Terminal sets WT_SESSION; ConPTY apps may still work without it.
	if os.Getenv("WT_SESSION") != "" || os.Getenv("WT_PROFILE_ID") != "" {
		return
	}
	// PowerShell 7+ often fine; skip noisy warnings unless FAMAN_WINDOWS_HINT=1
	if strings.ToLower(os.Getenv("FAMAN_WINDOWS_HINT")) != "1" &&
		strings.ToLower(os.Getenv("FAMAN_WINDOWS_HINT")) != "true" {
		return
	}
	msg := "faman: برای فارسی در ویندوز Windows Terminal + فونت Cascadia/Noto را پیشنهاد می‌کنیم."
	hint := "  docs/windows.md  |  $env:FAMAN_PLAIN=1"
	printWarn(msg, hint)
}

func printWarn(msg, hint string) {
	if ColorEnabled() && !plainMode() {
		fmt.Fprintln(os.Stderr, WarningStyle.Render(msg))
		fmt.Fprintln(os.Stderr, DimStyle.Render(hint))
		return
	}
	fmt.Fprintln(os.Stderr, msg)
	fmt.Fprintln(os.Stderr, hint)
}
