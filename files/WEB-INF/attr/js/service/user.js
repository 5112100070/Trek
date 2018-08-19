function makeLogin(username,secret){
    var url = base_url + '/login';
    var data = {
        username:username,
        secret:secret
    };

    var promise = $.ajax({
        url: url,
        type: 'POST',
        data: data
    });

    return promise;
}

function makeLogout(){
    var url = base_url + '/logout';
    
    var promise = $.ajax({
        url: url,
        type: 'POST',
        withCredentials: true,
        headers: {"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "access-control-allow-origin, access-control-allow-headers"},
    });

    return promise;
}