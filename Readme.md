Weight Random
======

Item's Data
-----------
date is interface so it can assign any value to it

## Example

```go
import (
	"fmt"
	"github.com/liontail/weightrandom"
)

.....

allItems := []weightrandom.Items{
  weightrandom.Items{
    Weight: 10,
    Data:   uint(1),
  },
  weightrandom.Items{
    Weight: 10,
    Data:   uint(2),
    },
  weightrandom.Items{
    Weight: 80,
    Data:   uint(3),
    },
  }

result := weightrandom.DoRandom(allItems, 1)

fmt.Println("Result: ", *result)
.....
```
output will be like
- [{80 3}]

## Example 2

```go
import (
	"fmt"
	"github.com/liontail/weightrandom"
)

.....

allItems := []weightrandom.Items{
  weightrandom.Items{
      Weight: 60,
      Data:   "Hello",
   	},
  weightrandom.Items{
   		Weight: 70,
   		Data:   "Gooooo",
  	},
  weightrandom.Items{
   		Weight: 20,
   		Data:   "Language",
  	},
  }

result := weightrandom.DoRandom(allItems, 2)

fmt.Println("Result: ", *result)

.....
```
output will be like
- [{70 "Gooooo"} ,{60 "Hello"}]
