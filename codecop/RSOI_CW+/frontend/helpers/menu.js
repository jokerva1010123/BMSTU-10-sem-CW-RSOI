window.onload = addMenu;

const defaultAddr = "http://3.67.182.34"//"http://3.67.182.34";
const defaultOrigin = "http://3.67.182.34:8887" //"http://3.67.182.34:8887"; //
const buttons = ["Рейсы", "Аэропорты"];
const functions = ["ShowAllFlights()", "ShowAllAirports()"];
async function addMenu() {
    main = document.getElementById('main')
    main.innerHTML = '';

    d = document.createElement('div');
    d.setAttribute("id", "navbar");

    for (let i = 0; i < buttons.length; i++) {
        a = document.createElement('a');
        a.setAttribute("onclick", functions[i]);
        a.innerText = buttons[i];
        d.appendChild(a);
    }

    let check = await isAuthed();
    if (check === false) {
        a = document.createElement('a');
        a.innerText = 'Вход';
        a.setAttribute('id', 'login');
        a.setAttribute('onclick', 'addLoginBlock()');
        d.appendChild(a);
    } else {
        a = document.createElement('a');
        a.innerText = 'User :)';
        a.setAttribute('onclick', 'addUserBlock()');
        a.setAttribute("class", "user button");
        d.appendChild(a);
    }
    main.appendChild(d);
    d = document.createElement('div');
    d.setAttribute("id", "container");
    main.appendChild(d);
}

function getCookie(name) {
    if (document.cookie.split(';').filter((item) => item.trim().startsWith(name)).length) {
        return 'User';
    }
    return undefined;
}

function navigate(path) {
    window.history.pushState(path, path, path);
}

window.onpopstate = function (event) {
    event.preventDefault();
    state = event.state === null? 'state' : event.state;

    render(state);
};


function render(state) {
    switch (state) {
        case '/user':
            addUserBlock();
            break;
        case '/airports':
            ShowAllAirports();
            break;
        case '/admin':
            addAdminBlock();
            break;
        case '/tickets':
            GetAll();
            break;
        default:
            ShowAllFlights();
            break;
    }
}

function addError(text) {
    err = document.getElementById('error');
    if (err === undefined || err === null) {
        d = document.createElement('div');
        d.innerText = text;
        d.setAttribute("id", "hideMe");
        main = document.getElementById('container');
        main.prepend(d);
    } else {
        err.innerText = text;
    }
}