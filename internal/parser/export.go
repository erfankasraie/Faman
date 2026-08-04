package parser

// PagesDir returns the resolved pages/fa directory used by LoadPage/ListPages.
func PagesDir() (string, error) {
	return pagesDir()
}
