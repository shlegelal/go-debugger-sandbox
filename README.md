# _Go_ debugger sandbox

A repo with examples of parallel programming problems debugging in _**Go**_.

## Contents

Descriptions of code files containing the examples.

### Deadlocks

Deadlock detection.
_**Go**_ is able to detect deadlocks at runtime if one of the locked goroutines is
main.

- **channelDeadlock.go**

  Main goroutine self-deadlock with channel waiting.

  _**Go**_ detects this at runtime.

- **crossChannelDeadlock.go**

  A classic deadlock example: two goroutines (one of which is main) lock each other waiting on two channels.

  _**Go**_ detects this at runtime.

- **crossMutexDeadlock.go**

  Another classic deadlock example: two goroutines (one of which is main) lock each other by locking two mutexes in
  different order.

  _**Go**_ detects this at runtime.

- **mutexDeadlock.go**

  One goroutine forgets to unlock a mutex making the other one wait on it forever and never complete.

  _**Go**_ cannot detect this _(but it can if the waiting goroutine is main)_.

### Data races

Data race detection.
_**Go**_ is able to detect data races at runtime when built with a flag `-race`.

- **dataRace.go**

  Two goroutines write to the same `map` causing a data race.

  _**Go**_ detects this at runtime.
