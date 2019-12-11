$(document).ready(function () {
    let B = document.querySelector(".basket");
    B.onclick = function () {
        let message = {
            'state': 'InitGame',
            'game': 'Баскетбол',
        };
        $.postJSON('/api/', message, function (resp) {
            if (resp === "OK") {
                window.open('/B/', "_self");
            } else {
                alert("что-то пошло не так");
            }
        });
    };

    let F = document.querySelector(".football");
    F.onclick = function () {
        let message = {
            'state': 'InitGame',
            'game': 'Футбол',
        };
        $.postJSON('/api/', message, function (resp) {
            if (resp === "OK") {
                window.open('/B/', "_self");
            } else {
                alert("что-то пошло не так");
            }
        });
    };

    let H = document.querySelector(".history");
    H.onclick = function () {
        let message = {
            'state': 'ResetGame',
        };
        $.postJSON('/api/', message, function (resp) {
            if (resp === "OK") {
                window.open('/H/', "_self");
            } else {
                alert("что-то пошло не так");
            }
        });
    };

});