function ProcessRegisterCompany(){
    var companyType = $("#company-type").val();
    var companyName = $("#company-name").val();

    var promise = RegisterCompany(companyType, companyName);
    promise.done(function(response){
        window.location.href = base_url + "/dashboard";
    }).fail(function(response){
        if(response.status >= 500){
            $("#request-alert-div").css("display", "block");
            $("#request-alert").html("silahkan mencoba sekali lagi");
        } else {
            $("#request-alert-div").css("display", "block");
            $("#request-alert").html(response.responseJSON.data);
        }      
    });

    $("#company-name").val("");
}