let OK = "OK";

jQuery.DoQuery = function(method) {
    return function (url, args, callback) {
        $.ajax({
            url: url, data: $.param(args), dataType: "text", type: method,
            success: function (response) {
                response = eval('(' + response + ')');
                if (callback) callback(response);
            }, error: function (response) {
                console.log("ERROR:", response);
            }
        });
    };
};

jQuery.postJSON = jQuery.DoQuery("POST");
jQuery.getJSON = jQuery.DoQuery("GET");

function insertAfter(newNode, referenceNode) {
    referenceNode.parentNode.insertBefore(newNode, referenceNode.nextSibling);
}
