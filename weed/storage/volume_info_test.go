package storage

import "testing"
import (
	"github.com/stretchr/testify/assert"
	"github.com/rs/xid"
)

func TestSortVolumeInfos(t *testing.T) {
	v2vid, err := NewVolumeId(xid.New().String())
	if err != nil {
		t.Fatal(err)
	}

	v1vid1 := NewVolumeIdV1(1)
	v1vid2 := NewVolumeIdV1(2)
	v1vid3 := NewVolumeIdV1(3)


	vis := []*VolumeInfo{
		{
			Id: v2vid,
		},
		{
			Id: v1vid2,
		},
		{
			Id: v1vid3,
		},
		{
			Id: v1vid1,
		},
	}
	sortVolumeInfos(vis)
	assert.Equal(t, v1vid1.String(), vis[0].Id.String())
	assert.Equal(t, v1vid2.String(), vis[1].Id.String())
	assert.Equal(t, v1vid3.String(), vis[2].Id.String())
	assert.Equal(t, v2vid.String(), vis[3].Id.String())
	assert.Equal(t, VolumeIdVersion(2), vis[3].Id.Version)
}
