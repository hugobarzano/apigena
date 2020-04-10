## Python Use Case

This use case aims to show api-rest generation capabilities of this project. 

To easy work with and successfully performs this use case
This use case provides 3 scenarios to cover the 3 technologies supported by.  

### Scenario 1: Python

To the first scenario we are going to use this data model:

```
{
   "name":"appel",
   "price":1.5,
   "inStock":true
}
```

```
./generator.linux -input resources/input/fruits.json \
                  -output=resources/output/ \
                  -name fruits \
                  -port 3001 \
                  -tech python
```
##### Scenario 1: spec
##### Scenario 1: source code
##### Scenario 1: Api Interface


### Scenario 2: Javascript

To the first scenario we are going to use this data model:

```
{
   "name":"appel",
   "price":1.5,
   "inStock":true
}
```

##### Scenario 2: spec
##### Scenario 2: source code
##### Scenario 2: Api Interface



### Scenario 3: Golang

To the first scenario we are going to use this data model:

```
{
   "name":"appel",
   "price":1.5,
   "inStock":true
}
```

##### Scenario 3: spec
##### Scenario 3: source code
##### Scenario 3: Api Interface


