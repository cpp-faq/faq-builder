---
display-name: A quoi correspondent les attributs [[likely]] et [[unlikely]] ?
end-links:
- descr: "cppreference.com – C++ attribute: likely, unlikely"
  url: https://en.cppreference.com/w/cpp/language/attributes/likely
  options:
    prefix: EN
- descr: open-std.org | p0479r0 'Attributes for Likely and Unlikely Branches'
  url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2016/p0479r0.html
  options:
    prefix: EN   
options:
  since: C++20 
---
Ces deux attributs ont été introduits avec **C++20** pour permettre d’indiquer au compilateur qu’une branche est plus probable que l’autre.

```[[likely]]```, respectivement ```[[unlikely]]```, peut être assigné à une étiquette ou une expression pour indiquer que le résultat sera le plus souvent  ```true```, respectivement ```false``` :

```cpp
int foo(int i) {
    switch(i) {
                    case 1: return 1;
      [[likely]]    case 2: return 2;
      [[unlikely]]  default: return 3;
    }
    return 2;
}

void bar(int i) {
    if ([[unlikely]] !foo(i)) {
      do_something();
    }
}
```

Ici, le compilateur sait que la valeur 2 est la plus probable, et que les valeurs autres que 1 et 2 sont peu probable. Il peut alors effectuer des optimisations avec ces informations supplémentaires sur le flot de contrôle.

Ces attributs sont à utiliser après une étude d'un code existant à l'aide d'un *profiler* par exemple, lorsqu'il est clair qu'une branche est plus souvent accédée que l'autre. L'usage prématuré de ces attributs peut fortement détériorer les performances
