package mail

// Copyright 2023 CUE Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"io"
	"net/mail"
	"time"
)

// ParseDate parses an RFC 5322 date string.
func ParseDate(date string) (time.Time, error) {
	return mail.ParseDate(date)
}

// ParseAddress parses a single RFC 5322 address, e.g. "Barry Gibbs <bg@example.com>"
func ParseAddress(address string) (*mail.Address, error) {
	return mail.ParseAddress(address)
}

// ParseAddressList parses the given string as a list of addresses.
func ParseAddressList(list string) ([]*mail.Address, error) {
	return mail.ParseAddressList(list)
}

// ReadMessage reads a message from r.
// The headers are parsed, and the body of the message will be available
// for reading from msg.Body.
func ReadMessage(r io.Reader) (msg *mail.Message, err error) {
	return mail.ReadMessage(r)
}
