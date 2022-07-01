package _interface

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Sort() {
	var names = []string{"bb", "a"}
	sort.Sort(StringSlice(names))
	fmt.Println(names)
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var Tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type ByYear []*Track

func (x ByYear) Len() int           { return len(x) }
func (x ByYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x ByYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func SortByArtist() {
	sort.Sort(ByYear(Tracks))
	PrintTracks(Tracks)
}
