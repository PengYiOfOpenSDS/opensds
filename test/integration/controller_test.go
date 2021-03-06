// Copyright 2017 The OpenSDS Authors.
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

// +build integration

package integration

import (
	"encoding/json"
	"testing"

	"github.com/opensds/opensds/pkg/controller/volume"
	pb "github.com/opensds/opensds/pkg/dock/proto"
	"github.com/opensds/opensds/pkg/model"
)

var vc = volume.NewController()

var dckInfo = &model.DockSpec{
	Endpoint:   "localhost:50050",
	DriverName: "default",
}

func TestControllerCreateVolume(t *testing.T) {
	vc.SetDock(dckInfo)

	vol, err := vc.CreateVolume(&pb.CreateVolumeOpts{})
	if err != nil {
		t.Error("create volume in controller failed:", err)
		return
	}

	volBody, _ := json.MarshalIndent(vol, "", "	")
	t.Log(string(volBody))
}

func TestControllerDeleteVolume(t *testing.T) {
	vc.SetDock(dckInfo)

	err := vc.DeleteVolume(&pb.DeleteVolumeOpts{})
	if err != nil {
		t.Error("delete volume in controller failed:", err)
	}
}

func TestControllerCreateVolumeAttachment(t *testing.T) {
	vc.SetDock(dckInfo)

	atc, err := vc.CreateVolumeAttachment(&pb.CreateAttachmentOpts{})
	if err != nil {
		t.Error("create volume attachment in controller failed:", err)
		return
	}

	atcBody, _ := json.MarshalIndent(atc, "", "	")
	t.Log(string(atcBody))
}

func TestControllerDeleteVolumeAttachment(t *testing.T) {
	vc.SetDock(dckInfo)

	err := vc.DeleteVolumeAttachment(&pb.DeleteAttachmentOpts{})
	if err != nil {
		t.Error("delete volume attachment in controller failed:", err)
	}
}

func TestControllerCreateVolumeSnapshot(t *testing.T) {
	vc.SetDock(dckInfo)

	snp, err := vc.CreateVolumeSnapshot(&pb.CreateVolumeSnapshotOpts{})
	if err != nil {
		t.Error("create volume snapshot in controller failed:", err)
		return
	}

	snpBody, _ := json.MarshalIndent(snp, "", "	")
	t.Log(string(snpBody))
}

func TestControllerDeleteVolumeSnapshot(t *testing.T) {
	vc.SetDock(dckInfo)

	err := vc.DeleteVolumeSnapshot(&pb.DeleteVolumeSnapshotOpts{})
	if err != nil {
		t.Error("delete volume snapshot in controller failed:", err)
	}
}
