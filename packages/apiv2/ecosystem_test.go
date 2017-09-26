// Copyright 2016 The go-daylight Authors
// This file is part of the go-daylight library.
//
// The go-daylight library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-daylight library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-daylight library. If not, see <http://www.gnu.org/licenses/>.

package apiv2

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/EGaaS/go-egaas-mvp/packages/converter"
	"github.com/EGaaS/go-egaas-mvp/packages/crypto"
)

func TestNewEcosystem(t *testing.T) {
	if err := keyLogin(1); err != nil {
		t.Error(err)
		return
	}
	form := url.Values{`Name`: {``}}
	if _, result, err := postTxResult(`NewEcosystem`, &form); err != nil {
		t.Error(err)
		return
	} else {
		var ret ecosystemsResult
		err := sendGet(`ecosystems`, nil, &ret)
		if err != nil {
			t.Error(err)
			return
		}
		if int64(ret.Number) != converter.StrToInt64(result) {
			t.Error(fmt.Errorf(`Ecosystems %d != %s`, ret.Number, result))
			return
		}
	}

	form = url.Values{`Name`: {crypto.RandSeq(13)}}
	if err := postTx(`NewEcosystem`, &form); err != nil {
		t.Error(err)
		return
	}
}

func TestEcosystemParams(t *testing.T) {
	if err := keyLogin(1); err != nil {
		t.Error(err)
		return
	}
	var ret ecosystemParamsResult
	err := sendGet(`ecosystemparams`, nil, &ret)
	if err != nil {
		t.Error(err)
		return
	}
	if len(ret.List) < 5 {
		t.Error(fmt.Errorf(`wrong count of parameters %d`, len(ret.List)))
	}
	err = sendGet(`ecosystemparams?names=ecosystem_name,new_table&idstate=1`, nil, &ret)
	if err != nil {
		t.Error(err)
		return
	}
	if len(ret.List) != 2 {
		t.Error(fmt.Errorf(`wrong count of parameters %d`, len(ret.List)))
	}
}

func TestEcosystemParam(t *testing.T) {
	if err := keyLogin(1); err != nil {
		t.Error(err)
		return
	}
	var ret paramValue
	err := sendGet(`ecosystemparam/changing_menu`, nil, &ret)
	if err != nil {
		t.Error(err)
		return
	}
	if ret.Value != "ContractConditions(`MainCondition`)" {
		t.Error(err)
		return
	}
	err = sendGet(`ecosystemparam/myval`, nil, &ret)
	if err != nil {
		t.Error(err)
		return
	}
	if len(ret.Value) != 0 {
		t.Error(err)
		return
	}
}
