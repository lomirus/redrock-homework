const username = document.querySelector('#username');
const password = document.querySelector('#password');
const rememberMe = document.querySelector('#rememberMe');
const login = document.querySelector('#login');

if (localStorage.getItem('token') || sessionStorage.getItem('token')) {
    window.location.href = "/app"
}


login.addEventListener('click', () => {
    fetch('/api/login', {
        method: 'POST',
        body: JSON.stringify({
            username: username.value,
            password: password.value,
        })
    })
        .then(data => data.json())
        .then(json => {
            alert(json.message)
            if (json.status) {
                if (rememberMe.checked)
                    localStorage.setItem('token', json.token);
                else
                    sessionStorage.setItem('token', json.token);

                window.location.href = "/app";
            }
        })
})