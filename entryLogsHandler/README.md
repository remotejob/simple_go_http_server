# entryLogsHandler
--
    import "github.com/remotejob/simple_go_http_server/entryLogsHandler"

used to keep records slice of struc in pointer it ad new records and delete not
used records in pointer keep clean Database (.csv file but it can be other
backend for example redis etc..)

## Usage

#### type EntryLog

```go
type EntryLog struct {
	EntryLog []domains.Log
}
```


#### func  NewEntryLog

```go
func NewEntryLog() *EntryLog
```

#### func (*EntryLog) AddLastRecords

```go
func (logs *EntryLog) AddLastRecords(file string, deltaTime time.Duration, init bool)
```
AddLastRecords used mostly for control database (in our case simple .csv file)
it keep only needed records. 1. It used at spart up (init func) init paramaters
must be set TRUE 2. Used as goroutines periodicly cleapUp Database init
paramaters must be set FALSE

#### func (*EntryLog) AddNewHit

```go
func (logs *EntryLog) AddNewHit(logstr string)
```
AddNewHit add last click hit in pointer it used as goroutines so it help for
speed improvement

#### func (*EntryLog) DeleteExtraRecords

```go
func (logs *EntryLog) DeleteExtraRecords(index int)
```
DeleteExtraRecords so it delete not used records from slice as well used in
goroutines
