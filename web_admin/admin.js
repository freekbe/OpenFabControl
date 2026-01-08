// api helper functions
async function post_json_api(data, route) {
  const headers = new Headers();
  headers.set('Content-type', 'application/json');
  headers.set('Accept', 'application/json');

  const request = new Request(route, {
    method: 'POST',
    headers: headers,
    body: JSON.stringify(data)
  });

  let res;
  try {
    res = await fetch(request);
  } catch (e) {
    return { json_data: null, ok: false, status_code: 0 }
  }

  try {
    return {
      json_data: await res.json(),
      ok: res.ok,
      status_code: res.status
    }
  } catch (error) {
    return {
      json_data: null,
      ok: res.ok,
      status_code: res.status
    }
  }
}
async function del_json_api(data, route) {
  const headers = new Headers();
  headers.set('Content-type', 'application/json');
  headers.set('Accept', 'application/json');

  const request = new Request(route, {
    method: 'DELETE',
    headers: headers,
    body: JSON.stringify(data)
  });

  let res;
  try {
    res = await fetch(request);
  } catch (e) {
    return { json_data: null, ok: false, status_code: 0 }
  }

  try {
    return {
      json_data: await res.json(),
      ok: res.ok,
      status_code: res.status
    }
  } catch (error) {
    return {
      json_data: null,
      ok: res.ok,
      status_code: res.status
    }
  }
}
async function get_api(route) {
  const headers = new Headers();
  const request = new Request(route, {
    method: 'GET',
    headers: headers,
    body: null
  });

  let res;
  try {
    res = await fetch(request);
  } catch (e) {
    return { json_data: null, ok: false, status_code: 0 }
  }

  return {
    json_data: await res.json(),
    ok: res.ok,
    status_code: res.status
  }
}
async function create_child_from_file(parent, file) {
  try {
    const response = await fetch(file);
    const html = await response.text();

    const el = createElementFromString(html);
    parent.appendChild(el);
    return (el);
  }
  catch (err) { console.error('Error loading HTML:', err);}
  return (null);
}
function createElementFromString(htmlString) {
  const container = document.createElement('div');
  container.innerHTML = htmlString;
  return (container.firstElementChild);
}

function accept_new_controler(uuid, element) {
  post_json_api({ "uuid": uuid }, '/web-admin-api/approve_machine_controler').then(data => {
    if (data.ok) {
      document.getElementById('machine_controlers_to_approve').removeChild(element);
      element.querySelector('.accept').classList.add("hidden");
      document.getElementById('machine_controlers_approved').appendChild(element);
    }
  });
}
function delete_machine_controler(uuid, element) {
  del_json_api({"uuid": uuid}, '/web-admin-api/delete_machine_controler').then (data => {
    if (data.ok) {
        element.remove();
    }
  })
}

// get the list of machine controlers (approved or not)
async function machine_controlers_list(approved) {

  if (approved == true) { route = '/web-admin-api/ofc_admin/get_machine_controler_list_approved'}
  else                  { route = '/web-admin-api/ofc_admin/get_machine_controler_list_to_approve' }

  get_api(route).then(data => {
    if (!data.ok) { return; }
    console.log(data.json_data);
    if (data.json_data != null) {
      data.json_data.forEach(async (mc) => {
        console.log(mc);
        if (approved == true) {
          element = await create_child_from_file(document.getElementById('machine_controlers_approved'), 'modules/machine_controler.html');
          if (element == null) { return; }
        } else {
          element = await create_child_from_file(document.getElementById('machine_controlers_to_approve'), 'modules/machine_controler.html');
          if (element == null) { return; }
          element.querySelector('.accept').addEventListener('click', () => { accept_new_controler(mc.uuid, element) })
          element.querySelector('.accept').classList.remove("hidden");
        }
        element.querySelector('.name').innerText = mc.name;
        element.querySelector('.uuid').innerText = mc.uuid;
        element.querySelector('.delete').addEventListener('click', () => { delete_machine_controler(mc.uuid, element) })

        // set machine controller logo
        switch (mc.type) {
          case "fm-bv2":
            element.querySelector('.machine_controler_logo').src = "assets/fm-bv2.svg";
            break;
          default:
            element.querySelector('.machine_controler_logo').src = "assets/question-mark-icon.svg";
            break;
        }
      });
    }
  });
}

machine_controlers_list(false)
machine_controlers_list(true)
