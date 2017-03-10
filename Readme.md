	
## Example

	.....

	import (
	"fmt"
	"github.com/liontail/weightrandom"
	)

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

## Example 2

.....

	import (
	"fmt"
	"github.com/liontail/weightrandom"
	)

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
  
	result := weightrandom.DoRandom(allItems, 1)

	fmt.Println("Result: ", *result)

	.....
