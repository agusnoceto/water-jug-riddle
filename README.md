# Water Jug Riddle
Build an application that solves the Water Jug Riddle for dynamic inputs (X, Y, Z).

# Overview
Build an application that solves the Water Jug Riddle for dynamic inputs (X, Y, Z). The simulation
should have a UI to display state changes for each state for each jug (Empty, Full or Partially Full).
You have an X-gallon and a Y-gallon jug that you can fill from a lake. (You should assume the lake
has unlimited amounts of water.) By using only an X-gallon and Y-gallon jug (no third jug),
measure Z gallons of water

# Goals
1. Measure Z gallons of water in the most efficient way.
2. Build a UI where a user can enter any input for X, Y, Z and see the solution.
3. If there is no solution, display “No Solution”.

# Limitations
- No partial measurement. Each jug can be empty or full.
- Actions allowed: Fill, Empty, Transfer.
- Use the following programming language: Go

# Deliverables

The application source code should be on Github and a link should be provided. If this is not an
option, a public link to the application source code or a zip archive is also acceptable. Any
executable must be provided, no assumptions must be made about the environment or so this
will be tested on. Ideally executed inside a docker image.

# Evaluation Criterias
- Functionality
- Efficiency (Time, Space)
- Code Quality / Design / Patterns
- Testability
- UI/UX design

# Instructions
Download the repo:
```
git clone https://github.com/agusnoceto/water-jug-riddle
```

And run the following script:
```
$> ./run-riddle.sh
```
Or alternatively, build & run docker image:

```
docker build --file Dockerfile --tag water-jug-riddle .
docker run -it water-jug-riddle
```

Note: you will need [docker](https://www.docker.com/) installed in order to run the application. 