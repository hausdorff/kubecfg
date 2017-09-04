// Copyright 2017 The kubecfg authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ksonnet/kubecfg/pkg/kubecfg"
)

const (
	flagFormat = "format"
)

func init() {
	RootCmd.AddCommand(showCmd)
	addJsonnetFlagsToCmd(showCmd)
	addKubectlFlagsToCmd(showCmd)
	addEnvCmdFlags(showCmd)
	showCmd.PersistentFlags().StringP(flagFormat, "o", "yaml", "Output format.  Supported values are: json, yaml")
}

var showCmd = &cobra.Command{
	Use:   "show [<env>|-f <file-or-dir>]",
	Short: "Render resources and print to stdout",
	RunE: func(cmd *cobra.Command, args []string) error {
		flags := cmd.Flags()
		var err error

		c := kubecfg.ShowCmd{}

		c.Format, err = flags.GetString(flagFormat)
		if err != nil {
			return err
		}

		objs, err := readObjs(cmd, args)
		if err != nil {
			return err
		}

		return c.Run(objs, cmd.OutOrStdout())
	},
	Long: `Render resources and print them to stdout. This subcommand accepts
YAML, JSON, or Jsonnet files, as well as ksonnet applications. In the case of
YAML and JSON, the rendering step is a no-op, and they are printed out
verbatim. In the case of Jsonnet (and, hence, ksonnet applications), the render
step involves expanding the file and printing out the results.`,
}
