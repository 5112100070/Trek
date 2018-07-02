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