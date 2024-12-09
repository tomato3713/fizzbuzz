package main

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"rsc.io/script"
	"rsc.io/script/scripttest"
)

func execMain() script.Cmd {
	return script.Command(
		script.CmdUsage{
			Summary: "run the program",
			Args:    "args...",
		},
		func(s *script.State, args ...string) (script.WaitFunc, error) {
			if len(args) != 1 {
				return nil, script.ErrUsage
			}

			return func(state *script.State) (string, string, error) {
				r, w, err := os.Pipe()
				if err != nil {
					return "", "", err
				}

				// Redirect stdout to a pipe.
				stdout := os.Stdout
				os.Stdout = w

				os.Args = []string{"fizzbuzz", args[0]}
				main()

				// restore stdout
				os.Stdout = stdout
				w.Close()

				// Read the output from the pipe.
				var buf bytes.Buffer
				io.Copy(&buf, r)

				s.Logf("args: %s, stdout: %s\n", os.Args, buf.String())

				return buf.String(), "", nil
			}, nil
		},
	)
}

func customEngine(engine *script.Engine) map[string]script.Cmd {
	cmds := script.DefaultCmds()
	cmds["run"] = execMain()
	return cmds
}

func TestAll(t *testing.T) {
	ctx := context.Background()
	engine := script.NewEngine()
	engine.Cmds = customEngine(engine)
	env := os.Environ()
	scripttest.Test(t, ctx, engine, env, "testdata/*.txt")
}
