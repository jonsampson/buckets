# buckets

Fill up a set of buckets with structs based on:
- fill level (`buckets.FillLeastFull`)
- how many items already exist (`buckets.FillLeastContent`)

Make the struct to be bucketed implement a `GetFilling() [FillType]` method, construct a `buckets.BucketSet` with
one of the above types and a number of buckets, then start adding. See the tests for examples.