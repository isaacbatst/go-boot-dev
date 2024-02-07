# Boot Dev Go

[Boot.dev Go](boot.dev) is a platform with a go course that aims to learn it by applying language features into little assignments to build Textio - Twilio's little brother. 

- [x] Ch 1: Introduction
- [x] Ch 2: Variables
- [x] Ch 3: Functions
- [x] Ch 4 Structs
- [x] Ch 5: Interfaces
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

Textio's software architects may have overcomplicated the requirements from the last coding assignment... oops. All we needed was a new generic error message that returns the string no dividing by 0 when a user attempts to get us to perform the taboo.

Complete the divide function. Use the errors.New() function to return an error when y == 0 that reads "no dividing by 0".