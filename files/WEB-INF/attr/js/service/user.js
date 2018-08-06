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