# Jeu du pendu

Le programme est un Jeu du pendu codé en golang 

# Prérequis
En premier lieu, il faudra installer les logiciels, VScode, Git, Go
VsCode: https://code.visualstudio.com/download
Git: https://git-scm.com/install/
 Go: https://go.dev/doc/install

Il faudra ensuite cloner le projet sur votre ordinateur.
https://github.com/MatteoSlz/Jeu_du_pendu.git

# Comment marche la séléction des mots?

le programme corrige l'ensemble des mots du fichier mot.txt afin de rendre tout les mots à l'interieur utilisable dans le jeu. Il les envoit dans le fichier mot_correct.txt.
Ensuite en fonction de la longueur du mot le mot selectionné est envoyé dnas un des 3 fichier mot_facile.txt , mot_normal.txt ou mot_difficile.txt


# Lancer une game

Une fois que vous avez lancer VScode ou votre terminal de commande et ouvert Le dossier du programme "Jeu_du_pendu", lancez le programme en ecrivant "go run . jouer" ou "go run . reprendre" ou "go run . score".
Entrer votre Pseudo et choisissez un niveau de difficulté.
à partir de là un mot plus ou moins long et un nombre d'essai vous sera attribué en fonction de la difficulté.
Vous pouvez également sauvegarder votre partie à tout moment en écrivant "!save", et reprendre votre partie plus tard en écrivant "go run . reprendre".
Vous pouvez également consulter le top 5 des joueurs en tapant oui apres le "go run . score"

# Fonctionnement des parties

A chaque partie on effectue une boucle tant que le joueur n'as n'y perdu ni gagné alors on continue le jeu.
On affiche le nombre de lettre du mot sous forme de _ et on donne une lettre au hasard dans le mot dont on affichera toute ses apparitions dans le mot.
A chaque lancement le joueur peux choisir de donner une lettre ou un mot complet.
Si la lettre n'est pas dans le mot alors le joueur perd 1 essai. Si il essaye un mot complet il perd 2 essais.Une lettre déjà essayé ne peux pas être re-essayé.
Lorseque le joueur trouve toute les lettres ou trouve le bon mot avant de ne plus avoir d'essais alors il gagne. Si en revanche il tombe à 0 essai restant alors il perd.

# Résultat 

A la fin d'une partie, on stock le resultat dans un dossier score.csv sous forme :
 "pseudo,status,essai(s)_restant,difficulte,mot,lettre_propose,date"

Si le joueur décide de !save alors on stock toute les données de la partie en cours dans un dossier save.json qui peut être utilisé à tout moment pour reprendre une partie.



