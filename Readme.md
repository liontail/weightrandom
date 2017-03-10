	
## Example

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
