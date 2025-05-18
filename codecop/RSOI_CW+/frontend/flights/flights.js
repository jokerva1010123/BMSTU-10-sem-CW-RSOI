async function ShowAllFlights() {
    main = document.getElementById('container')
    main.innerHTML = '';

    navigate('/tickets');
    const url = 'http://3.67.182.34:8030/api/v1/flights';

    let response = await fetch(url, {
        method: 'GET',
        credentials: 'same-origin',
        headers: {
            'Access-Control-Allow-Origin': 'http://3.67.182.34:8887/'
        },
    });
    if (response.ok) {
        main = document.getElementById('container')
        main.innerHTML = '';
        let body = await response.json();
        for (let i = 0; i < body.length; i++) {
            d = document.createElement('div');
            d.setAttribute("id", body[i].id);
            d.setAttribute("class", "flight-item");

            a = document.createElement('p');
            a.innerText = 'Город отправления: ' + body[i].from_city;
            a.setAttribute("id", body[i].from);
            //a.setAttribute('onclick', 'ShowOneFlight("' +body[i].from + '")')
            d.appendChild(a);

            a = document.createElement('p');
            a.innerText = 'Город прибытия: ' + body[i].to_city;
            a.setAttribute("id", body[i].to);
            //a.setAttribute('onclick', 'ShowOneFlight("' +body[i].to + '")');
            d.appendChild(a);

            a = document.createElement('p');
            a.innerText = 'Дата вылета: ' + body[i].date.slice(0, 11);
            d.appendChild(a);
            let check = await isAuthed();
            if (check === true) {
                a = document.createElement('button');
                a.setAttribute("class", "signin-button");
                a.setAttribute('onclick', 'ShowBuyBlock("' + body[i].id +'","' +body[i].date +  '")');
                a.innerText = 'Купить';
                d.appendChild(a);
            }

            main.appendChild(d);
        }

    } else {
        addError('Рейсы не доступны');
    }

}


async function PatchFlight() {
    let id = document.getElementById("flight-id").value;
    let from_id = document.getElementById("from-id").value;
    let from_city = document.getElementById("from-city").value;

    let to_id = document.getElementById("to-id").value;
    let to_city = document.getElementById("to-city").value;

    let date = document.getElementById("date").value;

    const url = 'http://3.67.182.34:8030/api/v1/flights/'+id;

    const data = {from: from_id, from_city: from_city,
        to: to_id, to_city: to_city, date: date+ 'T00:00:00.000Z'};

    let response = await fetch(url, {
        method: 'PATCH',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
    });
    if (response.ok) {
        render('/');
    } else {
       addError("Полёт не найден")
    }

}


async function AddFlight() {
    let from_id = document.getElementById("from-id").value;
    let from_city = document.getElementById("from-city").value;

    let to_id = document.getElementById("to-id").value;
    let to_city = document.getElementById("to-city").value;

    let date = document.getElementById("date").value;

    const url = 'http://3.67.182.34:8030/api/v1/flights';
    const data = {from: from_id, from_city: from_city,
        to: to_id, to_city: to_city, date: date+ 'T00:00:00.000Z'};

    let response = await fetch(url, {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
    });
    if (response.ok) {
        render("/");
    } else {
        addError("Не получилось создать рейс")
    }
}
