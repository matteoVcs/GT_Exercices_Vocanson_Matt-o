let img = document.getElementsByClassName('sliderImage');

let etape = 0;

let nb_img = img.length;

let precedent = document.querySelector('.precedent');
let suivant = document.querySelector('.suivant');

function enleverActiveImages() {
    for(let i = 0 ; i < nb_img ; i++) {
        img[i].classList.remove('active');
    }
}

suivant.addEventListener('click', function() {
    etape++;
    if(etape >= nb_img) {
        etape = 0;
    }
    enleverActiveImages();
    img[etape].classList.add('active');
})

precedent.addEventListener('click', function() {
    etape--;
    if(etape < 0) {
        etape = nb_img - 1;
    }
    enleverActiveImages();
    img[etape].classList.add('active');
})

setInterval(function() {
    etape++;
    if(etape >= nb_img) {
        etape = 0;
    }
    enleverActiveImages();
    img[etape].classList.add('active');
}, 3000)