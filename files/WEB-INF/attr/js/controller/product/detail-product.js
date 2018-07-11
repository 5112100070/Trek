$(document).ready(function(){
    window.dataLayer = window.dataLayer || [];
    function gtag(){dataLayer.push(arguments);}
    gtag('js', new Date());
    gtag('config', 'UA-112381262-1');
    
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
        console.log($(this).val());
        var date = parseDate($(this).val());
        console.log(date);

        //if (! isValidDate(date)) {
            //create date based on momentjs (we have that)
            date = moment(date).format('YYYY-MM-DD');
        //}

        $(this).val(date);
    });
    loadProductDetail()
});

function SendRequestQuot(){
    var productName = $("#product-name").html();
    var productID = $("#product-id").val();
    var typeDuration = $("#type-duration").val();
    var duration = $("#duration").val();
    var total = $("#total").val();
    var startDate = $("#start-date").val();
    var userEmail = $("#user-email").val();
    var projectAddress = $("#project-address").val();

    promise = sendRequestProduct(productID, productName, typeDuration, duration, total, startDate, userEmail, projectAddress);

    promise.done(function(response){
        window.location.href = base_url + '/thank-you'
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