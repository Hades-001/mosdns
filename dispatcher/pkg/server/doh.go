//     Copyright (C) 2020-2021, IrineSistiana
//
//     This file is part of mosdns.
//
//     mosdns is free software: you can redistribute it and/or modify
//     it under the terms of the GNU General Public License as published by
//     the Free Software Foundation, either version 3 of the License, or
//     (at your option) any later version.
//
//     mosdns is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU General Public License for more details.
//
//     You should have received a copy of the GNU General Public License
//     along with this program.  If not, see <https://www.gnu.org/licenses/>.

package server

import (
	"net/http"
	"time"
)

// startDoH always returns a non-nil error.
func (s *Server) startDoH() error {
	return s.buildHttpServer().ServeTLS(s.listener, s.cert, s.key)
}

// startHttp always returns a non-nil error.
func (s *Server) startHttp() error {
	return s.buildHttpServer().Serve(s.listener)
}

func (s *Server) buildHttpServer() *http.Server {
	return &http.Server{
		Handler:        s.httpHandler,
		TLSConfig:      s.tlsConfig,
		ReadTimeout:    time.Second * 5,
		WriteTimeout:   time.Second * 5,
		IdleTimeout:    s.getIdleTimeout(),
		MaxHeaderBytes: 2048,
	}
}
