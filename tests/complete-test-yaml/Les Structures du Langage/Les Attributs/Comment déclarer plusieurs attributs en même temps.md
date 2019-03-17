---
display-name: Comment déclarer plusieurs attributs en même temps ?
end-links: []
options:
since: C++11
---
Il est possible de spécifier une liste d'attributs dans une même déclaration en les séparant à l'aide d'une virgule. Les attributs de contrat (```[[assert]]```, ```[[ensures]]``` et ```[[expects]]```) font exception et doivent être déclarés seuls :

```cpp
[[nodiscard, deprecated, gnu::always_inline, gnu::hot]]
int foo();
```
