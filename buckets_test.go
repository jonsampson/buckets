package buckets_test

import (
	"reflect"
	"testing"

	"github.com/jonsampson/buckets"
)

type MockIntBucketFilling struct {
	filling int
	label   string
}

func (bf MockIntBucketFilling) GetFilling() int {
	return bf.filling
}

func Test_IntBucketSet_With_LeastContent(t *testing.T) {
	bucketSet := buckets.NewBucketSet[int, MockIntBucketFilling](buckets.FillLeastContent, 2)

	bucketSet.AddFilling(MockIntBucketFilling{
		filling: 1,
		label:   "one",
	})
	bucketSet.AddFilling(MockIntBucketFilling{
		filling: 2,
		label:   "two",
	})
	bucketSet.AddFilling(MockIntBucketFilling{
		filling: 3,
		label:   "three",
	})
	bucketSet.AddFilling(MockIntBucketFilling{
		filling: 4,
		label:   "four",
	})
	bucketSet.AddFilling(MockIntBucketFilling{
		filling: 5,
		label:   "five",
	})
	bucketSet.AddFilling(MockIntBucketFilling{
		filling: 6,
		label:   "six",
	})

	startingOneBucket := buckets.Bucket[int, MockIntBucketFilling]{
		TotalFill: 10,
		Contents: []MockIntBucketFilling{
			{
				filling: 1,
				label:   "one",
			},
			{
				filling: 4,
				label:   "four",
			},
			{
				filling: 5,
				label:   "five",
			},
		},
	}
	startingOneCorrect := reflect.DeepEqual(bucketSet.Buckets[0], startingOneBucket) ||
		reflect.DeepEqual(bucketSet.Buckets[1], startingOneBucket)

	if !startingOneCorrect {
		t.Errorf("no bucket contained expected starting one items %v", bucketSet)
	}

	startingTwoBucket := buckets.Bucket[int, MockIntBucketFilling]{
		TotalFill: 11,
		Contents: []MockIntBucketFilling{
			{
				filling: 2,
				label:   "two",
			},
			{
				filling: 3,
				label:   "three",
			},
			{
				filling: 6,
				label:   "six",
			},
		},
	}
	startingTwoCorrect := reflect.DeepEqual(bucketSet.Buckets[0], startingTwoBucket) ||
		reflect.DeepEqual(bucketSet.Buckets[1], startingTwoBucket)

	if !startingTwoCorrect {
		t.Errorf("no bucket contained expected starting two items %v", bucketSet)
	}
}
