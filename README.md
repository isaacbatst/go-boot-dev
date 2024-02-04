# Boot Dev Go

[Boot.dev Go](boot.dev) is a platform with a go course that aims to learn it by applying language features into little assignments to build Textio - Twilio's little brother. 

- [x] Ch 1: Introduction
- [x] Ch 2: Variables
- [ ] Ch 3: Functions
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

Let's clean up Textio's authentication logic. We store our user's authentication data inside an authenticationInfo struct. We need a method that can take that data and return a basic authorization string.

The format of the string should be:

Authorization: Basic USERNAME:PASSWORD
Copy icon
Create a method on the authenticationInfo struct called getBasicAuth that returns the formatted string.