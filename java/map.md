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
