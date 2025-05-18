
function addUserBlock() {
    navigate('/user');
    main = document.getElementById('container')
    main.innerHTML = '';

    d = document.createElement('div');
    d.setAttribute("class", "user-item");
    a = document.createElement('p');
    name = localStorage.getItem('user');
    a.innerText = "Добрый день, " + name;
    d.appendChild(a);

    a = document.createElement('div');
    a.innerText = 'Посмотреть билеты';
    a.setAttribute('onclick', 'GetMyTickets("b1b0b38c43f24cce8ccff4b1e343903c")')
    a.setAttribute("class", "signin-button ");
    d.appendChild(a);

    a = document.createElement('div');
    a.innerText = 'Посмотреть бонусы';
    a.setAttribute('onclick', 'GetMyBonus()')
    a.setAttribute("class", "signin-button ");
    d.appendChild(a);

    main.appendChild(d);
}


function addLoginBlock() {

    navigate('/login');
    main = document.getElementById('container')
    main.innerHTML = '';
    d = document.createElement('div');
    d.setAttribute("class", "platform");
    main.appendChild(d);
    main = d;

    d = document.createElement('div');
    d.setAttribute("class", "username");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "username");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "username");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('div');
    d.setAttribute("class", "password");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "password");
    inputfield.setAttribute("id", "password");
    inputfield.setAttribute("class", "pass-input");
    inputfield.setAttribute("placeholder", "password");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('button');
    d.setAttribute("class", "signin-button");
    d.setAttribute('onclick', 'login()');
    d.innerText = 'Войти';
    main.appendChild(d);
}


