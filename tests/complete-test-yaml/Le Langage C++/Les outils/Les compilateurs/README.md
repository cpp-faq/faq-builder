# Les compilateurs

### Sommaire

**En cours d'écriture**

## Est-ce que mon compilateur supporte la dernière norme C++ ?

Tous les compilateurs n'ont pas un support complet du standard C++, cependant les compilateurs majeurs ont un très bon support des dernières normes et certains proposent des implémentations de certaines fonctionnalités à venir, proposées sous le namespace ```std::experimental```.

cppreference.com maintient [cette page](http://en.cppreference.com/w/cpp/compiler_support) qui tente de référencer le support des différentes parties du langage par les compilateurs majeurs.

Dans tous les cas, il est souhaitable de mettre à jour son compilateur quel qu'il soit à la dernière version de façon à profiter de toutes les dernières fonctionnalités et résolutions de bugs.

## Existe-t-il des compilateurs C++ en ligne ?

Il existe en effet plusieurs plateformes webs permettant de compiler du code C++ en ligne :
 - [http://coliru.stacked-crooked.com/](http://coliru.stacked-crooked.com/).
 - [http://cpp.sh/](http://cpp.sh/).
 - [https://www.onlinegdb.com/online_c++_compiler](https://www.onlinegdb.com/online_c++_compiler).
 - [https://www.jdoodle.com/online-compiler-c++](https://www.jdoodle.com/online-compiler-c++).
 - [https://www.tutorialspoint.com/compile_cpp_online.php](https://www.tutorialspoint.com/compile_cpp_online.php).
 - [https://www.onlinegdb.com/online_c++_compiler](https://www.onlinegdb.com/online_c++_compiler).
 - [http://rextester.com/l/cpp_online_compiler_gcc](http://rextester.com/l/cpp_online_compiler_gcc) et [http://rextester.com/l/cpp_online_compiler_visual](http://rextester.com/l/cpp_online_compiler_visual).

Les compilateurs en lignes constituent un bon moyen d'essayer un morceau de code ou de tester des fonctionnalité sans nécessiter la mise en place d'un environnement de programmation. Néanmoins, ils souffrent de certaines limitations et ne sont pas aussi configurable qu'un environnement personnel. Les bibliothèques tierces sont notamment difficile voir impossible à utiliser et la subdivision du programme en fichiers n'est que partiellement disponible. Évitez donc d'utiliser un compilateur en ligne pour un véritable projet.

## Quel est le meilleur compilateur C++ ?

Tout comme la question [Quel est le meilleur IDE C++ ?](https://github.com/cpp-faq/cpp-faq/tree/develop/faq/fr-FR/.faq/404.md), cette question n'a pas vraiment de sens.

Chaque compilateur a des avantages et des inconveignents qui font que le choix du compilateur dépends d'un grand nombre de facteur : architecture, plateforme, convention de l'entreprise, habitude des développeurs...

De plus, certaines qualités et défaut sont ammenés à changer très rapidement, trop rapidement pour qu'une réponse objective sur le sujet soit valable plus de quelques années (voire quelques mois). La qualité du support du standard par exemple change d'une version à l'autre.

En conclusion, il n'y a pas vraiment de réponse à cette question. Utilisez le compilateur qui est le plus adapté à vos besoins.

## Qu'est ce que g++ ?

## Qu'est-ce que gcc et GCC ?

En général, **GCC** fait référence à *GNU Compiler Collection*, la suite de compilateurs du projet *GNU (GNU's Not Unix)*, tandis que **gcc** correspond à *GNU C Complier*.

En pratique **gcc** et **g++** sont des front-end de **GCC**, respectivement pour les langages C et C++.

Ces définitions ne sont pas universelles. GCC signifiait également *GNU C Compiler* à l'origine, lorsqu'il consistait simplement en un compilateur C. GCC inclue maintenant des front-end pour les langages C, Fortran, C++, Ada et Go.

## Qu'est-ce que MSVC ?

**MSVC** (MicroSoft Visual C++) est le compilateur de Microsoft, il est fourni par défaut avec *Visual Studio*.

Tout comme *Visual Studio*, **MSVC** n'est disponible que sur Windows. Il permet en revanche de cibler d'autres architectures.

**MSVC** a eu pendant longtemps la réputation de ne pas très bien supporter le standard C++. De nos jour, la conformité au standard du compilateur de Microsoft s'est grandement améliorée.
