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

We need better error logs for our backend developers to help them debug their code.

Complete the getSMSErrorString() function. It should return a string with this format:

```
SMS that costs $COST to be sent to 'RECIPIENT' can not be sent
```

- COST is the cost of the SMS, always showing the price formatted to 2 decimal places.
- RECIPIENT is the stringified representation of the recipient's phone number

Be sure to include the $ symbol and the single quotes