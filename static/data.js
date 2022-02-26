
/* -----   récupération des données   ----- */
let total = []

setTimeout(function() {
console.log(total.length)
let compteur;
let bandname = []
let members = []
let creationDate = []
let firstAlbum = []
for (let i = 0; i < total.length-1; i++) {
    if (i == 1) {
        bandname += total[i]
        bandname += " "
        compteur = 0;
    }
    if (compteur == 5) {
        bandname += total[i]
        bandname += " "
        compteur = 0
    }
    compteur++
}
compteur = 0;
for (let i = 0; i < total.length-1; i++) {
    if (i == 2) {
        members += total[i]
        members += " "
        compteur = 0;
    }
    if (compteur == 5) {
        members += total[i]
        members += " "
        compteur = 0
    }
    compteur++
}
compteur = 0;
for (let i = 0; i < total.length-1; i++) {
    if (i == 3) {
        creationDate += total[i]
        creationDate += " "
        compteur = 0;
    }
    if (compteur == 5) {
        creationDate += total[i]
        creationDate += " "
        compteur = 0
    }
    compteur++
}
compteur = 0;
for (let i = 0; i < total.length-1; i++) {
    if (i == 4) {
        firstAlbum += total[i]
        firstAlbum += " "
        compteur = 0;
    }
    if (compteur == 5) {
        firstAlbum += total[i]
        firstAlbum += " "
        compteur = 0
    }
    compteur++
}

total = []
total += bandname + members + creationDate + firstAlbum

}, 200);

