---
display-name: A quoi correspond l'attribut [[maybe_unused]] ?
end-links:
- descr: "cppreference.com – C++ attribute: maybe_unused"
  url: https://en.cppreference.com/w/cpp/language/attributes/maybe_unused
  options:
    prefix: EN
- descr: open-std.org | p0212r1 'Wording for [[maybe_unused]] attribute.'
  url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2016/p0212r1.pdf
  options:
    prefix: EN    
---
L’attribut ```[[maybe_unused]]``` (depuis **C++17**), signale au compilateur qu’une variable peut être inutilisée et qu’il n’y a pas lieu de s’inquiéter. Le compilateur ne signalera pas d’avertissement si la variable est effectivement inutilisée.

```cpp
void f([[maybe_unused]] bool thing1,
                        [[maybe_unused]] bool thing2)
{
   [[maybe_unused]] bool b = thing1 && thing2;
   assert(b);
}
```

Dans cet exemple issu de *cppreference*, l’attribut est appliqué aux arguments de ```f``` et à la variable locale ```b```.

Ici, dans le cas d’une compilation en mode *release*, la macro ```assert``` va disparaître et par conséquent la variable ```b``` sera inutilisée. Puisque l’attribut ```[[maybe_unused]]``` est spécifié, le compilateur ne générera pas d’avertissement. Par ricochet, les arguments de ```f ```seront eux aussi inutilisés, d’où l’application de l’attribut à nouveau.

```[[maybe_unused]]``` peut aussi s'appliquer à d'autres entités, notamment les fonction et les classes :

```cpp
[[maybe_unused]] class C {}
[[maybe_unused]] void foo()
```

On peut comparer ```[[maybe_unused]]``` à l’annotation **Java** ```@SupressWarning("unused")```, qui fait sensiblement la même chose.
