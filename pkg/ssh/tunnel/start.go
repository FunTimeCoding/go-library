package tunnel

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ssh/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	systemConstant "github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"log"
	"net"
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

		r := run.New()

		if !t.NoOutput {
			r = r.WithStdio()
		}

		t.process = r.Open(parts...)

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

		waitFail := t.process.Wait()

		if t.CleanStop && waitFail != nil {
			errors.PanicOnError(waitFail)

			if exit := t.process.ExitCode(); exit != 0 {
				log.Panicf(
					"non-zero exit code: %d\noutput: %s\n",
					exit,
					t.process.ProcessMessage(),
				)
			}
		}

		close(t.stopped)
	}()
	<-t.started
	<-t.listening

	return &Result{LocalPort: localPort}
}
