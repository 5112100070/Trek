$(document).ready(function(){
    window.dataLayer = window.dataLayer || [];
    function gtag(){dataLayer.push(arguments);}
    gtag('js', new Date());
    gtag('config', 'UA-112381262-1');
    
    $("#start-date").datetimepicker({
        format: 'dd/mm/yyyy'
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