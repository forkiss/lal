// Copyright 2021, Chef.  All rights reserved.
// https://github.com/forkiss/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package remux

import (
	"testing"

	"github.com/forkiss/lal/pkg/base"
	"github.com/forkiss/lal/pkg/mpegts"
	"github.com/forkiss/naza/pkg/assert"
)

var (
	fh    []byte
	poped []base.RtmpMsg
)

type qo struct {
}

func (q *qo) onPatPmt(b []byte) {
	fh = b
}

func (q *qo) onPop(msg base.RtmpMsg) {
	poped = append(poped, msg)
}

func TestRtmp2MpegtsFilter(t *testing.T) {
	goldenRtmpMsg := []base.RtmpMsg{
		{
			Header: base.RtmpHeader{
				MsgTypeId: base.RtmpTypeIdAudio,
			},
			Payload: []byte{0xAF},
		},
		{
			Header: base.RtmpHeader{
				MsgTypeId: base.RtmpTypeIdVideo,
			},
			Payload: []byte{0x17},
		},
	}

	q := &qo{}
	f := newRtmp2MpegtsFilter(8, q)
	for i := range goldenRtmpMsg {
		f.Push(goldenRtmpMsg[i])
	}
	assert.Equal(t, mpegts.FixedFragmentHeader, fh)
	assert.Equal(t, goldenRtmpMsg, poped)
}
