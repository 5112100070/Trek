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

function ProcessChangePassword(){
    var tokenOld = $("#token-old").val();
    var token = $("#token").val();
    var tokenVerification = $("#token-verification").val();

    if(tokenOld == ""){
        SimpleSwal("Ganti Password", "Masukkan Password lama anda", type_error, "Tutup");
        return 
    }

    if(token != tokenVerification){
        SimpleSwal("Ganti Password", "Password Baru dan Password Verifikasi salah", type_error, "Tutup");
        return 
    }

    var swal = SwalConfimationProcess("Ganti Password", "Lanjutkan ?", type_question, "Ganti Password", "Batal");
    swal.then(function(){
        promise = changePassword(tokenOld, token, tokenVerification);
        promise.done(function(response){
            SimpleSwal("Ganti Password", "Berhasil melakukan penggantian password", type_success, "OK");
        }).fail(function(response){
            if(response.responseJSON != undefined){
                SimpleSwal("Ganti Password", response.responseJSON.data.message, type_error, "Tutup");
            } else {
                SimpleSwal("Ganti Password", "silahkan mencoba sekali lagi", type_error, "Tutup");
            }
        });
    });


    $("#token-old").val("");
    $("#token").val("");
    $("#token-verification").val("");
}