/*
   Dirty Step-CA certificate decoder
   Copyright (C) 2025  Georg Pfuetzenreuter

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"bufio"
	"bytes"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

type lines [][]byte

func main() {
	for _, l := range readCertificates() {
		parseCertificate(l)
	}
}

func readCertificates() lines {
	scanner := bufio.NewScanner(os.Stdin)
	var input lines
	for scanner.Scan() {
		input = append(input, scanner.Bytes())
	}

	return input
}

func parseCertificate(data []byte) {
	// in PostgreSQL the value is prefixed with "/x"
	if bytes.Equal([]byte{92, 120}, data[0:2]) {
		data = data[2:]
	}

	crtPre := make([]byte, hex.DecodedLen(len(data)))
	_, err := hex.Decode(crtPre, data)
	if err != nil {
		panic(err)
	}

	crt, err := x509.ParseCertificate(crtPre)
	if err != nil {
		panic(err)
	}

	// the below is just for pretty-printing

	jcrt, err := json.MarshalIndent(*crt, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jcrt))
}
