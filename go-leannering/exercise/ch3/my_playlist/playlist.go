// Copyright © 2011-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Song struct {
	Title    string
	Filename string
	Seconds  int
}

type PlsType int

const (
	M3U PlsType = iota
	PLS
)

func (ptype PlsType) String() string {
	switch ptype {
	case M3U:
		return "M3U"
	case PLS:
		return "PLS"
	default:
		return "NULL"
	}
}

func main() {
	if len(os.Args) == 1 || !(strings.HasSuffix(os.Args[1], ".m3u") || strings.HasSuffix(os.Args[1], ".pls")) {
		fmt.Printf("usage: %s <file.m3u|pls>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	pType := M3U
	if strings.HasSuffix(os.Args[1], ".pls") {
		pType = PLS
	}

	if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		if songs, err := readPlaylist(pType, string(rawBytes)); err != nil {
			fmt.Printf("Error : %s", err)
		} else {
			writePlaylist(pType, songs)
		}
	}
}

func readPlaylist(pType PlsType, data string) (songs []Song, err error) {
	var song Song
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#EXTM3U") {
			if pType != M3U {
				return songs, fmt.Errorf("The file should be a type of %s", M3U)
			}
			continue
		} else if line == "" || strings.HasPrefix(line, "[playlist]") {
			if pType != PLS {
				return songs, fmt.Errorf("The file should be a type of %s", PLS)
			}
			continue
		}
		if pType == M3U {
			if strings.HasPrefix(line, "#EXTINF:") {
				song.Title, song.Seconds = parseExtinfLine(line)
			} else {
				song.Filename = strings.Map(mapPlatformDirSeparator, line)
			}
		} else if pType == PLS {
			if strings.HasPrefix(line, "File") {
				song.Filename = strings.Map(mapPlatformDirSeparator, strings.Split(line, "=")[1])
			} else if strings.HasPrefix(line, "Title") {
				song.Title = strings.Split(line, "=")[1]
			} else if strings.HasPrefix(line, "Length") {
				if sec, err := strconv.Atoi(strings.Split(line, "=")[1]); err != nil {
					song.Seconds = -1
				} else {
					song.Seconds = sec
				}
			}
		}
		//
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs, nil
}

func parseExtinfLine(line string) (title string, seconds int) {
	if i := strings.IndexAny(line, "-0123456789"); i > -1 {
		const separator = ","
		line = line[i:]
		if j := strings.Index(line, separator); j > -1 {
			title = line[j+len(separator):]
			var err error
			if seconds, err = strconv.Atoi(line[:j]); err != nil {
				log.Printf("failed to read the duration for '%s': %v\n",
					title, err)
				seconds = -1
			}
		}
	}
	return title, seconds
}

func mapPlatformDirSeparator(char rune) rune {
	if char == '/' || char == '\\' {
		return filepath.Separator
	}
	return char
}

func writePlaylist(pType PlsType, songs []Song) {
	switch pType {
	case PLS:
		writeM3UPlaylist(songs)
	case M3U:
		writePlsPlaylist(songs)
	}
}

func writePlsPlaylist(songs []Song) {
	fmt.Println("[playlist]")
	for i, song := range songs {
		i++
		fmt.Printf("File%d=%s\n", i, song.Filename)
		fmt.Printf("Title%d=%s\n", i, song.Title)
		fmt.Printf("Length%d=%d\n", i, song.Seconds)
	}
	fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
}

func writeM3UPlaylist(songs []Song) {
	fmt.Println("#EXTM3U")
	for _, song := range songs {
		fmt.Printf("EXTINF:%d,%s\n", song.Seconds, song.Title)
		fmt.Printf("%s\n", song.Filename)
	}
}
