// Copyright 2019, Chef.  All rights reserved.
// https://github.com/forkiss/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/forkiss/lal/pkg/base"
	"github.com/forkiss/lal/pkg/logic"
	"github.com/forkiss/naza/pkg/bininfo"
	"github.com/forkiss/naza/pkg/nazalog"
)

func main() {
	defer nazalog.Sync()

	confFilename := parseFlag()
	lals := logic.NewLalServer(func(option *logic.Option) {
		option.ConfFilename = confFilename
	})
	err := lals.RunLoop()
	nazalog.Infof("server manager done. err=%+v", err)
}

func parseFlag() string {
	binInfoFlag := flag.Bool("v", false, "show bin info")
	cf := flag.String("c", "", "specify conf file")
	p := flag.String("p", "", "specify current work directory")
	flag.Parse()

	if *binInfoFlag {
		_, _ = fmt.Fprint(os.Stderr, bininfo.StringifyMultiLine())
		_, _ = fmt.Fprintln(os.Stderr, base.LalFullInfo)
		os.Exit(0)
	}
	if *p != "" {
		os.Chdir(*p)
	}

	return *cf
}
