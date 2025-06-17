- this is useful when counting strings or characters or similar


```java
        for(char c: string.toCharArray()){
            map.put(c, map.getOrDefault(c, 0)+1);
        }
```

- Iteration over map

```java

  // using keySet() for iteration over keys
        for (String name : m.keySet()) 
            System.out.println("key: " + name);
        
        // using values() for iteration over values
        for (String url : m.values()) 
            System.out.println("value: " + url);
    }

```

- Useful map methods

1. `put(key, value)`
2. `get(key)`
3. `remove(key)`
4. `size()`
5. `isEmpty()`

6. `clear()`
7. `containsKey(key)`
8. `containsValue(value)`

9. `getOrDefault(key, setDefaultValue)`
10. `putIfAbsent(key, value)`
