$(document).ready(function () {

});

function magic(placeholder) {
    let u = document.getElementsByName("username")[0];
    let p = document.getElementsByName("password")[0];
    if (u.value === '' || p.value === '') {
        alert('Необходимо заполнить форму');
        return;
    }
    $.postJSON('/app/login/', {
        'username': u.value,
        'password': p.value,
    }, function (resp) {
        resp = resp['resp'];
        if (resp === 2) {
            do_query(2)
        } else if (resp === 1) {
            do_query(1)
        } else {
            alert(401);
        }
    });
}

function do_query(right) {
    message = {
        'query': document.getElementsByName('query')[0].value,
        'right': right,
    };
    $.postJSON('/app/row_sql/', message, function (resp) {
        alert('Ответ запроса: ' + Object.entries(resp));
    })
}
