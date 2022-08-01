// Copyright 2021, Chef.  All rights reserved.
// https://github.com/forkiss/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package hls

import (
	"github.com/forkiss/naza/pkg/mock"
	"github.com/forkiss/naza/pkg/nazalog"
)

var (
	PathStrategy IPathStrategy = &DefaultPathStrategy{}

	Clock = mock.NewStdClock()

	Log = nazalog.GetGlobalLogger()
)
