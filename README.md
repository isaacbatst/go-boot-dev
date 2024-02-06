# Boot Dev Go

[Boot.dev Go](boot.dev) is a platform with a go course that aims to learn it by applying language features into little assignments to build Textio - Twilio's little brother. 

- [x] Ch 1: Introduction
- [x] Ch 2: Variables
- [x] Ch 3: Functions
- [x] Ch 4 Structs
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

After submitting our last snippet of code for review, a more experienced gopher told us to use a type switch instead of successive assertions. Let's make that improvement!

Implement the getExpenseReport function using a type switch.

If the expense is an email then it should return the email's toAddress and the cost of the email.
If the expense is an sms then it should return the sms's toPhoneNumber and its cost.
If the expense has any other underlying type, just return an empty string and 0.0 for the cost.