let container = null;

$(document).ready(function () {

    container = document.querySelector('.main');
    new_elem();

});

function new_elem() {
    let new_item = document.createElement("input");
    container.insertBefore(new_item, null);
}

function del_elem() {
    container.lastChild.remove();
}

function make_new_team(members) {
    let forName = document.querySelector('input[name$="Name"]');
    let message = {
        'Members': members,
        'Name': forName.value,
    };
    message = JSON.stringify(message);
    console.log(message);
    $.postJSON('/insert/team/', message, function (resp) {
        if (resp === 'OK') {
            alert('Успех');
            window.open('/A/', '_self')
        } else {
            alert('Что-то пошло не так: ');
            alert(resp);
        }
    });
}

function send_and_exit() {
    let result = [];
    let n = container.children.length;
    for (let i = 0; i < n; i++) {
        if (container.children[i].value) {
            result[result.length] = container.children[i].value;
        }
    }
    make_new_team(result);
}

document.addEventListener('keydown', (event) => {
    const keyName = event.key;
    if (keyName === '+') {
        new_elem();
    } else if (keyName === '-') {
        del_elem();
    } else if (keyName === 'Enter') {
        send_and_exit();
    }
}, false);
