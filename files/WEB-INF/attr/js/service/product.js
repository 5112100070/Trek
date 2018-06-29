function getProductForMP(){
    var url = 'http://127.0.0.1:3000/product';

    var promise = $.ajax({
        url: url,
        contentType: 'application/json; char-set=utf-8',
        type: 'GET',
        data: {
            start:0,
            rows:8,
            sort:'ASC'
        },
        headers: {"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "access-control-allow-origin, access-control-allow-headers"},
    });

    return promise;
}