package translator

import "path/filepath"

// narrowPath reduces full path to file name and parent dir only.
func narrowPath(fp string) string {
	if !filepath.IsAbs(fp) {
		if abs, err := filepath.Abs(fp); err != nil {
			// seems to be reduced already
			return fp
		} else {
			fp = abs
		}
	}
	return filepath.Join(filepath.Base(filepath.Dir(fp)), filepath.Base(fp))
}
