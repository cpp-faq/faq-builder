# Les Attributs

### Sommaire

 - [Qu'est-ce qu'un attribut ?](#Qu%27est-ce%20qu%27un%20attribut%20?).
 - [Quels sont les attributs standards en C++ ?](#Quels%20sont%20les%20attributs%20standards%20en%20C++%20?).

## Qu'est-ce qu'un attribut ?

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

#### Liens et compléments
 - **[EN]** [cppreference.com – C++ attributes](https://en.cppreference.com/w/cpp/language/attributes).
 - **[EN]** [open-std.org | n2761 "General Attributes for C++](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2761.pdf).

## Quels sont les attributs standards en C++ ?

Les attributs standards (à l'heure de **C++20**) sont au nombre de treize :

 - *(C++11)* ```[[noreturn]]```.
 - *(C++11)* ```[[carries_dependency]]```.
 - *(C++14)* ```[[deprecated]]```.
 - *(C++17)* ```[[fallthrough]]```.
 - *(C++17)* ```[[nodiscard]]```.
 - *(C++17)* ```[[maybe_unused]]```.
 - *(C++20)* ```[[likely]]``` et ```[[unlikely]]```.
 - *(C++20)* ```[[no_unique_address]]```.
 - *(C++20)* ```[[expects]]```.
 - *(C++20)* ```[[assert]]```.
 - *(C++20)* ```[[ensures]]```.
 - *(C++20)* ```[[expects]]```.

#### Liens et compléments
 - **[EN]** [cppreference.com – C++ attributes](https://en.cppreference.com/w/cpp/language/attributes).
 - **[EN]** [open-std.org | n2761 "General Attributes for C++](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2761.pdf).
