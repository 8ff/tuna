package tuna

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

const minSize = 10000 // Minimum acceptable size for the downloaded file

func GetVersion(versionParam string) ([]byte, error) {
	args := os.Args
	out, err := exec.Command(args[0], versionParam).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s => %s", err, out)
	}
	return out, nil
}

func SelfUpdate(url string) error {
	// get path to running executable
	path, err := os.Executable()
	if err != nil {
		return err
	}

	// download new version
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}

	buf := new(bytes.Buffer)
	size, err := io.Copy(buf, resp.Body)
	if err != nil {
		return err
	}

	if size < minSize {
		return errors.New("returned response too small")
	}

	originalFile, err := os.Stat(path)
	if err != nil {
		return err
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, buf)
	if err != nil {
		return err
	}

	if err := os.Chmod(path, originalFile.Mode().Perm()); err != nil {
		return err
	}

	return nil
}
