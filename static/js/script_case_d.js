let sel_pers = null;
let sel_cur_teams = null;
let QUES = "Сколько прошло с начала матча";

$(document).ready(function () {

    let container = document.querySelector('.main');
    let yes = document.querySelector('#yes');
    let no = document.querySelector('#no');

    push_items(container);

    config_imgs(yes, no);

});


function push_items(placeholder) {
    $.getJSON('/queries/items/', {}, function (resp) {
        for (let i = 0; i < resp.length; i++) {

            let new_item = document.createElement("div");
            new_item.className = "item";
            new_item.go_id = resp[i].Id;
            placeholder.insertBefore(new_item, null);

            let header = document.createElement("p");
            header.innerText = resp[i].Alias;
            new_item.insertBefore(header, null);

            let gap = make_gap(resp[i].Type);
            new_item.insertBefore(gap, null);

        }
    });
}

function make_gap(type) {

    let res = eval(type);

    let new_el = document.createElement({
        PERS: 'select',
        INT: 'input',
        ENUM: 'select',
        STR: 'input',
        TEAM: 'select',
    }[res[0]]);
    new_el.className = "__val";

    _ = {
        PERS: __select_of_pers,
        INT: __input_int,
        ENUM: __select_enum,
        STR: __input_str,
        TEAM: __select_team,
    }[res[0]](new_el, res[1]);

    return new_el;

}

function __select_of_pers(new_el, _) {
    function wrapper() {
        for (let i = 0; i < sel_pers.length; i++) {
            let no = document.createElement("option");
            no.innerText = sel_pers[i].Name;
            new_el.add(no);
        }
    }

    if (sel_pers === null) {
        $.getJSON('/queries/persons/', {}, function (resp) {
            sel_pers = resp;
            wrapper();
        });
    } else {
        wrapper();
    }
}

function __input_int(new_el, _) {
    new_el.type = 'number';
}

function __input_str(new_el, _) {
    new_el.type = 'text';
}

function __select_enum(new_el, params) {

    for (let i = 0; i < params.length; i++) {
        let no = document.createElement('option');
        no.innerText = params[i];
        new_el.add(no);
    }
}

function __select_team(new_el, params) {

    function wrapper() {
        let no = document.createElement("option");
        no.innerText = sel_cur_teams['Team1'];
        new_el.add(no);

        no = document.createElement("option");
        no.innerText = sel_cur_teams['Team2'];
        new_el.add(no);
    }

    if (sel_cur_teams === null) {
        $.getJSON('/queries/cur-teams/', {}, function (resp) {
            sel_cur_teams = resp;
            wrapper();
        });
    } else {
        wrapper();
    }
}

function config_imgs(yes, no) {
    no.onclick = () => {
        window.open('/C/', '_self')
    };

    yes.onclick = () => {
        let time = prompt(QUES);
        let js = collect_json();
        alert(js);
        post_json(js, time);
    };
}

function collect_json() {
    let message = {};
    let main = document.querySelector('.main');
    for (let i = 0; i < main.children.length; i++) {
        message[main.children[i].go_id] = main.children[i]
            .querySelector('.__val')
            .value
    }
    return JSON.stringify(message);
}

function post_json(json, time) {
    let message = {
        'apply': json,
        'time': time,
    };
    $.postJSON('/insert/events/', message, function (resp) {
        if (resp === 'OK') {
            alert('Успех!');
            window.open('/C/', '_self');
        } else {
            alert('Неведомая ошибка. Ответ сервера: ' + resp)
        }
    });
}
