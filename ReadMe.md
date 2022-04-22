# Alien Invasion 

There is always a posibility of alien attack from outside world. The united nations defense wanted to build a simulator to be better prepared for the attack. Please refer the problem statement in `Challenge.md`.

The simulator is written in golang. 

## Build and Run 

To run the `alien-invasion` have a [working Golang environment](https://golang.org/doc/install) installed. If you are all set, just run the following:

```
$ go run main.go
```
This will run the simulation using all defaults and current Unix time as a random source of entropy.

To list all `cli` options ask for help:
```
~/go/❯ go run main.go -help                                                       Py base
Usage of /var/folders/83/dkktwqks635gtt_nd8m2yq900000gn/T/go-build1732504791/b001/exe/main:
  -aliens int
    	number of aliens invading (default 10)
  -iterations int
    	number of iterations (default 10000)
  -names string
    	a file used as alien names input (default "./data/alien_names.txt")
  -world string
    	a file used as world map input (default "./data/world-example-1.txt")
```

## Tests

To run the tests for `alien-invasion` run the following from the root of the repo:

```
$ go test ./... -v
```

# Assumptions 

1. City name does not have any spaces in them. 
2. In the city file, No city is repeated. 
3. If more than one alien are found at the city. The City will be destroyed and all the alien names will be printed. 
4. Only 4 directions are valid, east west and north south. 
5. The city roads are two way path. If City X is connected to City Y, this implies city Y will also be connected to City X.  
6. The code autocompletes the paths for the cities so you may see infomation which is not diretly given by user but is implied. For example, If user just gives a link between the city X and Y, Automatically the link between city Y and X will be maade. 



