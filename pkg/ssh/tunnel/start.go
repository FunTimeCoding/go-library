package tunnel

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ssh/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	systemConstant "github.com/funtimecoding/go-library/pkg/system/constant"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

func (t *Tunnel) Start(
	host string,
	targetHost string,
	targetPort int,
	localPort int,
) *Result {
	if localPort == 0 {
		localPort = system.FindUnusedPort(8080)
	}

	go func() {
		parts := []string{
			constant.Command,
			constant.NoRemoteCommandArgument,
			constant.BackgroundArgument,
			constant.LocalPortForwardArgument,
			fmt.Sprintf("%d:%s:%d", localPort, targetHost, targetPort),
			host,
		}

		if t.Verbose {
			fmt.Printf("Command: %s\n", parts)
		}

		t.command = exec.Command(parts[0], parts[1:]...)

		if !t.NoOutput {
			t.command.Stdout = os.Stdout
			t.command.Stderr = os.Stderr
		}

		errors.PanicOnError(t.command.Start())

		for {
			d, dialFail := net.DialTimeout(
				systemConstant.Transmission,
				fmt.Sprintf("localhost:%d", localPort),
				time.Second*5,
			)

			if dialFail == nil {
				errors.LogClose(d)
				close(t.listening)

				break
			}

			time.Sleep(time.Second * 5)
		}

		close(t.started)

		if t.Verbose {
			fmt.Printf(
				"Forward up: %s:%d to %s:%d\n",
				"localhost",
				localPort,
				targetHost,
				targetPort,
			)
		}

		waitFail := t.command.Wait()

		if t.CleanStop && waitFail != nil {
			errors.PanicOnError(waitFail)

			if exit := t.command.ProcessState.ExitCode(); exit != 0 {
				log.Panicf(
					"non-zero exit code: %d\n"+
						"output: %s\n",
					exit,
					t.command.ProcessState.String(),
				)
			}
		}

		close(t.stopped)
	}()

	<-t.started
	<-t.listening

	return &Result{LocalPort: localPort}
}
