# Les littéraux

**En cours d'écriture**

#### Sommaire

**En cours d'écriture**

## Qu'est-ce qu'un littéral ?

Un littéral (une *constante littérale*) correspond a une donnée constante dans le code source du programme. Un littéral, comme une variable, a un type.

```
int i = 22; // la variable i est initialisée à partir de la constante littérale entière 22.
auto salut = "Salut"; // la variable salut est initialisée à partir de la chaîne de caractère littérale "Salut".
foo(true); // la constante littéral booléenne true est passée en tant qu'argument à la fonction foo.
```

Notez qu'une variable constante et un littéral sont deux choses distinctes. Une variable constante est une variable dont la valeur ne peut être modifiée, mais elle occupe tout de même de la place en mémoire. Un littéral n'est qu'une valeur fixe qui sera utilisée dans le programme.

Le langage **C++** défini des constantes littérales pour les types primitifs, par exemple les constantes entières (`88`, `92UL`, `0b0010110`...), les littéraux caractères (`'H'`, `'\n'`...), `true`, `false`, `nullptr`... **C++11** a également introduit les littéraux utilisateurs, qui permettent de définir des littéraux de types personalisés. La bibliothèque standard définis d'ailleurs des littéraux utilisateurs pour certains types standards tels que `std::string` [Quels sont les littéraux utilisateurs de la bibliothèque standard ?](404.md).

