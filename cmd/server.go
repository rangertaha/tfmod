/* The MIT License (MIT)

Copyright © 2019 rangertaha@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE. */

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/rangertaha/tfmod/pkg"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the private terraform module registry",
	Long:  `Start the api server with the private terraform module registry.`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Server()
	},
}

func init() {

	// Address host/port to serve the API
	serverCmd.Flags().String("http.host", "0.0.0.0", "Host to expose the api server")
	serverCmd.Flags().StringP("http.port", "p", "8080", "Port to expose the api server")

	serverCmd.Flags().String("log.level", "debug", "Log level to use")
	serverCmd.Flags().String("log.file", "/var/log/tfmod.log", "TFMod application log file")

	// Bind Cobra flags to Viper configuations
	viper.BindPFlag("http.host", serverCmd.Flags().Lookup("http.host"))
	viper.BindPFlag("http.port", serverCmd.Flags().Lookup("http.port"))
	// Set viper default values
	viper.SetDefault("http.host", serverCmd.Flags().Lookup("http.host"))
	viper.SetDefault("http.port", serverCmd.Flags().Lookup("http.port"))

	// Prometheus metrics
	serverCmd.Flags().String("metrics.host", "0.0.0.0", "Host to expose the metrics")
	serverCmd.Flags().String("metrics.port", "8000", "Port to expose metrics")
	serverCmd.Flags().String("metrics.path", "/metrics", "Metrics path to expose metrics")

	// Bind Cobra flags to Viper configuations
	viper.BindPFlag("metrics.host", serverCmd.Flags().Lookup("metrics.host"))
	viper.BindPFlag("metrics.port", serverCmd.Flags().Lookup("metrics.port"))
	viper.BindPFlag("metrics.path", serverCmd.Flags().Lookup("metrics.path"))
	// Set viper default values
	viper.SetDefault("metrics.host", serverCmd.Flags().Lookup("metrics.host"))
	viper.SetDefault("metrics.port", serverCmd.Flags().Lookup("metrics.port"))
	viper.SetDefault("metrics.path", serverCmd.Flags().Lookup("metrics.path"))

	rootCmd.AddCommand(serverCmd)
}
