---
display-name: A quoi correspond l'attribut [[fallthrough]] ?
end-links:
- descr: "cppreference.com – C++ attribute: fallthrough"
  url: https://en.cppreference.com/w/cpp/language/attributes/fallthrough
  options:
    prefix: EN
- descr: open-std.org | p0188r1 'Wording for [[fallthrough]] attribute.'
  url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2016/p0188r1.pdf
  options:
    prefix: EN
---
L’attribut ```[[fallthrough]]``` (depuis **C++17**) est destiné aux instructions ```switch```. L’objectif est de préciser au compilateur qu’une absence de saut du flot de contrôle est volontaire (que ce soit avec ```break``` ou ```return```).

Les compilateurs signalent souvent les *fallthroughs*, c’est à dire lorsque le programme passe d’une ```case``` à l’autre sans saut. Si il s’agit en effet d’une erreur de programmation assez classique, il est également possible que celle-ci soit intentionnelle.

```[[fallthrough]]``` indique au compilateur que le *fallthrough* est intentionnel et ne doit pas être considéré comme une erreur. Il permet ainsi d’éviter un *warning* de la part du compilateur. Celui-ci doit précéder un  case :

```cpp
action handle_event(event my_event) {
    switch(my_event) {
      case event::mouse_click: [[fallthrough]]; // fallthrough explicite, pas de warning.
      case event::mouse_wheel:          
          auto mouse_button = get_mouse_button();
          return action(action::event::mouse, mouse_button);
      case event::button_click:
          auto& b = get_clicked_button();
          // fallthrough : warning possible du compilateur.
      default:
         return action(action::event::unknown);
    }
}
```
