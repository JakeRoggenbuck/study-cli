# study-cli

# Question Set Schema
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
