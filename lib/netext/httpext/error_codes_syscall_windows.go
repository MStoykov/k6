/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2021 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package httpext

import (
	"fmt"
	"net"
	"os"
	"syscall"

	"github.com/davecgh/go-spew/spew"
)

func getOSSyscallErrorCode(e *net.OpError, se *os.SyscallError) (errCode, string) {
	spew.Dump(se.Unwrap())
	spew.Dump(uintptr(se.Unwrap().(syscall.Errno)))
	spew.Dump(int(se.Unwrap().(syscall.Errno)))
	switch se.Unwrap() {
	case syscall.WSAECONNRESET:
		return tcpResetByPeerErrorCode, fmt.Sprintf(tcpResetByPeerErrorCodeMsg, e.Op)
	default:
	}
	return 0, ""
}
