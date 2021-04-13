# poc for an async library

Just some PoC code to think about how I should design an async library that would buffer messages.

The rough idea is to use expose some very primitive API calls into a library that all send messages over a channel. The
actual sending happen in a goroutine which will buffer stuff locally.

This will serve as the basis for some mqtt stuff.