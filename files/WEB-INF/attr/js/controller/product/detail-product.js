$(document).ready(function(){
    window.dataLayer = window.dataLayer || [];
    function gtag(){dataLayer.push(arguments);}
    gtag('js', new Date());
    gtag('config', 'UA-112381262-1');
    
    $(".form_datetime").datetimepicker({
        format: "dd MM yyyy",
        viewMode: 'months',
        autoclose: true
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
    var userEmail = $("#user-email").val()

    promise = sendRequestProduct(productID, productName, typeDuration, duration, total, startDate, userEmail)

    promise.done(function(){
        console.log("sukses");
    });
}