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

package cmd

import (
	sensortag_temperature_example "github.com/kupihse/go-bluetooth/examples/sensortag_temperature"
	"github.com/spf13/cobra"
)

// sensortagTemperatureCmd represents the sensortagTemperature command
var sensortagTemperatureCmd = &cobra.Command{
	Use:   "sensortag-temperature",
	Short: "Receives SensorTag temperature updates",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		adapterID, err := cmd.Flags().GetString("adapterID")
		if err != nil {
			fail(err)
		}

		if len(args) < 1 {
			failArgs([]string{"sensortag_address"})
		}

		fail(sensortag_temperature_example.Run(args[0], adapterID))
	},
}

func init() {
	rootCmd.AddCommand(sensortagTemperatureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sensortagTemperatureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sensortagTemperatureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
