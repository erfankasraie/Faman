package app

import (
	"github.com/spf13/cobra"

	"github.com/faman-project/faman/internal/update"
)

var (
	updateCheckOnly bool
	updatePagesOnly bool
	updateForce     bool
	updateVerify    bool
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "بررسی نسخه و به‌روزرسانی صفحات راهنما",
	Long: `بررسی آخرین انتشار روی GitHub و (اختیاری) تازه‌سازی صفحات فارسی.

  faman update                  # وضعیت نسخه + راهنما
  faman update --check          # فقط مقایسه نسخه
  faman update --pages          # صفحات از شاخه main (هش SHA256 چاپ می‌شود)
  faman update --pages --verify # آرشیو آخرین Release + تأیید SHA256SUMS

باینری را این دستور عوض نمی‌کند.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return update.Run(update.Options{
			CurrentVersion: Version(),
			CheckOnly:      updateCheckOnly,
			PagesOnly:      updatePagesOnly,
			Force:          updateForce,
			Verify:         updateVerify,
		})
	},
}

func init() {
	updateCmd.Flags().BoolVar(&updateCheckOnly, "check", false, "فقط بررسی نسخه روی GitHub")
	updateCmd.Flags().BoolVar(&updatePagesOnly, "pages", false, "دانلود و نصب صفحات")
	updateCmd.Flags().BoolVar(&updateForce, "force", false, "بازنویسی صفحات حتی اگر مسیر سیستم باشد")
	updateCmd.Flags().BoolVar(&updateVerify, "verify", false, "تأیید SHA256 در برابر SHA256SUMS آخرین Release")
}
