package config

const MaxRecords = 50000

type StatFile struct {
	Records    uint64
	Bytes      int
	Partitions int
}
