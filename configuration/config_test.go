/*
   Copyright 2022 The Queer Global Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package configuration

import (
	"math"
	"os"
	"testing"
)

func TestConfigFile(t *testing.T) {

	initPath := "unittest-init.json"

	initObj, err := NewInitParams(&initPath)
	if err != nil {
		t.Error(initObj)
	}

	configObj, err := NewConfigFetcher(initObj)
	if err != nil {
		t.Error(err)
	}
	conf, err := configObj.FetchConfig()
	if err != nil {
		t.Error(err)
	}
	confvals := conf.GetMap()

	inner, ok := (*confvals)["testvalues"]
	if !ok {
		t.Error()
	}
	testvaluesMap, ok := inner.(map[string]interface{})
	if !ok {
		t.Error()
	}
	test1, ok := testvaluesMap["test1"]
	if !ok {
		t.Error()
	}
	if !(test1 == "test value 1") {
		t.Error()
	}
	var i64 *int64
	var f64 *float64
	var i32 *int32

	i64, err = conf.GetInt64("test5")
	if err != nil {
		t.Error(err)
	} else {
		if !(*i64 == 5) {
			t.Error()
		}
	}

	f64, err = conf.GetFloat64("test4")
	if err != nil {
		t.Error(err)
	} else {
		if !(math.Round(*f64) == 4) {
			t.Error(err)
		}
	}

	i32, err = conf.GetInt32("test5")
	if err != nil {
		t.Error(err)
	} else {
		if !(*i32 == 5) {
			t.Error()
		}
	}
}

func TestConfigEnvVar(t *testing.T) {

	os.Setenv("TEST_ENV_VAR", "test value 1")
	os.Setenv("TEST_ENV_VAR2", "2")

	initPath := "unittest-init-envvars.json"

	initObj, err := NewInitParams(&initPath)
	if err != nil {
		t.Error(initObj)
	}

	conf, err := GetConfig(&initPath)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	test1, err := conf.Get("test1")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if !(test1 == "test value 1") {
		t.Log(err)
		t.Fail()
	}
	var i64 *int64

	i64, err = conf.GetInt64("test2")
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		if !(*i64 == 2) {
			t.Log(err)
			t.Fail()
		}
	}

}
