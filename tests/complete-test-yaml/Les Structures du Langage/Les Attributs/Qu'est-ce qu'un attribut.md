---
display-name: Qu'est-ce qu'un attribut ?
end-links:
- descr: cppreference.com – C++ attributes
  url: https://en.cppreference.com/w/cpp/language/attributes
  options:
    prefix: EN
- descr: open-std.org | n2761 "General Attributes for C++
  url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2761.pdf
  options:
    prefix: EN
---
Les attributs ont été introduit avec **C++11**. Ils fournissent une syntaxe unifiée pour les attributs des compilateurs (```__attribute__``` sur GCC et Clang, ```__declspec()``` pour MSVC). Le standard propose également des attributs standards.

Si un attribut n'est pas reconnu par le compilateur, celui-ci est ignoré. Les attributs sont déclarés entre deux crochets ```[[attribute]]``` et peuvent être associés à la majorité des constructions du langage. Ci suit un exemple avec quelques attributs standards :

```cpp
[[noreturn, deprecated("Superseded by foo(bool)")]]
void foo();

[[nodiscard]] error_code bar([[carries_dependency]] int* x)
    [[expects: *x >= 0]]
    [[ensures ec : ec != error_code::unknown]]
{
    [[maybe_unused]] error_code ec{};
    if([[likely]] *x == 0)
        ec = do_something();
    else
        switch(*x) {
            [[unlikely]] case 1:
                x = 0;
                [[fallthrough]]                    
             default:
                x = 1;                        
        }

    [[assert: x == 0 || x == 1]]

    ec = do_something_else(x);
    return ec;
}
```
