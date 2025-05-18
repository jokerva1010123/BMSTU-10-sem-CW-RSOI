async function seeAllUsers() {
    const url = 'http://3.67.182.34:8010/api/v1/users';

    let response = await fetch(url, {
        method: 'GET',
        credentials: 'include',
        headers: {
            'Access-Control-Allow-Origin': 'http://3.67.182.34:8887/'
        },
    });
    if (response.ok) {

        main = document.getElementById('container')
        main.innerHTML = '';

        let airports = await response.json();
        for (let i = 0; i < airports.length; i++) {
            d = document.createElement('div');
            d.setAttribute("id", airports[i].id);
            d.setAttribute("class", "user-item ");

            link = document.createElement('p');
            link.innerText = 'ID пользователя: ' + airports[i].id;
            d.appendChild(link);

            link = document.createElement('p');
            link.innerText = 'Логин пользователя: ' +  airports[i].login;
            d.appendChild(link);

            link = document.createElement('p');
            link.innerText = 'Роль пользователя: ' +  airports[i].role;
            d.appendChild(link);

            main.appendChild(d);
        }

    } else {
        addError("Возникла ошибка");
    }
}

async function createUserBlock() {
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
    d.setAttribute("class", "username");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "role");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "user role");
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
    d.setAttribute('onclick', 'signup()');
    d.innerText = 'Создать';
    main.appendChild(d);
}

async function getStatistics() {
    const url = 'http://3.67.182.34:8060/api/v1/reports/flights';

    let response = await fetch(url, {
        method: 'GET',
        credentials: 'include',
        headers: {
            'Access-Control-Allow-Origin': 'http://3.67.182.34:8887/'
        },
    });
    if (response.ok) {
        main = document.getElementById('container')
        main.innerHTML = '';
        let reports = await response.json();
        for (let i = 0; i < reports.length; i++) {
            d = document.createElement('div');
            d.setAttribute("id", reports[i].id);
            d.setAttribute("class", "flight-item");

            a = document.createElement('p');
            a.innerText = 'ID пользователя: ' + reports[i].user_uuid;
            d.appendChild(a);
            a = document.createElement('p');
            a.innerText = 'Совершенно перелетов: ' + reports[i].flights_made;
            d.appendChild(a);
            main.appendChild(d);

        }

    } else {
        addError("Не удалось получить статистику");
    }
}

async function getFillingStatistics() {
    const url = 'http://3.67.182.34:8060/api/v1/reports/flights-filling';

    let response = await fetch(url, {
        method: 'GET',
        credentials: 'include',
        headers: {
            'Access-Control-Allow-Origin': 'http://3.67.182.34:8887/'
        },
    });
    if (response.ok) {
        main = document.getElementById('container')
        main.innerHTML = '';
        let reports = await response.json();
        console.log(reports)
        for (let i = 0; i < reports.length; i++) {
            d = document.createElement('div');
            d.setAttribute("id", reports[i].id);
            d.setAttribute("class", "flight-item");

            a = document.createElement('p');
            a.innerText = 'ID полета: ' +reports[i].flight_uuid;
            d.appendChild(a);
            a = document.createElement('p');
            a.innerText = 'Билетов куплено: : ' + reports[i].tickets;
            d.appendChild(a);
            main.appendChild(d);

        }

    } else {
        addError("Не удалось получить статистику");
    }
}

async function addFlightBlock() {
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
    inputfield.setAttribute("id", "from-id");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "id точки отправления");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('div');
    d.setAttribute("class", "username");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "from-city");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "Город точки отправления");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('div');
    d.setAttribute("class", "username");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "to-id");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "id точки цели");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('div');
    d.setAttribute("class", "username");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "to-city");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "Город точки цели");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('div');
    d.setAttribute("class", "username");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "date");
    inputfield.setAttribute("id", "date");
    inputfield.setAttribute("class", "user-input");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('button');
    d.setAttribute("class", "signin-button");
    d.setAttribute('onclick', 'AddFlight()');
    d.innerText = 'Создать';
    main.appendChild(d);
}

function ShowBuyButton() {
    let token = localStorage.getItem("token");
    if (token !== null || token !== "") {
        return true;
    }
    return false;
}

async function addFlightPatchBlock() {
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
    inputfield.setAttribute("id", "flight-id");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "id рейса");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('div');
    d.setAttribute("class", "username");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "from-id");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "id точки отправления");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('div');
    d.setAttribute("class", "username");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "from-city");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "Город точки отправления");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('div');
    d.setAttribute("class", "username");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "to-id");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "id точки цели");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('div');
    d.setAttribute("class", "username");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "to-city");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "Город точки цели");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('div');
    d.setAttribute("class", "username");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "date");
    inputfield.setAttribute("id", "date");
    inputfield.setAttribute("class", "user-input");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('button');
    d.setAttribute("class", "signin-button");
    d.setAttribute('onclick', 'PatchFlight()');
    d.innerText = 'Изменить рейс';
    main.appendChild(d);
}
