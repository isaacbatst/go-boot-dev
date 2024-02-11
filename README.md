# Boot Dev Go

[Boot.dev Go](boot.dev) is a platform with a go course that aims to learn it by applying language features into little assignments to build Textio - Twilio's little brother. 

- [x] Ch 1: Introduction
- [x] Ch 2: Variables
- [x] Ch 3: Functions
- [x] Ch 4 Structs
- [x] Ch 5: Interfaces
- [x] Ch 6: Errors
- [ ] Ch 7: Loops
- [ ] Ch 8: Slices
- [ ] Ch 9: Maps
- [ ] Ch 10: Advanced Functions
- [ ] Ch 11: Pointers
- [ ] Ch 12: Local Development
- [ ] Ch 13: Channels
- [ ] Ch 14: Mutexes
- [ ] Ch 15: Generics

## Assignment

We have an interesting new cost structure from our SMS vendor. They charge exponentially more money for each consecutive text we send! Let's write a function that calculates how many messages we can send in a given batch given a costMultiplier and a maxCostInPennies.

In a nutshell, the first message costs a penny, and each message after that costs the same as the previous message multiplied by the costMultiplier.

There is a bug in the code! Add a condition to the for loop to fix the bug. The loop should terminate should stop before actualCostInPennies is greater than or equal to the maxCostInPennies.