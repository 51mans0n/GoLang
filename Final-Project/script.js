function toggleForms() {
    var loginForm = document.getElementById('login-form');
    var registrationForm = document.getElementById('registration-form');

    if (loginForm.style.display === 'none') {
        registrationForm.style.display = 'none';
        loginForm.style.display = 'block';
    } else {
        loginForm.style.display = 'none';
        registrationForm.style.display = 'block';
    }
}