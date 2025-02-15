// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

type Artifact struct {
	// The name of the snapshot
	snapshotName string

	// The ID of the image
	snapshotId int

	// The hcloudClient for making API calls
	hcloudClient *hcloud.Client

	// StateData should store data such as GeneratedData
	// to be shared with post-processors
	StateData map[string]interface{}
}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (*Artifact) Files() []string {
	return nil
}

func (a *Artifact) Id() string {
	return strconv.Itoa(a.snapshotId)
}

func (a *Artifact) String() string {
	return fmt.Sprintf("A snapshot was created: '%v' (ID: %v)", a.snapshotName, a.snapshotId)
}

func (a *Artifact) State(name string) interface{} {
	return a.StateData[name]
}

func (a *Artifact) Destroy() error {
	log.Printf("Destroying image: %d (%s)", a.snapshotId, a.snapshotName)
	_, err := a.hcloudClient.Image.Delete(context.TODO(), &hcloud.Image{ID: a.snapshotId})
	return err
}
