/*
Copyright © 2024 Brevis Network
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/cobra"
)

var (
	ChainCache = make(map[uint64]*OneChain)
	runPort    int
)

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

		router := httprouter.New()
		router.GET("/eigen/node/health", handleHealth)
		// block forever
		err := http.ListenAndServe(fmt.Sprintf(":%d", runPort), router)
		if err != nil {
			log.Println(err)
		}
	},
}

func handleHealth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().IntVar(&runPort, "port", 8081, "listen port for rpc")
}
