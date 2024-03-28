package util

import (
	"bufio"
	"context"
	"os/exec"
	pb "tuyoops/go/terraform-plugin/api"
)

func RunTerraformCommand(stream pb.TerraformService_ExecuteTerraformCmdServer, args ...string) error {
	cmd := exec.Command("terraform", args...)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	scannerOut := bufio.NewScanner(stdoutPipe)
	scannerErr := bufio.NewScanner(stderrPipe)

	errCh := make(chan error, 1)
	doneCh := make(chan struct{})

	go func() {
		defer func() {
			doneCh <- struct{}{}
		}()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if scannerOut.Scan() {
					if err := stream.Send(&pb.ExecuteTerraformCmdResponse{SuccessMessage: scannerOut.Text()}); err != nil {
						errCh <- err
						cancel()
						return
					}
				}
			}
		}
	}()

	go func() {
		defer func() {
			doneCh <- struct{}{}
		}()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if scannerErr.Scan() {
					if err := stream.Send(&pb.ExecuteTerraformCmdResponse{ErrorMessage: scannerErr.Text()}); err != nil {
						errCh <- err
						cancel()
						return
					}
				}
			}
		}
	}()

	cmdErrCh := make(chan error, 1)
	cmdDoneCh := make(chan struct{})
	go func() {
		defer close(cmdDoneCh)
		if err := cmd.Run(); err != nil {
			cmdErrCh <- err
			cancel()
			return
		}

		if err := stream.Send(&pb.ExecuteTerraformCmdResponse{StreamEOF: "EOF"}); err != nil {
			cmdErrCh <- err
			cancel()
			return
		}
	}()

	select {
	case <-cmdDoneCh:
		// 执行完成
	case err := <-cmdErrCh:
		return err
	}

	close(errCh)

	// 等待从 errCh 中读取错误
	var lastErr error
	for err := range errCh {
		lastErr = err
	}
	return lastErr

}
