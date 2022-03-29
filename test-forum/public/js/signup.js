const form = document.querySelector('.form-signin');
const login = document.querySelector('#username');
const email = document.querySelector('#email');
const pass1 = document.querySelector('#password');
const pass2 = document.querySelector('#password2');
const msg = document.querySelector('.msg');

form.addEventListener('submit', onSubmit);

function onSubmit(e) {
    e.preventDefault();
    if (login.value === '' || email.value === '' || pass1.value === '' || pass2.value === '') {
        raiseErrorMsg('Please fill in all fields');
    } else if (pass1.value != pass2.value) {
        raiseErrorMsg('Provided passwords don´t match');
    } else if (!validateEmail(email.value)) {
        raiseErrorMsg('Please check e-mail');
    } else if (!passwordStrength(pass1.value)) {
        raiseErrorMsg('Password is too weak. Needed:<br>8 symbols, lowercase letter, capital letter, number')
    } else {
        form.submit();
    }
}

function validateEmail(email) {
    const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}

function raiseErrorMsg(message) {
    msg.style.display = "block";
    msg.classList.add('error');
    msg.innerHTML = message;
    setTimeout(() => msg.style.display = "none", 3000);
}

function passwordStrength(password) {
    return password.match(/^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])([a-zA-Z0-9]{8,16})$/);
}