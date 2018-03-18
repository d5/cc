// +build windows

package cc

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	colorable "github.com/mattn/go-colorable"
)

var (
	osStdout = os.Stdout
	osStderr = os.Stderr

	stdoutPipeR, stdoutPipeW   *os.File
	stderrPipeR, stderrPipeW   *os.File
	stdoutPipeCh, stderrPipeCh chan error
)

func doEnable() error {
	var err error

	// create pipe for stdout
	stdoutPipeR, stdoutPipeW, err = os.Pipe()
	if err != nil {
		return fmt.Errorf("error creating pipe for stdout: %s", err.Error())
	}
	os.Stdout = stdoutPipeW
	go func() {
		stdoutPipeCh = make(chan error)
		_, err := io.Copy(colorable.NewColorableStdout(), stdoutPipeR)
		stdoutPipeCh <- err
	}()

	// create pipe for stderr
	stderrPipeR, stderrPipeW, err = os.Pipe()
	if err != nil {
		return fmt.Errorf("error creating pipe for stderr: %s", err.Error())
	}
	os.Stderr = stderrPipeW
	go func() {
		stderrPipeCh = make(chan error)
		_, err := io.Copy(colorable.NewColorableStderr(), stderrPipeR)
		stderrPipeCh <- err
	}()

	return nil
}

func doDisable() error {
	// close pipe for stdout
	if stdoutPipeR != nil {
		if err := stdoutPipeR.Close(); err != nil {
			return fmt.Errorf("error closing pipe for stdout: %s", err.Error())
		}

		if stdoutPipeCh != nil {
			select {
			case err := <-stdoutPipeCh:
				if err != nil {
					return err
				}
			case <-time.After(5 * time.Second):
				return errors.New("timed out closing stdout pipe")
			}
		}

	}
	os.Stdout = osStdout

	// close pipe for stderr
	if stderrPipeR != nil {
		if err := stderrPipeR.Close(); err != nil {
			return fmt.Errorf("error closing pipe for stderr: %s", err.Error())
		}

		if stderrPipeCh != nil {
			select {
			case err := <-stderrPipeCh:
				if err != nil {
					return err
				}
			case <-time.After(5 * time.Second):
				return errors.New("timed out closing stderr pipe")
			}
		}
	}
	os.Stderr = osStderr

	return nil
}
