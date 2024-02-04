# Boot Dev Go

[Boot.dev Go](boot.dev) is a platform with a go course that aims to learn it by applying language features into little assignments to build Textio - Twilio's little brother. 

- [x] Ch 1: Introduction
- [x] Ch 2: Variables
- [x] Ch 3: Functions
- [ ] Ch 4 Structs
- [ ] Ch 5: Interfaces
- [ ] Ch 6: Errors
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

The birthdayMessage and sendingReport structs have already implemented the getMessage methods. The getMessage method simply returns a string, and any type that implements the method can be considered a message.

First, add the getMessage() method as a requirement on the message interface.

Second, complete the sendMessage function. It should print a message's message, which it obtains through the interface method. Notice that your code doesn't need to worry at all about whether a specific message is a birthdayMessage or a sendingReport!