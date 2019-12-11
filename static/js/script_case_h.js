$(document).ready(function () {

    let container = document.querySelector('.main');

    push_links(container);

});

function push_links(placeholder) {
    $.getJSON('/queries/games/', {}, function (resp) {
        for (let i = 0; i < resp.length; i++) {

            let new_item = document.createElement("a");

            new_item.innerText = resp[i].GameDate.replace('T00:00:00Z', '') +
                ' матч по ' + resp[i].Kind + 'у команды `' + resp[i].Teams[0] +
                '` против команды `' + resp[i].Teams[1] + '`';
            new_item.href = '/G/?id=' + resp[i].GameId;

            placeholder.insertBefore(new_item, null);

            placeholder.insertBefore(document.createElement("br"), null);
        }
    });
}
