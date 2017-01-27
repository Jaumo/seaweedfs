package storage

import (
	"strconv"
	"github.com/rs/xid"
	"encoding/binary"
)

type VolumeIdVersion uint8;

type VolumeId struct {
	v1Id    uint64
	v2Id    xid.ID
	Version VolumeIdVersion
}

func NewVolumeIdV1(volumeId uint64) VolumeId {
	return VolumeId{v1Id: volumeId, Version: 1}
}

func EmptyVolumeId() VolumeId {
	return VolumeId{Version: 0}
}

func NewVolumeId(vid string) (VolumeId, error) {
	if (len(vid) == 20) {
		id, err := xid.FromString(vid)
		return VolumeId{v2Id: id, Version: 2}, err
	}

	volumeId, err := strconv.ParseUint(vid, 10, 64)
	return VolumeId{v1Id: volumeId, Version: 1}, err
}

func (vid *VolumeId) Comparable() bool {
	return vid.Version == 1
}

func (vid *VolumeId) Less(cmp *VolumeId) bool {
	if (vid.Version < cmp.Version) {
		return true
	}

	if (vid.Version == cmp.Version) {
		switch (vid.Version) {
		case 1:
			return vid.v1Id < cmp.v1Id
		case 2:
			return vid.String() < cmp.String()
		}
	}

	return false
}

func (vid *VolumeId) Int() uint64 {
	switch vid.Version {
	case 2:
		return binary.BigEndian.Uint64(vid.v2Id[0:7])
	default:
	}
	return vid.v1Id
}

func (vid *VolumeId) String() string {
	switch vid.Version {
	case 2:
		return vid.v2Id.String()
	default:
	}
	return strconv.FormatUint(vid.v1Id, 10)
}

func (vid *VolumeId) Next() VolumeId {
	switch vid.Version {
	case 2:
		return VolumeId{v2Id:xid.New(), Version: 2}
	default:
	}
	return VolumeId{v1Id: vid.v1Id + 1, Version: 1}
}
