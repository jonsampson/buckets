package buckets

import "sort"

type FillType interface {
	~int | ~float32
}

type hasFilling[F FillType] interface {
	GetFilling() F
}

type Bucket[F FillType, H hasFilling[F]] struct {
	TotalFill F
	Contents  []H
}

func NewBucket[F FillType, H hasFilling[F]]() *Bucket[F, H] {
	return &Bucket[F, H]{
		TotalFill: 0,
		Contents:  make([]H, 0),
	}
}

func (b *Bucket[F, H]) AddFilling(h H) {
	b.TotalFill = b.TotalFill + h.GetFilling()
	b.Contents = append(b.Contents, h)
}

type BucketSetType int64

const (
	FillLeastFull BucketSetType = iota
	FillLeastContent
)

type BucketSet[F FillType, H hasFilling[F]] struct {
	Buckets []Bucket[F, H]
	Type    BucketSetType
}

func NewBucketSet[F FillType, H hasFilling[F]](t BucketSetType, numberOfBuckets int) *BucketSet[F, H] {
	buckets := make([]Bucket[F, H], numberOfBuckets)
	for i, _ := range buckets {
		buckets[i] = *NewBucket[F, H]()
	}
	return &BucketSet[F, H]{
		Buckets: buckets,
		Type:    t,
	}
}

func (bs *BucketSet[F, H]) GetLeastFullBucket() *Bucket[F, H] {
	sort.Slice(bs.Buckets, func(i, j int) bool {
		return bs.Buckets[i].TotalFill < bs.Buckets[j].TotalFill
	})
	return &bs.Buckets[0]
}

func (bs *BucketSet[F, H]) GetLeastContentBucket() *Bucket[F, H] {
	sort.Slice(bs.Buckets, func(i, j int) bool {
		return len(bs.Buckets[i].Contents) < len(bs.Buckets[j].Contents)
	})
	return &bs.Buckets[0]
}

func (bs *BucketSet[F, H]) AddFilling(h H) {
	switch bs.Type {
	case FillLeastFull:
		bs.GetLeastFullBucket().AddFilling(h)
	case FillLeastContent:
		bs.GetLeastContentBucket().AddFilling(h)
	}
}
