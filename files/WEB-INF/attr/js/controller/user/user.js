var login_error_field = document.getElementById("login_error");

function ProcessLogin(){
    StartLoading();
    email = $("#email").val();
    password = $("#password").val();

    promise = makeLogin(email, password);

    promise.done(function(response){
        FinishLoading();
        if(response.error!=null){
            setErrorLogin(true, response.error.massage);
            $("#password").val("");
        } else {
            var swal = SimpleSwalWithoutDesc("Login Berhasil", type_success, "Ke Beranda");
            swal.then(function(){
                window.location.href = base_url + '/dashboard';
            });
        }
    }).fail(function(response){
        FinishLoading();
        $("#password").val("");
        SimpleSwal("Login Gagal", "Mohon maaf sedang terjadi kesalahan. Silahkan coba lagi.", type_error, "Tutup");
    });
}

function setErrorLogin(isShow, message = "") {
    $("#login_error").html(message);
  if (isShow) {
    $("#login_error").addClass("error");
    $("#login_error").removeClass("hide");
  } else {
    $("#login_error").addClass("hide");
    $("#login_error").removeClass("error");
  }
}

function ProcessLogout(){
    StartLoading();
    promise = makeLogout();

    promise.done(function(response){
        FinishLoading();
        window.location.href = base_url;
    }).fail(function(response){
        FinishLoading();
        location.reload();
    });
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

    StartLoading();
    promise = registerUser(fullname, email, password);
    
    var swal = SwalConfimationProcess("Daftar Menjadi Anggota", "Pastikan data yang akan anda daftarkan benar", type_warning, "Daftar!", "Batal");

    swal.then(function(){
        promise.done(function(response){
            FinishLoading();
            var swalDone = SimpleSwal("Daftar Menjadi Anggota", "Silahkan cek email untuk konfirmasi pendaftaran", type_success, "Pindah ke beranda");
            swalDone.then(function(){
                window.location.href = base_url;
            });
        }).fail(function(response){
            FinishLoading();
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