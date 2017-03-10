Example

.....

allItems := []Items{
    Items{
      Weight: 10,
      Data:   uint(1),
    },
    Items{
      Weight: 10,
      Data:   uint(2),
    },
    Items{
      Weight: 80,
      Data:   uint(3),
    },
  }
result := DoRandom(allItems, 1)

fmt.Println("Result: ", *result)

.....