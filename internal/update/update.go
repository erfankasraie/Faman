package update

import (
	"fmt"

	"github.com/faman-project/faman/internal/renderer"
)

// Run performs the update operation.
// Currently a placeholder for future online update support.
func Run() error {
	useColor := renderer.ColorEnabled()

	title := "∧  به‌روزرسانی"
	if useColor {
		fmt.Println(renderer.TitleStyle.Render(title))
		fmt.Println(renderer.DimStyle.Render(stringsRepeat("─", 36)))
		fmt.Println()
		fmt.Println(renderer.MetaStyle.Render("  به‌روزرسانی آنلاین در نسخه فعلی پشتیبانی نمی‌شود."))
		fmt.Println()
		fmt.Println(renderer.DimStyle.Render("  در نسخه‌های آینده (v0.4+) امکان دریافت آخرین صفحات"))
		fmt.Println(renderer.DimStyle.Render("  از مخزن رسمی faman اضافه خواهد شد."))
		fmt.Println()
		fmt.Println(renderer.SuccessStyle.Render("  فعلاً می‌توانید پروژه را از GitHub به‌روز کنید:"))
		fmt.Println(renderer.DimStyle.Render("    git pull origin main"))
	} else {
		fmt.Println("به‌روزرسانی آنلاین در نسخه فعلی پشتیبانی نمی‌شود.")
		fmt.Println()
		fmt.Println("در نسخه‌های آینده (v0.4+) امکان دریافت آخرین صفحات")
		fmt.Println("از مخزن رسمی faman اضافه خواهد شد.")
		fmt.Println()
		fmt.Println("فعلاً می‌توانید پروژه را از GitHub به‌روز کنید:")
		fmt.Println("  git pull origin main")
	}
	return nil
}

func stringsRepeat(s string, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += s
	}
	return out
}
