/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// TODO flag
const max int = 6

func getNotes() map[string][]int {

	var notes = make(map[string][]int)
	notes["Ab"] = []int{4, 9, 1, 6, 11, 4}
	notes["A"] = []int{5, 10, 2, 7, 12, 5}
	notes["A#"] = []int{6, 11, 3, 8, 1, 6}
	notes["Bb"] = []int{6, 11, 3, 8, 1, 6}
	notes["B"] = []int{7, 12, 4, 9, 2, 7}
	notes["C"] = []int{8, 1, 5, 10, 3, 8}
	notes["C#"] = []int{9, 2, 6, 11, 4, 9}
	notes["Db"] = []int{9, 2, 6, 11, 4, 9}
	notes["D"] = []int{10, 3, 7, 12, 5, 10}
	notes["D#"] = []int{11, 4, 8, 1, 6, 11}
	notes["Eb"] = []int{11, 4, 8, 1, 6, 11}
	notes["E"] = []int{12, 5, 9, 2, 7, 12}
	notes["F"] = []int{1, 6, 10, 3, 8, 1}
	notes["F#"] = []int{2, 7, 11, 4, 9, 2}
	notes["Gb"] = []int{2, 7, 11, 4, 9, 2}
	notes["G"] = []int{3, 8, 12, 5, 10, 3}
	notes["G#"] = []int{4, 9, 1, 6, 11, 4}

	return notes
}

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train [bpm] [note...]",
	Short: "A brief description of your command",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		bpm, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return fmt.Errorf("converting bpm input to int: %v", err)

		}

		requestedNotes := args[1:]
		notes := getNotes()

		for _, note := range requestedNotes {
			if len(notes[note]) == 0 {
				return fmt.Errorf("invalid note: %s", note)
			}
		}

		if len(requestedNotes) == 0 {
			requestedNotes = []string{"Ab", "A", "A#", "Bb", "B", "C", "C#", "Db", "D", "D#", "Eb", "E", "F", "F#", "Gb", "G", "G#"}
		}

		interval := float64(60.0/bpm) * 1000

		for i := 3; i > 0; i-- {
			fmt.Println(i)
			time.Sleep(time.Second)
		}

		for {
			guitarString := rand.Intn(max) + 1
			note := requestedNotes[rand.Intn(len(requestedNotes))]
			fmt.Print(note + strconv.Itoa(guitarString))
			time.Sleep(time.Duration(interval) * time.Millisecond)
			fmt.Printf(" -> %d\n", notes[note][guitarString-1])

		}

	},
}

func init() {
	rootCmd.AddCommand(trainCmd)
	// trainCmd.Flags().IntP("bpm", "b", 40, "bpm to run at")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
