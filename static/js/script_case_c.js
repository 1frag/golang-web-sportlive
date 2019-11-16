$(document).ready(function () {

    let container = document.querySelector('.main');

    push_events(container);

});


function push_events(placeholder) {
    $.getJSON('/queries/events/', {}, function (resp) {
        for (let i = 0; i < resp.length; i++) {

            let new_item = document.createElement("div");
            new_item.className = "item";
            new_item.innerText = resp[i].Name;
            new_item.onclick = make_ocl(resp[i]['Id']);
            placeholder.insertBefore(new_item, null);
        }
    });
}

function make_ocl(id) {
    return function () {
        let message = {
            'state': 'ChangeEvent',
            'event': id,
        };
        $.postJSON('/api/', message, function (resp) {
            if (resp === 'OK') {
                window.open('/D/', '_self')
            } else {
                alert('Case-c failed')
            }
        })
    }
}
