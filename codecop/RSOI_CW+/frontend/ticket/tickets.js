async function GetAll() {
    navigate('/tickets');
    const url = 'http://3.67.182.34:8040/api/v1/tickets';

    let response = await fetch(url, {
        method: 'GET',
        credentials: 'same-origin',
        headers: {
            'Access-Control-Allow-Origin': '*'
        },
    });
    if (response.ok) {
        main = document.getElementById('container')
        main.innerHTML = '';
        for (let i = 0; i < tickets.all.length; i++) {
            d = document.createElement('div');
            d.setAttribute("id", tickets.all[i].id);
            d.setAttribute("class", "ticket-item");

            id = document.createElement('p');
            id.innerText = tickets.all[i].id;
            namef = document.createElement('p');
            namef.innerText = tickets.all[i].flight;
            city = document.createElement('p');
            city.innerText = tickets.all[i].city;

            link = document.createElement('button');
            link.innerText = 'Return';
            link.setAttribute('onclick', 'Return(' + d.id + ')');
            if (tickets.all[i].date < Date()) {
                link.disabled = false;
            }


            d.appendChild(id);
            d.appendChild(namef);
            d.appendChild(city);
            d.appendChild(link);
            main.appendChild(d);
        }

    } else {
        console.log("Ошибка Билетов HTTP: " + response.status);
    }


}

async function Return(id) {
    main = document.getElementById(id);
    main.childNodes[main.childElementCount - 1].disabled = true;

    const url = 'http://3.67.182.34:8040/api/v1/tickets/' + id.toString();
    let response = await fetch(url, {
        method: 'DELETE',
        credentials: 'include',
        headers: {
            'Authorization': `Bearer ${user}`,
            'Access-Control-Allow-Origin': 'http://3.67.182.34:8887'
        },
    });
    if (response.ok) {
        main.parentNode.removeChild(main);
    } else {
        main.childNodes[main.childElementCount - 1].disabled = false;
    }
}

async function GetMyTickets(id) {
    const url = 'http://3.67.182.34:8040/api/v1/tickets';

    user  = localStorage.getItem("token");

    let response = await fetch(url, {
        method: 'GET',
        credentials: 'include',
        headers: {
            'Authorization': `Bearer ${user}`,
            'Access-Control-Allow-Origin': 'http://3.67.182.34:8887'
        },
    });
    if (response.ok) {
        main = document.getElementById('tickets-holder')
        if (main === null) {
            main = document.getElementById('container')
            d = document.createElement('div');
            d.setAttribute("id", "tickets-holder");
            main.appendChild(d);
            main = d;
        } else {
            main.innerHTML = '';
        }
        let tickets = await response.json();
        for (let i = 0; i < tickets.length; i++) {
            d = document.createElement('div');
            d.setAttribute("id", tickets[i].ticket_uuid);
            d.setAttribute("class", "ticket-item");


            //link.addEventListener('onclick', ShowAirport.bind(airports.all[i].id), false);
            id = document.createElement('p');
            id.innerText = 'ID билета: ' + tickets[i].ticket_uuid;
            namef = document.createElement('p');
            namef.innerText = 'ID рейса: ' + tickets[i].flight_uuid;

            link = document.createElement('div');
            link.innerText = 'Return';
            link.setAttribute('onclick', 'Return("' + tickets[i].ticket_uuid + '")');
            link.setAttribute("class", "btn");
            if (tickets[i].date < Date()) {
                link.disabled = false;
            }


            d.appendChild(id);
            d.appendChild(namef);
            d.appendChild(link);
            main.appendChild(d);
        }

    } else {
        console.log("Ошибка Билетов HTTP: " + response.status);
    }


}

async function ShowBuyBlock(id, date) {
    const url = 'http://3.67.182.34:8040/api/v1/tickets';

    user  = localStorage.getItem("token");
    console.log(user);

    let response = await fetch(url, {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
            'Authorization': `Bearer ${user}`,
            'Access-Control-Allow-Origin': 'http://3.67.182.34:8887/'
        },
        body: JSON.stringify({flight_uuid: id, date: date}),
    });
    if (response.ok) {
        render("/user");

    } else {
        addError("Невозможно купить билеты")
    }
}