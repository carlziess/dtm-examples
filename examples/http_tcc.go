/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/go-resty/resty/v2"
)

func init() {
	addSample("http_tcc_nested", func() string {
		gid := dtmcli.MustGenGid(dtmutil.DefaultHttpServer)
		err := dtmcli.TccGlobalTransaction(dtmutil.DefaultHttpServer, gid, func(tcc *dtmcli.Tcc) (*resty.Response, error) {
			resp, err := tcc.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TransOut", busi.Busi+"/TransOutConfirm", busi.Busi+"/TransOutRevert")
			if err != nil {
				return resp, err
			}
			return tcc.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TransInTccParent", busi.Busi+"/TransInConfirm", busi.Busi+"/TransInRevert")
		})
		logger.FatalIfError(err)
		return gid
	})
	addSample("http_tcc", func() string {
		logger.Debugf("tcc simple transaction begin")
		gid := dtmcli.MustGenGid(dtmutil.DefaultHttpServer)
		err := dtmcli.TccGlobalTransaction(dtmutil.DefaultHttpServer, gid, func(tcc *dtmcli.Tcc) (*resty.Response, error) {
			resp, err := tcc.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TransOut", busi.Busi+"/TransOutConfirm", busi.Busi+"/TransOutRevert")
			if err != nil {
				return resp, err
			}
			return tcc.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TransIn", busi.Busi+"/TransInConfirm", busi.Busi+"/TransInRevert")
		})
		logger.FatalIfError(err)
		return gid
	})
}
