package pty

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/u-root/u-root/pkg/termios"
)

func New() (*Pty, error) {
	tty, err := termios.New()
	if err != nil {
		return nil, err
	}
	restorer, err := tty.Get()
	if err != nil {
		return nil, err
	}

	ptm, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	sname, err := ptsname(ptm)
	if err != nil {
		return nil, err
	}

	if err := grantpt(ptm); err != nil {
		return nil, err
	}

	pts, err := os.OpenFile(sname, os.O_RDWR|syscall.O_NOCTTY, 0)
	if err != nil {
		return nil, err
	}
	return &Pty{Ptm: ptm, Pts: pts, Sname: sname, Kid: -1, TTY: tty, Restorer: restorer}, nil
}

type ptmget struct {
	Cfd int32
	Sfd int32
	Cn  [16]int8
	Sn  [16]int8
}

func ptsname(f *os.File) (string, error) {
	/*
	 * from ptsname(3): The ptsname() function is equivalent to:
	 * struct ptmget pm;
	 * ioctl(fd, TIOCPTSNAME, &pm) == -1 ? NULL : pm.sn;
	 */
	var ptm ptmget
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), syscall.TIOCPTSNAME, uintptr(unsafe.Pointer(&ptm)))
	if err != 0 {
		return "", fmt.Errorf("failed")
	}
	name := make([]byte, len(ptm.Sn))
	for i, c := range ptm.Sn {
		name[i] = byte(c)
		if c == 0 {
			return string(name[:i]), nil
		}
	}
	return "", errors.New("TIOCPTSNAME string not NUL-terminated")
}

func grantpt(ptm *os.File) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, ptm.Fd(), syscall.TIOCGRANTPT, 0)
	if err != 0 {
		return fmt.Errorf("grant failed")
	}
	return nil
}

func sysNetBSD(p *Pty) {
	p.C.SysProcAttr = &syscall.SysProcAttr{Setctty: true, Setsid: true}
}

func init() {
	sys = sysNetBSD
}
