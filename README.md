# GOBOX my golang sandbox

![](https://github.com/haagor/gobox/blob/main/img/line.png)

## Intro - motivation

Suite à un entretien j'ai reçu un feedback. Ce feedback n'était pas négatif, il était correctif. Il m'a permis de regarder en face une réalité que je ne voulais pas voir : mes compétences techniques relatives au dev backend ne sont pas au niveau.
En effet je me suis un peu dispersé. J'adore les problématiques liées à l'humain, aux process, à l'organisation et aussi l'architecture... J'y ai investi du temps, notamment chez Leboncoin. Mais j'ai perdu le focus sur le code. Cela ne serait pas un problème si je souhaitais naviguer vers d'autre métier, mais ce n'ai pas le cas. Je veux faire du dev back.  

A la suite de cet entretien, je dois réagir. Et je veux garder une trace de cette expérience parce qu'elle me donne un exemple de la valeur du feedback. Comment suite à un échec, un feedback peut me donner de l'énergie pour corriger et m'améliorer.  

Ma volonté: je veux affiner mes connaissances techniques, les approfondir. Ce n'est pas normal de passer 2 ans à travailler en tant que dev back Go sur du micro service et des API chez Leboncoin et arriver plus d'un ans après en entretien et ne pas savoir expliquer ce qu'est un "context" dans Go.  

Mes actions: installer une habitude de lecture, d'expérimentation de code et d'outil. Faire des katas. Les ressources ne manquent pas! Aujourd'hui j'ai du temps, c'est le moment. Le bon moment aurait été avant de passer à côté d'entretiens. Un petit auto shame et au travail!  
Si j'arrive à faire régulièrement du sport, je dois réussir à investir régulièrement du temps sur ma technique, indépendamment de mes expériences professionnelles.

## Liens

- [Dave.Cheney](https://dave.cheney.net/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Katas](https://codingdojo.org/kata/)

## Katas

Je vais faire un premier Kata sans objectif particulier. J'ai seulement lu l'introduction de [Practical-Go](https://dave.cheney.net/practical-go/presentations/qcon-china.html
) et il y a beaucoup de guide line dans ce texte. Je vais essayer à chacun de mes Katas d'ajouter des guide lines. Ainsi Mon premier sera libre, puis il y aura de plus en plus de point. Je devrais pouvoir observer une evolution interressante de mon code.

### ![Bank OCR](https://github.com/haagor/gobox/tree/main/katas/bankOCR#readme)
![](https://github.com/haagor/gobox/blob/main/img/digits.png)

Pour ce premier Kata je n'ai pas grand chose à ajouter. Je pense que le point central est la manière de gérer les digits. J'ai traduit un digits en 3 lignes de chacune 3 valeurs qui se lit de haut en bas et de gauche à droite. Ainsi un 1 signifie qu'il y a un `|` si on est en colonne 1 ou 3, ou un `_` en colonne 2. Cette string est coûteuse et convertir ça en binaire puis en entier serait plus efficace.
 
Dans ce premier kata je n'ai pas pris en compte de guide line. En revanche c'est l'occasion de creuser certaines définitions et concepts que je ne maîtrise pas.
Le `checksum` est une string de chiffres et de lettres qui indique le nombre de bits qui doit être présent dans une transmission de données. Si le nombre de bits indiqué par le checksum et le nombre de bits réel sont différents, il y a un problème. Le `checksum` est utilisé avec les protocoles `TCP` et `UDP`. Le `checksum` value est assigné en exécutant une `cryptographic hash function`. C'est un algorithme qui traduit un message en bit d'une taille fixe appelé `hash-value` ou message digest. Ce mapping est déterministe et quasi impossible à inverser. Il existe plusieurs exemples de fonctions comme les `SHA` ou `DSA` ou encore `MD5`.

Un des contextes où j'ai rencontré ces algorithmes est la génération de clés ssh. J'ai voulu testé leur déterminisme en rentrant la même passphrase, sauf que je n'obtiens pas le même résultat. L'occasion de découvrir l'existence de la couche de `RNG (Random Number Generator)` dans l'algorithme de génération de clé. Plus le `RNG` est exécuté dans un environnement avec de l'`entropy` plus il est efficace. L'occasion d'essayer de comprendre le principe d'`entropy` avec une [video](https://www.youtube.com/watch?v=YM-uykVfq_E).
 
Une `hash table` ou `hash map` fonctionne sur le même principe: on map un clé en la donnant à un `hash fonction` avec un index derrière lequel se trouve la valeur que l'on souhaite trouver. Il est donc bien question d'une map mais avec une complexité de recherche et d'insertion en moyenne de Θ(1).
La 'hashmap' est utilisée quand on crée des `index` en base de données. Il y a un coût en espace mémoire de Θ(n). C'est ce point qui oblige à utiliser les `index` de manière réfléchie.

### ![Cupecake](https://github.com/haagor/gobox/tree/main/katas/cupecake#readme)
![](https://github.com/haagor/gobox/blob/main/img/cupecake.png)

Dans ce Kata j'ai voulu mettre en application des guide line lu dans [Practical-go](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_identifiers).
Ce Kata invite aussi à découvrir le `Decorator pattern`. Je ne connaissais pas. Il s'appuie sur la récursivité. Il permet d'ajouter dynamiquement des comportements à un objet. Je n'ai pas d'exemple d'utilisation dans un cas réel.

Il y a beaucoup de duplication dans mon code, j'assume que ce pattern induit cela et que le fait que je fasse ce Kata dans un seul fichier accentue ce point. L'exemple est artificiel et je vais mettre la factorisation de côté.

### [Birthday Greetings](https://github.com/haagor/gobox/tree/main/katas/birthdayGreetings#readme)
![](https://github.com/haagor/gobox/blob/main/img/hb.png)

Avec ce kata c'est l'occasion de mettre en place une base de données. J'utilise `PostreSQL` et la lib `database/sql`. Les tutorials ne manque pas à ce niveau, je suis [go-database-sql](http://go-database-sql.org/overview.html). Coté API j'utilise gin et je me suis appuyé sur [go.dev](https://go.dev/doc/tutorial/web-service-gin). Je vais séparer mon code en 2 modules : Identity et Contact. Ainsi c'est l'occasion d'implémenter des adapter pour ma base de donnée, mais aussi pour l'envoi d'email ou sms. C'est aussi un bon prétexte pour implémenter des API. Je structure cela en suivant les principes de la `Clean Architecture`, en tout cas j'essaye. Je ne pense pas que ma première implémentation soit fidèle à cette architecture mais cela va me permettre de voir clairement les points que je ne comprends pas.

Et en effet je suis loin d'une clean archi! Je n'ai pas implémenté d'interface pour isoler chaque couche. Mon objet 'friend.Friend' se retrouve un peu partout dans chaque module. Mais ce Kata m'a permis de mettre en place une API avec le lib `github.com/gin-gonic/gin"` et de mettre en place un adapter de db avec `database/sql`. J'ai aussi fait un schéma pour mieux visualiser certains points de mon implémentation.
![](https://github.com/haagor/gobox/blob/main/img/arch1.png)


`NEXT add log, read/code error handling`

### [Christmas Delivery](https://github.com/haagor/gobox/tree/main/katas/christmasdelivery#readme)
![](https://github.com/haagor/gobox/blob/main/img/christmasdelivery.png)

Dans ce Kata il est question de programmation concurrente. Je vais itérer sur les différents usecases sans essayer de les anticiper. Pour le premier usecase ma machine à joué et mon elf fonctionne comme attendu :

```
$ go run main.go 
"Gift n° 7eea5c73-146b-4703-9cc6-3593ada5db83 handle captain !"
* 1 gift created *
* 1 gift created *
"Gift n° 198c5ea2-5965-459d-a6a8-f981009d90ff handle captain !"
* 1 gift created *
* 1 gift created *
"Gift n° f454aa4f-e781-44dc-9811-da2823f37469 handle captain !"
* 1 gift created *
* 1 gift created *
* 1 gift created *
"Gift n° cd073103-772e-43fd-a667-40e6d241c4a5 handle captain !"
* 1 gift created *
* 1 gift created *
"Gift n° e42d629e-63f8-4b86-bb26-6b71f169830f handle captain !"
* 1 gift created *
* Creating gifts finish *
```

Pour le usecase n°2 je n'ai pas vraiment cette notion d'elf disponible mais plutôt de chan à 1 élément sur lequel écoute tous les elfs. Ainsi Mrs Claus les distribuent un par un et si aucun elf est disponible il va attendre.

### [Peak](https://github.com/haagor/gobox/tree/main/katas/christmasdelivery#readme)
![](https://github.com/haagor/gobox/blob/main/img/peak.png)

Dans ce kata il est question d’un simple algorithme pour trouver le nombre max d’interval qui se touchent. J’ai rencontré cette problématique pour calculer le nombre max de connexion simultanée et le code déjà écrit que j’ai lu m’a fait mal à la tête. Je me suis dit que j’allais écrire quelque chose de plus simple et efficace. J’ai donc écrit une version. Bon pas forcement plus simple mais plus efficace. Puis je suis allé chercher ce que je pouvais trouver sur internet.
L’algo que j’ai trouvé sur internet est assez intéressant. J’aurai pu y penser pour des intervalles composés de nombres, mais cela ne m’est pas venu à l’esprit pour mon cas composé de date. Il est question d’écrire une nouvelle liste de doublé; la valeur de début d'intervalle avec la lettre `x`. La valeur de fin d’un intervalle avec la lettre `y`. On tri cette liste. Puis on la parcourt, quand il y à `x` on ajoute 1 à notre résultat, quand il y à un `y` on soustrait 1 à notre résultat.
Efficace je trouve.

### [Game of life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)
![](https://github.com/haagor/gobox/blob/main/img/gol.png)

Ici les premiers points à traiter sont les structures de données pour représenter mon monde, et l'outil d'affichage 2D. Une fois ces points eclaircis je pourrais aller plus loin dans les experimentations. J'utilise la lib [pixel](https://github.com/faiface/pixel) pour avoir un quadrillage animé. C'est ma premiere utilisation de cette lib et j'ai envie d'aller beaucoup plus loin dans son utilisation pour avoir des visuels de certains algos.
Voici la premiere version de mon monde:
![](https://github.com/haagor/gobox/blob/main/img/golv1.png)

### Jungle
![](https://github.com/haagor/gobox/blob/main/img/jungle.png)

Ici il n'est pas vraiment question de katas mais d'un dossier ou j'empile des petits tests ou des choses que je veux verifier.

  - Suite à un test technique je me retrouve à devoir faire une fonction qui donne la somme des 2 plus grands entiers de ma slice. Je fais une première implémentation naïve qui trouve la solution en 1 parcours de slice. Puis une autre implémentation, j'utilise le sort natif de Go et j'addition les 2 derniers éléments. Sur cette implémentation les performances sont faibles. /!\ le sort natif utilise des interfaces pour des questions de généricité, cela à un coup et rend ce tri moins performant qu'un quicksort écrit à la main qui n'utilise pas d'interface.
  - Un point sur les microservices. Dans une vidéo d'histoire sur Napoléon je me suis retrouvé avec un parallèle assez intéressant entre l'organisation de l'armée et celui d'un logiciel. Évidemment ce genre de parallèle sont nombreux et l'informatique n'a pas inventé grand-chose selon moi. Je vais creuser ce parallèle, cela peut être un excellent outil pédagogique! Mon point de depart a été cette [video](https://www.youtube.com/watch?v=bhQe2cjr5XQ).

![](https://github.com/haagor/gobox/blob/main/img/napoleon.png)



## TODO

- array / slice
- API
  - JWT
  - Token / certificat
  - http / https
  - context
- journalisation BD


---

Simon P
