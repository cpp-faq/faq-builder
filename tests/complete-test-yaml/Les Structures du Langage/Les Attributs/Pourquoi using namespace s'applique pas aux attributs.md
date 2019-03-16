---
display-name: Pourquoi using namespace ne s'applique pas aux attributs ?
end-links: []
---
Les espaces de noms d'attributs (*attribute-namespace*) sont distincts des espaces de nom habituels et donc la directive ```using namespace``` ne s'applique pas aux attributs.

Cependant, il est possible d'extraire tous les attributs d'un *attribute-namespace* à l'aide de la directive ```using``` dans la déclaration des attributs ```[[ using attribute-namespace : attribute-list ]]``` :

```cpp
using namespace gnu; // Ici, l'espace de nom gnu est importé.

[[unused, const]] int x; // les attributs sont non reconnus et donc ignorés.
[[gnu::unused, gnu::const]] int y; // OK
[[using gnu : unused, const]] int z; // OK
```
