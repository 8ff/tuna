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

func GetVersion(versionParam string) ([]byte, error) {
	args := os.Args
	out, err := exec.Command(args[0], versionParam).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s => %s", err, out)
	}
	return out, nil
}

func SelfUpdate(url string) error {
	args := os.Args
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	size, err := io.Copy(buf, resp.Body)
	if err != nil {
		return err
	}

	if size < 10000 {
		return errors.New("returned response too small")
	}

	originalFile, _ := os.Stat(args[0])
	if err := os.Remove(args[0]); err != nil {
		return err
	}

	out, err := os.Create(args[0])
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, buf)
	if err != nil {
		return err
	}
	if err := os.Chmod(args[0], originalFile.Mode().Perm()); err != nil {
		return err
	}

	return nil
}
