function getProductForMP(totalRequested = 8){
    var url = product_url + '/product';

    var promise = $.ajax({
        url: url,
        type: 'GET',
        data: {
            start:0,
            rows:totalRequested,
            sort:'ASC'
        },
        xhrFields: {
            withCredentials: true
         },
    });

    return promise;
}

function getProductDetailByID(productID){
    var url = product_url + '/product/detail-by-id';

    var promise = $.ajax({
        url: url,
        type: 'GET',
        data: {
            "product-id":productID
        },
        xhrFields: {
            withCredentials: true
         },
    });

    return promise;
}

function getProductDetail(){
    var url = product_url + '/product/detail';

    var promise = $.ajax({
        url: url,
        type: 'GET',
        data: {
            path:productPath
        },
        xhrFields: {
            withCredentials: true
         },
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

function sendNewProduct(productName, type, priceRentDaily, priceRentWeekly, priceRentMonthly, path){
    return sendProduct(0, productName, type, 1, priceRentDaily, priceRentWeekly, priceRentMonthly, path);
}

function sendUpdateProduct(productID, productName, type, status, priceRentDaily, priceRentWeekly, priceRentMonthly, path){
    return sendProduct(productID, productName, type, status, priceRentDaily, priceRentWeekly, priceRentMonthly, path);
}

function sendProduct(productID, productName, type, status, priceRentDaily, priceRentWeekly, priceRentMonthly, path){
    var url = product_url + '/product/save';
    var data = {
        product_id:productID,
        product_name:productName,
        type: type,
        status: status,
        price_rent_daily: priceRentDaily,
        price_rent_weekly: priceRentWeekly,
        price_rent_monthly: priceRentMonthly,
        path: path
    };

    var promise = $.ajax({
        url: url,
        type: 'POST',
        data: data,
        xhrFields: {
            withCredentials: true
         },
    });

    return promise;
}

function sendUpdateImgProduct(productID, img){
    var url = product_url + '/product/upload-image';
    
    var data = new FormData();
    data.append("product_id", productID)
    data.append("product_img", img)

    var promise = $.ajax({
        url: url,
        type: 'POST',
        data: data,
        contentType: false,
        processData: false,
        xhrFields: {
            withCredentials: true
         },
     });

     return promise;
}