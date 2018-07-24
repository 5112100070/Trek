function getProductForMP(totalRequested = 8){
    var url = product_url + '/product';

    var promise = $.ajax({
        url: url,
        contentType: 'application/json; char-set=utf-8',
        type: 'GET',
        data: {
            start:0,
            rows:totalRequested,
            sort:'ASC'
        },
        headers: {"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "access-control-allow-origin, access-control-allow-headers"},
    });

    return promise;
}

function getProductDetail(){
    var url = product_url + '/product/detail';

    var promise = $.ajax({
        url: url,
        contentType: 'application/json; char-set=utf-8',
        type: 'GET',
        data: {
            path:productPath
        },
        headers: {"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "access-control-allow-origin, access-control-allow-headers"},
    });

    return promise;
}

function sendRequestProduct(productID, productName, typeDuration, duration, total, startDate, email, projectAddress){
    var url = base_url + '/send-request-item';
    var data = {
        product_id:productID,
        product_name:productName,
        type_duration: typeDuration,
        duration: duration,
        total: total,
        start_date: startDate,
        email: email,
        project_address: projectAddress
    };

    var promise = $.ajax({
        url: url,
        type: 'POST',
        data: data
    });

    return promise;
}