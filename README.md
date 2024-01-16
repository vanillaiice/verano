# Verano

Package to manage activities in a project.
This project is inspired by the 'Primavera' or P6 software (Spring in Spanish), hence the name 'Verano' (Summer in Spanish).

# Motivation

- Understand how P6 works under the hood.
- Adding activities one by one in P6 is quite time consuming.
- Provide a Free and Open Source Alternative to P6.
- Enhance my Go coding skills.

# Installation

```sh
# inside a Go module, execute the following command:
$ go get github.com/vanillaiice/verano
```

# Features

- Sort activities in a project based on their relationships
(only start to finish relationships supported for now).
- Compute the start and finish times of all activities.
- Render a graph (with graphviz) image file showing the activities
and their relationships.
- Parse and process lists of activities in JSON, CSV, and XLSX formats
- Storage of the activities in a SQLite database.

> Please check the 'examples' directory in this repo to see these features in action.

# Structure of an Activity

For reference, here is the data structure of activities used by this package:

```go
type Activity struct {
	Id             int           // Unique identifier of the activity
	Description    string        // description of the activity
	Duration       time.Duration // duration of the activity
	Start          time.Time     // Start time of the activity
	Finish         time.Time     // Finish time of he activity
	PredecessorsId []int         // ID of the activities that precede
	SuccessorsId   []int         // ID of the activities that come after
	Cost           float64       // Cost of the activity
}
```

# Author

Vanillaiice

# Licence

GPLv3
