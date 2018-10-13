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

function getCompanyMember(){
    var url = product_url + '/company/get-member';

    var data = {
        rows : -1,
    };

    var promise = $.ajax({
        url: url,
        type: 'GET',
        xhrFields: {
            withCredentials: true
        }
    });

    return promise;
}