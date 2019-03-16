# Les littéraux utilisateurs

**En cours d'écriture**

#### Sommaire

**En cours d'écriture**

## Qu'est ce que les littéraux utilisateurs ?

Les littéraux utilisateur (*user-defined literals*) correspondent à des suffixes littéraux définis par l'utilisateur permettant de créer des littéraux personnalisés. Ils ont été introduits avec **C++11**.

Ci suit l'exemple d'un littéral utilisateur permettant d'obtenir un entier `short` :
```cpp
constexpr unsigned short operator "" _us(unsigned long long l)
{
    return static_cast<unsigned short>(l);
}

void foo() {
    auto s = 42_us; // s est de type unsigned short.
}
```

Les littéraux utilisateurs doivent commencer par un underscore `_` de manière à ne pas entrer en conflit avec les littéraux utilisateurs de la bibliothèque standard [Quels sont les littéraux utilisateurs de la bibliothèque standard ?](404.md).

Les littéraux utilisateur peuvent aussi être utilisés pour leurs effets de bord. Ci-dessous, le littéral `print` permet l'affichage d'un entier ou d'un flottant :

```cpp
void operator "" _print(const char* s)
{
    std::cout << s << '\n';
}

void foo() {
    1_print;
    2039.0_print;
}
```

#### Liens et compléments
 - **[EN]** [cppreference.com - user literals](https://en.cppreference.com/w/cpp/language/user_literal)
 - [Quelles sont les surcharges possibles pour un littéral utilisateur ?](404.md)

## Quelles sont les surcharges possibles pour un littéral utilisateur ?

Si le type de retour est libre, il existe 13 surcharges possibles de l'opérateur ̀`""`, permettant chacune de traiter un type différent de littéral :

- surcharge `unsigned long long`, pour les littéraux entiers. _exemple_ : `0x123456LL_foo`.
- surcharge ̀`long double`, pour les littéraux flottant. _exemple_ : `321.2e-3_foo`.
- la surcharge `const char*`, correspondant aux *raw literals* qui sont utilisé pour les entiers et les flottants lorsqu'il est nécessaire de traiter l'entrée à partir d'une chaîne de caractère.  _exemple_ : `255_print;`.
- littéraux caractères :
 - surcharge `char`.  _exemple_ : `'a'_foo`.
 - surcharge `char8_t` (**C++20**).  _exemple_ : `u8'a'_foo`.
 - surcharge `char16_t`.  _exemple_ : `u'a'_foo`.
 - surcharge `char32_t`.  _exemple_ : `U'a'_foo`.
 - surcharge `wchar_t`.  _exemple_ : `'L'a'_foo`.
- chaînes de caractères littérales, qui prennent deux paramètres `operator"" foo(const CharT*, std::size_t)`, le second correspondant à la taille de la chaîne :
 - surcharge `(const char*, size_t)`.  _exemple_ : `"Hello"_foo`.
 - surcharge `(const char8_t*, size_t)`.  _exemple_ : `u8"Hello"_foo`.
 - surcharge `(const char16_t*, size_t)`.  _exemple_ : `u"Hello"_foo`.
 - surcharge `(const char32_t*, size_t)`.  _exemple_ : `U"Hello"_foo`.
 - surcharge `(const wchar_t*, size_t)`.  _exemple_ : `L"Hello"_foo`.

## Quels sont les littéraux utilisateurs de la bibliothèque standard ?

La bibliothèque standard défini les littéraux utilisateurs suivants depuis **C++14** :
- les littéraux complexes : définis dans le header *<complex>* et au sein de l'espace de nom `std::literals::complex_literals`. Ils se construisent à partir de littéraux entiers ou flottants et correspondent a un complexe avec uniquement une partie imaginaire.
 - suffixe `i` (`std::complex<double>`) : `auto c = 3. + 2if;`.
 - suffixe `if` (`std::complex<float>`) : `auto c = 3. + 2if;`.
 - suffixe `il` (`std::complex<long double>`) : `auto c = 3. + 2il;`.
- les littéraux de durées : définis dans l'entête *<chrono>* et au sein de l'espace de nom `std::literals::chrono_literals`. Ils se construisent à partir de littéraux entiers ou flottants (à partir d'un flottant, crée la valeu `std::chrono::duration`):
 - les nanosecondes `ns` (`std::chrono::nanoseconds`): `auto ns = 21ns;`.
 - les microsecondes `us` (`std::chrono::microseconds`): `auto us = 21us;`.
 - les millisecondes `ms` (`std::chrono::milliseconds`): `auto ms = 21ms;`.
 - les secondes `s` (`std::chrono::seconds`): `auto s = 21s;`.
 - les minutes `min` (`std::chrono::seconds`): `auto min = 21s;`.
 - les heures `h` (`std::chrono::seconds`): `auto h = 21s;`.
- le littéral ̀`std::string`, définis dans l'entête *<string>* et au sein de l'espace de nom `std::literals::string_literals`, qui se construit à partir de chaînes de caractères littérales : `auto s = "Hello"s; auto us = u"Hello"s;`.
- le littéral ̀`std::string_view` (**C++17**), définis dans l'entête *<string>* et au sein de l'espace de nom `std::literals::string_literals`, qui se construit à partir de chaînes de caractères littérales : `auto s = "Hello"s; auto us = u"Hello"s;`.
- les littéraux de dates (**C++20**) : définis également dans l'entête *<chrono>* et au sein de l'espace de nom `std::literals::chrono_literals`. Ils ls se construisent à partir de littéraux entiers :
 - les mois `d` (`std::chrono::day`): `auto sixth = 6d;`.
 - les années `y` (`std::chrono::year`): `auto y = 2019y;`.

Notez que l'espace de nom `std::literals` et les namespaces internes (`string_literals`, `chrono_literals`...) sont `inline` ([Qu'est ce qu'un namespace inline ?](404.md)) et permettent de sélectionner les littéraux à importer :

```cpp
{
    using namespace std::literals::string_literals;
    auto x = "Hello"s; // x est de type std::string.
}
{
    using namespace std::literals::chrono_literals;
    auto x = "Hello"s; // x est de type std::chrono::seconds.    
}
```

#### Liens et compléments
 - **[EN]** [cppreference.com - user literals](https://en.cppreference.com/w/cpp/language/user_literal)
 - **[EN]** [cppreference.com - complex literals](https://en.cppreference.com/w/cpp/numeric/complex/operator%22%22i)
 - **[EN]** [cppreference.com - nanoseconds literals](https://en.cppreference.com/w/cpp/chrono/operator%22%22ns)
 - **[EN]** [cppreference.com - microseconds literals](https://en.cppreference.com/w/cpp/chrono/operator%22%22us)
 - **[EN]** [cppreference.com - milliseconds literals](https://en.cppreference.com/w/cpp/chrono/operator%22%22ms)
 - **[EN]** [cppreference.com - seconds literals](https://en.cppreference.com/w/cpp/chrono/operator%22%22s)
 - **[EN]** [cppreference.com - minutes literals](https://en.cppreference.com/w/cpp/chrono/operator%22%22min)
 - **[EN]** [cppreference.com - hours literals](https://en.cppreference.com/w/cpp/chrono/operator%22%22h)
 - **[EN]** [cppreference.com - string literals](https://en.cppreference.com/w/cpp/string/basic_string/operator%22%22s)
 - **[EN]** [cppreference.com - string_view literals](https://en.cppreference.com/w/cpp/string/basic_string_view/operator%22%22sv)
 - **[EN]** [cppreference.com - day literals](https://en.cppreference.com/w/cpp/chrono/operator%22%22d)
 - **[EN]** [cppreference.com - year literals](https://en.cppreference.com/w/cpp/chrono/operator%22%22y)
