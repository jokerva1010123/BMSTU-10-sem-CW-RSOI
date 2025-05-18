async function CheckAdmin() {
    const url = defaultAddr+ ':8010/api/v1/admin';

    token = localStorage.getItem('token');
    if (token === null) {
        return false;
    }
    let response = await fetch(url, {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
            'Access-Control-Allow-Origin': 'http://3.67.182.34:8887',
            'Authorization': `Bearer ${(token)}`
        }
    });
    return response.ok;
}

function addAdminBlock() {
    navigate('/admin');
    main = document.getElementById('container')
    main.innerHTML = '';

    //6 кнопок
    d = document.createElement('div');
    d.innerText = 'Показать пользователей';
    d.setAttribute("class", "airport-items");
    d.setAttribute('onclick', 'seeAllUsers()');
    main.appendChild(d);

    d = document.createElement('div');
    d.innerText = 'Добавить пользователя';
    d.setAttribute("class", "airport-items");
    d.setAttribute('onclick', 'createUserBlock()');
    main.appendChild(d);

    d = document.createElement('div');
    d.innerText = 'Добавить рейс';
    d.setAttribute("class", "airport-items");
    d.setAttribute('onclick', 'addFlightBlock()');
    main.appendChild(d);

    d = document.createElement('div');
    d.innerText = 'Изменить рейс';
    d.setAttribute("class", "airport-items");
    d.setAttribute('onclick', 'addFlightPatchBlock()');
    main.appendChild(d);

    d = document.createElement('div');
    d.innerText = 'Посмотреть статистику по пользователям';
    d.setAttribute("class", "airport-items");
    d.setAttribute('onclick', 'getStatistics()');
    main.appendChild(d);

    d = document.createElement('div');
    d.innerText = 'Посмотреть статистику по полетам';
    d.setAttribute("class", "airport-items");
    d.setAttribute('onclick', 'getFillingStatistics()');
    main.appendChild(d);
}
