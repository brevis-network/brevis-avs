/*
Copyright © 2024 Brevis Network
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var ChainCache = make(map[uint64]*OneChain)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "monitor Brevis requests and send signed to gateway",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		mcc := GetMcc(kMultichain)
		for _, onecfg := range mcc {
			onc, err := NewOneChain(onecfg)
			if err != nil {
				log.Println("setup err:", err, "config:", *onecfg)
				continue
			}
			ChainCache[onc.ChainID] = onc
			onc.MonBrevisReq()
			defer onc.Close()
		}
		// block forever
		<-make(chan bool)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
