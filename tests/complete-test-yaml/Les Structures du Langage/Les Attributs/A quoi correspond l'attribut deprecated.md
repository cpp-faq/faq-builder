---
display-name: A quoi correspond l'attribut [[deprecated]] ?
end-links:
- descr: "cppreference.com – C++ attribute: deprecated"
  url: https://en.cppreference.com/w/cpp/language/attributes/deprecated
  options:
    prefix: EN
- descr: open-std.org | n3760 'General Attributes for C++''
  url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2013/n3760.html
  options:
    prefix: EN
options:
  since: C++14
---
L'attribut ```[[deprecated]]```, depuis **C++14**, permet d'indiquer qu'un élément est déprécié et optionnellement d'en indiquer la raison.

Ci suit des exemples d'utilisation de cet attribut :

```cpp
namespace [[deprecated("old API")]] ns { // depuis C++17

template<typename T>
struct [[deprecated]] A
{
    using [[deprecated]] size_type = typename T::size_type;
    [[deprecated] T t;

    enum class [[deprecated]] color {
        BLUE,
        [[deprecated]] INDIGO = 2 // depuis C++17
    };
};

[[deprecated]] constexpr auto k = 22;

template<>
struct [[deprecated]] A<void> {
    [[deprecated("1234")]] void foo(int& a) {
        [[deprecated]] auto i = 1 + a;
        a = i % 22;
    }
};

}
```
