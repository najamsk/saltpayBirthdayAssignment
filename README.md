Given a JSON file with a list of people and their dates of birth, write a program to print out the people whose birthday is today.

If a person was born on Feb 29th, then their birthday during a non-leap year should be considered to be Feb 28th.

Input sample file:

```
[
    ["Doe", "John", "1982/10/08"],
    ["Wayne", "Bruce", "1965/01/30"],
    ["Gaga", "Lady", "1986/03/28"],
    ["Curry", "Mark", "1988/02/29"]
]
```

You can use whichever programming language you like. The assignment should take 60 to 90 min. If it’s taking longer, consider whether you’re complicating things.

If you make any assumptions, trade-offs or de-prioritise features for timeliness, please document these decisions.

Your submission must have:

* Instructions to run the code

* Tests to check if your code is correct, robust and complete

* Edge cases handled and tested

* Iterative development, backed by commit history

* Modular, cohesive code with sensible separation of concerns

Bonus points for following Test-Driven Development.

Please do not overcomplicate the code. You don’t need a web framework, database or message queues for this submission. Keep it simple!


## How to run Binary
* make sure you have data.json on same level as main.go
* go run main.go
* you can supply a custom data json file using go run main.go -file=data.json if file is not availble program will exit
* reading today date and returing all the records with birthday matching todays date. We can later take date from user in case they want to get a list against any date. Core service layer does support this and you can see in tests I am matching past and future dates.
