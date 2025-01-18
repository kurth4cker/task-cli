// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import "log"

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
