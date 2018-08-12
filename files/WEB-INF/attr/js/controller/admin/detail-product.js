$(document).ready(function(){
    $("#datetimepicker1").datetimepicker({
        format: 'dd-mm-yyyy',
			icons: {
				date: "fa fa-calendar",
				up: "fa fa-arrow-up",
				down: "fa fa-arrow-down"
			}
		}).find('#start-date').on("blur",function () {
        // check if the date is correct. We can accept dd-mm-yyyy and yyyy-mm-dd.
        // update the format if it's yyyy-mm-dd
        var date = parseDate($(this).val());

        //if (! isValidDate(date)) {
            //create date based on momentjs (we have that)
            date = moment(date).format('YYYY-MM-DD');
        //}

        $(this).val(date);
    });

    var url = new URL(window.location.href);
    loadProductDetailByID(url.searchParams.get("product-id"));
});

function ProcessUpdate(){
    var url = new URL(window.location.href);

    var productID = url.searchParams.get("product-id");
    var productName = $("#product-name").val();
    var typeProduct = $("#type").val();
    var status = $("#status").val();
    var priceRentDaily = $("#price-rent-daily").val();
    var priceRentWeekly = $("#price-rent-weekly").val();
    var priceRentMonthly = $("#price-rent-monthly").val();
    var path = $("#path").val();

    promise = sendUpdateProduct(productID, productName, typeProduct, status, priceRentDaily, priceRentWeekly, priceRentMonthly, path);

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