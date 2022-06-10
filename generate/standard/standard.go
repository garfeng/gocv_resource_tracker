package standard

import "os"

func ScanDirs(name string) ([]string, error) {
	b, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer b.Close()

	names, err := b.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	return names, nil
}
