# Assumptions 

1. City name does not have any spaces in them. 
2. In the city file, No city is repeated. 
3. If more than one alien are found at the city. The City will be destroyed and all the alien names will be printed. 
4. Only 4 directions are valid, east west and north south. 
5. Only 1 Simulation can be run at a single time. 
6. All cities even with no link will appear in the map. For ex- Input like this will be invalid 
```
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee north=Lee
```
The right input will be 
```
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee north=Lee
Baz
Qu-ux
Bee
Lee
```

