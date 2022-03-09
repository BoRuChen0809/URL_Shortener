package my_hashids_test

import (
	my_hashids "URL_Shortener/package/hashids"
	"testing"
)

func TestHashIDs(t *testing.T) {
	var i int64
	for i = 1; i < 10000; i++ {
		hash_id, err := my_hashids.NewHashID(i)
		if err != nil {
			t.Error(err)
		}
		id, err := my_hashids.ParseHashID(hash_id)
		if err != nil {
			t.Error(err)
		}
		if i != id {
			t.Errorf("i:%d != id:%d\n", i, id)
		}
		t.Logf("%d : %s\n", i, hash_id)
	}
}
