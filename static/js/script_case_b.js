$(document).ready(function () {

    let t1 = document.querySelector('#select-team1');
    let t2 = document.querySelector('#select-team2');
    let date = document.querySelector('#date-val');
    let go = document.querySelector('#go-next');

    push_teams(t1);
    push_teams(t2);

    go.onclick = (e) => setup_go(e, t1, t2, date);

});


function push_teams(placeholder) {
    $.getJSON('/queries/teams/', {}, function (resp) {
        for (let i = 0; i < resp.length; i++) {
            let new_item = document.createElement("option");
            new_item.innerText = resp[i]['Name'];
            new_item.go_id = resp[i]['Id'];
            placeholder.add(new_item);
        }
    });
}

function setup_go(e, t1, t2, date) {
    e.preventDefault();

    let message = {
        'state': 'DateAndTeams',
        'date': date.value,
    };
    for (let i = 0; i < t1.options.length; i++) {
        if (t1.options[i].selected) {
            message['team1'] = t1.options[i].go_id;
        }
        if (t2.options[i].selected) {
            message['team2'] = t2.options[i].go_id;
        }
    }
    if (message['team1'] === message['team2']) {
        alert('Сами с собой?');
        alert('Извините но нет');
        return;
    }

    $.postJSON('/api/', message, function (resp) {
        if (resp === 'OK') {
            $.postJSON('/insert/game/', message, function (resp) {
                if (resp === 'OK')
                    window.open('/C/', "_self");
                else
                    alert('Unexpected response')
            });
        } else {
            alert('Unexpected response')
        }
    });

}
