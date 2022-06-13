# study-cli

## Question Set Schema
```go
type Single struct {
    Name string `json:"name"`
    Correct int `json:"correct"`
    Answers []string `json:"answers"`
    Question string `json:"question"`
}
```

```json
[
    {
        "name": "question one",
        "correct": 0,
        "answers": [ "A. yes", "B. no" ],
        "question": "Which of the two options is a three letter world?",
    }
]
```

## TODO
- Add delay or user interaction for learn mode
- Add set filename for dataset
- Add dataset manager
- Add random quiz mode (allow user to run the command a lot, and only quiz a fraction of the time, e.g. run at bash start)
