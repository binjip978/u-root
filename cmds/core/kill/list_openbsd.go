// Copyright 2024 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"syscall"
)

var (
	signames = []string{
		"SIGHUP",
		"SIGINT",
		"SIGQUIT",
		"SIGILL",
		"SIGTRAP",
		"SIGABRT",
		"SIGEMT",
		"SIGFPE",
		"SIGKILL",
		"SIGBUS",
		"SIGSEGV",
		"SIGSYS",
		"SIGPIPE",
		"SIGALRM",
		"SIGTERM",
		"SIGURG",
		"SIGSTOP",
		"SIGTSTP",
		"SIGCONT",
		"SIGCHLD",
		"SIGTTIN",
		"SIGTTOU",
		"SIGIO",
		"SIGXCPU",
		"SIGXFSZ",
		"SIGVTALRM",
		"SIGPROF",
		"SIGWINCH",
		"SIGINFO",
		"SIGUSR1",
		"SIGUSR2",
		"SIGTHR",
	}

	signums = map[string]os.Signal{
		"SIGHUP":    syscall.Signal(1),
		"SIGINT":    syscall.Signal(2),
		"SIGQUIT":   syscall.Signal(3),
		"SIGILL":    syscall.Signal(4),
		"SIGTRAP":   syscall.Signal(5),
		"SIGABRT":   syscall.Signal(6),
		"SIGEMT":    syscall.Signal(7),
		"SIGFPE":    syscall.Signal(8),
		"SIGKILL":   syscall.Signal(9),
		"SIGBUS":    syscall.Signal(10),
		"SIGSEGV":   syscall.Signal(11),
		"SIGSYS":    syscall.Signal(12),
		"SIGPIPE":   syscall.Signal(13),
		"SIGALRM":   syscall.Signal(14),
		"SIGTERM":   syscall.Signal(15),
		"SIGURG":    syscall.Signal(16),
		"SIGSTOP":   syscall.Signal(17),
		"SIGTSTP":   syscall.Signal(18),
		"SIGCONT":   syscall.Signal(19),
		"SIGCHLD":   syscall.Signal(20),
		"SIGTTIN":   syscall.Signal(21),
		"SIGTTOU":   syscall.Signal(22),
		"SIGIO":     syscall.Signal(23),
		"SIGXCPU":   syscall.Signal(24),
		"SIGXFSZ":   syscall.Signal(25),
		"SIGVTALRM": syscall.Signal(26),
		"SIGPROF":   syscall.Signal(27),
		"SIGWINCH":  syscall.Signal(28),
		"SIGINFO":   syscall.Signal(29),
		"SIGUSR1":   syscall.Signal(30),
		"SIGUSR2":   syscall.Signal(31),
		"SIGTHR":    syscall.Signal(31),
		"1":         syscall.Signal(1),
		"2":         syscall.Signal(2),
		"3":         syscall.Signal(3),
		"4":         syscall.Signal(4),
		"5":         syscall.Signal(5),
		"6":         syscall.Signal(6),
		"7":         syscall.Signal(7),
		"8":         syscall.Signal(8),
		"9":         syscall.Signal(9),
		"10":        syscall.Signal(10),
		"11":        syscall.Signal(11),
		"12":        syscall.Signal(12),
		"13":        syscall.Signal(13),
		"14":        syscall.Signal(14),
		"15":        syscall.Signal(15),
		"16":        syscall.Signal(16),
		"17":        syscall.Signal(17),
		"18":        syscall.Signal(18),
		"19":        syscall.Signal(19),
		"20":        syscall.Signal(20),
		"21":        syscall.Signal(21),
		"22":        syscall.Signal(22),
		"23":        syscall.Signal(23),
		"24":        syscall.Signal(24),
		"25":        syscall.Signal(25),
		"26":        syscall.Signal(26),
		"27":        syscall.Signal(27),
		"28":        syscall.Signal(28),
		"29":        syscall.Signal(29),
		"30":        syscall.Signal(30),
		"31":        syscall.Signal(31),
		"32":        syscall.Signal(32),
	}
)