#### Liens et compléments
  - **[EN]** [wikipedia.org - Literal (Computer programming)](https://en.wikipedia.org/wiki/Literal_%28computer_programming%29)


## Quels sont les littéraux booléens en C++ ?

`true` et `false` sont les deux littéraux booléens en C++. Cependant, en raison des règles de [conversion implicites](404), des constantes entières peuvent être utilisée pour initialiser des booléens. Toute valeur différente est converti en `true` et une valeur nulle est convertie en `false` :

```cpp
bool b = 0;
bool b2 = 42;
bool b3 = -33;

std::cout << std::boolalpha << b << " " << b2 << " " << b3 << '\n';
// affiche : false true true
```

A l'inverse, `true` est converti en `1` dans un contexte entier et `false` est converti en `0`.

#### Liens et compléments
  - **[EN]** [cppreference.com - Boolean literals](https://en.cppreference.com/w/cpp/language/bool_literal)

## Quels sont les littéraux entiers autorisés en C++ ?

Les suffixes (*integer-suffix*) `u`, `l` et `ll` (et leurs équivalents en majuscule) permettent de définir le type correspondant du littéral entier. Ils peuvent être combinés :

```cpp
auto i   = 13;      // int
auto ui  = 13u;     // unsigned int
auto l   = 13l;     // long int
auto ul  = 13ul;   // unsigned long int
auto ll  = 13ll;   // long long int
auto ull = 13ull;  // unsigned long long int
```

Notez qu'il n'existe pas de littéral entier pour le type `short` en **C++** (`static_cast<short>(13)` est de type `short` en revanche).

Il est également possible d'utiliser les préfixes `0`, ̀`0x`, `0b` (**C++14**) pour modifier la base :

```cpp
auto i = 42;    // littéral en base 10 (par défaut).
auto i = 0b10;  // littéral en base 2  (binaire). depuis C++14
auto i = 07182;   // littéral en base 8  (octale).
auto i = 0xAB33;  // littéral en base 16 (hexadécimale).
```

Les chiffres ̀`a`, `b`, `c`, `d`, `e`, `f` peuvent être écrits indifféremment en minuscule ou en majuscule. En raison de la notation de la base 8, `0123` et `123` ne sont pas équivalents en **C++**.

Enfin, il est possible d'insérer des quotes `'`, non consécutives, dans une constante littérale entière depuis **C++14** pour simplifier la lisibilité des très grands nombres : `std::cout << "J'ai gagné : " << 18'207'395'723'507ll << " euros !\n";`

Notez que le **C++** ne possède pas de littéraux négatifs. `-213` est une expression comportant le littéral entier `213`, de type `int` et l'opérateur unaire `-`.

Le **C++** autorise également les [littéraux multi-caractères](404.md), qui sont de type `int`.

#### Liens et compléments
  - **[EN]** [cppreference.com - Integer literals](https://en.cppreference.com/w/cpp/language/integer_literal)

## Quels sont les littéraux flottants autorisés en C++ ?

Un littéral flottant est composé d'une partie entière (avec le signe) et d'une partie réelle séparé par un point ainsi que d'un exposant :

```
auto d1 = 0.1;
audo d2 = .5e-1;
auto d3 = -10.;
auto d4 = 123e2;
```

Seule l'une des deux parties (entière et réelle) est obligatoire. Le point n'est pas obligatoire lorsque l'exposant est spécifié. L'exposant correspond au symbole `e` ou `E` (ou `p` et `P` en notation hexadécimale depuis **C++17**) et de la valeur de l'exposant. La partie avant l'exposant est la *mantisse*. Le nombre a pour valeur `n ~= mantisse * 10 ^ exposant` (en réalité ce n'est pas un égalité stricte puisqu'il peut s'agir d'une valeur approchée, le nombre étant stocké en binaire) :

```
auto d1 = 1e3;      // d1 ~= 1   * 10^3   = 1000.
auto d2 = -0.5e-10;  // d2 ~= 0.5 * 10^-10 = -0.00000000005;
auto d3 = 6.02214086e23 // constante d'Avogadro.
```

Par défaut, un littéral flottant est de type `double`, les suffixes suffixes `f` et `l` (et leurs équivalents en majuscule) permettent d'obtenir respectivement un littéral de type `float` et `long double`.

```cpp
auto f   = 42.f;    // float
auto d   = 42;      // double
auto ld  = 42.l;    // long double
auto e   = 42f;     // erreur, litéral entier avec suffixe 'f'
auto l   = 42L;     // long
```

Tout comme pour les littéraux entiers, il est possible d'utiliser des *simple quotes* comme séparateur de chiffre, à la fois pour la mantisse et pour l'exposant :

```
auto ld = 142'849.124'774'1e1'1L;
```

Depuis **C++17**, il est également possible d'écrire les littéraux flottant en [notation hexadécimale](faq://hexa-float) avec le préfixe `0x`.

#### Liens et compléments
- **[EN]** [cppreference.com - floating point literal](https://en.cppreference.com/w/cpp/language/floating_literal)

## Comment fonctionne un littéral flottant hexadécimal ?

Le suffixe `0x` peut-être utilisé pour écrire un littéral flottant en notation hexadécimale depuis **C++17**. Cette notation a l'avantage de permettre d'éviter les conversions dues à l’imprécision des flottants.

Étant donné que `e` est un chiffre hexadécimal, `p` (et `P`) est utilisé pour indiquer l'exposant dans ce cas. En notation hexadécimale, l'exposant est obligatoire :

```cpp
auto d = 0xA0p2;
```

La **mantisse** est écrite en base hexadécimale, l'**exposant** quant à lui est en décimal. Le nombre vaut `n ~= mantisse * 2^exposant` :
 - mantisse = `0xA0` = `0b10100000` = `160`.
 - exposant = `2`.
 - d = `160 * 2^2 = 640`.

Ci-suit quelques exemples de flottant en hexadécimal :

```cpp
std::cout << std::setprecision(30); // #include <iomanip>
std::cout << 0xFFp0 << '\n';                    
std::cout << 0x0.1p0 << '\n';
std::cout << 0x100p6 << '\n';
std::cout << 0x1p-1 << '\n';
std::cout << 0xae3p2 << '\n';
std::cout << 0x7'123'456'1BC'DEF << '\n';
std::cout << 142'849e20l;
/* affiche :
255
0.0625
16384
0.5
11148
1990340829236719
14284900000000000000000000
*/
```

#### Liens et compléments
- **[EN]** [cppreference.com - floating point literal](https://en.cppreference.com/w/cpp/language/floating_literal)

## Quels sont les littéraux caractères en C++ ?

Un littéral caractère en C++ est entouré de simple quote `'` :

```cpp
auto c = 'a';
std::cout << c;
```

Par défaut, un littéral caractère est de type `char`. Cependant, des préfixes permettent de modifier le type de la constante littérale :

```cpp
auto a = u8'a'; // char8_t  : encodage UTF8
auto b = u'a';  // char16_t : encodage UCS-2
auto c = U'a';  // char32_t : encodage UCS-4
auto d = L'a';  // wchar_t
```

Les préfixes `u` et `U` ont été introduits avec **C++11**. `u8`, introduit avec **C++17**, correspondait originellement au type `char` et est de type `char8_t` depuis **C++20**.

Le **C++** autorise également les [littéraux multi-caractères](404.md).

#### Liens et compléments
- **[EN]** [cppreference.com - character literal](https://en.cppreference.com/w/cpp/language/character_literal  )

## Qu'est ce que les littéraux char multicaractères ?

Le **C++** autorise un cas assez particulier de littéraux de type `int`. Il s'agit d'une suite de caractères délimité par une simple quote `'`. Le standard indique que la valeur de l'`int` résultant est laissée au choix du compilateur (*implementation-defined*) et que cette fonctionnalité est *conditionnaly supported*, c'est-à-dire qu'un compilateur n'est pas obligé de la proposer.

```cpp
auto a = 'Oui';
auto b = 'C++';     
auto c = 'ABCDEFG';  

std::cout << a << '\n'; // GCC 8.1 : 5207401
std::cout << b << '\n'; // GCC 8.1 : 4401963
std::cout << c << '\n'; // GCC 8.1 : 1145390663
```

Notez que cette fonctionnalité déclenche souvent des avertissements sur un compilateur bien configuré, de manière à pouvoir signaler les erreurs lorsque l'utilisateur voulait en réalité déclarer une chaîne de caractères littérale.

Les *multicharacter literals* trouvent cependant leurs utilités, notamment pour fixer des valeurs entières facilement reconnaissables en binaire :

```
enum class Side { left = 'left', right = 'righ', top = ' top', bottom = ' bot' }

auto side = Side::left;
```

Dans l'exemple ci-dessus, la valeur entières de `side` correspondra à `0x6C656674`, qui correspond à `left` en ASCII et permet donc de faciliter le debugage par exemple.  

#### Liens et compléments
- **[EN]** [cppreference.com - character literal](https://en.cppreference.com/w/cpp/language/character_literal  )
- **[EN]** [wikipedia.org - ASCII](https://en.wikipedia.org/wiki/ASCII)

## Quelles sont les séquences d'échappement autorisées en C++ ?

Les séquences d'échappement (utilisables dans les littéraux caractères et chaînes de caractères) définies par le standard :

|Séquence | Caractère correspondant|
|:-------:|:----------------------:|
|`\'`| `'` (ASCII : 0x27)|
|`\"`| `"` (ASCII : 0x22)|
|`\?`| `?` (ASCII : 0x3f)|
|`\\`| `\` (ASCII : 0x5c)|
|`\a`| alerte (ASCII : 0x07)|
|`\b`| retour arrière (ASCII : 0x08)|
|`\f`| nouvelle page (ASCII : 0x0c)|
|`\n`| saut de ligne (ASCII : 0x0a)|
|`\r`| retour chariot (ASCII : 0x0d)|
|`\t`| tabulation horizontale (ASCII : 0x09)|
|`\v`| tabulation verticale (ASCII : 0x0b)|
|`\nnn`| valeur octale (byte)|
|`\Xnn`| valeur hexadécimale (byte) |
|`\unnnn`| valeur Unicode (**C++11**) |
|`\Unnnnnnnn`| valeur Unicode (**C++11**) |

Pour information, `\a` correspond au signal d'alerte (la carte mère fait un 'bip' d'alarme le plus souvent). `\?` permet d'échapper le caractère `?` lorsqu'on souhaite éviter les [trigraphes](404.md), qui ont été retirés du langage avec **C++17**.

#### Liens et compléments
- **[EN]** [cppreference.com - escape sequences](https://en.cppreference.com/w/cpp/language/escape)

## Quels sont les littéraux chaîne de caractères en C++ ?

Les chaînes de caractères littérales se différencient des caractères par l'usage de double quote `"` pour la délimitation, plutôt que des simples `'` :

```cpp
auto s = "Hello World !";
std::cout << s << '\n';
```

Les chaînes de caractères littérales sont de type `const char*`, lorsque rien n'est précisé. Les suffixes `u8`, `u`, `U` et `L` permettent de définir des littérals basés sur des types caractères différents :

```cpp
auto a = u8"Hello world!"; // const char8_t[N]  : encodage UTF8
auto b = u"Hello world!";  // const char16_t[N] : encodage UCS-2
auto c = U"Hello world!";  // const char32_t[N] : encodage UCS-4
auto d = L"Hello world!";  // const wchar_t[N]
```

Les préfixes `u8`, `u` et `U` sont disponibles depuis **C++11**. Avant **C++20**, le préfixe `u8` correspondait au type `const char*`.

Le caractère nul `\0` est automatiquement ajouté à la fin de chaque littéral. Ainsi, `"Hello world!"` est de type `const char[13]` ([Comment créer une chaîne de caractère contenant le caractère nul '\0' ?](404.md)).
Notez toutefois que les chaînes de caractères juxtaposées sont regroupées juste avant la compilation, le caractère nul n'est alors ajouté qu'à la fin de la chaîne composée.

Les [séquence d'échappement](Quelles sont les séquences d'échappement autorisées en C++ ?) peuvent être utilisées dans les chaînes de caractères littérales.

#### Liens et compléments
- **[EN]** [cppreference.com - string literal](https://en.cppreference.com/w/cpp/language/string_literal)
- [Comment créer une string contenant un '\0' ?](404.md)

## Qu'est-ce que les raw string literals ?

**C++11** a introduit un nouveau type de littéraux chaînes de caractères permettant de simplifier l'écriture de chaînes contenant des caractères à échapper : les *raw string literals*.

La syntaxe est de la forme suivante : `Rdelimiteur(séquence de caractères)delimiteur`. Le délimiteur est généralement `"`, mais n'importe quelle séquence de caractère (\*) peut être utilisée comme délimiteur en fonction du contenu :

```cpp
auto email_regex = u8R"((?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\]))"; // according to RFC 5322.

auto str = Rdel(("()"))del;
```

Contrairement aux chaînes de caractères littérales classiques, il n'est pas nécessaire d'échapper les caractères tels que `"` ou `\`. Seul la chaîne constituant le délimiteur (respectivement `"` et `del` dans les exemples ci-dessus) ne doivent pas être contenus dans la chaîne. De plus, le caractère de saut de ligne dans une chaîne brute est simplement inclus dans la chaîne.

(\*) la norme indique que les caractères suivant ne sont pas autorisés dans le délimiteur : `(`, `)`, `\`, ainsi que les caractères de contrôle `\n`, `\t`, `\v` et `\f`. Le délimiteur ne peut excéder 16 caractères.

#### Liens et compléments
- **[EN]** [cppreference.com - string literal](https://en.cppreference.com/w/cpp/language/string_literal)

## Peut-on modifier une chaîne de caractère littérale en C++ ?

La modification d'un littéral chaîne de caractère est un **comportement indeterminé** en C++ :

```cpp
auto str = const_cast<char*>("Hello");
str[0] = 'U'; // Undefined behavior.
```

L'exemple précédent provoque une *erreur de segmentation* avec GCC 8.1 et clang 5.0 sur Ubuntu 18.1. Il s'agit du comportement le plus commun, même si certaines architectures autorisent cette modification.

## Où sont stockée les chaînes de caractères littérales en C++ ?

Le choix est laissé au compilateur. Très souvent, les chaînes littérales sont stockées dans une section de la mémoire en lecture seule (le segment *rodata*).

Le standard impose que ces littéraux aient une durée de stockage statique (*static storage duration*), c'est-à-dire qu'ils existent pendant toute la durée du programme.

## Le compilateur peut-il regrouper les chaînes de caractères littérales identiques ?

Le compilateur est effectivement capable d'utiliser la même adresse pour des chaînes identiques. Si ces chaînes sont dans des unités de traduction différentes, l'éditeur de liens (*linker*) peut lui aussi choisir d'effectuer le regroupement :

```cpp
auto s = "Hello";
auto s2 = "Hello";
auto s3 = "World";

std::cout << (void*)s << " " << (void*)s2 << " " << (void*)s3 << '\n';
// GCC 8.1 :   0x400944 0x400944 0x40094c
// Clang 5.0 : 0x400924 0x400924 0x40092a
```

#### Liens et compléments
- **[EN]** [stackoverflow.com - How Do C++ Compilers Merge Identical String Literals](https://stackoverflow.com/questions/6281835/how-do-c-compilers-merge-identical-string-literals)

## Le compilateur peut-il regrouper les chaînes de caractères qui se superposent ?

Le compilateur est effectivement autorisé (sans obligation) à le faire. Dans l'exemple suivant, on observe que la chaîne `"valid"` est en fait une sous-partie de la chaîne `"invalid"` car le compilateur a effectivement fait ce choix :

```cpp
auto s = "invalid";
auto s2 = "valid";

std::cout << (void*)s << " " << (void*)s2 << '\n';
std::cout << std::boolalpha << (s + 2 == s2) << '\n';
// Clang 5.0 :
// 0x400994 0x400996
// true
```

#### Liens et compléments
- **[EN]** [cppreference.com - string literal](https://en.cppreference.com/w/cpp/language/string_literal)

## Dois-je préférer NULL ou nullptr ?

`nullptr` a été introduit avec **C++11** et présente des avantages par rapport à la macro `NULL`. Il s'agit du littéral représentant le pointeur nul et est de type `std::nullptr_t`.

`nullptr` n'est pas implicitement convertible en entier, contrairement à `NULL` qui est le plus souvent défini comme `(void*)0` (certains compilateurs récents utilisent aussi `#define NULL nullptr`). Cela permet d'éviter des problèmes de surcharge :

```cpp
foo(int*);
foo(int);

foo(NULL); // Appelle foo(int).
foo(nullptr); // Appelle foo(int*).
```

Lorsqu'on travaille avec des templates, la constante `NULL` peut également provoquer des comportements inattendus, comme le montre cet exemple issu de [cppreference](https://en.cppreference.com/w/cpp/language/nullptr) :

```cpp
#include <iostream>

template<class F, class A>
void fwd(F f, A a)
{
    f(a);
}

void g(int* i)
{
    std::cout << "Function g called\n";
}

int main()
{
    g(NULL);           // Ok
    g(0);              // Ok

    fwd(g, nullptr);   // Ok
//  fwd(g, NULL);  // ERROR: No function g(int)
}
```

#### Liens et compléments
- **[EN]** [cppreference.com - nullptr, the pointer literal](https://en.cppreference.com/w/cpp/language/nullptr)
