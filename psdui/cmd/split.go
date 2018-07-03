// Copyright © 2018 Li Tao<taodev@qq.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"log"
	"path/filepath"

	"github.com/koalakit/psdui"

	_ "github.com/koalakit/psdui/egret"
	"github.com/spf13/cobra"
)

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "切分图素",
	Long:  `切分图素`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		parser := psdui.NewPSDParser()
		err = parser.Load(args[0])
		if err != nil {
			log.Fatal(args[0], err)
		}

		for _, v := range parser.Nodes {
			absPath, err := filepath.Abs(args[1])
			if err != nil {
				log.Fatal("输出路径出错:", err)
			}

			psdui.Export(v, "egret", absPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(splitCmd)
}
