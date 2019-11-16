$(document).ready(function () {

    let container = document.querySelector('.main');

    push_links(container);

});

function push_links(placeholder) {
    let id = placeholder.getAttribute("__tag");
    $.getJSON('/queries/history-events/?id=' + id, {}, function (resp) {
        for (let i = 0; i < resp.length; i++) {

            let new_item = document.createElement("p");
            new_item.innerText = make(resp[i]);
            placeholder.insertBefore(new_item, null);
        }
    });
}

function my_join(sep, m, list) {
    let result = "";
    for (let i = 0; i < list.length; i++) {
        if (list[i].substr(0, 2) === ';;') {
            result = result + m[list[i].substr(2)] + sep;
        } else
            result = result + list[i] + sep;
    }
    return result;
}

function make(resp) {
    let m = redialect(resp.Detail);
    let prefix = resp.Event.Time + '. ';
    return prefix + my_join(' ', m, {
        'Гол': [';;Кто', 'забил гол, когда на воротах был', ';;Вратарь'],
        'Замена': ['Замена в команде', ';;Команда', "Кто встает на замену",
            "вместо", ";;Кто уходит"],
        'Бросок': ["Успешный бросок игрока", ";;Кто", "получил",
            ";;Количество очков", "очков"],
        'Пенальти': ["Назначено пенальти, бьёт", ";;Кто бьёт",
            "на воротах", ";;Вратарь"],
        'Карточка': [";;Цвет карточки", "карточка получена игроком",
            "Кому", "по причине:", ";;Причина"],
        'Победа': ["Победила команда", ";;Победившая Команда", "со счетом",
            ";;Счет"],
        'Штрафной бросок': ["Из-за фола игрока", ";;Чей фол",
            "назначен штрафной бросок, который исполнял", ";;Кто бросает",
            "реализовав возможность со счетом", ";;Результат"],
    }[resp.Event.Name]);
}

function redialect(list) {
    let retval = {};
    for (let i = 0; i < list.length; i++) {
        alert(list[i].Alias + ' - ' + list[i].Value);
        retval[list[i].Alias] = list[i].Value;
    }
    return retval;
}
