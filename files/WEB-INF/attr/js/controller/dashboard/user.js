function ProcessLogout(){
    StartLoading();
    promise = makeLogout();

    promise.done(function(response){
        FinishLoading();
        if(response.error!=null){
            setErrorMessage(true, response.error.massage);
        } else {
            setErrorMessage(false, "");
            window.location.href = base_url;
        }
    }).fail(function(response){
        FinishLoading();
        setErrorMessage(true,  "Mohon maaf sedang terjadi kesalahan. Silahkan coba lagi.");
    });
}

function setErrorMessage(isShow, message = "") {
    if (isShow) {
        $("#error_message").addClass("form_error");
        $("#error_message").removeClass("hide");
        $("#error_message").html(message);
    } else {
        $("#error_message").addClass("hide");
        $("#error_message").removeClass("form_error");
        $("#error_message").html("");
    }
}

function ProcessRegister(){
    if($("#password").val() != $("#password-ver").val()){
        SimpleSwal("Daftar Menjadi Anggota", "kesalahan pada verifikasi password", type_error, "Tutup");
        return;
    } 

    fullname = $("#fullname").val();
    email = $("#email").val();
    password = $("#password").val();

    if(password.length <= 9){
        SimpleSwal("Daftar Menjadi Anggota", "Password kurang dari 10 karakter", type_error, "Tutup");
        return;
    }

    if(!validateEmail(email)){
        SimpleSwal("Daftar Menjadi Anggota", "Email tidak valid", type_error, "Tutup");
        return;
    }

    promise = registerUser(fullname, email, password);
    
    var swal = SwalConfimationProcess("Daftar Menjadi Anggota", "Pastikan data yang akan anda daftarkan benar", type_warning, "Daftar!", "Batal");

    swal.then(function(){
        promise.done(function(response){
            var swalDone = SimpleSwal("Daftar Menjadi Anggota", "Konfirmasi pendaftaran telah dikirim ke " + email, type_success, "Selesai");
            swalDone.then(function(){
                $("#fullname").val("");
                $("#email").val("");
                $("#password").val("");
                $("#password-ver").val("");
            });
        }).fail(function(response){
            $("#fullname").val("");
            $("#email").val("");
            $("#password").val("");
            $("#password-ver").val("");
            if(response.responseJSON != undefined){
                SimpleSwal("Daftar Menjadi Anggota", response.responseJSON.data.error_message, type_error, "Tutup");                
            } else {
                SimpleSwal("Daftar Menjadi Anggota", "silahkan mencoba sekali lagi", type_error, "Tutup");
            }
        });
    });
    
}

function ProcessResetPassword(){
    email = $("#email").val();
    if(!validateEmail(email)){
        SimpleSwal("Atur Ulang Password", "Email tidak valid", type_error, "Tutup");
        return;
    }

    promise = resetPasssword(email);
    promise.done(function(response){
        var swal = SimpleSwal("Atur Ulang Password", "Sukses mengatur ulang passsword. Silahkan cek email", type_success, "Kembali ke halaman utama");
        swal.then(function(){
            window.location.href = base_url;
        });
    }).fail(function(response){
        $("#email").val("");
        if(response.status >= 500){
            SimpleSwal("Atur Ulang Password", "Terjadi kesalahan. Silahkan coba lagi. Silahkan cek email", type_error, "Tutup");
        } else {
            SimpleSwal("Atur Ulang Password", response.responseJSON.data.error_message, type_error, "Tutup");
        }
    });
    
}

function validateEmail(email) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}