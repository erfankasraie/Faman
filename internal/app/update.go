package app

import (
	"github.com/spf13/cobra"

	"github.com/faman-project/faman/internal/update"
)

var (
	updateCheckOnly bool
	updatePagesOnly bool
	updateForce     bool
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "بررسی نسخه و به‌روزرسانی صفحات راهنما",
	Long: `بررسی آخرین انتشار روی GitHub و (اختیاری) تازه‌سازی صفحات فارسی.

  faman update           # وضعیت نسخه + راهنمای بعدی
  faman update --check   # فقط مقایسه نسخه
  faman update --pages   # دانلود pages/fa از شاخه main

باینری را این دستور عوض نمی‌کند؛ برای باینری دوباره get.sh / release را اجرا کنید.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return update.Run(update.Options{
			CurrentVersion: Version(),
			CheckOnly:      updateCheckOnly,
			PagesOnly:      updatePagesOnly,
			Force:          updateForce,
		})
	},
}

func init() {
	updateCmd.Flags().BoolVar(&updateCheckOnly, "check", false, "فقط بررسی نسخه روی GitHub")
	updateCmd.Flags().BoolVar(&updatePagesOnly, "pages", false, "دانلود و نصب صفحات از main")
	updateCmd.Flags().BoolVar(&updateForce, "force", false, "بازنویسی صفحات حتی اگر مسیر سیستم باشد")
}
