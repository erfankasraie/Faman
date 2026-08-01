package app

import (
	"github.com/spf13/cobra"

	"github.com/faman-project/faman/internal/update"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "به‌روزرسانی صفحات راهنما (در حال حاضر غیرفعال)",
	Long: `به‌روزرسانی صفحات راهنما از مخزن آنلاین.

در نسخه فعلی این دستور فقط یک placeholder است.
در نسخه‌های آینده امکان به‌روزرسانی آنلاین اضافه خواهد شد.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return update.Run()
	},
}
