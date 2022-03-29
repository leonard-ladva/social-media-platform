const form = document.querySelector('.form-signin');
const login = document.querySelector('#username');
const pass = document.querySelector('#password');
const msg = document.querySelector('.msg');

form.addEventListener('submit', onSubmit);

function onSubmit(e) {
    e.preventDefault();
    if (login.value === '' || pass.value === '') {
        msg.style.display = "block"
        msg.classList.add('error');
        msg.innerHTML = 'Please fill in all fields';
        setTimeout(() => msg.style.display = "none", 3000);
    } else {

        form.submit();
    }
}