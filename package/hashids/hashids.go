package my_hashids

import (
	"github.com/speps/go-hashids"
)

var hash *hashids.HashID

func init() {
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 6

	var err error
	hash, err = hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}
}

func NewHashID(id int64) (string, error) {
	nums := append([]int64{}, id)
	hash_id, err := hash.EncodeInt64(nums)
	if err != nil {
		return "", err
	}
	return hash_id, nil
}

func ParseHashID(hash_id string) (int64, error) {
	nums, err := hash.DecodeInt64WithError(hash_id)
	if err != nil {
		return 0, err
	}
	return nums[0], nil
}
