package cmd

import (
	"github.com/ljesparis/govm/pkg/utils"
	"github.com/spf13/cobra"
)

func getCorrectPackage(version string, cmd *cobra.Command) (p string, err error) {
	os, _ := cmd.Flags().GetString("os")
	arch, _ := cmd.Flags().GetString("arch")
	pt, _ := cmd.Flags().GetString("package")

	p, err = utils.PackageFilename2(version, os, arch, pt)
	return
}
