function ProcessSaveNew(){
    ProcessRequestProductSave()
}

function ProcessUpdate(){
    ProcessRequestProductSave(2)
}

/*
    1 = save new
    2 = update
*/
function ProcessRequestProductSave(command = 1){
    var productName = $("#product-name").val();
    var typeProduct = $("#type").val();
    var status = $("#status").val();
    var priceRentDaily = $("#price-rent-daily").val();
    var priceRentWeekly = $("#price-rent-weekly").val();
    var priceRentMonthly = $("#price-rent-monthly").val();
    var path = $("#path").val();

    if(command == 1){
        promise = sendNewProduct(productName, typeProduct, priceRentDaily, priceRentWeekly, priceRentMonthly, path);
    }
    if(command == 2){
        var url = new URL(window.location.href);
        var productID = url.searchParams.get("product-id");
        promise = sendUpdateProduct(productID, productName, typeProduct, status, priceRentDaily, priceRentWeekly, priceRentMonthly, path);
    }
    
    promise.done(function(response){
        location.reload();
    }).fail(function(response){
        if(response.status==400){
            $("#request-alert-div").css("display", "block");
            $("#request-alert").html("mohon isi seluruh data");
        } else {
            $("#request-alert-div").css("display", "block");
            $("#request-alert").html("silahkan mencoba sekali lagi");
        }
    });
}

function ProcessUploadImg(){
    var url = new URL(window.location.href);
    
    var productID = url.searchParams.get("product-id");
    var blobFile = $("#product-img-new")[0].files[0];
    
    promise = sendUpdateImgProduct(productID, blobFile);
    promise.done(function(response){
        console.log("selesai");
    })
    .fail(function(response){
        console.log("gagal");
    });
}

var isValidDate = function(value, format) {
    format = format || false;
    // lets parse the date to the best of our knowledge
    if (format) {
        value = parseDate(value);
    }

    var timestamp = Date.parse(value);

    return isNaN(timestamp) == false;
}

var parseDate = function(value) {
    var m = value.match(/^(\d{1,2})(\/|-)?(\d{1,2})(\/|-)?(\d{4})$/);
    if (m)
        value = m[5] + '-' + ("00" + m[3]).slice(-2) + '-' + ("00" + m[1]).slice(-2);

    return value;
}