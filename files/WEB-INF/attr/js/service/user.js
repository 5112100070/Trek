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

function getUserByID(userID){
    var url = product_url + '/user/detail';

    var promise = $.ajax({
        url: url,
        contentType: 'application/json; char-set=utf-8',
        type: 'GET',
        data: {
            user_id:userID
        },
        headers: {"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "access-control-allow-origin, access-control-allow-headers"},
    });

    return promise;
}

function getUserList(totalRequested = 8, typeSort = 'ASC'){
    var url = product_url + '/user';

    var promise = $.ajax({
        url: url,
        contentType: 'application/json; char-set=utf-8',
        type: 'GET',
        data: {
            start:0,
            rows:totalRequested,
            sort:typeSort
        },
        headers: {"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "access-control-allow-origin, access-control-allow-headers"},
    });

    return promise;
}

function sendNewUser(fullname, username, password, status, typeUser){
    return sendUser(0, fullname, username, password, status, typeUser);
}

function sendUpdateUser(userID, fullname, username, password, status, typeUser){
    return sendUser(userID, fullname, username, password, status, typeUser);
}

function sendUser(userID, fullname, username, password, status, typeUser){
    var url = product_url + '/user/save';
    var data = {
        user_id:userID,
        fullname:fullname,
        username:username,
        password:password,
        type: typeUser,
        status: status
    };
    
    var promise = $.ajax({
        url: url,
        type: 'POST',
        data: data
    });

    return promise;
}

function sendUpdateImgUser(userID, img){
    var url = product_url + '/user/upload-image';
    
    var data = new FormData();
    data.append("user_id", userID)
    data.append("user_img", img)

    var promise = $.ajax({
        url: url,
        type: 'POST',
        data: data,
        contentType: false,
        processData: false,
     });

     return promise;
}