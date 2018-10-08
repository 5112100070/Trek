function changePassword(tokenOld, token, tokenVerification){
    var url = product_url + '/user/change-password';
    var data = {
        tokenOld:tokenOld,
        token: token,
        tokenVerification: tokenVerification
    };
    
    var promise = $.ajax({
        url: url,
        type: 'POST',
        data: data,
        xhrFields: {
            withCredentials: true
        }
    });

    return promise;
}