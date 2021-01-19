package format

import "time"

/*
1
00:02:17,440 --> 00:02:20,375
Senator, we're making
our final approach into Coruscant.

2
00:02:20,476 --> 00:02:22,501
Very good, Lieutenant.
*/

type ItemSRT struct {
	Index int
	Start time.Duration
	End   time.Duration
	Text  string
}
