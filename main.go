// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

/*
PLAN:
- [TODO] tasks.json would not be an json array, it would be series of json objects
- [TODO] we need a Task structure for json
- [TODO] we need an Set structure
	- [TODO] Set should parse file which is a series of json Task objects
	- [TODO] Set should lookup for objects with given Id
	- [TODO] Methods:
		- [TODO] Add(id)
		- [TODO] Del(id)
		- [TODO] Update(id)
*/

/*
Task management CLI application

Created for a project idea taken from roadmap.sh
*/
package main

import "codeberg.org/kurth4cker/go-sample"

/*
TODO: write sub commands here
*/
func main() {
	sample.Helloln("world")
}
