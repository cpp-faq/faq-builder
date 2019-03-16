# Les énumérations

### Sommaire

- [Qu'est-ce qu'une énumération ?](https://github.com/cpp-faq/cpp-faq/tree/master/faq/fr-FR/02%20-%20Typage%20et%20variables/Les%20%C3%A9num%C3%A9rations#quest-ce-quune-%C3%A9num%C3%A9ration-)
- [Comment déclarer une enum en C++ ?](https://github.com/cpp-faq/cpp-faq/tree/master/faq/fr-FR/02%20-%20Typage%20et%20variables/Les%20%C3%A9num%C3%A9rations#comment-d%C3%A9clarer-une-enum-en-c-)
- [Qu’est-ce qu’une enum class (énumération fortement typée) ?](https://github.com/cpp-faq/cpp-faq/tree/master/faq/fr-FR/02%20-%20Typage%20et%20variables/Les%20%C3%A9num%C3%A9rations#quest-ce-quune-enum-class%C3%A9num%C3%A9ration-fortement-typ%C3%A9e-)
- [Comment choisir entre enum et enum class ?](https://github.com/cpp-faq/cpp-faq/tree/master/faq/fr-FR/02%20-%20Typage%20et%20variables/Les%20%C3%A9num%C3%A9rations#comment-choisir-entre-enum-et-enum-class)
- [Quel est l’avantage des énumérations par rapport aux constantes entières et aux constantes de préprocesseur ?](https://github.com/cpp-faq/cpp-faq/tree/master/faq/fr-FR/02%20-%20Typage%20et%20variables/Les%20%C3%A9num%C3%A9rations#quel-est-lavantage-des-%C3%A9num%C3%A9rations-par-rapport-aux-constantes-enti%C3%A8reset-aux-constantes-de-pr%C3%A9processeur-)
- [Quelle est le type et la taille d'une enum ?](https://github.com/cpp-faq/cpp-faq/tree/master/faq/fr-FR/02%20-%20Typage%20et%20variables/Les%20%C3%A9num%C3%A9rations#quelle-est-le-type-et-la-taille-dune-enum-)
- [Comment définir le type d'une énumération ?](https://github.com/cpp-faq/cpp-faq/tree/master/faq/fr-FR/02%20-%20Typage%20et%20variables/Les%20%C3%A9num%C3%A9rations#comment-d%C3%A9finir-le-type-dune-%C3%A9num%C3%A9ration-)
- [Est-il possible de créer des énumérations de types réels ou même de classes ?](https://github.com/cpp-faq/cpp-faq/tree/master/faq/fr-FR/02%20-%20Typage%20et%20variables/Les%20%C3%A9num%C3%A9rations#est-il-possible-de-cr%C3%A9er-des-%C3%A9num%C3%A9rations-de-types-r%C3%A9els-ou-m%C3%AAme-de-classes)
- [Est-ce que les énumérations peuvent définir des fonctions et des variables membres ?](https://github.com/cpp-faq/cpp-faq/tree/master/faq/fr-FR/02%20-%20Typage%20et%20variables/Les%20%C3%A9num%C3%A9rations#est-ce-que-les-%C3%A9num%C3%A9rations-peuvent-d%C3%A9finir-des-fonctions-et-des-variables-membres)

## Qu'est-ce qu'une énumération ?

Une énumération est un type entier qui ne contient qu'un nombre spécifié d'énumérateur, c'est à dire de valeurs (nommées) possibles :

```cpp
enum Day {Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday};

int main()
{
    std::cout << Sunday;
}
```

Ces valeurs constantes sont très utiles pour décrire un nombre limité d'état possible.

## Comment déclarer une enum en C++ ?

La déclaration d’une énumération se fait à l’aide du mot clef enum :

```cpp
enum MonEnum {ELEM_1, ELEM_2, ELEM_3};
```

Il est possible de préciser une valeur entière à un élément de l’énumération (appelé « énumérateur ») :
```cpp
enum Flags { FULLSCREEN = 0x1, WINDOWED = 0x2, WINDOWED_FULLESCREEN = 0x4 };
```
Si aucune valeur n’est précisée, le compilateur assigne aux énumérateurs des valeurs croissantes à partir de zéro.

Si seulement certaines valeurs sont précisées, le compilateur calculera les énumérateurs non spécifiés en incrémentant depuis ceux qui sont explicitement définies :

```cpp
enum MonEnum {I1 = -1, I2, I3, I4 = 12, I5, I6 = I4 + I1};

int main()
{
    std::cout << I1 << " " << I2 << " " << I3 << " " << I4 << " " << I5 << " " << I6 << '\n';    
}
```
Ce code affichera: ```-1 0 1 12 13 11```

## Qu’est-ce qu’une enum class (énumération fortement typée) ?

Les énumération fortement typées (*strong-typed enums*) ont été introduite avec le standard **C++11**.

L’objectif est de proposer un typage fort pour les énumérations de manière à les englober dans un espace de nom et à éviter les conversions automatiques non voulues.

Pour déclarer une énumération fortement typée, il faut utiliser le mot clef ```class``` (ou ```struct```, indifféremment) :
```cpp
enum Color { red, blue, green }; // déclaration d’une énumération.
enum class SColor { red, blue, green };  // déclaration d’une énumération fortement typée.
Une enum class n’est pas directement convertible en entier :
Color c;
int i = c; // OK.
bool b{red && blue ^ green}; // OK, même si ça n’a pas de sens.

SColor sc;
i = sc; // Erreur.
i = static_cast<int>(sc); // OK.
if(sc == 2) // Erreur.
    /*do something*/;
Bool b{red && blue ^ green}; // Erreur.
```

Cette sécurité du type à plusieurs avantage, notamment pour éviter d’avoir des valeurs ne correspondant à aucun énumérateur ou pour permettre des surcharges de fonction qui ne seraient pas toujours possible avec des énumérations du **C++03**.

Autre avantage des ```enum class``` est le fait que les énumérateurs appartiennent à un espace de nom (*namespace*) :
```cpp
enum class SColor { red, blue, green };

int main()
{   
    auto red{ SColor::red };
}
```

Il est nécessaire d’accéder aux énumérateurs à l’aide de l’opérateur de résolution de portée. Les identificateurs de énumérateurs sont donc libres pour le namespace global.

Avant **C++11**, il était commun d’envelopper des enum à l’intérieur de namespace ou de classes de manière à proposer cette syntaxe et éviter de polluer le namespace global.

## Comment choisir entre enum et enum class ?

Les ```enum class``` devraient être privilégiées autant que possible par rapport aux enum classiques parce qu’elles sont plus pratiques à utiliser (pas de risque de collision de nom, pas de comportement inattendu) et plus sécurisées (absence de conversion implicite notamment).

Il arrive que les énumérations simples soient préférées aux ```enum class``` lorsque l’arithmétique entier est nécessaire ou pour des raisons de compatibilité.

## Quel est l’avantage des énumérations par rapport aux constantes entières et aux constantes de préprocesseur ?

Comparons les deux approches :
```cpp
enum class Days1 : unsigned { Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday };

namespace Days2
{
    constexpr const unsigned Monday     { 0 };
    constexpr const unsigned Tuesday    { 1 };
    constexpr const unsigned Wednesday  { 2 };
    constexpr const unsigned Thursday   { 3 };
    constexpr const unsigned Friday     { 4 };
    constexpr const unsigned Saturday   { 5 };
    constexpr const unsigned Sunday     { 6 };
}

int main()
{
    std::cout << (unsigned)Days1::Monday << " " << Days2::Friday << std::endl;
}
```

Affiche : ```0 4```

Les deux constructions précédentes ```Days1``` (énumération) et ```Days2``` (constantes de compilation) sont équivalente. Les énumérateurs sont des constantes de compilation, sont déclarée ```const``` et sont implémentés en terme d’entiers non signé.

Cependant, on n’échappe pas aux conversions implicites avec ```Days2```. De plus, l'```enum``` nous garantit dans cet exemple que tous les énumérateurs ont une valeur différentes, mais si ```Thursday``` valait 1 par exemple, le compilateur n’aurait pas pu nous avertir de cette erreur.

On se rend compte que les énumérations sont plus pratiques à utiliser dans ce genre de cas. La syntaxe est plus courte et le compilateur nous propose une sécurité accrue qu’avec des constantes dont la valeur est déclarée à la main.

#### Liens et compléments
 - [Comment choisir entre const et #define ?](https://github.com/cpp-faq/cpp-faq/tree/develop/faq/fr-FR/.faq/404.md)

## Quelle est le type et la taille d'une enum ?

En **C**, le type d’une énumération était toujours ```int``` (la taille était donc souvent de 4 à 8 bytes).

En **C++**, il est possible de définir [le type d'une énumération](https://github.com/cpp-faq/cpp-faq/tree/develop/faq/fr-FR/.faq/404.md) et dans ce cas le type est celui précisé.

Si aucun type n’est précisé, la norme indique ceci :

> From Standard C++ 7.2/5:
The underlying type of an enumeration is an integral type that can represent all the enumerator values defined in the enumeration. It is implementation-defined which integral type is used as the underlying type for an enumeration except that the underlying type shall not be larger than int unless the value of an enu- merator cannot fit in an int or unsigned int. If the enumerator-list is empty, the underlying type is as if the enumeration had a single enumerator with value 0. The value of sizeof() applied to an enu- meration type, an object of enumeration type, or an enumerator, is the value of sizeof() applied to the underlying type.

En simplifiant, la taille d’une énumération est **implementation-defined**. La norme précise simplement que le type utilisé doit être un type entier et ne pas dépasser la taille d’un ```int``` sauf s’il n’est pas possible de stocker toutes les énumérateurs de l’énumération dans un ```int```.

## Comment définir le type d'une énumération ?

Depuis **C++11**, il est possible de préciser le type sous-jacent d'une énumération avec cette syntaxe :

```cpp
enum E1 : char {} ; // E1 est une énumération de type char.
enum class E2 : long {} ; E2 est une énumération fortement typée (enum class) de type long.
```

Cette syntaxe ressemble à de l'héritage, mais il **n'y a pas** d'héritage dans ce cas. Si la valeur d'un énumérateur ne peut tenir dans le type spécifié, le compilateur affichera une erreur. Dans l'exemple, 257 ne peut pas être encodé avec seulement 8 bits :

```cpp
#include <cstdint>

enum Test : std::int8_t {ERROR = 257}; // GCC (6.3.0) : error: enumerator value 257 is outside the range of underlying type 'int8_t {aka signed char}'
```

## Est-il possible de créer des énumérations de types réels ou même de classes ?

En C++, les énumérations ne peuvent être que de types entiers et il n’est pas possible de créer des énumérations de ```double``` ou de ```string``` par exemple :

```cpp
enum : double {}; // Erreur.
enum : std::string {}; // Erreur.
enum { STR_ENUMERATOR = "ENUM"}; // Erreur
```

Certaines méthodes permettent de simuler une énumération d’un autre type. Par exemple, il est possible de mapper une ```enum``` sur un tableau :
```cpp
enum QUALITY_VAR { VERYLOW, LOW, MEDIUM, HIGH, VERYHIGH };
std::array Quality = { "Very Low", "Low", "Medium", "High", "Very High" };

int main()
{
    std::cout << Quality[VERYLOW] << std::endl;
    std::cout << Quality[HIGH] << std::endl;
}
```

Affiche :
```
Very Low
High
```
Dans cet exemple les énumérateurs servent d’index au tableau contenant les valeurs des énumérateurs de notre « énumération de string ».

Les conteneurs tels que ```std::map``` peuvent également servir à obtenir un équivalent des ```enum```. La clef serait alors l’équivalent du nom d’un énumérateur qui aurait sa valeur associée. Cependant, l’utilisation de ```std::map``` serait bien plus coûteuse qu’un simple accès dans un tableau.

D’autres techniques plus complexes utilisent les macros et/ou les templates pour simuler des énumérations de types non entiers.
Une autre solution est d’utiliser une structure qui contient des constantes de compilation. Le nom de la constante correspond au nom d’un énumérateur :
```cpp
#include <iostream>
#include <string>

struct StrEnum
{
   static constexpr const char* HELLO{"Hello"};
   static constexpr const char* WORLD{"World"};
};

int main()
{
    std::cout << StrEnum::HELLO << " " << StrEnum::WORLD << std::endl;
}
```

Affiche : ```Hello World```.

Dans cette version simpliste, les énumérateurs sont déclaré comme variable membre statique (```constexpr``` également puisque c’est possible pour les chaînes de caractère). Cette solution perds certains des avantage de l'```enum``` ([Quel est l’avantage des énumérations par rapport aux constantes entières et aux constantes de préprocesseur ?](https://github.com/cpp-faq/cpp-faq/tree/develop/faq/fr-FR/.faq/404.md)).

## Est-ce que les énumérations peuvent définir des fonctions et des variables membres ?

Non, en *C++* les énumérations ne sont pas des classes (pas même les **enum class**) et il n’est pas possible de définir des variables/fonctions membres pour une enum comment on l'aurait fait en **Java** par exemple.

Cependant, il est tout à fait possible de fournir des fonctions libres qui agissent sur les énumérations. On peut même regrouper l’énumération et les fonctions relatives dans un espace de nom :

```cpp
namespace Direction
{
    enum Dir { Left, Right, Top, Bottom };

    Dir getOpposite(const Dir& dir);
    std::string toString(const Dir& dir);   
}

int main()
{
    auto right = Direction::getOpposite(Direction::Left);
    std::cout << Direction::toString(Direction::Top) << std::endl;
}
```

## Pourquoi les enum sont parfois utilisée à la place de constantes ?

En effet, vous verrez dans beaucoup d’ancien code C++ des codes comme ceux-ci :
```cpp
struct MyArray {
    enum : unsigned { size = 1000 };
    double tab[size];
};
```
Dans cet exemple, une énumération anonyme est utilisée pour définir la constante size, cette technique a été surnommé le ***enum hack***. Depuis **C++11**, ```constexpr``` permet souvent de remplacer une énumération :
```cpp
struct MyArray2 {
    static constexpr unsigned size{ 1000 };  
    double tab[size];
};
```

Cette seconde version conserve les propriétés qui nous intéressent : ```size``` est une constante **compile-time** et **immuable** (```constexpr```), ```size``` est une variable de classe (```static```) et ```size``` est directement convertible en un nombre entier non signé.

La raison est que dans de vielles versions du langage, les variables membres des classes ne pouvaient pas être à la fois ```static``` et ```const```. Le **enum hack** permettait donc d’avoir une variable constante (et également **compile-time** ce qui n’était pas possible autrement avant le C++11 et l’introduction du mot clef ```constexpr```).
```
