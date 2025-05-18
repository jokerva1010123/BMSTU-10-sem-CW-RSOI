async function login() {
   let user = document.getElementById("username").value;
   let user_pass = document.getElementById("password").value;
   const url = defaultAddr + ':8010/api/v1/auth';

   let auth = user +':' + user_pass;
   let response = await fetch(url, {
      method: 'POST',
      credentials: 'include',
      headers: {
         'Authorization': `Basic ${(btoa(auth))}`,
         'Access-Control-Allow-Origin': defaultOrigin,
      },
   });
   if (response.ok) {
      let body = await response.json();
      localStorage.setItem('token', body.Token);
      localStorage.setItem('user', user);


      b1 = document.getElementById('login');
      b1.parentNode.removeChild(b1);

      d = document.getElementById('navbar');

      let admin = await CheckAdmin();
      if (admin === true) {
         a = document.createElement('a');
         a.innerText = 'Админка';
         a.setAttribute('id', 'admin');
         a.setAttribute('onclick', 'addAdminBlock()');
         d.appendChild(a);
      }

      a = document.createElement('a');
      a.innerText =  user;
      a.setAttribute('onclick', 'addUserBlock()');
      d.appendChild(a);
      render('/');

      } else {
         addError("Unauth");
   }
}

async function signup() {
   let user = document.getElementById("username").value;
   let user_pass = document.getElementById("password").value;
   let role = document.getElementById("role").value;


      const url = defaultAddr +':8010/api/v1/users';
      const data = {login: user, password: user_pass, role: role};

      let response = await fetch(url, {
         method: 'POST',
         credentials: 'same-origin',
         headers: {
            'Content-Type': 'application/json',
         },
         body: JSON.stringify(data)
      });
      if (response.ok) {
         let user = await response.json();
         render("/");

      } else {
         addError("Не получилось создать пользователя")
      }

}

function utf8_to_b64(str) {
   return window.btoa(unescape(encodeURIComponent(str)));
}

function b64_to_utf8(str) {
   return decodeURIComponent(escape(window.atob(str)));
}

async function isAuthed() {

   const url = defaultAddr+ ':8010/api/v1/verify';

   token = localStorage.getItem('token');
   if (token == null) {
      return false;
   }
   let response = await fetch(url, {
      method: 'POST',
      credentials: 'same-origin',
      headers: {
         'Access-Control-Allow-Origin': defaultOrigin,
         'Authorization': `Bearer ${(token)}`
      }
   });
   return response.ok;
}
