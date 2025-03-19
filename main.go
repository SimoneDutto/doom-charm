package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"unsafe"

	gossh "golang.org/x/crypto/ssh"

	"github.com/creack/pty"
	"github.com/gliderlabs/ssh"
)

//go:embed embed/terminal-doom embed/doom1.wad
var embeddedFiles embed.FS

func extractFile(embeddedFS embed.FS, name, destDir string) (string, error) {
	data, err := embeddedFS.ReadFile(name)
	if err != nil {
		return "", err
	}

	destPath := filepath.Join(destDir, filepath.Base(name))
	err = os.WriteFile(destPath, data, 0755)
	if err != nil {
		return "", err
	}

	return destPath, nil
}

const (
	hostKey = `-----BEGIN OPENSSH PRIVATE KEY-----
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
fi543r7xawAAAAthbGV4QGxhcHRvcAE=
-----END OPENSSH PRIVATE KEY-----
`
)

func setWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}

func main() {
	tmpDir, err := os.MkdirTemp("", "doom-charm")
	if err != nil {
		log.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	terminalDoomPath, err := extractFile(embeddedFiles, "embed/terminal-doom", tmpDir)
	if err != nil {
		log.Fatalf("Failed to extract terminal-doom: %v", err)
	}

	_, err = extractFile(embeddedFiles, "embed/doom1.wad", tmpDir)
	if err != nil {
		log.Fatalf("Failed to extract doom1.wad: %v", err)
	}

	server := &ssh.Server{
		Addr: ":2223",
		PublicKeyHandler: func(ctx ssh.Context, key ssh.PublicKey) bool {
			return true
		},
		PasswordHandler: func(ctx ssh.Context, password string) bool {
			return true
		},
		Handler: func(s ssh.Session) {
			cmd := exec.Command(terminalDoomPath)
			cmd.Dir = tmpDir
			ptyReq, winCh, isPty := s.Pty()
			if isPty {
				cmd.Env = append(cmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
				f, err := pty.Start(cmd)
				if err != nil {
					panic(err)
				}
				go func() {
					for win := range winCh {
						setWinsize(f, win.Width, win.Height)
					}
				}()
				go func() {
					io.Copy(f, s) // stdin
				}()
				io.Copy(s, f) // stdout
				cmd.Wait()
			} else {
				io.WriteString(s, "No PTY requested.\n")
				s.Exit(1)
			}
		},
	}
	s, err := gossh.ParsePrivateKey([]byte(hostKey))
	if err != nil {
		log.Fatal(err)
	}
	server.AddHostKey(s)
	log.Println("Starting SSH controller server on :2223")
	log.Fatal(server.ListenAndServe())
}
