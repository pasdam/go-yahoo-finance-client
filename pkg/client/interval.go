package client

// Interval enumeration of all supported aggregation interval
type Interval uint8

const (
	// Interval1Minute indicates an interval of 1 minute
	Interval1Minute Interval = iota

	// Interval5Minutes indicates an interval of 5 minutes
	Interval5Minutes Interval = iota

	// Interval15Minutes indicates an interval of 15 minutes
	Interval15Minutes Interval = iota

	// Interval30Minutes indicates an interval of  30 minutes
	Interval30Minutes Interval = iota

	// Interval1Hour indicates an interval of 1 hour
	Interval1Hour Interval = iota

	// Interval6Hours indicates an interval of 6 hours
	Interval6Hours Interval = iota

	// Interval1Day indicates an interval of 1 day
	Interval1Day Interval = iota
)

var (
	intervalNamesMap = map[Interval]string{
		Interval1Minute:   "1m",
		Interval5Minutes:  "5m",
		Interval15Minutes: "15m",
		Interval30Minutes: "30m",
		Interval1Hour:     "1h",
		Interval6Hours:    "6h",
		Interval1Day:      "1d",
	}
)

func (val Interval) String() string {
	return intervalNamesMap[val]
}
